package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}

	if configuration.Global.ExtTeleportation {
		if inpututil.IsKeyJustPressed(ebiten.KeyT) {
			if len(g.floor.Portals) == 2 {
				g.floor.Portals = make([]floor.Portal, 0)
				g.floor.Portals = append(g.floor.Portals, floor.Portal{
					X: g.character.X,
					Y: g.character.Y,
				})
				fmt.Println("Portal restarted")

			}else if len(g.floor.Portals) == 1 {
				g.floor.Portals = append(g.floor.Portals, floor.Portal{
					X: g.character.X,
					Y: g.character.Y,
				})
				fmt.Println("Portal created")
			}else if len(g.floor.Portals) == 0 {
				g.floor.Portals = append(g.floor.Portals, floor.Portal{
					X: g.character.X,
					Y: g.character.Y,
				})
				fmt.Println("Portal started")
			}
		}

		if len(g.floor.Portals) == 2 && g.floor.Portals[0].X == g.character.X && g.floor.Portals[0].Y == g.character.Y { 
			g.character.X = g.floor.Portals[1].X
			g.character.Y = g.floor.Portals[1].Y
		}
	
	
	}

	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y)

	return nil
}
