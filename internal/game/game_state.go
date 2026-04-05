package game

type GameWeather string

const (
	GameWeatherNone     GameWeather = "none"
	GameWeatherRain     GameWeather = "rain"
	GameWeatherSunlight GameWeather = "sunlight"
)

/**
 * [GameState]
 * This object exists to bypass a complicated system of
 * checks to see if more than one modifier of a given
 * type can exist at once.
 *
 * Instead, just modifity this object so only one value can exist
 * examples can be weather, terrain,
 *
 * the design space does overlap a bit with modifiers if not for the singular nature
 */
type GameState struct {
	weather GameWeather
}
