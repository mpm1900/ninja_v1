import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { useStore } from '@tanstack/react-store'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from './ui/alert-dialog'
import type { Context } from '#/lib/game/context'
import { sendContextMessage, socketStore } from '#/lib/stores/socket'
import { useQuery } from '@tanstack/react-query'
import { isActionContextValidQuery } from '#/lib/queries/is-action-context-valid'
import { actionTargetsQuery } from '#/lib/queries/action-targets'
import { TargetButton } from './target-button'
import { Button } from './ui/button'
import { useEffect, useState } from 'react'
import type { Action } from '#/lib/game/action'

function PromptControl({
  action,
  context,
  onContextChange,
}: {
  action?: Action
  context: Context
  onContextChange: (context: Context) => void
}) {
  const instanceID = useStore(socketStore, (s) => s.instanceID!)
  const game = useStore(gameStore, (g) => g)
  const valid = useQuery(isActionContextValidQuery(action?.ID, context))
  const client = useStore(clientsStore, (c) => c.me!)
  const actionTargets = useQuery(
    actionTargetsQuery(instanceID, action?.ID, context)
  )
  const loading = valid.isFetching || actionTargets.isFetching


  return (
    <>
      {action && <div className="p-2 flex gap-2">
        {game.actors
          .filter((a) => actionTargets.data?.includes(a.ID))
          .map((a) => (
            <TargetButton
              key={a.ID}
              actor={a}
              enabled
              loading={loading}
              contextValid={!!valid.data}
              context={context}
              onContextChange={onContextChange}
              targetType={action.target_type}
            />
          ))}
      </div>}
      <AlertDialogFooter>
        <AlertDialogAction asChild>
          <Button disabled={loading || !valid.data} onClick={() => {
            sendContextMessage({
              type: 'resolve-prompt',
              client_ID: client.ID,
              action_ID: action?.ID,
              context,
            })
          }}>
            Select
          </Button>
        </AlertDialogAction>
      </AlertDialogFooter>
    </>
  )
}

function PromptController() {
  const game = useStore(gameStore, (g) => g)
  const prompt = game.prompt

  const [context, setContext] = useState(prompt?.context)

  useEffect(() => {
    setContext(prompt?.context)
  }, [prompt?.ID])

  return (
    <AlertDialog open={!!prompt}>
      {prompt && context && (
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>{prompt.mutation.config.name}</AlertDialogTitle>
            <AlertDialogDescription>Select Shinobi to switch in.</AlertDialogDescription>
          </AlertDialogHeader>
          <PromptControl
            action={prompt.mutation}
            context={context}
            onContextChange={setContext}
          />
        </AlertDialogContent>
      )}
    </AlertDialog>
  )
}

export { PromptController }
