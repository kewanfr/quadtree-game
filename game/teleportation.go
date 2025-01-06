package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func (g *Game) UpdateTeleporation() error {


	if configuration.Global.ExtTeleportation {
		if inpututil.IsKeyJustPressed(ebiten.KeyT) {
			if len(g.Portals) == 2 {
				g.Portals = make([]Portal, 0)
				g.Portals = append(g.Portals, Portal{
					X: g.character.X,
					Y: g.character.Y,
				})
				fmt.Println("Portal restarted")

			}else if len(g.Portals) == 1 {
				g.Portals = append(g.Portals, Portal{
					X: g.character.X,
					Y: g.character.Y,
				})
				fmt.Println("Portal created")
			}else if len(g.Portals) == 0 {
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
	
	
	}


	return nil
}
