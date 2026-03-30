import { gameStore } from '#/lib/stores/game'
import { useStore } from '@tanstack/react-store'
import { ActorCard } from './actor-card'
import { clientsStore } from '#/lib/stores/clients'
import { Item } from './ui/item'
import { cn } from '#/lib/utils'

function PlayerPositions({
  flip,
  player_ID,
  selected,
  onSelectedChange,
}: {
  flip: boolean
  player_ID: string
  selected?: string
  onSelectedChange?: (selected: string) => void
}) {
  const game = useStore(gameStore, (g) => g)
  const client = useStore(clientsStore, (c) => c.me)
  const player = game.players.find((p) => p.ID === player_ID)

  if (!player) return null

  return (
    <div className="flex gap-8 py-4 px-4">
      {player.positions.map((pos) => (
        <div
          key={pos.ID}
          className={cn('flex items-end', flip && 'items-start')}
        >
          {pos.actor_ID ? (
            <ActorCard
              key={pos.ID}
              actor={game.actors.find((a) => a.ID === pos.actor_ID)}
              client_ID={client?.ID}
              game={game}
              selected={selected === pos.actor_ID}
              onClick={() => onSelectedChange?.(pos.actor_ID ?? '')}
              className={flip ? 'flex-col-reverse' : ''}
            />
          ) : (
            <Item className="p-6 w-72">
              <span className="text-center w-full text-muted-foreground">
                Empty
              </span>
            </Item>
          )}
        </div>
      ))}
    </div>
  )
}

export { PlayerPositions }
