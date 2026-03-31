import { NULL_CONTEXT, type Context } from '#/lib/game/context'
import { useQuery } from '@tanstack/react-query'
import { Button } from './ui/button'
import { useStore } from '@tanstack/react-store'
import { sendContextMessage, socketStore } from '#/lib/stores/socket'
import { gameStore } from '#/lib/stores/game'
import { actionTargetsQuery } from '#/lib/queries/action-targets'
import { isActionContextValidQuery } from '#/lib/queries/is-action-context-valid'
import { clientsStore } from '#/lib/stores/clients'
import { TargetButton } from './target-button'
import type { Action, ActionTransaction } from '#/lib/game/action'
import { setActionID } from '#/lib/stores/battle-context'

function ActionControl({
  action,
  queued,
  enabled,
  context,
  onContextChange,
}: {
  action?: Action
  queued?: ActionTransaction
  enabled: boolean
  context: Context
  onContextChange: (context: Context) => void
}) {
  const instanceID = useStore(socketStore, (s) => s.instanceID!)
  const game = useStore(gameStore, (g) => g)
  const valid = useQuery(isActionContextValidQuery(context))
  const client = useStore(clientsStore, (c) => c.me!)
  const actionTargets = useQuery(actionTargetsQuery(instanceID, context))
  const loading = valid.isFetching || actionTargets.isFetching
  const actors = game.actors.filter((a) => actionTargets.data?.includes(a.ID))

  if (!!queued) {
    return (
      <div className="flex flex-col py-4 items-center">
        <Button
          disabled={loading || !enabled}
          onClick={() => {
            sendContextMessage({
              type: 'remove-action',
              client_ID: client.ID,
              context: {
                ...NULL_CONTEXT,
                action_ID: queued.ID,
              },
            })
          }}
        >
          Cancel {queued.mutation.config.name}
        </Button>
      </div>
    )
  }

  return (
    <div className="flex flex-col items-center gap-2">
      {action && (
        <div className="p-2 flex gap-2 justify-center">
          {actors.map((a) => (
            <TargetButton
              key={a.ID}
              actor={a}
              enabled={enabled}
              loading={loading}
              contextValid={!!valid.data}
              targetType={action.target_type}
              context={context}
              onContextChange={onContextChange}
            />
          ))}
          {!loading && actors.length == 0 && valid.data === false && (
            <span className="text-muted-foreground text-sm">
              no targets available
            </span>
          )}
        </div>
      )}
      {enabled && action && valid.data && (
        <Button
          disabled={loading}
          onClick={() => {
            sendContextMessage({
              type: 'push-action',
              client_ID: client.ID,
              context,
            })

            setActionID(context.source_actor_ID!, context.action_ID!, game)
          }}
        >
          Select
        </Button>
      )}
    </div>
  )
}

export { ActionControl }
