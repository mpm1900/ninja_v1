import type { Actor } from '#/lib/game/actor'
import type { Game } from '#/lib/game/game'
import { sendContextMessage } from '#/lib/stores/socket'
import { useStore } from '@tanstack/react-store'
import { Button } from './ui/button'
import { ButtonGroup } from './ui/button-group'
import { clientsStore } from '#/lib/stores/clients'

function PositionSelect({ actor, game }: { actor: Actor; game: Game }) {
  const player = game.players.find((p) => p.ID == actor.player_ID)
  const client = useStore(clientsStore, (c) => c.me!)

  const capacity = player?.positions_capacity ?? 0
  const options = Array.from({ length: capacity })
  const positionIndex = player?.positions
    .map((p) => p.ID)
    .indexOf(actor.state.position_ID)

  return (
    <div className="flex items-center justify-end gap-2">
      {player ? (
        <>
          Position Index:
          <ButtonGroup>
            {options.map((_, i) => (
              <Button
                key={i}
                size="icon"
                disabled={game.status == 'running'}
                variant={i === positionIndex ? 'default' : 'outline'}
                onClick={() => {
                  sendContextMessage({
                    type: 'set-actor-position',
                    client_ID: client.ID,
                    position_index: i,
                    context: {
                      parent_actor_ID: null,
                      source_actor_ID: null,
                      source_player_ID: client.ID,
                      target_actor_IDs: [actor.ID],
                      target_position_IDs: [],
                    },
                  })
                }}
              >
                {i + 1}
              </Button>
            ))}
          </ButtonGroup>
        </>
      ) : (
        <>Player Not Found</>
      )}
    </div>
  )
}

export { PositionSelect }
