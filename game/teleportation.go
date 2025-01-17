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

	g.character.X = x
	g.character.Y = y
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y)
}

// Cette fonction est exécutée dans le game/update si l'extension est activée
// Elle gère l'ajout de portails avec la touche T
// Et la téléportation du personnage si il se trouve sur un portail
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

	// Si il y a deux portails
	if len(g.Portals) == 2 {

		for i := 0; i < len(g.Portals); i++ {
			// Si le perso est sur un portail

			if g.Portals[i].X == g.character.X && g.Portals[i].Y == g.character.Y {
				// Il faut qu'il n'ait pas été récemment téléporté et qu'il ne soit pas en mouvement
				if !g.justTeleported && !g.character.GetIsMoving() {
					// On le téléporte sur l'autre portail
					g.TeleportTo(g.Portals[(i+1)%2].X, g.Portals[(i+1)%2].Y)
					g.justTeleported = true
				}

				// S'il se met à bouger et vient d'être téléporté, on réinitialise la variable
				if g.character.GetIsMoving() && g.justTeleported {
					g.justTeleported = false
				}
			}
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
