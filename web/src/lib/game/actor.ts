import type { Modifier } from './modifier'

type ActorBaseStat =
  | 'accuracy'
  | 'evasion'
  | 'genjutsu'
  | 'hp'
  | 'ninjutsu'
  | 'speed'
  | 'stamina'
  | 'taijutsu'

type Nature = 'earth' | 'fire' | 'lightning' | 'water' | 'wind' | 'yang' | 'yin'


const natureNames: Partial<Record<Nature, string>> = {
  fire: '火',
  wind: '風',
  lightning: '雷',
  earth: '土',
  water: '水',
  yin: '陰',
  yang: '陽',
}


type ActorStats<T> = Record<ActorBaseStat, T>
type NatureStats<T> = Record<Nature, T>

type Actor = {
  ID: string
  actor_ID: string
  player_ID: string
  name: string
  level: number
  experience: number
  action_count: number
  alive: true
  active: true
  critical: number
  base_stats: ActorStats<number>
  staged_stats: ActorStats<number>
  pre_stats: ActorStats<number>
  stats: ActorStats<number>
  natures: Array<Nature>
  nature_damage: NatureStats<number>
  nature_resistance: NatureStats<number>
  applied_modifiers: Record<string, number>
  innate_modifiers: Array<Modifier>
}

function checkActorStat(actor: Actor, key: ActorBaseStat) {
  const stat = actor.stats[key]
  const pre = actor.pre_stats[key]
  return stat === pre ? 0 : stat > pre ? 1 : -1
}

function getTotalBaseStats(actor: Actor) {
  const stats: Actor['base_stats'] = {
    ...actor.base_stats,
    accuracy: 0,
    evasion: 0,
  }
  return Object.values(stats).reduce((p, c) => p + c, 0)
}

export type { Nature, Actor, ActorBaseStat }
export { natureNames, checkActorStat, getTotalBaseStats }
