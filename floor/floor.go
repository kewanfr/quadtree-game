package floor

import (
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
	animCounter int // compteur de frame pour l'animation du sol
	animStep    int // pas de l'animation
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

// GetHeightFullContent retourne la hauteur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetHeightFullContent() (height int) {
	return len(f.fullContent)
}

// GetWidthFullContent retourne la largeur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetWidthFullContent() (width int) {
	if len(f.fullContent) > 0 {
		width = len(f.fullContent[0])
	}
	return
}

// GetBlockFullContent retourne la largeur (en cases) du terrain
// à partir du tableau content, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetHeightContent() (height int) {
	return len(f.content)
}

// GetWidthContent retourne la largeur (en cases) du terrain
// à partir du tableau content, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetWidthContent() (width int) {
	if len(f.content) > 0 {
		width = len(f.content[0])
	}
	return
}

func (f Floor) GetBlockContent(x, y int) int {
	return f.content[x][y]
}

func (f Floor) GetContent() (content [][]int) {
	return f.content
}
