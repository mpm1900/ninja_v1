import { useStore } from '@tanstack/react-store'
import { InstanceCombobox } from './instance-combobox'
import {
  connectSocket,
  sendContextMessage,
  socketStore,
} from '#/lib/stores/socket'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { Loader, Signal, Unplug } from 'lucide-react'
import { Tabs, TabsList, TabsTrigger } from './ui/tabs'
import { Link, useRouterState } from '@tanstack/react-router'
import { NULL_CONTEXT } from '#/lib/game/context'
import { Button } from './ui/button'

function AppHeader() {
  const instanceID = useStore(socketStore, (s) => s.instanceID)
  const status = useStore(socketStore, (s) => s.status)
  const client = useStore(clientsStore, (c) => c.me)
  const game = useStore(gameStore, (g) => g)
  const pathname = useRouterState({
    select: (state) => state.location.pathname,
  })
  const activeTab = pathname === '/battle' ? 'battle' : 'setup'
  return (
    <header className="flex justify-between p-2">
      <div className="flex items-center gap-2">
        <code className="px-4">
          {status}/{game.status}
        </code>
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
        <Tabs value={activeTab}>
          <TabsList>
            <TabsTrigger value="setup" asChild>
              <Link to="/">Setup</Link>
            </TabsTrigger>
            <TabsTrigger value="battle" asChild>
              <Link to="/battle">Battle</Link>
            </TabsTrigger>
          </TabsList>
        </Tabs>
        {client && (
          <div className="flex gap-2">
            <Button
              disabled={game.actions.length == 0 || game.status === 'running'}
              onClick={() => {
                sendContextMessage({
                  type: 'run-game-actions',
                  client_ID: client.ID,
                  context: NULL_CONTEXT,
                })
              }}
            >
              Run
            </Button>
            <Button
              disabled={!client || game.status === 'running'}
              onClick={() => {
                sendContextMessage({
                  type: 'validate-state',
                  client_ID: client.ID,
                  context: NULL_CONTEXT,
                })
              }}
            >
              Validate
            </Button>
          </div>
        )}
      </div>
      <div className="font-mono text-sm">ME: {client?.ID}</div>
    </header>
  )
}

export { AppHeader }
