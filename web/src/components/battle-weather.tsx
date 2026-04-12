import { gameStore } from "#/lib/stores/game"
import { useStore } from "@tanstack/react-store"

function BattleWeather() {
  const game = useStore(gameStore, g => g)
  return (
    <div>
      <div>Weather: {game.state.weather}</div>
      <div>Terrain: {game.state.terrain}</div>
      <div>
        Modifiers:{' '}
        {game.modifiers
          .filter((tx) =>
            game.applied_game_state_tx.includes(tx.ID)
          )
          .map((tx) => tx.mutation.name)
          .join(',') || 'none'}
      </div>
    </div>
  )
}


export { BattleWeather }
