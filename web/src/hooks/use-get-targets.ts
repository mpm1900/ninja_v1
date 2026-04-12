import type { Context } from '#/lib/game/context'
import { clientsStore } from '#/lib/stores/clients'
import {
  sendContextMessage,
  subscribeSocketMessages,
} from '#/lib/stores/socket'
import { useStore } from '@tanstack/react-store'
import { useEffect, useMemo, useRef, useState } from 'react'

type GetTargetsState = {
  context: Context | null
  loading: boolean
}

function contextScopeKey(context: Context): string {
  const targetPositionIDs = context.target_position_IDs ?? []
  return [
    context.action_ID ?? '',
    context.source_player_ID ?? '',
    context.parent_actor_ID ?? '',
    context.source_actor_ID ?? '',
    targetPositionIDs.join('|'),
  ].join('::')
}

function useGetTargets(context: Context, prompt_ID?: string) {
  const client = useStore(clientsStore, (c) => c.me)

  const [state, setState] = useState<GetTargetsState>({
    context: null,
    loading: false,
  })

  const latestScopeRef = useRef<string>('')
  const latestRequestRef = useRef(0)

  const requestScope = useMemo(() => contextScopeKey(context), [context])

  useEffect(() => {
    return subscribeSocketMessages((_, message) => {
      if (message?.type !== 'target-IDs' || !message.context) return

      // Ignore responses from unrelated get-targets requests.
      const responseScope = contextScopeKey(message.context)
      if (responseScope !== latestScopeRef.current) return

      setState({
        context: message.context,
        loading: false,
      })
    })
  }, [])

  useEffect(() => {
    if (!client) return

    const requestID = latestRequestRef.current + 1
    latestRequestRef.current = requestID
    latestScopeRef.current = requestScope

    setState((s) => ({
      ...s,
      loading: true,
    }))

    const sent = sendContextMessage({
      type: 'get-targets',
      context,
      client_ID: client.ID,
      prompt_ID,
    })

    if (!sent && latestRequestRef.current === requestID) {
      setState((s) => ({
        ...s,
        loading: false,
      }))
    }
  }, [client?.ID, context, prompt_ID, requestScope])

  return state
}

export { useGetTargets }
