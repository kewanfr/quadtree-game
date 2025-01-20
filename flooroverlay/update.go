package flooroverlay

// Update met à jour l'état de l'overlay de tuile. Elle incrémente le compteur de frames d'animation,
// réinitialise l'étape d'animation si elle atteint 3, et vérifie si la durée de l'animation est atteinte.
// Si la durée de l'animation est atteinte, le compteur de frames d'animation est réinitialisé.
func (p *TileOverlay) Update() {
	// Incrémente le compteur de frames d'animation
	p.animationFrameCount++

	// Réinitialise l'étape d'animation si elle atteint 3 (correspondance avec le sheet de sprites)
	if p.AnimationStep == 3 {
		p.AnimationStep = 0
	}

	// Incrémente l'étape d'animation en fonction de la durée de l'étape
	if p.animationFrameCount%p.StepDuration == 0 {
		p.AnimationStep++
	}

	// Vérifie si la durée de l'animation est atteinte
	if p.animationFrameCount >= p.AnimationDuration {
		p.animationFrameCount = 0
	}
}
