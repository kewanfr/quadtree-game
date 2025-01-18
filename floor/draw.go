package floor

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
func (f Floor) Draw(screen *ebiten.Image) {
	// // Effectuer une action toutes les 2 updates (modifiable)
	// if f.updateCounter%2 == 0 {
	// 	fmt.Println("Action effectuée à update :", f.updateCounter)
	// }

	for y := range f.content {
		for x := range f.content[y] {
			if f.content[y][x] != -1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize))

				shiftX := f.content[y][x] * configuration.Global.TileSize

				if configuration.Global.ExtFloorAnimation && f.content[y][x] == 4 {
					screen.DrawImage(assets.WaterImage.SubImage(
						image.Rect(0, (f.animStep*configuration.Global.TileSize)-configuration.Global.TileSize, configuration.Global.TileSize, f.animStep*configuration.Global.TileSize),
					).(*ebiten.Image), op)
				} else {
					screen.DrawImage(assets.FloorImage.SubImage(
						image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
					).(*ebiten.Image), op)
				}

			}
		}
	}

}
