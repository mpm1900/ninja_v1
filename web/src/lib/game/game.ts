import type { Actor } from './actor'

type Context = {
  sourcePlayerID: string | null

  parentActorID: string | null
  sourceActorID: string | null

  targetActorIDs: Array<string>
  targetPositionIDs: Array<string>
}

type Game = {
  actors: Actor[]

  // TODO
}

export type { Game, Context }
