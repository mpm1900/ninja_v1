import { cn } from '#/lib/utils'
import { useStore } from '@tanstack/react-store'
import { ActorThumbnail } from './actor-thumbnail'
import { gameStore } from '#/lib/stores/game'
import { clientsStore } from '#/lib/stores/clients'

function PlayerThumbnails({ player_ID }: { player_ID: string }) {
  const client = useStore(clientsStore, (s) => s.me)
  const actors = useStore(gameStore, (g) => g.actors)
  const is_player = client?.ID == player_ID
  return (
    <div className="flex gap-2 py-2">
      {actors
        .filter((a) => a.player_ID === player_ID)
        .map((a, i) => (
          <ActorThumbnail
            key={a.ID}
            actor={a}
            index={i}
            showHealthBar={is_player}
            showRing
            hidden={!is_player && !a.seen && !a.position_ID}
            className={cn({
              'opacity-40': !a.position_ID,
            })}
          />
        ))}
    </div>
  )
}

export { PlayerThumbnails }
