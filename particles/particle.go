package particles

type Particle struct {
	X, Y                int
	AnimationStep       int
	Moving              bool
	shift               int
	animationFrameCount int
}
