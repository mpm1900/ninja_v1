import type { Actor } from './actor'
import type { Modifier } from './modifier'

type Context = {
  sourcePlayerID: string | null

  parentActorID: string | null
  sourceActorID: string | null

  targetActorIDs: Array<string>
  targetPositionIDs: Array<string>
}

type Game = {
  actors: Actor[]
  modifiers: Modifier[]

  // TODO
}

export type { Game, Context }
