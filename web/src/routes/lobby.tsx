import { useStore } from '@tanstack/react-store'
import {
  ClientOnly,
  createFileRoute,
  Link,
  redirect,
} from '@tanstack/react-router'
import { gameStore } from '#/lib/stores/game'
import { clientsStore } from '#/lib/stores/clients'
import { PromptController } from '#/components/prompt-controller'
import { AppHeader } from '#/components/app-header'
import { PlayerThumbnails } from '#/components/player-thumbnails'
import {
  Card,
  CardAction,
  CardContent,
  CardHeader,
  CardTitle,
} from '#/components/ui/card'
import { sendContextMessage } from '#/lib/stores/socket'
import { NULL_CONTEXT } from '#/lib/game/context'
import { Button } from '#/components/ui/button'
import { useEffect } from 'react'

export const Route = createFileRoute('/lobby')({
  beforeLoad: ({ context }) => {
    if (!context.auth.user) {
      throw redirect({ to: '/login' })
    }
  },
  component: App,
})

function App() {
  const client = useStore(clientsStore, (c) => c.me)
  const game = useStore(gameStore, (g) => g)
  const players = game.players.filter((p) => p.ID === client?.ID)
  const enemies = game.players.filter((p) => p.ID !== client?.ID)
  const ready = players.length > 0 && enemies.length > 0
  const unstarted = game.status !== 'running' && game.turn.count == 0
  const nav = Route.useNavigate()

  useEffect(() => {
    if (game.status === 'running') {
      nav({ to: '/battle' })
    }
  }, [game.status])

  return (
    <ClientOnly>
      <PromptController />
      <main className="min-w-0 overflow-x-hidden">
        <AppHeader />
        <div className="flex min-w-0">
          <div className="min-w-0 space-y-2 flex-1 overflow-auto">
            <Card className="m-6">
              <CardHeader>
                <CardTitle>Lobby</CardTitle>
                <CardAction>
                  {client && ready && (
                    unstarted ?
                      <Button
                        asChild
                        onClick={() => {
                          sendContextMessage({
                            type: 'validate-state',
                            client_ID: client!.ID,
                            context: NULL_CONTEXT,
                          })
                        }}
                      >
                        <Link to="/battle">Start Battle</Link>
                      </Button> : <Button
                        onClick={() => {
                          sendContextMessage({
                            type: 'reset',
                            client_ID: client!.ID,
                            context: NULL_CONTEXT,
                          })
                        }}
                      >
                        Reset
                      </Button>

                  )}
                </CardAction>
              </CardHeader>
              <CardContent>
                <div className="flex justify-between">
                  <div className="flex flex-col">
                    <div className="left-0 z-10">
                      {players.map((player) => (
                        <PlayerThumbnails
                          key={player.ID}
                          player_ID={player.ID}
                        />
                      ))}
                    </div>
                  </div>
                  <div className="flex flex-col items-end">
                    <div className="left-0 z-10">
                      {enemies.map((player) => (
                        <PlayerThumbnails
                          key={player.ID}
                          player_ID={player.ID}
                        />
                      ))}
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </main>
    </ClientOnly>
  )
}
