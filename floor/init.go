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
	var err error // Stocke l'erreur potentielle

	floorContent = [][]int{}

	// Ouvre le fichier
	floorFile, err := os.Open(fileName)
	
	if err != nil {
		return nil
	}

	var scanner *bufio.Scanner  = bufio.NewScanner(floorFile)

	// On scanne le fichier ligne par ligne
	for scanner.Scan(){
		line := scanner.Text()
		var lineArr []int
		
		// On itère sur chaque élément de la ligne
		for _, r := range line {
			elInt, err :=  strconv.Atoi(string(r))
			if err != nil {
				return floorContent
			}
			// On l'ajoute à la slice
			lineArr = append(lineArr, elInt)
		}

		// On ajoute la slice au tableau
		floorContent = append(floorContent, lineArr)

	}

	err = floorFile.Close()
	if err != nil {
		return floorContent
	}

	return floorContent
}
