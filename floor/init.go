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
	f.animStep = 1
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	var fileContent [][]int

	// Si on veut générer un terrain aléatoire, on utilise la fonction generateRandomFloorContent comme contenu du tableau
	if configuration.Global.ExtRandomFloorGeneration {
		fileContent = generateRandomFloorContent()
		f.fullContent = fileContent
	} else {
		fileContent = readFloorFromFile(configuration.Global.FloorFile)
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = fileContent
	case QuadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(fileContent)
	default:
		panic("Type de sol invalide")
	}
}

/*
*
readFloorFromFile lit le contenu d'un fichier de terrain et en génère un tableau en 2 dimensions représentant le terrain.

Entrées:
- fileName: le nom du fichier à lire (dans le dossier floor-files)

Sorties:
  - floorContent: tableau en 2 dimensions de int représentant les cases du terrain (les valeurs sont entre -1 et 4 pour le terrain de l'exemple)
    -1: en dehors du terrain, la case est pas accessible
    Au dessus de 0: un type de sol

*
*/
func readFloorFromFile(fileName string) (floorContent [][]int) {
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
