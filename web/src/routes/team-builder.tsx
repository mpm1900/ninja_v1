import { ActionsTable } from '#/components/actions-table'
import { ActorCombobox } from '#/components/actor-combobox'
import { ActorsTable } from '#/components/actors-table'
import { AppHeader } from '#/components/app-header'
import { FocusSelect } from '#/components/focus-select'
import { Button } from '#/components/ui/button'
import { Field, FieldContent } from '#/components/ui/field'
import { Input } from '#/components/ui/input'
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarInset,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarProvider,
  SidebarSeparator,
} from '#/components/ui/sidebar'
import { actionsQuery } from '#/lib/queries/actions'
import { actorsQuery } from '#/lib/queries/actors'
import { instancesQuery } from '#/lib/queries/instances'
import type { ActorConfig } from '#/lib/stores/socket'
import { useForm } from '@tanstack/react-form'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute, redirect } from '@tanstack/react-router'

type TeamActor = {
  actor_ID: string
  config: ActorConfig
}

export const Route = createFileRoute('/team-builder')({
  component: RouteComponent,
  beforeLoad: ({ context }) => {
    if (!context.auth.user) {
      throw redirect({ to: '/login' })
    }
  },
  loader: async ({ context }) => {
    await context.queryClient.ensureQueryData(actionsQuery)
    await context.queryClient.ensureQueryData(actorsQuery)
    await context.queryClient.ensureQueryData(instancesQuery)
  },
})

function RouteComponent() {
  const actors = useSuspenseQuery(actorsQuery)
  const actions = useSuspenseQuery(actionsQuery)
  const form = useForm({
    defaultValues: {
      selected_index: 0,
      team_actors: [null, null, null, null, null, null] as (TeamActor | null)[],
    },
  })
  return (
    <main className="flex h-screen flex-col overflow-hidden">
      <AppHeader />

      <section className="flex flex-1 min-h-0 p-4 md:p-6">
        <SidebarProvider className="h-full min-h-0 w-full overflow-hidden rounded-xl border bg-background shadow-sm">
          <Sidebar collapsible="none" className="border-r">
            <SidebarHeader>
              <SidebarMenu>
                <SidebarMenuItem>
                  <SidebarMenuButton size="lg" asChild>
                    <a href="#">
                      <div className="flex flex-col gap-0.5 leading-none">
                        <span className="font-medium">Team Builder</span>
                        <span className="text-xs text-muted-foreground">
                          v1.0.0
                        </span>
                      </div>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              </SidebarMenu>
            </SidebarHeader>

            <SidebarSeparator />

            <SidebarContent>
              <SidebarGroup>
                <SidebarGroupLabel>Teams</SidebarGroupLabel>
                <SidebarGroupContent>
                  <SidebarMenu className="gap-1">
                    <SidebarMenuItem>
                      <SidebarMenuButton>
                        <span>Team 1</span>
                      </SidebarMenuButton>
                    </SidebarMenuItem>
                    <SidebarMenuItem>
                      <SidebarMenuButton>
                        <span>Team 2</span>
                      </SidebarMenuButton>
                    </SidebarMenuItem>
                  </SidebarMenu>
                </SidebarGroupContent>
              </SidebarGroup>
            </SidebarContent>
          </Sidebar>

          <SidebarInset className="min-h-0">
            <div className="flex h-full min-h-0 flex-row-reverse items-start p-4 gap-8">
              <div>
                <div className="mb-4 flex items-center justify-between gap-6">
                  <Field>
                    <FieldContent>
                      <Input placeholder="Team Name" />
                    </FieldContent>
                  </Field>
                  <div>
                    <Button>Save Team</Button>
                  </div>
                </div>
                <form.Field name="team_actors" mode="array">
                  {(field) => (
                    <form.Subscribe
                      selector={(state) =>
                        state.values.team_actors
                          .filter((a) => !!a?.actor_ID)
                          .map((a) => a?.actor_ID!)
                      }
                    >
                      {(selected) => (
                        <div className="space-y-2">
                          {field.state.value.map((_, i) => (
                            <form.Field name={`team_actors[${i}].actor_ID`}>
                              {(actor_ID) => (
                                <ActorCombobox
                                  selected={selected}
                                  value={actor_ID.state.value}
                                  onValueChange={(v) => {
                                    actor_ID.handleChange(v)
                                    form.setFieldValue('selected_index', i)
                                  }}
                                  onClick={() =>
                                    form.setFieldValue('selected_index', i)
                                  }
                                />
                              )}
                            </form.Field>
                          ))}
                        </div>
                      )}
                    </form.Subscribe>
                  )}
                </form.Field>
              </div>

              <form.Subscribe
                selector={(state) => ({
                  selected_index: state.values.selected_index,
                  team_actors: state.values.team_actors,
                })}
              >
                {({ selected_index, team_actors }) => {
                  const teamActor =
                    selected_index === null
                      ? null
                      : (team_actors[selected_index] ?? null)

                  if (selected_index === null || !teamActor) {
                    return (
                      <div className="flex-1 text-sm text-muted-foreground">
                        Select a shinobi portrait to edit config
                      </div>
                    )
                  }

                  const baseActor = actors.data.find(
                    (a) => a.actor_ID === teamActor.actor_ID
                  )

                  return (
                    <div className="flex-1 space-y-4">
                      <div className="text-sm font-medium">
                        Editing slot {selected_index + 1}
                      </div>

                      <div className="grid grid-cols-2">
                        <div>
                          <FocusSelect
                            value={teamActor.config?.focus ?? 'none'}
                            onValueChange={(focus) => {
                              form.setFieldValue(
                                `team_actors[${selected_index}].config.focus`,
                                focus
                              )
                            }}
                          />
                        </div>
                        {baseActor && (
                          <div>
                            <ActionsTable
                              data={actions.data.filter((a) =>
                                baseActor.action_IDs.includes(a.ID)
                              )}
                              enabled
                              rowSelection={{}}
                              onRowSelectionChange={() => {}}
                            />
                          </div>
                        )}
                      </div>
                    </div>
                  )
                }}
              </form.Subscribe>
            </div>
          </SidebarInset>
        </SidebarProvider>
      </section>
    </main>
  )
}
