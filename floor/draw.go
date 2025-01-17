package floor

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
func (f Floor) Draw(screen *ebiten.Image, imageClock int) {

	f.updateCounter++ // Incrémenter le compteur à chaque update

	// // Effectuer une action toutes les 2 updates (modifiable)
	// if f.updateCounter%2 == 0 {
	// 	fmt.Println("Action effectuée à update :", f.updateCounter)
	// }

	// Réinitialiser le compteur après 60 updates (optionnel, utile pour garder un cycle simple)
	if f.updateCounter >= 60 {
		f.updateCounter = 0
	}

	for y := range f.content {
		for x := range f.content[y] {
			if f.content[y][x] != -1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize))

				shiftX := f.content[y][x] * configuration.Global.TileSize

				if configuration.Global.ExtWaterAnimation && f.content[y][x] == 4 {

					// shiftY depending on x, y position and actual counter
					// for a counter, an x and a y, we can have a unique shiftY

					// shiftY := (f.updateCounter + x + y) % 16 * 16

				
					screen.DrawImage(assets.WaterImage.SubImage(
						image.Rect(0, imageClock, 16, imageClock+16),
					).(*ebiten.Image), op)
				}else {
					screen.DrawImage(assets.FloorImage.SubImage(
						image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
					).(*ebiten.Image), op)
				}

			}
		}
	}

}
