import type { Actor } from './actor'

type Context = {
  sourcePlayerID: string

  parentActorID: string
  sourceActorID: string

  targetActorIDs: Array<string>
  targetPositionIDs: Array<string>
}

type Game = {
  actors: Actor[]

  // TODO
}

export type { Game, Context }
