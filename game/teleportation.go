package game

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func (g *Game) TeleportTo(x, y int) {

	blocking := g.floor.Blocking(x, y, g.camera.X, g.camera.Y)
	
	// Regarde si le personnage peut se déplacer sur la case souhaitée
	// (Marche bien pour les bordures de la map, ou pour les blocking blocks si y'en a que 1)
	// (Mais pas idéal si plusieurs blocking blocks sont côte à côte)

	if blocking[0] {
		x = x + 1
	} else if blocking[1] {
		x = x - 1
	} else if blocking[2] {
		y = y + 1
	} else if blocking[3] {
		y = y - 1
	}else {
		x = x + 1
	}
	
	g.character.X = x
	g.character.Y = y
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y)
}

func (g *Game) UpdateTeleport() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		if len(g.Portals) == 2 {
			g.Portals = make([]Portal, 0)
			g.Portals = append(g.Portals, Portal{
				X: g.character.X,
				Y: g.character.Y,
			})
			if configuration.Global.DebugMode {
				fmt.Println("Portal reset")
			}

		} else if len(g.Portals) == 1 {
			g.Portals = append(g.Portals, Portal{
				X: g.character.X,
				Y: g.character.Y,
			})
			if configuration.Global.DebugMode {
				fmt.Println("Portal ended")
			}
		} else if len(g.Portals) == 0 {
			g.Portals = append(g.Portals, Portal{
				X: g.character.X,
				Y: g.character.Y,
			})
			if configuration.Global.DebugMode {
				fmt.Println("Portal started")
			}
		}
	}



	if len(g.Portals) == 2 {
		if (g.Portals[0].X == g.character.X && g.Portals[0].Y == g.character.Y) {
			g.TeleportTo(g.Portals[1].X, g.Portals[1].Y)
		} else if (g.Portals[1].X == g.character.X && g.Portals[1].Y == g.character.Y) {
			g.TeleportTo(g.Portals[0].X, g.Portals[0].Y)
	}
	}
	return nil
}

func (g Game) DrawTeleport(screen *ebiten.Image, camX, camY int) {

	topLeftX := configuration.Global.ScreenCenterTileX - camX
	topLeftY := configuration.Global.ScreenCenterTileY - camY

	if len(g.Portals) > 0 && topLeftX+g.Portals[0].X >= 0 && topLeftX+g.Portals[0].X < configuration.Global.NumTileX && topLeftY+g.Portals[0].Y >= 0 && topLeftY+g.Portals[0].Y < configuration.Global.NumTileY {

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(topLeftX*configuration.Global.TileSize+g.Portals[0].X*configuration.Global.TileSize), float64(topLeftY*configuration.Global.TileSize+g.Portals[0].Y*configuration.Global.TileSize))

		screen.DrawImage(assets.TeleporterImage.SubImage(image.Rect(0, 0, configuration.Global.TileSize, configuration.Global.TileSize)).(*ebiten.Image), op)

	}
	if len(g.Portals) > 1 && topLeftX+g.Portals[1].X >= 0 && topLeftX+g.Portals[1].X < configuration.Global.NumTileX && topLeftY+g.Portals[1].Y >= 0 && topLeftY+g.Portals[1].Y < configuration.Global.NumTileY {

		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(float64(topLeftX*configuration.Global.TileSize+g.Portals[1].X*configuration.Global.TileSize), float64(topLeftY*configuration.Global.TileSize+g.Portals[1].Y*configuration.Global.TileSize))

		screen.DrawImage(assets.Teleporter_endImage.SubImage(image.Rect(0, 0, configuration.Global.TileSize, configuration.Global.TileSize)).(*ebiten.Image), op2)
	}

}
