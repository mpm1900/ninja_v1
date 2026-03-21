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
import { useState } from 'react'
import type { RowSelectionState } from '@tanstack/react-table'
import { clientsStore, type Client } from '#/lib/stores/clients'
import { ActorCard } from '#/components/actor-card'

export const Route = createFileRoute('/')({
  component: App,
  loader: async ({ context }) => {
    await context.queryClient.ensureQueryData(actorsQuery)
    await context.queryClient.ensureQueryData(instancesQuery)
  },
})

function App() {
  const actors = useSuspenseQuery(actorsQuery)
  const instanceID = useStore(socketStore, (s) => s.instanceID)
  const status = useStore(socketStore, (s) => s.status)
  const client = useStore(clientsStore, (c) => c[0] as Client | undefined)
  const game = useStore(gameStore, (g) => g)
  const [rows, setRows] = useState<RowSelectionState>({})

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
      <div>
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
          rowSelection={rows}
          onRowSelectionChange={(r) => {
            if (!client) return

            const actorIDs = Object.entries(r)
              .filter((r) => r[1])
              .map((r) => r[0])

            sendContextMessage({
              type: 'set-actors',
              clientID: client!.ID,
              context: {
                sourcePlayerID: null,
                sourceActorID: null,
                parentActorID: null,
                targetActorIDs: actorIDs,
                targetPositionIDs: [],
              },
            })

            setRows(r)
          }}
        />
        <pre>{JSON.stringify(game, null, 2)}</pre>
      </div>
    </main>
  )
}
