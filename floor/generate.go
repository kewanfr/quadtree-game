package floor

import (
	"math/rand"
	"time"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

/**
generateRandomFloorContent génère un terrain aléatoire en générant un tableau de tableau d'entiers
**/
func generateRandomFloorContent() (floorContent [][]int) {
	for x := 0; x < configuration.Global.NumTileX; x++ {
		var line []int
		for y := 0; y < configuration.Global.NumTileY; y++ {
			rand.Seed(time.Now().UnixNano())
			line = append(line, rand.Intn(5))
		}
		floorContent = append(floorContent, line)
	}

	return floorContent
}
