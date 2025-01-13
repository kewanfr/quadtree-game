package floor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	switch configuration.Global.FloorKind {
	case GridFloor:
		f.updateGridFloor(topLeftX, topLeftY)
	case FromFileFloor:
		f.updateFromFileFloor(topLeftX, topLeftY)
	case QuadTreeFloor:
		f.updateQuadtreeFloor(topLeftX, topLeftY)
	}

	if configuration.Global.ExtZoom {
		if ebiten.IsKeyPressed(ebiten.KeyEqual) && ((configuration.Global.NumTileX < configuration.Global.MaxZoom) || (configuration.Global.NumTileY < configuration.Global.MaxZoom)) {
			configuration.Global.NumTileX += 1
			configuration.Global.NumTileY += 1

			configuration.SetComputedFields()

			f.content = make([][]int, configuration.Global.NumTileY)
			for y := 0; y < len(f.content); y++ {
				f.content[y] = make([]int, configuration.Global.NumTileX)
			}
		}

		//hard limit à 6 pour éviter tout problème avec blocking.go
		if ebiten.IsKeyPressed(ebiten.KeyMinus) && (!(configuration.Global.NumTileX <= configuration.Global.MinZoom) || !(configuration.Global.NumTileY <= configuration.Global.MinZoom)) {
			configuration.Global.NumTileX -= 1
			configuration.Global.NumTileY -= 1

			configuration.SetComputedFields() // update l'ui

			f.content = make([][]int, configuration.Global.NumTileY)
			for y := 0; y < len(f.content); y++ {
				f.content[y] = make([]int, configuration.Global.NumTileX)
			}
		}

	}

}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(topLeftX, topLeftY int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absX := topLeftX
			if absX < 0 {
				absX = -absX
			}
			absY := topLeftY
			if absY < 0 {
				absY = -absY
			}
			f.content[y][x] = ((x + absX%2) + (y + absY%2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
func (f *Floor) updateFromFileFloor(topLeftX, topLeftY int) {
	for y := 0; y < configuration.Global.NumTileY; y++ {
		for x := 0; x < configuration.Global.NumTileX; x++ {
			// Calculer les coordonnées dans fullContent
			var fullContentX int = topLeftX + x
			var fullContentY int = topLeftY + y

			// Vérifier si les coordonnées sont dans les limites de fullContent (évite que le programme crash avec des sizes de tile trop grandes)
			if fullContentY >= 0 && fullContentY < len(f.fullContent) &&
				fullContentX >= 0 && fullContentX < len(f.fullContent[fullContentY]) {
				f.content[y][x] = f.fullContent[fullContentY][fullContentX]
			} else {
				f.content[y][x] = -1 // indique le manque de sol, bloque le bonhomme
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(topLeftX, topLeftY int) {
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
