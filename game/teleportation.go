package game

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func (g *Game) UpdateTeleport() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		if len(g.Portals) == 2 {
			g.Portals = make([]Portal, 0)
			g.Portals = append(g.Portals, Portal{
				X: g.character.X,
				Y: g.character.Y,
			})
			fmt.Println("Portal restarted")

		} else if len(g.Portals) == 1 {
			g.Portals = append(g.Portals, Portal{
				X: g.character.X,
				Y: g.character.Y,
			})
			fmt.Println("Portal created")
		} else if len(g.Portals) == 0 {
			g.Portals = append(g.Portals, Portal{
				X: g.character.X,
				Y: g.character.Y,
			})
			fmt.Println("Portal started")
		}
	}

	if len(g.Portals) == 2 && g.Portals[0].X == g.character.X && g.Portals[0].Y == g.character.Y {
		g.character.X = g.Portals[1].X
		g.character.Y = g.Portals[1].Y
	}

	return nil
}

func (g Game) DrawTeleport(screen *ebiten.Image, camX, camY int) {

	topLeftX := configuration.Global.ScreenCenterTileX - camX
	topLeftY := configuration.Global.ScreenCenterTileY - camY

	if len(g.Portals) > 0 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(topLeftX*configuration.Global.TileSize+g.Portals[0].X*configuration.Global.TileSize), float64(topLeftY*configuration.Global.TileSize+g.Portals[0].Y*configuration.Global.TileSize))

		screen.DrawImage(assets.TeleporterImage.SubImage(image.Rect(0, 0, configuration.Global.TileSize, configuration.Global.TileSize)).(*ebiten.Image), op)

	}
	if len(g.Portals) > 1 {

		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(float64(topLeftX*configuration.Global.TileSize+g.Portals[1].X*configuration.Global.TileSize), float64(topLeftY*configuration.Global.TileSize+g.Portals[1].Y*configuration.Global.TileSize))

		screen.DrawImage(assets.Teleporter_endImage.SubImage(image.Rect(0, 0, configuration.Global.TileSize, configuration.Global.TileSize)).(*ebiten.Image), op2)
	}

}
