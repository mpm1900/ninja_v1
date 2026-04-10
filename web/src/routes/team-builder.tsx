import { AbilitySelect } from '#/components/ability-select'
import { ActionsTable } from '#/components/actions-table'
import { ActorCombobox } from '#/components/actor-combobox'
import { AppHeader } from '#/components/app-header'
import { FocusSelect } from '#/components/focus-select'
import { ItemSelect } from '#/components/item-select'
import { Button } from '#/components/ui/button'
import { Field, FieldContent, FieldLabel } from '#/components/ui/field'
import { Input } from '#/components/ui/input'
import { SidebarInset, SidebarProvider } from '#/components/ui/sidebar'
import { NULL_CONTEXT } from '#/lib/game/context'
import { actionsQuery } from '#/lib/queries/actions'
import { actorsQuery } from '#/lib/queries/actors'
import { instancesQuery } from '#/lib/queries/instances'
import { clientsStore } from '#/lib/stores/clients'
import {
  TeamSchema,
  type TeamActor,
  type TeamConfig,
} from '#/lib/stores/config'
import {
  SAVED_TEAMS_KEY,
  parseSavedTeamsJSON,
  serializeSavedTeams,
  upsertSavedTeam,
} from '#/lib/team-storage'
import { sendContextMessage } from '#/lib/stores/socket'
import { useForm, useStore } from '@tanstack/react-form'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute, redirect } from '@tanstack/react-router'
import { useEffect, useState } from 'react'
import { type ActorDef } from '#/lib/game/actor'
import { TeamBuilderStats } from '#/components/team-builder-stats'
import { Save, Swords } from 'lucide-react'
import { TeamBuilderSidebar } from '#/components/team-builder-sidebar'

function makeConfigFromDef(def: ActorDef): TeamActor {
  return {
    actor_ID: def.actor_ID,
    config: {
      ability_ID: def.abilities[0]?.ID ?? null,
      item_ID: null,
      action_IDs: [],
      focus: 'none',
      aux_stats: {
        hp: 0,
        stamina: 0,
        speed: 0,
        attack: 0,
        defense: 0,
        chakra_attack: 0,
        chakra_defense: 0,
      },
    },
  }
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
  const nav = Route.useNavigate()
  const client = useStore(clientsStore, (s) => s.me)
  const query = useSuspenseQuery(actorsQuery)
  const actions = useSuspenseQuery(actionsQuery)
  const form = useForm({
    defaultValues: {
      name: 'Team',
      selected_index: 0,
      actors: [] as TeamActor[],
    },
    validators: {
      onMount: TeamSchema,
      onChange: TeamSchema,
    },
    onSubmit: ({ value }) => {
      if (!client) return
      sendContextMessage({
        type: 'set-team',
        client_ID: client.ID,
        team_config: value,
        context: NULL_CONTEXT,
      })
      nav({ to: '/battle' })
    },
  })

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

  const loadSavedTeam = (team: TeamConfig) => {
    form.setFieldValue('name', team.name)
    form.setFieldValue('actors', team.actors)
    form.setFieldValue(
      'selected_index',
      Math.min(team.selected_index, Math.max(team.actors.length - 1, 0))
    )
  }

  return (
    <main className="flex h-screen flex-col overflow-hidden bg-zinc-800">
      <AppHeader />

      <section className="flex flex-1 min-h-0 p-4 md:p-6">
        <SidebarProvider className="h-full min-h-0 w-full overflow-hidden rounded-xl border bg-background shadow-sm">
          <TeamBuilderSidebar
            onLoadTeam={loadSavedTeam}
            savedTeams={savedTeams}
          />

          <SidebarInset className="min-h-0">
            <div className="flex h-full min-h-0 flex-row-reverse items-stretch p-4 gap-8">
              <div>
                <div className="mb-4 flex items-center justify-end gap-6">
                  <form.Subscribe>
                    {({ isValid, isSubmitting, isValidating, values }) => {
                      return (
                        <div className="flex gap-2">
                          <Button
                            size="icon"
                            disabled={!values.name.trim()}
                            onClick={() => saveTeam(values as TeamConfig)}
                          >
                            <Save />
                          </Button>
                          <Button
                            size="icon"
                            disabled={!isValid || isSubmitting || isValidating}
                            onClick={form.handleSubmit}
                          >
                            <Swords />
                          </Button>
                        </div>
                      )
                    }}
                  </form.Subscribe>
                </div>
                <form.Field name="actors" mode="array">
                  {(field) => (
                    <form.Subscribe
                      selector={(state) => ({
                        selected: state.values.actors.map((a) => a.actor_ID!),
                        active:
                          state.values.actors[state.values.selected_index]
                            ?.actor_ID,
                      })}
                    >
                      {({ selected, active }) => (
                        <div className="flex flex-col gap-2 min-w-sm">
                          <div className="">Team: {selected.length}/6</div>
                          {field.state.value.map((_, i) => (
                            <form.Field key={i} name={`actors[${i}]`}>
                              {(actorID) => (
                                <ActorCombobox
                                  active={active}
                                  selected={selected}
                                  value={actorID.state.value?.actor_ID}
                                  onValueChange={(actor) => {
                                    actorID.handleChange(
                                      makeConfigFromDef(actor)
                                    )
                                    form.setFieldValue('selected_index', i)
                                  }}
                                  onClick={() =>
                                    form.setFieldValue('selected_index', i)
                                  }
                                />
                              )}
                            </form.Field>
                          ))}
                          {selected.length < 6 && (
                            <ActorCombobox
                              active={active}
                              selected={selected}
                              value={undefined}
                              onValueChange={(actor) => {
                                form.pushFieldValue(
                                  'actors',
                                  makeConfigFromDef(actor)
                                )
                                form.setFieldValue(
                                  'selected_index',
                                  selected.length
                                )
                              }}
                            />
                          )}
                        </div>
                      )}
                    </form.Subscribe>
                  )}
                </form.Field>
              </div>

              <form.Subscribe
                selector={(state) => ({
                  selected_index: state.values.selected_index,
                  actors: state.values.actors,
                })}
              >
                {({ selected_index, actors }) => {
                  const actor =
                    selected_index === null
                      ? null
                      : (actors[selected_index] ?? null)

                  if (selected_index === null || !actor) {
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
                    <div className="flex-1 space-y-4 overflow-auto">
                      <div className="flex justify-between items-end gap-6">
                        <form.Field name="name">
                          {(field) => (
                            <Field>
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
                          variant="destructive"
                          onClick={() => {
                            form.removeFieldValue('actors', selected_index)
                          }}
                        >
                          Remove {def.name}
                        </Button>
                      </div>

                      <div className="flex gap-4">
                        <div className="flex flex-col gap-2 min-w-1/4">
                          <img
                            src={def.sprite_url}
                            className="object-cover size-16"
                          />
                          <FocusSelect
                            value={actor.config?.focus ?? 'none'}
                            onValueChange={(focus) => {
                              form.setFieldValue(
                                `actors[${selected_index}].config.focus`,
                                focus
                              )
                            }}
                          />
                          <AbilitySelect
                            options={def.abilities}
                            value={actor.config?.ability_ID ?? null}
                            onValueChange={(ability_ID) => {
                              form.setFieldValue(
                                `actors[${selected_index}].config.ability_ID`,
                                ability_ID
                              )
                            }}
                          />
                          <ItemSelect
                            value={actor.config?.item_ID ?? null}
                            onValueChange={(item_ID) => {
                              form.setFieldValue(
                                `actors[${selected_index}].config.item_ID`,
                                item_ID
                              )
                            }}
                          />
                        </div>

                        <form.Subscribe
                          selector={(s) =>
                            Object.values(
                              s.values.actors[s.values.selected_index]?.config
                                .aux_stats ?? {}
                            ).reduce((sum, value) => sum + value, 0)
                          }
                        >
                          {(total) => (
                            <TeamBuilderStats
                              total={total}
                              config={actor.config}
                              def={def}
                              onConfigChange={(config) => {
                                form.setFieldValue(
                                  `actors[${selected_index}].config`,
                                  config
                                )
                              }}
                            />
                          )}
                        </form.Subscribe>
                      </div>
                      <div>
                        <form.Subscribe
                          selector={(state) => ({
                            selected_index: state.values.selected_index,
                            actors: state.values.actors,
                          })}
                        >
                          {({ selected_index, actors }) => (
                            <ActionsTable
                              total={def.action_count}
                              data={actions.data.filter((a) =>
                                def.action_IDs.includes(a.ID)
                              )}
                              rowSelection={Object.fromEntries(
                                actors[selected_index]?.config.action_IDs?.map(
                                  (id) => [id, true]
                                ) ?? []
                              )}
                              onRowSelectionChange={(selection) => {
                                form.setFieldValue(
                                  `actors[${selected_index}].config.action_IDs`,
                                  Object.keys(selection)
                                )
                              }}
                            />
                          )}
                        </form.Subscribe>
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
