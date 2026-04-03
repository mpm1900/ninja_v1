import type { Actor, ActorBaseStat } from '#/lib/game/actor'
import { ActorStat, NatureDamageStat, NatureResistanceStat } from './actor-stat'
import { NatureBadge } from './nature-badge'

const STAT_LABELS: Record<ActorBaseStat, string> = {
  hp: 'HP',
  stamina: 'Stamina',
  speed: 'Speed',
  accuracy: 'Accuracy',
  evasion: 'Evasion',
  attack: 'Attack',
  chakra_attack: 'Chakra Attack',
  defense: 'Defense',
  chakra_defense: 'Chakra Defense',
}

function ActorStats({ actor }: { actor: Actor }) {
  return (
    <div className="flex items-start gap-4">
      <div>
        {Object.keys(actor.stats).map((key) => (
          <div key={key} className="flex justify-between gap-2">
            <div className="text-muted-foreground">
              {STAT_LABELS[key as ActorBaseStat]}:{' '}
            </div>
            <div>
              <ActorStat
                actor={actor}
                showBase={false}
                stat={key as keyof typeof actor.stats}
              />
            </div>
          </div>
        ))}
      </div>
      <div>
        <div className='uppercase text-muted-foreground font-bold text-center'>DMG</div>
        {Object.keys(actor.nature_damage).map((key) => (
          <div key={key} className="flex justify-between gap-1">
            <div className="text-muted-foreground">
              <NatureBadge
                nature={key as keyof typeof actor.nature_damage}
              />
            </div>
            <div>
              <NatureDamageStat
                actor={actor}
                nature={key as keyof typeof actor.nature_damage}
              />
            </div>
          </div>
        ))}
      </div>
      <div>
        <div className='uppercase text-muted-foreground font-bold text-center'>RES</div>
        {Object.keys(actor.resolved_nature_resistance).map((key) => (
          <div key={key} className="flex justify-between gap-1">
            <div className="text-muted-foreground">
              <NatureBadge
                nature={key as keyof typeof actor.resolved_nature_resistance}
              />
            </div>
            <div>
              <NatureResistanceStat
                actor={actor}
                nature={key as keyof typeof actor.resolved_nature_resistance}
              />
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export { ActorStats }
