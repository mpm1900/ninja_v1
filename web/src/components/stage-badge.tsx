import type { ActorBaseStat } from '#/lib/game/actor'
import { cn } from '#/lib/utils'

const stageCoef = {
  [-6]: 'x0.25',
  [-5]: 'x0.28',
  [-4]: 'x0.33',
  [-3]: 'x0.4',
  [-2]: 'x0.5',
  [-1]: 'x0.66',
  [0]: 'x1',
  [1]: 'x1.5',
  [2]: 'x2.0',
  [3]: 'x2.5',
  [4]: 'x3.0',
  [5]: 'x3.5',
  [6]: 'x4.0',
}

const statName: Record<ActorBaseStat, string> = {
  hp: 'HP',
  stamina: 'STA',
  accuracy: 'ACC',
  attack: 'P.ATK',
  chakra_attack: 'C.ATK',
  defense: 'P.DEF',
  chakra_defense: 'C.DEF',
  speed: 'SPE',
  evasion: 'EVA',
}

function StageBadge({
  stat,
  stage,
}: {
  stat: ActorBaseStat
  stage: keyof typeof stageCoef
}) {
  if (stage === 0) return null
  if (stage > 6) stage = 6
  if (stage < -6) stage = -6
  return (
    <span
      className={cn(
        'border border-transparent leading-3 text-[10px] font-bold px-1 rounded',
        {
          'border-green-600 bg-green-200 text-green-900': stage > 0,
          'border-red-600 bg-red-200 text-red-900': stage < 0,
        }
      )}
    >
      {stageCoef[stage]} {statName[stat]}
    </span>
  )
}

export { StageBadge }
