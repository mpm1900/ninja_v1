import type { ActorAttackStat } from './actor'
import type { Context } from './context'
import type { NatureSet } from './nature'

type Action = {
  ID: string
  config: {
    name: string
    nature: NatureSet | null
    cost: number | null
    accuracy: number | null
    stat: ActorAttackStat | null
    power: number | null
    recoil: number | null
  }
  priority: number
  target_type: 'target-actor-id' | 'target-position-type'
}

type ActionTransaction = {
  ID: string
  context: Context
  mutation: Action
}

export type { Action, ActionTransaction }
