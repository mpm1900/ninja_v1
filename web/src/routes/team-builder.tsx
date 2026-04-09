import { AbilitySelect } from '#/components/ability-select'
import { ActionsTable } from '#/components/actions-table'
import { ActorCombobox } from '#/components/actor-combobox'
import { AppHeader } from '#/components/app-header'
import { FocusSelect } from '#/components/focus-select'
import { ItemSelect } from '#/components/item-select'
import { TeamBuilderStat } from '#/components/team-builder-stat'
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
} from '#/components/ui/sidebar'
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
  cloneTeamConfig,
  parseSavedTeamsJSON,
  serializeSavedTeams,
  upsertSavedTeam,
} from '#/lib/team-storage'
import { sendContextMessage } from '#/lib/stores/socket'
import { useForm, useStore } from '@tanstack/react-form'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute, redirect } from '@tanstack/react-router'
import { useEffect, useState } from 'react'

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
    const loadedTeam = cloneTeamConfig(team)
    form.setFieldValue('name', loadedTeam.name)
    form.setFieldValue('actors', loadedTeam.actors)
    form.setFieldValue(
      'selected_index',
      Math.min(
        loadedTeam.selected_index,
        Math.max(loadedTeam.actors.length - 1, 0)
      )
    )
  }

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
                          <SidebarMenuButton
                            onClick={() => loadSavedTeam(team)}
                          >
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

          <SidebarInset className="min-h-0">
            <div className="flex h-full min-h-0 flex-row-reverse items-stretch p-4 gap-8">
              <div>
                <div className="mb-4 flex items-center justify-between gap-6">
                  <form.Field name="name">
                    {(field) => (
                      <Field>
                        <FieldContent>
                          <Input
                            placeholder="Team Name"
                            value={field.state.value}
                            onChange={(e) => field.handleChange(e.target.value)}
                          />
                        </FieldContent>
                      </Field>
                    )}
                  </form.Field>
                  <form.Subscribe>
                    {({ isValid, isSubmitting, isValidating, values }) => {
                      return (
                        <div className="flex gap-2">
                          <Button
                            disabled={!values.name.trim()}
                            onClick={() => saveTeam(values as TeamConfig)}
                          >
                            Save Team
                          </Button>
                          <Button
                            disabled={!isValid || isSubmitting || isValidating}
                            onClick={form.handleSubmit}
                          >
                            Load Team
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
                        <div className="space-y-2">
                          <div className="">Team: {selected.length}/6</div>
                          {field.state.value.map((_, i) => (
                            <form.Field key={i} name={`actors[${i}]`}>
                              {(actorID) => (
                                <ActorCombobox
                                  active={active}
                                  selected={selected}
                                  value={actorID.state.value?.actor_ID}
                                  onValueChange={(actor) => {
                                    actorID.handleChange({
                                      actor_ID: actor.actor_ID,
                                      config: {
                                        ability_ID:
                                          actor.abilities[0]?.ID ?? null,
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
                                    })
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
                                form.pushFieldValue('actors', {
                                  actor_ID: actor.actor_ID,
                                  config: {
                                    ability_ID: null,
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
                                })
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
                      <div className="text-sm font-medium">
                        Editing {def.name}
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
                            <table className="flex-1">
                              <tbody>
                                <tr>
                                  <td colSpan={3}>Stats</td>
                                  <td
                                    colSpan={2}
                                    className={
                                      total > 64 ? 'text-destructive' : ''
                                    }
                                  >
                                    {total}
                                    /64
                                  </td>
                                </tr>

                                <TeamBuilderStat
                                  total={total}
                                  focus={actor.config.focus}
                                  base={def}
                                  stat="hp"
                                  config={actor.config}
                                  onConfigChange={(config) => {
                                    form.setFieldValue(
                                      `actors[${selected_index}].config`,
                                      config
                                    )
                                  }}
                                />
                                <TeamBuilderStat
                                  total={total}
                                  focus={actor.config.focus}
                                  base={def}
                                  stat="stamina"
                                  config={actor.config}
                                  onConfigChange={(config) => {
                                    form.setFieldValue(
                                      `actors[${selected_index}].config`,
                                      config
                                    )
                                  }}
                                />

                                <TeamBuilderStat
                                  total={total}
                                  focus={actor.config.focus}
                                  base={def}
                                  stat="attack"
                                  config={actor.config}
                                  onConfigChange={(config) => {
                                    form.setFieldValue(
                                      `actors[${selected_index}].config`,
                                      config
                                    )
                                  }}
                                />
                                <TeamBuilderStat
                                  total={total}
                                  focus={actor.config.focus}
                                  base={def}
                                  stat="defense"
                                  config={actor.config}
                                  onConfigChange={(config) => {
                                    form.setFieldValue(
                                      `actors[${selected_index}].config`,
                                      config
                                    )
                                  }}
                                />
                                <TeamBuilderStat
                                  total={total}
                                  focus={actor.config.focus}
                                  base={def}
                                  stat="chakra_attack"
                                  config={actor.config}
                                  onConfigChange={(config) => {
                                    form.setFieldValue(
                                      `actors[${selected_index}].config`,
                                      config
                                    )
                                  }}
                                />
                                <TeamBuilderStat
                                  total={total}
                                  focus={actor.config.focus}
                                  base={def}
                                  stat="chakra_defense"
                                  config={actor.config}
                                  onConfigChange={(config) => {
                                    form.setFieldValue(
                                      `actors[${selected_index}].config`,
                                      config
                                    )
                                  }}
                                />
                                <TeamBuilderStat
                                  total={total}
                                  focus={actor.config.focus}
                                  base={def}
                                  stat="speed"
                                  config={actor.config}
                                  onConfigChange={(config) => {
                                    form.setFieldValue(
                                      `actors[${selected_index}].config`,
                                      config
                                    )
                                  }}
                                />
                              </tbody>
                            </table>
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
                              enabled
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
