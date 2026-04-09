import {
  type ActorDef,
  type ActorBaseStat,
  type ActorStats,
  type ActorFocus,
  ACTOR_FOCUS_DETAILS,
} from '#/lib/game/actor'
import type { ActorConfig } from '#/lib/stores/socket'
import { useEffect, useState } from 'react'
import { useDebouncedCallback } from 'use-debounce'
import { cn } from '#/lib/utils'
import { Input } from './ui/input'
import { Slider } from './ui/slider'

const CAP = 64
const PER_STAT_MAX = 31

const STAT_NAMES: ActorStats<string> = {
  hp: 'HP',
  stamina: 'Stamina',
  speed: 'Speed',
  accuracy: 'Accuracy',
  evasion: 'Evasion',
  attack: 'Attack',
  defense: 'Defense',
  chakra_attack: 'Chakra Attack',
  chakra_defense: 'Chakra Defense',
}

function TeamBuilderStatGuage({ baseStat }: { baseStat: number }) {
  return (
    <div className="relative bg-gray-600 rounded-md h-6 w-60">
      <div
        className="absolute top-0 left-0 rounded-md h-6 bg-white"
        style={{
          width: `${(baseStat * 100) / 200}%`,
        }}
      />
    </div>
  )
}

function TeamBuilderStat({
  total,
  focus = 'none',
  base,
  stat,
  config,
  onConfigChange,
}: {
  total: number
  base: ActorDef
  stat: ActorBaseStat
  focus?: ActorFocus
  config: ActorConfig
  onConfigChange: (config: ActorConfig) => void
}) {
  const detail = ACTOR_FOCUS_DETAILS[focus]
  const up = detail.up === stat
  const down = detail.down === stat
  const aux = (config.aux_stats as any)[stat] ?? 0
  const [localAux, setLocalAux] = useState(aux)

  useEffect(() => {
    setLocalAux(aux)
  }, [aux])

  const debouncedConfigChange = useDebouncedCallback((value: number) => {
    onConfigChange({
      ...config,
      aux_stats: {
        ...config.aux_stats,
        [stat]: value,
      },
    })
  }, 120)

  function handleConfigChange(value: number) {
    const otherStatsTotal = total - aux
    const maxAllowedByCap = CAP - otherStatsTotal
    const clamped = Math.max(0, Math.min(PER_STAT_MAX, maxAllowedByCap, value))
    setLocalAux(clamped)
    debouncedConfigChange(clamped)
  }

  return (
    <tr>
      <td
        className={cn('w-24 text-muted-foreground whitespace-nowrap text-xs', {
          'text-green-300': up,
          'text-red-300': down,
        })}
      >
        {STAT_NAMES[stat]}
        {up && ' ( + )'}
        {down && ' ( - )'}
      </td>
      <td className="w-8 text-right p-2 whitespace-nowrap font-black">
        {base.stats[stat]}
      </td>
      <td>
        <TeamBuilderStatGuage baseStat={base.stats[stat]} />
      </td>
      <td className="px-2">
        <Input
          className="w-16"
          value={localAux}
          type="number"
          onChange={(e) => {
            const v = parseInt(e.target.value)
            if (!isNaN(v)) {
              handleConfigChange(v)
            }
          }}
        />
      </td>
      <td>
        <Slider
          value={[localAux]}
          max={PER_STAT_MAX}
          step={1}
          className="min-w-40"
          onValueChange={(v) => handleConfigChange(v[0])}
        />
      </td>
    </tr>
  )
}

export { TeamBuilderStat }
