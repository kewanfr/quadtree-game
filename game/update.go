package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
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

	if configuration.Global.ExtFloorSave && inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		g.floor.SaveFloor()
	}

	// Particule DEVANT le bonhomme
	if configuration.Global.ExtParticles {
		for i := 0; i < len(g.particles); i++ {
			g.particles[i].Update()

			//Si la particule a atteint ça durée maximale, on le retire de la liste des particules active dans le jeu
			if !g.particles[i].Alive {
				g.particles = append(g.particles[:i], g.particles[i+1:]...)
				i--
			}
		}
	}

	// g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))

	blocking := g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y)
	g.character.Update(blocking, &g.particles)

	if configuration.Global.ExtSpeedRun {
		g.character.Update(blocking, &g.particles)
	}

	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y)

	if configuration.Global.ExtTeleportation {
		err := g.UpdateTeleport()
		if err != nil {
			return err
		}
	}

	return nil
}
