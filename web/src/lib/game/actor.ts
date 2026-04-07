import type { Action } from './action'
import type { Modifier } from './modifier'
import type { Nature, NatureSet } from './nature'

type ActorAttackStat = 'attack' | 'chakra_attack'
type ActorDefenseStat = 'defense' | 'chakra_defense'

type ActorNatureStat =
  | ActorAttackStat
  | ActorDefenseStat
  | 'accuracy'
  | 'evasion'
  | 'speed'

type ActorBaseStat = ActorNatureStat | 'hp' | 'stamina'

type ActorStats<T> = Record<ActorBaseStat, T>
type NatureStats<T> = Record<Nature, T>

const actorFocuses = [
  'none',
  'aggressive',
  'relentless',
  'reckless',
  'heavy',
  'patient',
  'hardened',
  'tough',
  'steadfast',
  'intelligent',
  'volatile',
  'intense',
  'calculated',
  'agile',
  'hasty',
  'impulsive',
  'alert',
] as const

type ActorFocus = (typeof actorFocuses)[number]

type ActorDef = {
  actor_ID: string
  sprite_url: string
  name: string
  clan: string
  affiliations: Array<string>
  stats: ActorStats<number>
  natures: Partial<Record<NatureSet, Array<Nature>>>
  nature_damage: NatureStats<number>
  nature_resistance: NatureStats<number>
  ability: Modifier | null
  action_count: number
  action_IDs: Array<string>
}

type ActorState = {
  active_turns: number
  alive: boolean
  damage: number
  inactive_turns: number
  position_ID: string
  seen: boolean
  stamina_damage: number
  stunned: boolean
  statused: boolean
  burned: boolean
  paralyzed: boolean
  sleeping: boolean
}

type Actor = ActorDef &
  ActorState & {
    ID: string
    player_ID: string
    level: number
    experience: number
    focus: ActorFocus
    base_stats: ActorStats<number>
    staged_stats: ActorStats<number>
    pre_stats: ActorStats<number>
    applied_modifiers: Record<string, number>
    actions: Array<Action>
    resolved_nature_resistance: NatureStats<number>
    resolved_nature_damage: NatureStats<number>
    summon?: Actor & { proxy: boolean }
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
  return Math.floor(
    Object.values(stats).reduce((p, c) => p + c, 0) - stats.stamina
  )
}

export type {
  ActorDef,
  Actor,
  ActorFocus,
  ActorAttackStat,
  ActorDefenseStat,
  ActorBaseStat,
}
export { checkActorStat, getTotalBaseStats, actorFocuses }
