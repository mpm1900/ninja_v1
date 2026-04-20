import type { ActorBaseStat } from '#/lib/game/actor'
import { cn } from '#/lib/utils'

const stageCoef = {
  [-6]: '0.25',
  [-5]: '0.28',
  [-4]: '0.33',
  [-3]: '0.4',
  [-2]: '0.5',
  [-1]: '0.67',
  [0]: '1',
  [1]: '1.5',
  [2]: '2.0',
  [3]: '2.5',
  [4]: '3.0',
  [5]: '3.5',
  [6]: '4.0',
}

const statName: Record<ActorBaseStat, string> = {
  hp: 'HP',
  stamina: 'STA',
  accuracy: 'ACC',
  attack: 'ATK',
  chakra_attack: 'C.ATK',
  defense: 'DEF',
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
      x{stageCoef[stage]} <span className='opacity-70'>{statName[stat]}</span>
    </span>
  )
}

export { StageBadge }
