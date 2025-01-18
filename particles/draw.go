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
	xShift := 0
	yShift := 0

	// Calcule la position de la particule sur le terrain
	xTileForDisplay := p.X - camX + configuration.Global.ScreenCenterTileX
	yTileForDisplay := p.Y - camY + configuration.Global.ScreenCenterTileY

	// Calcule la position de la particule sur l'écran
	xPos := (xTileForDisplay)*configuration.Global.TileSize + xShift
	yPos := (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 + yShift

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
