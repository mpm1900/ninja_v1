import { gameStore } from "#/lib/stores/game"
import { useStore } from "@tanstack/react-store"

function BattleWeather() {
  const state = useStore(gameStore, g => g.state)
  const modifiers = useStore(gameStore, g => g.modifiers
    .filter((tx) =>
      g.applied_game_state_tx.includes(tx.ID)
    ))
  return (
    <div>
      <div>Weather: {state.weather}</div>
      <div>Terrain: {state.terrain}</div>
      <div>
        Modifiers:{' '}
        {modifiers
          .map((tx) => tx.mutation.name)
          .join(',') || 'none'}
      </div>
    </div>
  )
}


export { BattleWeather }
