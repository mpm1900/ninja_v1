import { STAT_ICONS } from '#/data/icons'
import type { Actor } from '#/lib/game/actor'
import type { Game } from '#/lib/game/game'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { cn } from '#/lib/utils'
import { cva } from 'class-variance-authority'
import { HealthBar } from './health-bar'
import { NatureBadge } from './nature-badge'
import { Item, ItemActions, ItemContent, ItemTitle } from './ui/item'

const actorVariants = cva(
  cn(
    'group/item flex flex-wrap items-center rounded-md border text-sm transition-colors duration-100 outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 [a]:transition-colors [a]:hover:bg-accent/50',
    'p-2 w-80 border-transparent cursor-pointer',
    'border-black border-2'
  ),
  {
    variants: {
      player: {
        player: 'bg-gray-800',
        enemy: '',
      },
      selected: {
        selected: 'scale-110 bg-gray-700!',
      },
    },
    defaultVariants: {},
  }
)

function ActorCard({
  actor,
  clientID,
  game,
  selected,
  className,
  ...props
}: React.ComponentProps<typeof Item> & {
  actor: Actor | undefined
  clientID?: string
  game: Game
  selected: boolean
}) {
  const modifiers = (game.modifiers ?? [])
    .map((m) => m.mutation)
    .concat(actor?.innate_modifiers ?? [])

  return (
    <div className={cn('flex flex-col', className)}>
      <div className="flex flex-wrap gap-3">
        {Object.entries(actor?.applied_modifiers ?? {}).map(([ID, count]) => (
          <span key={ID}>
            {modifiers.find((m) => m.group_ID === ID)?.name}
            {count > 1 ? ` (${count})` : null}
          </span>
        ))}
      </div>
      <div
        className={actorVariants({
          player: actor?.player_ID === clientID ? 'player' : 'enemy',
          selected: selected ? 'selected' : undefined,
        })}
        {...props}
      >
        <ItemContent>
          {actor && (
            <div className="flex justify-between items-end gap-4">
              <ItemTitle>
                <span className="text-muted-foreground text-sm">
                  Lv.{actor.level}
                </span>{' '}
                <span
                  className={cn({
                    'text-blue-400': actor.player_ID === clientID,
                    'text-red-400': actor.player_ID !== clientID,
                    'text-foregroud': selected,
                  })}
                >
                  {actor.name}
                </span>
              </ItemTitle>
              <ItemActions className="gap-0">
                {(Object.keys(actor.natures) as Array<NatureSet>)
                  .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                  .map((nature) => (
                    <NatureBadge key={nature} nature={nature} />
                  ))}
              </ItemActions>
            </div>
          )}
          <div className="space-y-2">
            {actor && <HealthBar actor={actor} selected={selected} />}
          </div>
        </ItemContent>
      </div>
    </div>
  )
}

export { ActorCard }
