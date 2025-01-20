package flooroverlay

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

func (p TileOverlay) Draw(screen *ebiten.Image, camX, camY int) {
	// Calcule la position de l'animation sur le terrain
	xTileForDisplay := p.X - camX + configuration.Global.ScreenCenterTileX
	yTileForDisplay := p.Y - camY + configuration.Global.ScreenCenterTileY

	xPos := (xTileForDisplay) * configuration.Global.TileSize
	yPos := (yTileForDisplay) * configuration.Global.TileSize

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xPos), float64(yPos))

	shiftX := p.AnimationStep * configuration.Global.TileSize

	var img *ebiten.Image
	switch p.Type {
	case 1:
		img = assets.FlowerImage
	case 2:
		img = assets.BuissonImage
	}
	screen.DrawImage(img.SubImage(
		image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
	).(*ebiten.Image), op)
}
