import { useStore } from '@tanstack/react-store'
import { InstanceCombobox } from './instance-combobox'
import {
  connectSocket,
  sendContextMessage,
  socketStore,
} from '#/lib/stores/socket'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { Check, Loader, Loader2, LogOut, Signal, TriangleAlert, Unplug } from 'lucide-react'
import { Tabs, TabsList, TabsTrigger } from './ui/tabs'
import { Link, useRouterState } from '@tanstack/react-router'
import { NULL_CONTEXT } from '#/lib/game/context'
import { Button } from './ui/button'
import { GiNinjaHead } from 'react-icons/gi'
import { useLogout } from '#/lib/mutations/logout'
import { useUser } from '#/lib/queries/auth'
import { Tooltip, TooltipContent, TooltipTrigger } from './ui/tooltip'

function AppHeader() {
  const { data: user } = useUser()
  const logout = useLogout()
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
        <Link to="/" className='pl-2'>
          <GiNinjaHead />
        </Link>
        <div className="flex items-center">
          {game.status === 'running' && <Loader2 className="animate-spin" />}
          {game.status === 'idle' && <Check />}
          {game.status === 'waiting' && <Loader2 className="animate-spin" />}
        </div>
        {user && (
          <InstanceCombobox
            icon={
              <>
                {status === 'idle' && <Unplug />}
                {status === 'connecting' && <Loader />}
                {status === 'open' && <Signal />}
                {status === 'closed' && <TriangleAlert />}
              </>
            }
            value={instanceID}
            onValueChange={(instanceID) => connectSocket(instanceID)}
          />
        )}
        <Tabs value={activeTab}>
          <TabsList>
            <TabsTrigger value="setup" asChild>
              <Link to="/setup">Setup</Link>
            </TabsTrigger>
            <TabsTrigger value="battle" asChild>
              <Link to="/battle">Battle</Link>
            </TabsTrigger>
          </TabsList>
        </Tabs>
        {game.turn.count}
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
      <div className="flex items-center gap-4 px-2">
        <div className="font-mono text-sm flex items-center">
          {user && (
            <div className="flex items-center gap-2">
              <Tooltip>
                <TooltipTrigger>
                  <span>{user.email}</span>
                </TooltipTrigger>
                <TooltipContent>{user.id}</TooltipContent>
              </Tooltip>
              <Button
                variant="ghost"
                size="icon"
                className="size-8"
                onClick={() => logout.mutate()}
                title="Logout"
              >
                <LogOut className="size-4" />
              </Button>
            </div>
          )}
        </div>
      </div>
    </header>
  )
}

export { AppHeader }
