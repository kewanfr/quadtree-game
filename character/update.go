package character

import "C"
import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/particles"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Character) Update(blocking [4]bool, particlesList *[]particles.Particle, currentTile int) {

	if !c.moving {
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			c.orientation = orientedRight
			if !blocking[1] {
				c.xInc = 1
				c.moving = true
				if configuration.Global.ExtParticles {
					addParticle(c, particlesList, currentTile)
				}
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			c.orientation = orientedLeft
			if !blocking[3] {
				c.xInc = -1
				c.moving = true
				if configuration.Global.ExtParticles {
					addParticle(c, particlesList, currentTile)
				}
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			c.orientation = orientedUp
			if !blocking[0] {
				c.yInc = -1
				c.moving = true
				if configuration.Global.ExtParticles {
					addParticle(c, particlesList, currentTile)
				}
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			c.orientation = orientedDown
			if !blocking[2] {
				c.yInc = 1
				c.moving = true
				if configuration.Global.ExtParticles {
					addParticle(c, particlesList, currentTile)
				}
			}
		}
	} else {
		c.animationFrameCount++
		if c.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			c.animationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			c.shift += shiftStep
			c.animationStep = -c.animationStep
			if c.shift > configuration.Global.TileSize-shiftStep {
				c.shift = 0
				c.moving = false
				c.X += c.xInc
				c.Y += c.yInc
				c.xInc = 0
				c.yInc = 0
			}
		}
	}

}

func addParticle(c *Character, particlesList *[]particles.Particle, currentTile int) {
	*particlesList = append(*particlesList, particles.Particle{
		X:                 c.X,
		Y:                 c.Y,
		AnimationStep:     1,
		Alive:             true,
		StepDuration:      10,
		AnimationDuration: 30,
		Type:              currentTile,
	})
}
