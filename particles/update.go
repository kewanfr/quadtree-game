package particles

// Update met à jour l'état de la particule. Elle incrémente le compteur de frames d'animation,
// met à jour l'étape d'animation et vérifie si la durée de l'animation de la particule est atteinte.
// Si la durée de l'animation est atteinte, la particule est marquée comme n'étant plus en vie et sera donc supprimée.
func (p *Particle) Update() {
	if p.Alive {
		// Incrémente le compteur de frames d'animation
		p.animationFrameCount++

		// Réinitialise l'étape d'animation si elle atteint 4 (correspondance avec le sheet de sprites)
		if p.AnimationStep == 4 {
			p.AnimationStep = 1
		}

		// Incrémente l'étape d'animation en fonction de la durée de l'étape
		if p.animationFrameCount%p.StepDuration == 0 {
			p.AnimationStep++
		}

		// Vérifie si la durée de l'animation est atteinte
		if p.animationFrameCount >= p.AnimationDuration {
			p.animationFrameCount = 0
			p.Alive = false
		}
	}
}
