import type { ActionTransaction } from '#/lib/game/action'
import type { Game } from '#/lib/game/game'
import { gameStore } from '#/lib/stores/game'
import { useStore } from '@tanstack/react-store'

function ActionItem({
  game,
  transaction,
}: {
  game: Game
  transaction: ActionTransaction
}) {
  const actor = game.actors.find(
    (a) => a.ID === transaction.context.source_actor_ID
  )
  return (
    <div>
      <div>
        {transaction.mutation.config.name} ({transaction.mutation.priority})
      </div>
      {actor && (
        <div>
          {actor.name} ({actor.stats.speed})
        </div>
      )}
    </div>
  )
}

function ActionQueue() {
  const game = useStore(gameStore, (g) => g)

  return (
    <div className="flex justify-between p-2">
      <div className="flex gap-6">
        {game.actions.map((t) => (
          <ActionItem key={t.ID} game={game} transaction={t} />
        ))}
      </div>
    </div>
  )
}

export { ActionQueue }
