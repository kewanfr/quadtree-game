package flooroverlay

// TileOverlay représente un overlay de tuile.
// X et Y sont les coordonnées de la tuile.
// AnimationStep est l'étape de l'animation.
// animationFrameCount est le nombre de frames de l'animation.
// StepDuration est la durée d'une étape.
// AnimationDuration est la durée totale de l'animation.
// Type est le type de l'overlay (1: fleurs, 2: buisson).
type TileOverlay struct {
	X, Y                int
	AnimationStep       int
	animationFrameCount int
	StepDuration        int
	AnimationDuration   int
	Type                int
}
