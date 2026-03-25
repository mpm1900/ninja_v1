import type { Actor } from '#/lib/game/actor'
import { Progress } from './ui/progress'

function HealthBar({ actor }: { actor: Actor }) {
  const maxHealth = actor.stats.hp
  const damage = actor.state.damage
  const ratio = (maxHealth - damage) / maxHealth

  return <Progress value={ratio * 100} />
}

export { HealthBar }
