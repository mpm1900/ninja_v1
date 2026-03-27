import type { ActionTransaction } from '#/lib/game/action'
import type { Game } from '#/lib/game/game'
import { gameStore } from '#/lib/stores/game'
import { useStore } from '@tanstack/react-store'
import { Button } from './ui/button'
import { sendContextMessage } from '#/lib/stores/socket'
import { clientsStore } from '#/lib/stores/clients'
import { NULL_CONTEXT } from '#/lib/game/context'

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
  const client = useStore(clientsStore, (c) => c.me!)

  return (
    <div className="flex justify-between p-2">
      <div className="flex gap-6">
        {game.actions.map((t) => (
          <ActionItem key={t.ID} game={game} transaction={t} />
        ))}
      </div>
      <div className="flex gap-2">
        <Button
          disabled={game.actions.length == 0 || game.status === 'running'}
          onClick={() => {
            sendContextMessage({
              type: 'run-game-actions',
              client_ID: client.ID,
              context: NULL_CONTEXT,
            })
          }}
        >
          Run
        </Button>
        <Button
          disabled={!client || game.status === 'running'}
          onClick={() => {
            sendContextMessage({
              type: 'validate-state',
              client_ID: client.ID,
              context: NULL_CONTEXT,
            })
          }}
        >
          Validate
        </Button>
      </div>
    </div>
  )
}

export { ActionQueue }
