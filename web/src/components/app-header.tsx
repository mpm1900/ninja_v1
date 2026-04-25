import { useStore } from '@tanstack/react-store'
import { InstanceCombobox } from './instance-combobox'
import {
  connectSocket,
  sendContextMessage,
  socketStore,
} from '#/lib/stores/socket'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import {
  Check,
  Loader,
  Loader2,
  LogOut,
  Signal,
  TriangleAlert,
  Unplug,
} from 'lucide-react'
import { Tabs, TabsList, TabsTrigger } from './ui/tabs'
import { Link, useRouterState } from '@tanstack/react-router'
import { NULL_CONTEXT } from '#/lib/game/context'
import { Button } from './ui/button'
import { GiNinjaHead } from 'react-icons/gi'
import { useLogout } from '#/lib/mutations/logout'
import { useUser } from '#/lib/queries/auth'
import { Tooltip, TooltipContent, TooltipTrigger } from './ui/tooltip'
import { Badge } from './ui/badge'

function getActiveTable(pathname: string) {
  switch (pathname) {
    case '/battle':
      return 'battle'
    case '/team-builder':
      return 'team-builder'
    default:
      return 'debug'
  }
}

function AppHeader() {
  const { data: user } = useUser()
  const logout = useLogout()
  const instanceID = useStore(socketStore, (s) => s.instanceID)
  const status = useStore(socketStore, (s) => s.status)
  const client = useStore(clientsStore, (c) => c.me)
  const turn = useStore(gameStore, (g) => g.turn)
  const game_status = useStore(gameStore, (g) => g.status)
  const actions = useStore(gameStore, g => g.actions)
  const pathname = useRouterState({
    select: (state) => state.location.pathname,
  })
  const activeTab = getActiveTable(pathname)
  return (
    <header className="flex justify-between p-2">
      <div className="flex items-center gap-2">
        <Link to="/" className="pl-2">
          <GiNinjaHead />
        </Link>
        <div className="flex items-center">
          {game_status === 'running' && <Loader2 className="animate-spin" />}
          {game_status === 'idle' && <Check />}
          {game_status === 'waiting' && <Loader2 className="animate-spin" />}
        </div>
        {user && (
          <InstanceCombobox
            icon={
              <>
                {status === 'idle' && <Unplug />}
                {(status === 'connecting' || status === 'reconnecting') && (
                  <Loader className="animate-spin" />
                )}
                {status === 'open' && <Signal />}
                {(status === 'closed' || status === 'error') && (
                  <TriangleAlert className="text-destructive" />
                )}
              </>
            }
            value={instanceID}
            onValueChange={(instanceID) => connectSocket(instanceID)}
          />
        )}
        <Tabs value={activeTab}>
          <TabsList>
            <TabsTrigger value="team-builder" asChild>
              <Link to="/team-builder">Team Builder</Link>
            </TabsTrigger>
            <TabsTrigger value="debug" asChild>
              <Link to="/debug">Debug</Link>
            </TabsTrigger>
            <TabsTrigger value="battle" asChild>
              <Link to="/battle">Battle</Link>
            </TabsTrigger>
          </TabsList>
        </Tabs>
        <Badge>
          {turn.count}: {turn.phase}
        </Badge>
        {client && (
          <div className="flex gap-2">
            <Button
              disabled={actions.length == 0 || game_status === 'running'}
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
              disabled={!client || game_status === 'running'}
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
