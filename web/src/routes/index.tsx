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
  const game = useStore(gameStore, (g) => g)
  const [rows, setRows] = useState<RowSelectionState>({})

  return (
    <main className="">
      <header className="flex justify-between p-2">
        <div></div>
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
      </header>
      <ActorsTable
        data={actors.data}
        rowSelection={rows}
        onRowSelectionChange={(r) => {
          const sourceActorID = Object.entries(r).find((r) => r[1])?.[0] ?? ''
          console.log(sourceActorID, r)
          sendContextMessage('add-actor', {
            sourcePlayerID: '',
            sourceActorID,
            parentActorID: sourceActorID,
            targetActorIDs: [],
            targetPositionIDs: [],
          })
          setRows(r)
        }}
      />
      <pre>{JSON.stringify(game, null, 2)}</pre>
    </main>
  )
}
