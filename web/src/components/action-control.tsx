import { NULL_CONTEXT, type Context } from '#/lib/game/context'
import { Button } from './ui/button'
import { useStore } from '@tanstack/react-store'
import { sendContextMessage } from '#/lib/stores/socket'
import { gameStore } from '#/lib/stores/game'
import { clientsStore } from '#/lib/stores/clients'
import { TargetButton } from './target-button'
import type { Action, ActionTransaction } from '#/lib/game/action'
import { setActionID } from '#/lib/stores/battle-context'
import { useValidateContext } from '#/hooks/use-validate-context'
import { useGetTargets } from '#/hooks/use-get-targets'
import { ChevronRight } from 'lucide-react'

function ActionControl({
  action,
  staged,
  enabled,
  context,
  onContextChange,
}: {
  action?: Action
  staged?: ActionTransaction
  enabled: boolean
  context: Context
  onContextChange: (context: Context) => void
}) {
  const game = useStore(gameStore, (g) => g)
  const { valid } = useValidateContext(context)

  const client = useStore(clientsStore, (c) => c.me!)
  const { context: t_context } = useGetTargets(context)
  const actors = game.actors.filter((a) =>
    t_context?.target_actor_IDs?.includes(a.ID)
  )
  const enemy_actors = actors.filter(a => a.player_ID !== client.ID)
  const player_actors = actors.filter(a => a.player_ID === client.ID)
  const has_queued_action = game.queued_actions[context.source_actor_ID ?? '']

  if (!!staged) {
    return (
      <div className="flex flex-col py-4 items-center">
        {has_queued_action ? (
          <span className="text-muted-foreground">
            {staged.mutation.config.name}
          </span>
        ) : (
          <Button
            disabled={!enabled}
            onClick={() => {
              sendContextMessage({
                type: 'remove-action',
                client_ID: client.ID,
                context: {
                  ...NULL_CONTEXT,
                  action_ID: staged.ID,
                },
              })
            }}
          >
            Cancel {staged.mutation.config.name}
          </Button>
        )}
      </div>
    )
  }

  return (
    <div className="flex flex-col items-center gap-2 min-w-xs">
      {action && (
        <div className='flex flex-col gap-2'>
          <div className="gap-2 grid grid-cols-2">
            {enemy_actors.map((a) => (
              <TargetButton
                key={a.ID}
                actor={a}
                enabled={enabled}
                loading={false}
                contextValid={!!valid}
                targetType={action.target_type}
                context={context}
                onContextChange={onContextChange}
              />
            ))}
          </div>
          <div className="gap-2 grid grid-cols-2">
            {player_actors.map((a) => (
              <TargetButton
                key={a.ID}
                actor={a}
                enabled={enabled}
                loading={false}
                contextValid={!!valid}
                targetType={action.target_type}
                context={context}
                onContextChange={onContextChange}
              />
            ))}
          </div>
          {actors.length == 0 && valid === false && (
            <span className="text-muted-foreground text-sm">
              no targets available
            </span>
          )}
          {actors.length == 0 && valid === true && (
            <span className="text-muted-foreground/50 text-sm">
              this action does not target
            </span>
          )}

        </div>
      )}

      <div className="flex w-full justify-end">
        <Button
          disabled={!(enabled && action && valid)}
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
          <ChevronRight />
        </Button>
      </div>
    </div>
  )
}

export { ActionControl }
