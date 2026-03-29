type PlayerPosition = {
  ID: string
  actor_ID: string | null
}

type PlayerUser = {
  id: string
  username: string
  email: string
}

type Player = {
  ID: string
  user: PlayerUser
  positions_capacity: number
  positions: Array<PlayerPosition>
  team_capacity: number
}

export type { Player, PlayerUser }
