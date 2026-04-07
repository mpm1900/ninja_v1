import type { Context } from './context'

type Modifier = {
  ID: string
  group_ID: string
  name: string
  duration: number | null
  icon?: string
}

type ModifierTransaction = {
  ID: string
  context: Context
  mutation: Modifier
}

export type { Modifier, ModifierTransaction }
