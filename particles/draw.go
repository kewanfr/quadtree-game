package particles

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

/*
*
Draw dessine les particules sur l'écran
*
*/
func (p Particle) Draw(screen *ebiten.Image, camX, camY int) {
	// Calcule la position de l'animation sur le terrain
	xTileForDisplay := p.X - camX + configuration.Global.ScreenCenterTileX
	yTileForDisplay := p.Y - camY + configuration.Global.ScreenCenterTileY

	xPos := (xTileForDisplay) * configuration.Global.TileSize
	yPos := (yTileForDisplay) * configuration.Global.TileSize

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xPos), float64(yPos))

	shiftX := configuration.Global.TileSize
	if p.Alive {
		shiftX = p.AnimationStep * configuration.Global.TileSize
	}

	var particleImage *ebiten.Image

	switch p.Type {
	case 0:
		particleImage = assets.DustGrassImage
	case 1:
		particleImage = assets.DustSandImage
	case 3:
		particleImage = assets.DustWoodImage
	default:
		particleImage = assets.DustImage
	}

	screen.DrawImage(particleImage.SubImage(
		image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
	).(*ebiten.Image), op)
}
