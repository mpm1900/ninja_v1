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
  CardDescription,
  CardHeader,
  CardTitle,
} from '#/components/ui/card'
import { sendContextMessage } from '#/lib/stores/socket'
import { NULL_CONTEXT } from '#/lib/game/context'
import { Button } from '#/components/ui/button'
import { useEffect, useState } from 'react'
import { LobbyThumbnails } from '#/components/lobby-thumbnails'
import { LobbyActorsList } from '#/components/lobby-actors-list'
import { Swords } from 'lucide-react'

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
  const ready = players.length > 0
  const unstarted = game.status !== 'running' && game.turn.count == 0
  const nav = Route.useNavigate()
  const pids = players.map((p) => p.ID)
  const default_enabled = game.actors.filter(
    (a) => pids.includes(a.player_ID) && a.enabled
  )
  const [enabled, setEnabled] = useState<string[]>(
    default_enabled.map((a) => a.ID)
  )

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
            <Card className="m-6 bg-stone-950 py-4">
              <CardHeader className='px-4'>
                <CardTitle>Pre-Game Lobby</CardTitle>
                <CardDescription>Select 4-shinobi your line-up</CardDescription>
                <CardAction>
                  {client &&
                    ready &&
                    (unstarted ? (
                      <Link to="/battle"><Button
                        disabled={
                          players.some((p) => !p.ready) ||
                          enemies.some((e) => !e.ready)
                        }
                        onClick={() => {
                          sendContextMessage({
                            type: 'start-battle',
                            client_ID: client!.ID,
                            context: NULL_CONTEXT,
                          })
                        }}
                      >
                        <Swords /> Start Battle
                      </Button>
                      </Link>
                    ) : (
                      <Button
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
                    ))}
                </CardAction>
              </CardHeader>
              <CardContent>
                <div className="flex justify-between">
                  {players.map((player) => (
                    <div key={player.ID} className="flex flex-col gap-2">
                      <h3 className="font-bold text-xl">
                        {player.user.email} (You)
                      </h3>
                      <LobbyThumbnails
                        player_ID={player.ID}
                        enabled={enabled}
                        onEnabledChange={setEnabled}
                      />
                      {game.turn.phase === 'init' &&
                        !player.ready && (
                          <Button
                            disabled={enabled.length !== 4}
                            onClick={() => {
                              sendContextMessage({
                                type: 'ready-team',
                                client_ID: client!.ID,
                                context: {
                                  ...NULL_CONTEXT,
                                  target_actor_IDs: enabled,
                                },
                              })
                            }}
                          >
                            Ready Team ({enabled.length}/4)
                          </Button>
                        )}
                      {game.turn.phase === 'init' && player.ready && (
                        <Button
                          variant="destructive"
                          onClick={() => {
                            sendContextMessage({
                              type: 'cancel-team',
                              client_ID: client!.ID,
                              context: NULL_CONTEXT,
                            })
                          }}
                        >
                          Cancel
                        </Button>
                      )}

                      <LobbyActorsList
                        player={player}
                        enabled={enabled}
                        onEnabledChange={setEnabled}
                      />
                    </div>
                  ))}
                  <div className="flex flex-col items-end">
                    {enemies.map((player) => (
                      <div key={player.ID}>
                        <h3 className="font-bold text-xl">
                          {player.user.email}
                        </h3>
                        <PlayerThumbnails player_ID={player.ID} />
                      </div>
                    ))}
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
