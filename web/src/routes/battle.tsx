import { ActorThumbnail } from '#/components/actor-thumbnail'
import { AppHeader } from '#/components/app-header'
import { BattleActions } from '#/components/battle-actions'
import { BattleContextController } from '#/components/battle-context-controller'
import { PlayerPositions } from '#/components/player-positions'
import { PromptController } from '#/components/prompt-controller'
import { battleContext, setContextSource } from '#/lib/stores/battle-context'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { cn } from '#/lib/utils'
import { createFileRoute, redirect } from '@tanstack/react-router'
import { useStore } from '@tanstack/react-store'

export const Route = createFileRoute('/battle')({
  beforeLoad: ({ context }) => {
    if (!context.auth.user) {
      throw redirect({ to: '/login' })
    }
  },
  component: RouteComponent,
})

function RouteComponent() {
  const game = useStore(gameStore, (g) => g)
  const client = useStore(clientsStore, (c) => c.me)
  const context = useStore(battleContext, (c) => c)
  const actor = game.actors.find((a) => a.ID === context.source_actor_ID)

  return (
    <>
      <PromptController />
      <BattleContextController />
      <main className="flex flex-col h-screen">
        <AppHeader />
        <div className="flex flex-col flex-1 relative overflow-auto">
          <div className="fixed top-10 w-full flex px-4 justify-between z-10">
            <div>
              {game.players
                .filter((p) => p.ID !== client?.ID)
                .map((player) => (
                  <div key={player.ID} className="flex gap-2 py-4">
                    {game.actors
                      .filter((a) => a.player_ID === player.ID)
                      .map((a, i) => (
                        <ActorThumbnail
                          key={a.ID}
                          actor={a}
                          index={i}
                          className={cn({
                            'opacity-70': !a.position_ID,
                          })}
                        />
                      ))}
                  </div>
                ))}
            </div>
            <div className="flex">
              {game.players
                .filter((p) => p.ID !== client?.ID)
                .map((player) => (
                  <PlayerPositions
                    key={player.ID}
                    flip={true}
                    player_ID={player.ID}
                  />
                ))}
            </div>
          </div>
          <div className="flex-1 grid place-items-center overflow-hidden">
            {actor && game.status === 'idle' && <BattleActions actor={actor} />}
          </div>
          <div className="fixed bottom-0 w-full px-4 flex items-end justify-between z-10">
            <div>
              {game.players
                .filter((p) => p.ID === client?.ID)
                .map((player) => (
                  <PlayerPositions
                    key={player.ID}
                    flip={false}
                    player_ID={player.ID}
                    selected={
                      game.status === 'idle'
                        ? (context.source_actor_ID ?? '')
                        : ''
                    }
                    onSelectedChange={(actor_ID) => setContextSource(actor_ID)}
                  />
                ))}
            </div>
            <div>
              {game.players
                .filter((p) => p.ID === client?.ID)
                .map((player) => (
                  <div key={player.ID} className="flex gap-2 py-2">
                    {game.actors
                      .filter((a) => a.player_ID === player.ID)
                      .map((a, i) => (
                        <ActorThumbnail
                          key={a.ID}
                          actor={a}
                          index={i}
                          className={cn({
                            'opacity-40': !a.position_ID,
                          })}
                        />
                      ))}
                  </div>
                ))}
            </div>
          </div>
        </div>
      </main>
    </>
  )
}
