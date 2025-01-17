package particles

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func (p *Particle) Update() {
	if p.Alive {
		p.animationFrameCount++
		if p.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			p.animationFrameCount = 0
			p.timeAlive += 1
			p.AnimationStep = -p.AnimationStep

			if p.timeAlive > p.Duration {
				p.timeAlive = 0
				p.Alive = false
			}
		}
	}

}
