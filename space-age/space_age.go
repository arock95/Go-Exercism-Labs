// Package space provides a function to calculate person's age in years on different
// planets, given their age in seconds
package space

type Planet string

// mapPlanetSeconds creates a map tying planets and their year (in seconds)
func mapPlanetSeconds() map[Planet]float64 {
	planetYears := make(map[Planet]float64)

	planetYears["Earth"] = 31557600
	planetYears["Mercury"] = 31557600 * 0.2408467
	planetYears["Venus"] = 31557600 * 0.61519726
	planetYears["Mars"] = 31557600 * 1.8808158
	planetYears["Jupiter"] = 31557600 * 11.862615
	planetYears["Saturn"] = 31557600 * 29.447498
	planetYears["Uranus"] = 31557600 * 84.016846
	planetYears["Neptune"] = 31557600 * 164.79132

	return planetYears
}

// Age calculates person's age in years on different planets, given their age in seconds
func Age(seconds float64, planet Planet) float64 {
	return seconds / mapPlanetSeconds()[planet]
}
