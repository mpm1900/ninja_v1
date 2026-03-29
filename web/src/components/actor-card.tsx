import { STAT_ICONS } from '#/data/icons'
import type { Actor } from '#/lib/game/actor'
import type { Game } from '#/lib/game/game'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { cn } from '#/lib/utils'
import { cva } from 'class-variance-authority'
import { ActorStat } from './actor-stat'
import { HealthBar } from './health-bar'
import { NatureBadge } from './nature-badge'
import { Item, ItemActions, ItemContent, ItemTitle } from './ui/item'

const actorVariants = cva(
  cn(
    'group/item flex flex-wrap items-center rounded-md border text-sm transition-colors duration-100 outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 [a]:transition-colors [a]:hover:bg-accent/50',
    'px-2 pb-0 pt-2 w-80 border-transparent cursor-pointer'
  ),
  {
    variants: {
      player: {
        player: 'bg-gray-700',
        enemy: '',
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
  const HpIcon = STAT_ICONS.hp
  const ChakraIcon = STAT_ICONS.chakra
  const SpeedIcon = STAT_ICONS.speed
  const NinjutsuIcon = STAT_ICONS.ninjutsu
  const GenjutsuIcon = STAT_ICONS.genjutsu
  const Taijutsu = STAT_ICONS.taijutsu

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
            {actor && (
              <div className="grid grid-cols-3 lg:grid-cols-6 space-x-2 hidden">
                <div className="flex items-center gap-2 justify-center">
                  {HpIcon && <HpIcon />}
                  <ActorStat actor={actor} stat="hp" showBase={false} />
                </div>
                <div className="flex items-center gap-2 justify-center">
                  {ChakraIcon && <ChakraIcon />}
                  <ActorStat actor={actor} stat="chakra" showBase={false} />
                </div>
                <div className="flex items-center gap-2 justify-center">
                  {SpeedIcon && <SpeedIcon />}
                  <ActorStat actor={actor} stat="speed" showBase={false} />
                </div>
                <div className="flex items-center gap-2 justify-center">
                  {NinjutsuIcon && <NinjutsuIcon />}
                  <ActorStat actor={actor} stat="ninjutsu" showBase={false} />
                </div>
                <div className="flex items-center gap-2 justify-center">
                  {GenjutsuIcon && <GenjutsuIcon />}
                  <ActorStat actor={actor} stat="genjutsu" showBase={false} />
                </div>
                <div className="flex items-center gap-2 justify-center">
                  {Taijutsu && <Taijutsu />}
                  <ActorStat actor={actor} stat="taijutsu" showBase={false} />
                </div>
              </div>
            )}
          </div>
        </ItemContent>
      </div>
    </div>
  )
}

export { ActorCard }
