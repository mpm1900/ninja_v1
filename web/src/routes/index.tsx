import { ActorsTable } from '#/components/actors-table'
import { InstanceCombobox } from '#/components/instance-combobox'
import { actorsQuery } from '#/lib/queries/actors'
import { instancesQuery } from '#/lib/queries/instances'
import {
  connectSocket,
  sendContextMessage,
  socketStore,
} from '#/lib/stores/socket'
import { useStore } from '@tanstack/react-store'
import { useSuspenseQuery } from '@tanstack/react-query'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { Loader, Signal, Unplug } from 'lucide-react'
import { gameStore } from '#/lib/stores/game'
import { clientsStore } from '#/lib/stores/clients'
import { ActorCard } from '#/components/actor-card'
import { modifiersQuery } from '#/lib/queries/modifiers'
import { ModifiersTable } from '#/components/modifiers-table'
import { actionsQuery } from '#/lib/queries/actions'
import { ActorControl } from '#/components/actor-control'
import { triggersQuery } from '#/lib/queries/trigger-types'
import { Button } from '#/components/ui/button'
import { ActionQueue } from '#/components/action-queue'
import { PromptController } from '#/components/prompt-controller'

export const Route = createFileRoute('/')({
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
  const triggers = useSuspenseQuery(triggersQuery)
  const instanceID = useStore(socketStore, (s) => s.instanceID)
  const status = useStore(socketStore, (s) => s.status)
  const client = useStore(clientsStore, (c) => c.me)
  const game = useStore(gameStore, (g) => g)

  return (
    <ClientOnly>
      <PromptController />
      <main className="">
        <header className="flex justify-between p-2">
          <div>
            <code className="px-4">{status}</code>
            <InstanceCombobox
              icon={
                <>
                  {status === 'idle' && <Unplug />}
                  {status === 'connecting' && <Loader />}
                  {status === 'open' && <Signal />}
                </>
              }
              value={instanceID}
              onValueChange={connectSocket}
            />
          </div>
          <div>ME: {client?.ID}</div>
        </header>
        <div className="flex">
          <div className="space-y-2 flex-1 overflow-auto">
            <ActionQueue />

            <div className="grid grid-cols-3 gap-2 p-2">
              <ActorCard
                actor={game.actors[0]}
                clientID={client?.ID}
                game={game}
              />
              <ActorCard
                actor={game.actors[1]}
                clientID={client?.ID}
                game={game}
              />
              <ActorCard
                actor={game.actors[2]}
                clientID={client?.ID}
                game={game}
              />
              <ActorCard
                actor={game.actors[3]}
                clientID={client?.ID}
                game={game}
              />
              <ActorCard
                actor={game.actors[4]}
                clientID={client?.ID}
                game={game}
              />
              <ActorCard
                actor={game.actors[5]}
                clientID={client?.ID}
                game={game}
              />
            </div>
            <ActorsTable
              data={actors.data}
              enabled={status === 'open' && game.status == 'idle'}
              rowSelection={Object.fromEntries(
                game.actors.map((a) => [a.actor_ID, true])
              )}
              onRowCheckedChange={(a, checked) => {
                if (!client) return
                const actor = game.actors.find(
                  (ga) => ga.actor_ID === a.actor_ID
                )
                const ID = checked ? a.actor_ID : actor!.ID

                sendContextMessage({
                  type: checked ? 'add-actor' : 'remove-actor',
                  client_ID: client.ID,
                  context: {
                    source_player_ID: client.ID,
                    source_actor_ID: ID,
                    parent_actor_ID: ID,
                    target_actor_IDs: [],
                    target_position_IDs: [],
                  },
                })
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
          <div className="w-sm">
            {game.log?.map((log, i) => (
              <div key={i}>{log}</div>
            ))}
          </div>
        </div>

        <pre>{JSON.stringify(game, null, 2)}</pre>
      </main>
    </ClientOnly>
  )
}
