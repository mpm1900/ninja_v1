import type { Modifier } from './modifier'
import type { Nature, NatureSet } from './nature'

type ActorBaseStat =
  | 'accuracy'
  | 'evasion'
  | 'genjutsu'
  | 'hp'
  | 'ninjutsu'
  | 'speed'
  | 'chakra'
  | 'taijutsu'

type ActorStats<T> = Record<ActorBaseStat, T>
type NatureStats<T> = Record<Nature, T>

type ActorDef = {
  actor_ID: string
  name: string
  action_count: number
  stats: ActorStats<number>
  natures: Partial<Record<NatureSet, Array<Nature>>>
  nature_damage: NatureStats<number>
  nature_resistance: NatureStats<number>
  innate_modifiers: Array<Modifier>
  action_IDs: Array<string>
}

type Actor = ActorDef & {
  ID: string
  player_ID: string
  level: number
  experience: number
  state: {
    alive: boolean
    damage: number
    position_ID: string
  },
  base_stats: ActorStats<number>
  staged_stats: ActorStats<number>
  pre_stats: ActorStats<number>
  applied_modifiers: Record<string, number>
  actions: Array<Action>
}

function checkActorStat(actor: Actor, key: ActorBaseStat) {
  const stat = actor.stats[key]
  const pre = actor.pre_stats[key]
  return stat === pre ? 0 : stat > pre ? 1 : -1
}

function getTotalBaseStats(actor: ActorDef) {
  const stats: ActorDef['stats'] = {
    ...actor.stats,
    accuracy: 0,
    evasion: 0,
  }
  return Object.values(stats).reduce((p, c) => p + c, 0)
}

export type { ActorDef, Actor, ActorBaseStat }
export { checkActorStat, getTotalBaseStats }
