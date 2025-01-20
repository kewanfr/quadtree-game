package flooroverlay

func (p *TileOverlay) Update() {
	p.animationFrameCount++

	if p.AnimationStep >= 4 {
		p.AnimationStep = 0
	}

	// Si on a atteint la durée d'une étape, on passe à l'étape suivante
	if p.animationFrameCount%p.StepDuration == 0 {
		p.AnimationStep++
	}

	if p.animationFrameCount >= p.AnimationDuration {
		p.animationFrameCount = 0
	}
}
