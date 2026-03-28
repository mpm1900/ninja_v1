type PlayerPosition = {
  ID: string
  actor_ID: string | null
}

type Player = {
  ID: string
  positions_capacity: number,
  positions: Array<PlayerPosition>
  team_capacity: number
}


export type { Player }
