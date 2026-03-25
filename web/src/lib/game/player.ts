type PlayerPosition = {
  ID: string
  actor_ID: string | null
}

type Player = {
  ID: string
  positions_capacity: number,
  positions: Array<PlayerPosition>
}


export type { Player }
