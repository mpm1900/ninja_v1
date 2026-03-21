import type { Context } from "./game"

type Modifier = {
  ID: string
  name: string
  duration: number | null
}

type ModifierTransaction = {
  ID: string
  context: Context
  mutation: Modifier
}

export type { Modifier, ModifierTransaction }
