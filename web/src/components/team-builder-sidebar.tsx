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
import type { TeamConfig } from '#/lib/stores/config'
import { cloneTeamConfig } from '#/lib/team-storage'

function TeamBuilderSidebar({
  savedTeams,
  onLoadTeam,
}: {
  savedTeams: TeamConfig[]
  onLoadTeam: (config: TeamConfig) => void
}) {
  const loadSavedTeam = (team: TeamConfig) => {
    const loadedTeam = cloneTeamConfig(team)
    onLoadTeam(loadedTeam)
  }

  return (
    <Sidebar collapsible="none" className="border-r">
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
          <SidebarGroupLabel>Teams</SidebarGroupLabel>
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
                  <SidebarMenuItem key={team.name}>
                    <SidebarMenuButton onClick={() => loadSavedTeam(team)}>
                      <span>{team.name}</span>
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
