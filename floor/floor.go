package floor

import (
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - content : partie du terrain qui doit être affichée à l'écran
//   - fullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")

type Floor struct {
	content         [][]int
	fullContent     [][]int
	quadtreeContent quadtree.Quadtree
	// P1X, P1Y				int
	// P2X, P2Y				int
	// PortalStarted		bool
	// PortalCreated		bool

}

func (f Floor) GetFullContent() [][]int {
	return f.fullContent
}

// types d'affichage du terrain disponibles
const (
	GridFloor     int = iota // = 0
	FromFileFloor            // = 1
	QuadTreeFloor            // 2
)

// GetHeight retourne la hauteur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetHeight() (height int) {
	return len(f.fullContent)
}

// GetWidth retourne la largeur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetWidth() (width int) {
	if len(f.fullContent) > 0 {
		width = len(f.fullContent[0])
	}
	return
}

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