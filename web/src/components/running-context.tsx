import { getTargets, type Context } from '#/lib/game/context'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { cn } from '#/lib/utils'
import { useStore } from '@tanstack/react-store'

function RunningContext({ context }: { context: Context }) {
  const game = useStore(gameStore, (g) => g)
  const client_ID = useStore(clientsStore, (s) => s.me?.ID)
  const source = game.actors.find((a) => a.ID === context.source_actor_ID)
  const source_action = source?.actions.find((a) => a.ID === context.action_ID)
  const targets = getTargets(source_action?.target_type, game, context)
  const has_targets =
    targets.length > 0 && targets[0].ID !== context.source_actor_ID
  return (
    <div className="absolute text-muted-foreground">
      {source && source_action && (
        <div>
          <span
            className={cn('font-black text-3xl', {
              'text-blue-300': client_ID === source?.player_ID,
              'text-red-300': client_ID !== source?.player_ID,
            })}
          >
            {source.name}
          </span>{' '}
          uses{' '}
          <span className="font-bold text-2xl text-foreground">
            {source_action.config.name}
          </span>
        </div>
      )}
      {has_targets && source && <div>on {targets.map((t) => t.name).join(', ')}</div>}
    </div>
  )
}

export { RunningContext }
