package floor

import (
	"bufio"
	"math/rand"
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

	var fileContent [][]int
	if configuration.Global.ExtRandomFloorGeneration {
		fileContent = generateRandomFloorContent()
	} else {
		fileContent = readFloorFromFile(configuration.Global.FloorFile)
	}


	if configuration.Global.ExtSmoothTerrain {
		fileContent = SmoothTerrain(fileContent)
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = fileContent
	case QuadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(fileContent)
	}
}


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


// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
/**
readFloorFromFile lit le contenu d'un fichier de terrain et en génère un tableau en 2 dimensions représentant le terrain.

Entrées:
- fileName: le nom du fichier à lire (dans le dossier floor-files)

Sorties:
- floorContent: tableau en 2 dimensions de int représentant les cases du terrain (les valeurs sont entre -1 et 4 pour le terrain de l'exemple)
	-1: en dehors du terrain, la case est pas accessible
	Au dessus de 0: un type de sol

**/
func readFloorFromFile(fileName string) (floorContent [][]int) {
	// TODO
	var err error // Stocke l'erreur potentielle

	// Le tableau pour stocker le terrain
	floorContent = [][]int{}

	// Ouvre le fichier
	floorFile, err := os.Open(fileName)

	// Si erreur (ex: Le fichier existe pas ou autre)
	if err != nil {
		return nil
	}

	var scanner *bufio.Scanner = bufio.NewScanner(floorFile)

	// On scanne le fichier ligne par ligne
	for scanner.Scan() {
		line := scanner.Text()
		var lineArr []int

		// On itère sur chaque élément de la ligne
		for _, r := range line {
			elInt, err := strconv.Atoi(string(r))

			// Si une erreur de conversion (ex: le contenu est pas un entier)
			if err != nil {
				return floorContent
			}

			// On l'ajoute à la slice
			lineArr = append(lineArr, elInt)
		}

		// On ajoute la slice au slice global
		floorContent = append(floorContent, lineArr)

	}

	// Ferme le fichier
	err = floorFile.Close()
	if err != nil {
		return floorContent
	}

	return floorContent
}
