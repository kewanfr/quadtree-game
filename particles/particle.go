package particles

/*
*

		Particle est la structure permettant de représenter une particule sur le terrain
		- X : la position en x de la particule
	   - Y : la position en y de la particule
	   - AnimationStep : l'étape d'animation de la particule
	   - Alive : un booléen qui indique si la particule est en "vie"
	   - timeAlive : le temps de vie de la particule
	   - animationFrameCount : le nombre de frames depuis la dernière étape d'animation
	   - Duration : la durée de vie de la particule

*
*/
type Particle struct {
	X, Y                int
	AnimationStep       int
	Alive               bool
	timeAlive           int
	animationFrameCount int
	Duration            int
}
