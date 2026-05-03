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
import { useQuery, useSuspenseQuery } from '@tanstack/react-query'
import { ClientOnly, createFileRoute, redirect } from '@tanstack/react-router'
import { Loader2, Save, Swords, Trash } from 'lucide-react'
import { TeamBuilderSidebar } from '#/components/team-builder-sidebar'
import { useTeamBuilderForm } from '#/hooks/use-team-builder-form'
import { TeamBuilderActionsTable } from '#/components/team-builder-actions-table'
import { TeamBuilderList } from '#/components/team-builder-list'
import { TeamBuilderActorConfig } from '#/components/team-builder-actor-config'
import { teamsQuery, type Team } from '#/lib/queries/teams'
import { useUpsertTeam } from '#/lib/mutations/upsert-team'
import { useState } from 'react'
import { useDeleteTeam } from '#/lib/mutations/delete-team'

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

  const savedQuery = useQuery(teamsQuery)
  const [id, setID] = useState<string | null>(null)
  const upsertMutation = useUpsertTeam()
  const deleteMutation = useDeleteTeam()

  const loadSavedTeam = (team: Team) => {
    setID(team.id)
    form.setFieldValue('name', team.team_config.name)
    form.setFieldValue('actors', team.team_config.actors)
    form.setFieldValue(
      'selected_index',
      Math.min(team.team_config.selected_index ?? 0, Math.max(team.team_config.actors.length - 1, 0))
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
              savedTeams={savedQuery.data ?? []}
              onDeleteTeam={t => {
                deleteMutation.mutate(t.id, {
                  onSuccess: () => {
                    savedQuery.refetch()
                  }
                })
              }}
            />

            <SidebarInset className="min-h-0 bg-stone-950">
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
                                disabled={upsertMutation.isPending || !values.name.trim() || !isValid}
                                onClick={() => {
                                  upsertMutation.mutate({
                                    id,
                                    config: values as TeamConfig
                                  }, {
                                    onSuccess: () => {
                                      savedQuery.refetch()
                                    }
                                  })
                                }}
                              >
                                {upsertMutation.isPending ? <Loader2 className='animate-spin' /> : <Save />}
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
