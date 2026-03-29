import type { Actor } from '#/lib/game/actor'
import { Progress } from './ui/progress'

function HealthBar({ actor }: { actor: Actor }) {
  const hp_ratio = (actor.stats.hp - actor.damage) / actor.stats.hp
  const chakra_ratio =
    (actor.stats.chakra - actor.chakra_damage) / actor.stats.chakra

  return (
    <div className="relative">
      <Progress
        className="absolute top-0.5 z-10 [&_[data-slot=progress]]:bg-red-300/20 [&_[data-slot=progress-indicator]]:bg-rose-800"
        value={hp_ratio * 100}
      />
      <Progress className="h-3" value={chakra_ratio * 100} />
    </div>
  )
}

export { HealthBar }
