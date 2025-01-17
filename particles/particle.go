package particles

type Particle struct {
	X, Y                int
	AnimationStep       int
	Alive               bool
	timeAlive           int
	animationFrameCount int
	Duration            int
}
