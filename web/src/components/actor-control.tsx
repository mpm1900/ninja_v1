import { Card, CardContent } from './ui/card'
import { useStore } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { ActionControl } from './action-control'
import { PositionSelect } from './position-select'
import type { Actor } from '#/lib/game/actor'
import { ActionsTable } from './actions-table'
import { useGameContext } from '#/hooks/use-game-context'
import { ActorStats } from './actor-stats'

function ActorControl({ actor, enabled }: { actor: Actor; enabled: boolean }) {
  const game = useStore(gameStore, (g) => g)
  const player = game.players.find((p) => p.ID == actor.player_ID)
  const { context, onContextChange } = useGameContext(actor, undefined, [game])

  return (
    <Card className="grid grid-cols-2 rounded-t-none border-t-0 mx-2 mb-2 py-2 gap-0">
      <div>
        <CardContent className="px-2 flex flex-col gap-2">
          <div className="flex gap-2">
            <div className="h-16 w-16 overflow-hidden">
              <img
                src={actor.sprite_url}
                className="h-full w-full object-cover"
                width={64}
                height={64}
              />
            </div>
            <div className="flex flex-col gap-2">
              <PositionSelect actor={actor} game={game} />
            </div>
            <ActorStats actor={actor} />
          </div>
        </CardContent>
      </div>
      <div>
        <CardContent className="px-2">
          <ActionsTable
            data={actor.actions}
            enabled={enabled && !!player && !!actor.position_ID}
            selected={context.action_ID ?? undefined}
            onSelectedChange={(aid) =>
              onContextChange({
                ...context,
                action_ID: aid,
              })
            }
            subRow={({ row }) => (
              <ActionControl
                action={actor.actions.find((a) => a.ID === context.action_ID)}
                enabled={
                  enabled &&
                  !!player &&
                  !!actor.position_ID &&
                  row.original.cooldown == null
                }
                context={context}
                onContextChange={onContextChange}
              />
            )}
          />
        </CardContent>
      </div>
    </Card>
  )
}

export { ActorControl }
