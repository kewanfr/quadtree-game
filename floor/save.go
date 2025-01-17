package floor

import (
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

/**
SaveFloor est exécuté dans game/update.go, lorsque l'on appuie sur la touche F5
Il sauvegarde le contenu du quadtree ou du tableau content dans un fichier texte
**/
func (f Floor) SaveFloor() {

	var floorContent [][]int = make([][]int, configuration.Global.NumTileX)
	for i := range floorContent {
		floorContent[i] = make([]int, configuration.Global.NumTileY)
	}

	topLeftX := 0
	topLeftY := 0

	if configuration.Global.FloorKind == QuadTreeFloor {
		f.quadtreeContent.GetContent(topLeftX, topLeftY, floorContent)
	} else {
		floorContent = f.fullContent
	}

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