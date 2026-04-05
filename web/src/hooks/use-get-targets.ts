import type { Context } from "#/lib/game/context"
import { clientsStore } from "#/lib/stores/clients"
import { sendContextMessage, subscribeSocketMessages } from "#/lib/stores/socket"
import { Store, useStore } from "@tanstack/react-store"
import { useEffect } from "react"

const getTargetsStore = new Store<{ context: Context | null, loading: boolean }>({
  context: null,
  loading: false,
})

function useGetTargets(context: Context, prompt_ID?: string) {
  const client = useStore(clientsStore, (c) => c.me)
  const store = useStore(getTargetsStore, s => s)

  useEffect(() => {
    getTargetsStore.setState(() => ({ context: null, loading: false }))
    return subscribeSocketMessages((_, message) => {
      if (message?.type !== 'target-IDs') return
      getTargetsStore.setState(() => ({
        context: message.context,
        loading: false,
      }))
    })
  }, [])

  useEffect(() => {
    if (!client || getTargetsStore.get().loading) return

    getTargetsStore.setState(s => ({
      ...s,
      loading: true
    }))
    sendContextMessage({
      type: 'get-targets',
      context: context,
      client_ID: client.ID,
      prompt_ID: prompt_ID
    })
  }, [context])


  return store
}

export { useGetTargets }
