import type { Actor } from '#/lib/game/actor'
import type { Game } from '#/lib/game/game'
import { Button } from './ui/button'
import { ButtonGroup } from './ui/button-group'

function PositionSelect({ actor, game }: { actor: Actor; game: Game }) {
  const player = game.players.find((p) => p.ID == actor.player_ID)

  const capacity = player?.positions_capacity ?? 0
  const options = Array.from({ length: capacity })
  const positionIndex = Object.keys(player?.positions ?? {}).indexOf(
    actor.state.position_ID
  )

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
                variant={i === positionIndex ? 'default' : 'outline'}
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
