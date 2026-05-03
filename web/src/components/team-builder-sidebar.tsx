import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from './ui/sidebar'
import { cloneTeamConfig } from '#/lib/team-storage'
import type { Team } from '#/lib/queries/teams'
import { Button } from './ui/button'
import { Trash } from 'lucide-react'

function TeamBuilderSidebar({
  savedTeams,
  onLoadTeam,
  onDeleteTeam,
}: {
  savedTeams: Team[]
  onLoadTeam: (team: Team) => void
  onDeleteTeam: (team: Team) => void
}) {
  const loadSavedTeam = (team: Team) => {
    const team_config = cloneTeamConfig(team.team_config)
    onLoadTeam({
      ...team,
      team_config,
    })
  }

  return (
    <Sidebar collapsible="none" className="border-r bg-stone-900">
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild>
              <a href="#">
                <div className="flex flex-col gap-0.5 leading-none">
                  <span className="font-medium">Team Builder</span>
                  <span className="text-xs text-muted-foreground">v1.0.0</span>
                </div>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>

      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel className="justify-between">
            Teams{' '}
            <Button
              size="xs"
              onClick={() => {
                onLoadTeam({
                  id: null,
                  team_config: {
                    name: '',
                    selected_index: 0,
                    actors: [],
                  },
                })
              }}
            >
              Add Team
            </Button>
          </SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu className="gap-1">
              {savedTeams.length === 0 ? (
                <SidebarMenuItem>
                  <SidebarMenuButton disabled>
                    <span className="text-muted-foreground">
                      No saved teams
                    </span>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ) : (
                savedTeams.map((team) => (
                  <SidebarMenuItem key={team.team_config.name}>
                    <SidebarMenuButton
                      className="justify-between"
                      onClick={() => loadSavedTeam(team)}
                    >
                      <span>{team.team_config.name}</span>
                      <Button
                        size="icon-xs"
                        variant="ghost"
                        onClick={(e) => {
                          e.stopPropagation()
                          e.preventDefault()
                          onDeleteTeam(team)
                        }}
                      >
                        <Trash />
                      </Button>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))
              )}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  )
}

export { TeamBuilderSidebar }
