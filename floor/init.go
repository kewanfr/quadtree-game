package floor

import (
	"bufio"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case QuadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	// TODO
	file, _ := os.Open("../floor-files/"+fileName)
	

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		var lineArr []int
		
		for _, r := range line {
			fInt, _ :=  strconv.Atoi(string(r))
			lineArr = append(lineArr, fInt)
		}

		floorContent = append(floorContent, lineArr)

	}

	return floorContent
}
