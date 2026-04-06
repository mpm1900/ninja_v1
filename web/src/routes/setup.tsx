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
import { ActorControl } from '#/components/actor-control'
import { ActionQueue } from '#/components/action-queue'
import { PromptController } from '#/components/prompt-controller'
import { ActorThumbnail } from '#/components/actor-thumbnail'
import { cn } from '#/lib/utils'
import { AppHeader } from '#/components/app-header'
import { actionsQuery } from '#/lib/queries/actions'
import { PlayerPositions } from '#/components/player-positions'
import { PlayerThumbnails } from '#/components/player-thumbnails'

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
    await context.queryClient.ensureQueryData(instancesQuery)
  },
})

function App() {
  const actors = useSuspenseQuery(actorsQuery)
  const status = useStore(socketStore, (s) => s.status)
  const client = useStore(clientsStore, (c) => c.me)
  const game = useStore(gameStore, (g) => g)
  const players = game.players.filter((p) => p.ID === client?.ID)
  const enemies = game.players.filter((p) => p.ID !== client?.ID)

  return (
    <ClientOnly>
      <PromptController />
      <main className="min-w-0 overflow-x-hidden">
        <AppHeader />
        <div className="flex min-w-0">
          <div className="min-w-0 space-y-2 flex-1 overflow-auto">
            <ActionQueue />

            <div className='flex justify-between'>
              <div className='flex flex-col'>
                <div className="px-4 left-0 z-10">
                  {players.map((player) => (
                    <PlayerThumbnails key={player.ID} player_ID={player.ID} />
                  ))}
                </div>
                <div className="px-4 flex right-0 z-10">
                  {players.map((player) => (
                    <PlayerPositions key={player.ID} flip player_ID={player.ID} />
                  ))}
                </div>
              </div>
              <div className='flex flex-col items-end'>
                <div className="px-4 left-0 z-10">
                  {enemies.map((player) => (
                    <PlayerThumbnails key={player.ID} player_ID={player.ID} />
                  ))}
                </div>
                <div className="px-4 flex right-0 z-10">
                  {enemies.map((player) => (
                    <PlayerPositions key={player.ID} flip player_ID={player.ID} />
                  ))}
                </div>
              </div>
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
          </div>
        </div>

        <pre className="max-w-full overflow-x-auto whitespace-pre-wrap wrap-break-word">
          {JSON.stringify(game, null, 2)}
        </pre>
      </main>
    </ClientOnly>
  )
}
