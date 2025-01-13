package floor

import (
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func (f Floor) SaveFloor() {

	var floorContent [][]int = make([][]int, configuration.Global.NumTileX)
	for i := range floorContent {
		floorContent[i] = make([]int, configuration.Global.NumTileY)
	}

	topLeftX := 0
	topLeftY := 0
	f.quadtreeContent.GetContent(topLeftX, topLeftY, floorContent)

	// fmt.Println("floorContent", floorContent)

	file, err := os.Create("../floor-files/floor")

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(floorContent); i++ {
		line := ""
		for j := 0; j < len(floorContent[i]); j++ {
			// fmt.Println(floorContent[i][j])
			line += strconv.Itoa(floorContent[i][j])
		}
		// fmt.Println(line)
		file.WriteString(line + "\n")

	}

	file.Close()
}