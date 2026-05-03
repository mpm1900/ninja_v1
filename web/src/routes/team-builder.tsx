import { AppHeader } from '#/components/app-header'
import { Button } from '#/components/ui/button'
import { Field, FieldContent, FieldLabel } from '#/components/ui/field'
import { Input } from '#/components/ui/input'
import { SidebarInset, SidebarProvider } from '#/components/ui/sidebar'
import { actionsQuery } from '#/lib/queries/actions'
import { actorsQuery } from '#/lib/queries/actors'
import { instancesQuery } from '#/lib/queries/instances'
import { clientsStore } from '#/lib/stores/clients'
import { type TeamConfig } from '#/lib/stores/config'
import { useStore } from '@tanstack/react-form'
import { useSuspenseQuery } from '@tanstack/react-query'
import { ClientOnly, createFileRoute, redirect } from '@tanstack/react-router'
import { Save, Swords, Trash } from 'lucide-react'
import { TeamBuilderSidebar } from '#/components/team-builder-sidebar'
import { useSavedTeams } from '#/hooks/use-saved-teams'
import { useTeamBuilderForm } from '#/hooks/use-team-builder-form'
import { TeamBuilderActionsTable } from '#/components/team-builder-actions-table'
import { TeamBuilderList } from '#/components/team-builder-list'
import { TeamBuilderActorConfig } from '#/components/team-builder-actor-config'

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
  const nav = Route.useNavigate()
  const client = useStore(clientsStore, (s) => s.me)
  const query = useSuspenseQuery(actorsQuery)
  const form = useTeamBuilderForm({
    clientID: client?.ID ?? '',
    onSubmit: () => {
      nav({
        to: '/lobby',
      })
    },
  })

  const { savedTeams, saveTeam } = useSavedTeams()

  const loadSavedTeam = (team: TeamConfig) => {
    form.setFieldValue('name', team.name)
    form.setFieldValue('actors', team.actors)
    form.setFieldValue(
      'selected_index',
      Math.min(team.selected_index, Math.max(team.actors.length - 1, 0))
    )
  }

  return (
    <ClientOnly>
      <main className="flex h-screen flex-col overflow-hidden">
        <AppHeader />

        <section className="flex flex-1 min-h-0 p-4 md:p-6">
          <SidebarProvider className="h-full min-h-0 w-full overflow-hidden rounded-xl border border-stone-300/30 ring ring-black bg-stone-950 shadow-sm">
            <TeamBuilderSidebar
              onLoadTeam={loadSavedTeam}
              savedTeams={savedTeams}
            />

            <SidebarInset className="min-h-0">
              <form.AppForm>
                <div className="flex h-full min-h-0 flex-row-reverse items-stretch p-4 gap-8">
                  <div>
                    <div className="mb-4 flex items-center justify-end gap-6">
                      <form.Subscribe>
                        {({ isValid, isSubmitting, isValidating, values }) => {
                          return (
                            <div className="flex gap-2">
                              <Button
                                size="icon"
                                disabled={!values.name.trim() || !isValid}
                                onClick={() => saveTeam(values as TeamConfig)}
                              >
                                <Save />
                              </Button>
                              <Button
                                size="icon"
                                disabled={
                                  !isValid ||
                                  isSubmitting ||
                                  isValidating ||
                                  !client
                                }
                                onClick={form.handleSubmit}
                              >
                                <Swords />
                              </Button>
                            </div>
                          )
                        }}
                      </form.Subscribe>
                    </div>
                    <TeamBuilderList form={form} />
                  </div>

                  <form.Subscribe
                    selector={(state) => ({
                      selected_index: state.values.selected_index,
                      actor: state.values.actors[state.values.selected_index],
                    })}
                  >
                    {({ selected_index, actor }) => {
                      if (!actor) {
                        return (
                          <div className="flex-1 text-sm text-muted-foreground">
                            Select a shinobi portrait to edit config
                          </div>
                        )
                      }

                      const def = query.data.find(
                        (a) => a.actor_ID === actor.actor_ID
                      )

                      if (!def) return null

                      return (
                        <div className="flex-1 flex flex-col gap-2 overflow-auto">
                          <div className="flex justify-between items-end gap-6">
                            <form.Field name="name">
                              {(field) => (
                                <Field className="max-w-sm">
                                  <FieldLabel>Team Name</FieldLabel>
                                  <FieldContent>
                                    <Input
                                      placeholder="Team Name"
                                      value={field.state.value}
                                      onChange={(e) =>
                                        field.handleChange(e.target.value)
                                      }
                                    />
                                  </FieldContent>
                                </Field>
                              )}
                            </form.Field>
                            <Button
                              size="icon"
                              variant="destructive"
                              onClick={() => {
                                form.removeFieldValue('actors', selected_index)
                              }}
                            >
                              <Trash />
                            </Button>
                          </div>

                          <TeamBuilderActorConfig form={form} def={def} />
                          <TeamBuilderActionsTable form={form} def={def} />
                        </div>
                      )
                    }}
                  </form.Subscribe>
                </div>
              </form.AppForm>
            </SidebarInset>
          </SidebarProvider>
        </section>
      </main>
    </ClientOnly>
  )
}
