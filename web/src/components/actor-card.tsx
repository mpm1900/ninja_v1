import { STAT_ICONS } from '#/data/icons'
import type { Actor } from '#/lib/game/actor'
import type { Game } from '#/lib/game/game'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { ActorStat } from './actor-stat'
import { HealthBar } from './health-bar'
import { NatureBadge } from './nature-badge'
import { Item, ItemActions, ItemContent, ItemTitle } from './ui/item'

function ActorCard({ actor, game }: { actor: Actor | undefined; game: Game }) {
  const modifiers = (game.modifiers ?? [])
    .map((m) => m.mutation)
    .concat(actor?.innate_modifiers ?? [])
  const HpIcon = STAT_ICONS.hp
  const StaminaIcon = STAT_ICONS.stamina
  const SpeedIcon = STAT_ICONS.speed
  const NinjutsuIcon = STAT_ICONS.ninjutsu
  const GenjutsuIcon = STAT_ICONS.genjutsu
  const Taijutsu = STAT_ICONS.taijutsu
  return (
    <Item variant='muted'>
      <ItemContent>
        {actor && (
          <div className='flex justify-between'>
            <ItemTitle>
              <span className="text-muted-foreground text-sm">
                Lv.{actor.level}
              </span>{' '}
              {actor.name}
            </ItemTitle>
            <ItemActions className='gap-0'>
              {(Object.keys(actor.natures) as Array<NatureSet>)
                .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                .map((nature) => (
                  <NatureBadge key={nature} nature={nature} />
                ))}
            </ItemActions>
          </div>
        )}
        <div className="space-y-2">
          {actor && <HealthBar actor={actor} />}
          {actor && (
            <div className="grid grid-cols-3 lg:grid-cols-6 space-x-2">
              <div className="flex items-center gap-2 justify-center">
                {HpIcon && <HpIcon />}
                <ActorStat actor={actor} stat="hp" showBase={false} />
              </div>
              <div className="flex items-center gap-2 justify-center">
                {StaminaIcon && <StaminaIcon />}
                <ActorStat actor={actor} stat="stamina" showBase={false} />
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
          <div className='flex flex-wrap gap-3'>
            {Object.entries(actor?.applied_modifiers ?? {}).map(([ID, count]) => (
              <span key={ID}>
                {modifiers.find((m) => m.group_ID === ID)?.name}
                {count > 0 ? ` (${count})` : null}
              </span>
            ))}
          </div>
        </div>
      </ItemContent>
    </Item>
  )
}

export { ActorCard }
