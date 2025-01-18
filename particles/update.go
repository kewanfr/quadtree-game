package particles

func (p *Particle) Update() {
	if p.Alive {
		p.animationFrameCount++

		if p.AnimationStep == 4 {
			p.AnimationStep = 1
		}

		if p.animationFrameCount%p.StepDuration == 0 {
			p.AnimationStep++
		}

		if p.animationFrameCount >= p.AnimationDuration {
			p.animationFrameCount = 0
			p.Alive = false

		}
	}

}
