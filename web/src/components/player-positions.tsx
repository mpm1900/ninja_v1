import { gameStore } from '#/lib/stores/game'
import { useStore } from '@tanstack/react-store'
import { ActorCard } from './actor-card'
import { clientsStore } from '#/lib/stores/clients'

function PlayerPositions({
  player_ID,
  selected,
  onSelectedChange,
}: {
  player_ID: string
  selected?: string
  onSelectedChange?: (selected: string) => void
}) {
  const game = useStore(gameStore, (g) => g)
  const client = useStore(clientsStore, (c) => c.me)
  const player = game.players.find((p) => p.ID === player_ID)

  if (!player) return null

  return (
    <div className="flex gap-2 py-2">
      {player.positions.map((pos) =>
        pos.actor_ID ? (
          <>
            <ActorCard
              key={pos.ID}
              actor={game.actors.find((a) => a.ID === pos.actor_ID)}
              clientID={client?.ID}
              game={game}
              selected={selected === pos.actor_ID}
              onClick={() => onSelectedChange?.(pos.actor_ID ?? '')}
            />
          </>
        ) : (
          <div>placeholder</div>
        )
      )}
    </div>
  )
}

export { PlayerPositions }
