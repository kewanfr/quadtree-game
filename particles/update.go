package particles

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func (p *Particle) Update() {
	if p.Moving {
		p.animationFrameCount++
		if p.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			p.animationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			p.shift += shiftStep
			p.AnimationStep = -p.AnimationStep

			if p.shift > configuration.Global.TileSize-shiftStep {
				p.shift = 0
				p.Moving = false
			}
		}
	}

}
