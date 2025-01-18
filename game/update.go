package game

import (
	"log"

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

	if g.messageFrames > 0 {
		g.messageFrames-- // Réduire la durée d'affichage du message
	}

	// Si on appuie sur la touche F5, on sauvegarde le terrain
	if configuration.Global.ExtFloorSave && inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		g.floor.SaveFloor()
		if configuration.Global.DebugMode {
			log.Println("Terrain sauvegardé")
		}
		g.message = "File saved"
		g.messageFrames = 120 // Affiche le message pendant 60 FPS * 2
	}

	// Particule AVANT le bonhomme
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

	var currentTile int

	relativeXPos := g.character.X - g.camera.X + configuration.Global.ScreenCenterTileX
	relativeYPos := g.character.Y - g.camera.Y + configuration.Global.ScreenCenterTileY

	if relativeXPos >= 0 && relativeXPos < len(g.floor.GetContent()[0]) && relativeYPos >= 0 && relativeYPos < len(g.floor.GetContent()) {
		currentTile = g.floor.GetContent()[relativeYPos][relativeXPos]
	} else {
		currentTile = -1
	}

	g.character.Update(blocking, &g.particles, currentTile)

	if configuration.Global.ExtSpeedRun {
		g.character.Update(blocking, &g.particles, currentTile)
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
