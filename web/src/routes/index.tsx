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
import { createFileRoute } from '@tanstack/react-router'
import { Loader, Signal, Unplug } from 'lucide-react'
import { gameStore } from '#/lib/stores/game'
import { clientsStore, type Client } from '#/lib/stores/clients'
import { ActorCard } from '#/components/actor-card'
import { modifiersQuery } from '#/lib/queries/modifiers'
import { Button } from '#/components/ui/button'
import { Card, CardContent } from '#/components/ui/card'
import { ModifiersTable } from '#/components/modifiers-table'

export const Route = createFileRoute('/')({
  component: App,
  loader: async ({ context }) => {
    await context.queryClient.ensureQueryData(actorsQuery)
    await context.queryClient.ensureQueryData(modifiersQuery)
    await context.queryClient.ensureQueryData(instancesQuery)
  },
})

function App() {
  const actors = useSuspenseQuery(actorsQuery)
  const modifiers = useSuspenseQuery(modifiersQuery)
  const instanceID = useStore(socketStore, (s) => s.instanceID)
  const status = useStore(socketStore, (s) => s.status)
  const client = useStore(clientsStore, (c) => c[0] as Client | undefined)
  const game = useStore(gameStore, (g) => g)

  return (
    <main className="">
      <header className="flex justify-between p-2">
        <div></div>
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
      </header>
      <div className="space-y-6">
        <div className="grid grid-cols-3 gap-2 p-2">
          <ActorCard actor={game.actors[0]} game={game} />
          <ActorCard actor={game.actors[1]} game={game} />
          <ActorCard actor={game.actors[2]} game={game} />
          <ActorCard actor={game.actors[3]} game={game} />
          <ActorCard actor={game.actors[4]} game={game} />
          <ActorCard actor={game.actors[5]} game={game} />
        </div>
        <ActorsTable
          data={actors.data}
          enabled={status === 'open'}
          rowSelection={Object.fromEntries(
            game.actors.map((a) => [a.actor_ID, true])
          )}
          onRowCheckedChange={(a, checked) => {
            if (!client) return
            const actor = game.actors.find((ga) => ga.actor_ID === a.actor_ID)
            const ID = checked ? a.actor_ID : actor!.ID

            sendContextMessage({
              type: checked ? 'add-actor' : 'remove-actor',
              clientID: client.ID,
              context: {
                source_player_ID: client.ID,
                source_actor_ID: ID,
                parent_actor_ID: ID,
                target_actor_IDs: [],
                target_position_IDs: [],
              },
            })
          }}
          subRow={({ row }) => (
            <Card className="rounded-t-none border-t-0 mx-2 mb-2">
              <CardContent>
                {modifiers.data.map((m) => (
                  <Button
                    key={m.ID}
                    onClick={() => {
                      if (!client) return
                      const actor = game.actors.find(
                        (a) => a.actor_ID === row.original.actor_ID
                      )
                      if (!actor) return

                      sendContextMessage({
                        type: 'add-modifier',
                        clientID: client.ID,
                        modifierID: m.ID,
                        context: {
                          source_player_ID: client.ID,
                          source_actor_ID: actor.ID,
                          parent_actor_ID: actor.ID,
                          target_actor_IDs: [],
                          target_position_IDs: [],
                        },
                      })
                    }}
                  >
                    {m.name}
                  </Button>
                ))}
              </CardContent>
            </Card>
          )}
        />
        <ModifiersTable
          data={game.modifiers}
          onRowRemove={(modifier) => {
            if (!client) return

            sendContextMessage({
              type: 'remove-modifier',
              clientID: client.ID,
              modifierID: modifier.ID,
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

        <pre>{JSON.stringify(game, null, 2)}</pre>
      </div>
    </main>
  )
}
