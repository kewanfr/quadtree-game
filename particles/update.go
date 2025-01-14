package particles

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func (p *Particle) Update() {
	if p.Moving {
		p.animationFrameCount++
		if p.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			p.animationFrameCount = 0
			p.shift += 1
			p.AnimationStep = -p.AnimationStep

			if p.shift > 4 {
				p.shift = 0
				p.Moving = false
			}
		}
	}

}
