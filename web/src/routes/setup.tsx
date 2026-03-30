import { ActorsTable } from '#/components/actors-table'
import { actorsQuery } from '#/lib/queries/actors'
import { instancesQuery } from '#/lib/queries/instances'
import { sendContextMessage, socketStore } from '#/lib/stores/socket'
import { useStore } from '@tanstack/react-store'
import { useSuspenseQuery } from '@tanstack/react-query'
import { ClientOnly, createFileRoute, redirect } from '@tanstack/react-router'
import { gameStore } from '#/lib/stores/game'
import { clientsStore } from '#/lib/stores/clients'
import { ActorCard } from '#/components/actor-card'
import { modifiersQuery } from '#/lib/queries/modifiers'
import { ModifiersTable } from '#/components/modifiers-table'
import { actionsQuery } from '#/lib/queries/actions'
import { ActorControl } from '#/components/actor-control'
import { triggersQuery } from '#/lib/queries/trigger-types'
import { ActionQueue } from '#/components/action-queue'
import { PromptController } from '#/components/prompt-controller'
import { ActorThumbnail } from '#/components/actor-thumbnail'
import { cn } from '#/lib/utils'
import { AppHeader } from '#/components/app-header'

export const Route = createFileRoute('/setup')({
  beforeLoad: ({ context }) => {
    if (!context.auth.user) {
      throw redirect({ to: '/login' })
    }
  },
  component: App,
  loader: async ({ context }) => {
    await context.queryClient.ensureQueryData(actionsQuery)
    await context.queryClient.ensureQueryData(actorsQuery)
    await context.queryClient.ensureQueryData(modifiersQuery)
    await context.queryClient.ensureQueryData(triggersQuery)
    await context.queryClient.ensureQueryData(instancesQuery)
  },
})

function App() {
  const actors = useSuspenseQuery(actorsQuery)
  const status = useStore(socketStore, (s) => s.status)
  const client = useStore(clientsStore, (c) => c.me)
  const game = useStore(gameStore, (g) => g)

  return (
    <ClientOnly>
      <PromptController />
      <main className="min-w-0 overflow-x-hidden">
        <AppHeader />
        <div className="flex min-w-0">
          <div className="min-w-0 space-y-2 flex-1 overflow-auto">
            <ActionQueue />

            <div>
              {game.players.map((player, i) => (
                <div
                  key={player.ID}
                  className={cn('space-y-2', {
                    'border-b pb-4 mb-4': i !== game.players.length - 1,
                  })}
                >
                  <div className="flex gap-2 px-4">
                    {game.actors
                      .filter((a) => a.player_ID === player.ID)
                      .map((a, i) => (
                        <ActorThumbnail key={a.ID} actor={a} index={i} />
                      ))}
                  </div>
                  <div className="flex items-end gap-2">
                    {game.actors
                      .filter(
                        (a) => !!a.position_ID && a.player_ID == player.ID
                      )
                      .map((a) => (
                        <ActorCard
                          key={a.ID}
                          actor={a}
                          clientID={client?.ID}
                          game={game}
                          selected={false}
                        />
                      ))}
                  </div>
                </div>
              ))}
            </div>

            <ActorsTable
              data={actors.data}
              enabled={!!client && status === 'open' && game.status == 'idle'}
              rowSelection={Object.fromEntries(
                game.actors
                  .filter((a) => a.player_ID === client?.ID)
                  .map((a) => [a.actor_ID, true])
              )}
              onRowCheckedChange={(a, checked) => {
                if (!client) return
                if (checked) {
                  sendContextMessage({
                    type: 'add-actor',
                    client_ID: client.ID,
                    context: {
                      action_ID: null,
                      source_player_ID: client.ID,
                      source_actor_ID: a.actor_ID,
                      parent_actor_ID: a.actor_ID,
                      target_actor_IDs: [],
                      target_position_IDs: [],
                    },
                  })
                } else {
                  const actor = game.actors.find(
                    (ga) =>
                      ga.player_ID === client.ID && ga.actor_ID === a.actor_ID
                  )!
                  sendContextMessage({
                    type: 'remove-actor',
                    client_ID: client.ID,
                    context: {
                      action_ID: null,
                      source_player_ID: client.ID,
                      source_actor_ID: actor.ID,
                      parent_actor_ID: actor.ID,
                      target_actor_IDs: [],
                      target_position_IDs: [],
                    },
                  })
                }
              }}
              subRow={({ row }) =>
                client && (
                  <ActorControl
                    actor={
                      game.actors.find(
                        (a) => a.actor_ID === row.original.actor_ID
                      )!
                    }
                    enabled={game.status == 'idle'}
                  />
                )
              }
            />
            <ModifiersTable
              data={game.modifiers ?? []}
              onRowRemove={(modifier) => {
                if (!client) return

                sendContextMessage({
                  type: 'remove-modifier',
                  client_ID: client.ID,
                  modifier_ID: modifier.ID,
                  context: {
                    action_ID: null,
                    source_player_ID: null,
                    source_actor_ID: null,
                    parent_actor_ID: null,
                    target_actor_IDs: [],
                    target_position_IDs: [],
                  },
                })
              }}
            />
          </div>
          <div className="w-sm hidden max-w-[200px] shrink-0 overflow-x-auto">
            {game.log?.map((log, i) => (
              <div key={i} className="break-words">
                {log}
              </div>
            ))}
          </div>
        </div>

        <pre className="max-w-full overflow-x-auto whitespace-pre-wrap break-words">
          {JSON.stringify(game, null, 2)}
        </pre>
      </main>
    </ClientOnly>
  )
}
