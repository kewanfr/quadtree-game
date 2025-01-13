package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"golang.org/x/exp/rand"
)

func generateRandomFloorContent() (floorContent [][]int) {
	for x := 0; x < configuration.Global.NumTileX; x++ {
		var line []int
		for y := 0; y < configuration.Global.NumTileY; y++ {
			line = append(line, rand.Intn(5))
		}
		floorContent = append(floorContent, line)
	}

	return floorContent
}
