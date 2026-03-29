import { ActionCard } from '#/components/action-card'
import { ActorThumbnail } from '#/components/actor-thumbnail'
import { AppHeader } from '#/components/app-header'
import { BattleActions } from '#/components/battle-actions'
import { PlayerPositions } from '#/components/player-positions'
import { PromptController } from '#/components/prompt-controller'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { createFileRoute, redirect } from '@tanstack/react-router'
import { useStore } from '@tanstack/react-store'
import { useState } from 'react'

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
  const actors = game.actors.filter(
    (a) => a.player_ID === client?.ID && !!a.position_ID
  )
  const [selected, setSelected] = useState<string>(actors[0]?.ID ?? '')
  const actor = game.actors.find((a) => a.ID === selected)

  return (
    <>
      <PromptController />
      <main className="flex flex-col h-screen">
        <AppHeader />
        <div className="flex flex-col flex-1">
          <div className="flex justify-between px-2">
            <div>
              {game.players
                .filter((p) => p.ID !== client?.ID)
                .map((player) => (
                  <div key={player.ID} className="flex gap-2 py-4">
                    {game.actors
                      .filter((a) => a.player_ID === player.ID)
                      .map((a, i) => (
                        <ActorThumbnail key={a.ID} actor={a} index={i} />
                      ))}
                  </div>
                ))}
            </div>
            <div className="flex">
              {game.players
                .filter((p) => p.ID !== client?.ID)
                .map((player) => (
                  <PlayerPositions key={player.ID} player_ID={player.ID} />
                ))}
            </div>
          </div>
          <div className="flex-1 grid place-items-center">
            {actor && game.status === 'idle' && <BattleActions actor={actor} />}
          </div>
          <div className="flex items-end justify-between px-2 z-10">
            <div>
              {game.players
                .filter((p) => p.ID === client?.ID)
                .map((player) => (
                  <PlayerPositions
                    key={player.ID}
                    player_ID={player.ID}
                    selected={game.status === 'idle' ? selected : ''}
                    onSelectedChange={setSelected}
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
                        <ActorThumbnail key={a.ID} actor={a} index={i} />
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
