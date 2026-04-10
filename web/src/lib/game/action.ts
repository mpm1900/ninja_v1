import type { ActorAttackStat } from './actor'
import type { Context } from './context'
import type { NatureSet } from './nature'

type Action = {
  ID: string
  config: {
    name: string
    jutsu: string
    nature?: NatureSet
    cost?: number
    cooldown?: number
    accuracy?: number
    stat?: ActorAttackStat
    power?: number
    recoil?: number
    description: string
    target_count?: number
  }
  locked: boolean
  cooldown: number | null
  disabled: boolean
  priority: number
  target_type: 'target-actor-id' | 'target-position-type'
}

type ActionTransaction = {
  ID: string
  context: Context
  mutation: Action
}

export type { Action, ActionTransaction }
