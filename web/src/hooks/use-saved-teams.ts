import type { TeamConfig } from "#/lib/stores/config"
import { parseSavedTeamsJSON, SAVED_TEAMS_KEY, serializeSavedTeams, upsertSavedTeam } from "#/lib/team-storage"
import { useEffect, useState } from "react"

function useSavedTeams() {
  const [savedTeams, setSavedTeams] = useState<TeamConfig[]>([])

  useEffect(() => {
    setSavedTeams(parseSavedTeamsJSON(localStorage.getItem(SAVED_TEAMS_KEY)))
  }, [])

  const persistSavedTeams = (teams: TeamConfig[]) => {
    setSavedTeams(teams)
    localStorage.setItem(SAVED_TEAMS_KEY, serializeSavedTeams(teams))
  }

  const saveTeam = (team: TeamConfig) => {
    const nextTeams = upsertSavedTeam(savedTeams, team)
    if (nextTeams === savedTeams) return
    persistSavedTeams(nextTeams)
  }


  return { savedTeams, saveTeam }
}


export { useSavedTeams }
