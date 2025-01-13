package particles

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

func (p Particle) Draw(screen *ebiten.Image, camX, camY int) {
	xShift := 0
	yShift := 0

	xTileForDisplay := p.X - camX + configuration.Global.ScreenCenterTileX
	yTileForDisplay := p.Y - camY + configuration.Global.ScreenCenterTileY
	xPos := (xTileForDisplay)*configuration.Global.TileSize + xShift
	yPos := (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 + yShift

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xPos), float64(yPos))

	shiftX := configuration.Global.TileSize
	if p.Moving {
		shiftX += p.AnimationStep * configuration.Global.TileSize
	}
	shiftY := 0 * configuration.Global.TileSize

	screen.DrawImage(assets.DustImage.SubImage(
		image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
	).(*ebiten.Image), op)
}
