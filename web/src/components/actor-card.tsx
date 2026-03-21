import { STAT_ICONS } from '#/data/icons'
import type { Actor } from '#/lib/game/actor'
import type { Game } from '#/lib/game/game'
import { ActorStat } from './actor-stat'
import { Card, CardContent, CardHeader, CardTitle } from './ui/card'

function ActorCard({ actor, game }: { actor: Actor | undefined; game: Game }) {
  const modifiers = (game.modifiers ?? []).map(m => m.mutation).concat(actor?.innate_modifiers ?? [])
  const HpIcon = STAT_ICONS.hp
  const StaminaIcon = STAT_ICONS.stamina
  const SpeedIcon = STAT_ICONS.speed
  const NinjutsuIcon = STAT_ICONS.ninjutsu
  const GenjutsuIcon = STAT_ICONS.genjutsu
  const Taijutsu = STAT_ICONS.taijutsu
  return (
    <Card>
      <CardHeader>
        <CardTitle>{actor?.name}</CardTitle>
      </CardHeader>
      <CardContent>
        {actor && (
          <div className="grid grid-cols-6 gap-2">
            <div className="flex items-center gap-2">
              {HpIcon && <HpIcon />}
              <ActorStat actor={actor} stat="hp" showBase={false} />
            </div>
            <div className="flex items-center gap-2">
              {StaminaIcon && <StaminaIcon />}
              <ActorStat actor={actor} stat="stamina" showBase={false} />
            </div>
            <div className="flex items-center gap-2">
              {SpeedIcon && <SpeedIcon />}
              <ActorStat actor={actor} stat="speed" showBase={false} />
            </div>
            <div className="flex items-center gap-2">
              {NinjutsuIcon && <NinjutsuIcon />}
              <ActorStat actor={actor} stat="ninjutsu" showBase={false} />
            </div>
            <div className="flex items-center gap-2">
              {GenjutsuIcon && <GenjutsuIcon />}
              <ActorStat actor={actor} stat="genjutsu" showBase={false} />
            </div>
            <div className="flex items-center gap-2">
              {Taijutsu && <Taijutsu />}
              <ActorStat actor={actor} stat="taijutsu" showBase={false} />
            </div>
          </div>
        )}
        <div>
          {Object.entries(actor?.applied_modifiers ?? {}).map(([ID, count]) => (
            <span key={ID}>{modifiers.find((m) => m.ID === ID)?.name}{count > 0 ? ` (${count + 1})` : null}</span>
          ))}
        </div>
      </CardContent>
    </Card>
  )
}

export { ActorCard }
