import { cn } from '#/lib/utils'
import { useStore } from '@tanstack/react-store'
import { ActorThumbnail } from './actor-thumbnail'
import { gameStore } from '#/lib/stores/game'
import { clientsStore } from '#/lib/stores/clients'

function LobbyThumbnails({
  player_ID,
  enabled,
  onEnabledChange,
}: {
  player_ID: string
  enabled: string[]
  onEnabledChange: (enabled: string[]) => void
}) {
  const client = useStore(clientsStore, (s) => s.me)
  const actors = useStore(gameStore, (g) => g.actors)
  const is_player = client?.ID == player_ID

  function toggleActor(ID: string) {
    if (enabled.length >= 4 && !enabled.includes(ID)) {
      return
    }
    if (enabled.includes(ID)) {
      onEnabledChange(enabled.filter((id) => id !== ID))
    } else {
      onEnabledChange([...enabled, ID])
    }
  }
  return (
    <div className="flex gap-2">
      {actors
        .filter((a) => a.player_ID === player_ID)
        .map((a, i) => (
          <ActorThumbnail
            key={a.ID}
            actor={a}
            active={enabled.includes(a.ID)}
            onClick={() => toggleActor(a.ID)}
            index={i}
            showHealthBar={is_player}
            showRing
            hidden={!is_player && !a.seen && !a.position_ID}
            className={cn({
              'cursor-pointer': enabled.length < 4 || enabled.includes(a.ID),
              'cursor-not-allowed':
                enabled.length >= 4 && !enabled.includes(a.ID),
              'opacity-40': !enabled.includes(a.ID),
            })}
          />
        ))}
    </div>
  )
}

export { LobbyThumbnails }
