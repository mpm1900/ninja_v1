import { ActorThumbnail } from '#/components/actor-thumbnail'
import { AppHeader } from '#/components/app-header'
import { BattleActions } from '#/components/battle-actions'
import { BattleContextController } from '#/components/battle-context-controller'
import { GameLogList } from '#/components/game-log'
import { PlayerPositions } from '#/components/player-positions'
import { PromptController } from '#/components/prompt-controller'
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '#/components/ui/accordion'
import { ScrollArea } from '#/components/ui/scroll-area'
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
      <main className="flex flex-col h-screen bg-zinc-800">
        <AppHeader />
        <div className="flex flex-col flex-1 relative overflow-auto">
          <div>
            <div className="fixed top-12 px-4 flex right-0 z-10">
              {game.players
                .filter((p) => p.ID !== client?.ID)
                .map((player) => (
                  <PlayerPositions key={player.ID} flip player_ID={player.ID} />
                ))}
            </div>
            <div className="fixed top-12 px-4 left-0 z-10">
              {game.players
                .filter((p) => p.ID !== client?.ID)
                .map((player) => (
                  <div key={player.ID} className="flex gap-2 py-2">
                    {game.actors
                      .filter((a) => a.player_ID === player.ID)
                      .map((a, i) => (
                        <ActorThumbnail
                          key={a.ID}
                          actor={a}
                          index={i}
                          hidden={!a.seen && !a.position_ID}
                          className={cn({
                            'opacity-40': !a.position_ID,
                          })}
                        />
                      ))}
                  </div>
                ))}
              <Accordion defaultValue={['log']} type="multiple" className='bg-black/20 px-3'>
                <AccordionItem value="log">
                  <AccordionTrigger>Log</AccordionTrigger>
                  <AccordionContent>
                    <ScrollArea className="h-40">
                      <GameLogList />
                    </ScrollArea>
                  </AccordionContent>
                </AccordionItem>
              </Accordion>
            </div>
          </div>
          <div className="flex-1 grid place-items-center overflow-hidden relative">
            {actor && <BattleActions actor={actor} />}
            {game.status === 'running' &&
              game.active_context?.source_actor_ID && (
                <div className="absolute">
                  <div>
                    {game.active_context?.source_actor_ID} uses{' '}
                    {game.active_context?.action_ID}
                  </div>
                  <div>
                    on{' '}
                    {game.active_context?.target_actor_IDs?.concat(
                      game.active_context?.target_position_IDs
                    )}
                  </div>
                </div>
              )}
          </div>
          <div className="fixed bottom-0 px-4 flex left-0 z-10">
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
                  onSelectedChange={(actor_ID) =>
                    setContextSource(actor_ID, game)
                  }
                />
              ))}
          </div>
          <div className="fixed bottom-0 px-4 flex right-0 z-10">
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
      </main>
    </>
  )
}
