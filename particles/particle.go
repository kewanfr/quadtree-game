package particles

/*
*

		Particle est la structure permettant de représenter une particule sur le terrain
		- X : la position en x de la particule
	   - Y : la position en y de la particule
	   - AnimationStep : compteur de l'étape d'animation de la particule
	   - Alive : un booléen qui indique si la particule est en "vie"
	   - animationFrameCount : le nombre de frames depuis la dernière étape d'animation
	   - StepDuration : la durée d'une étape d'animation
	   - AnimationDuration : la durée totale de l'animation
	   - Type : le type de sol sur lequel la particule est posée

*
*/
type Particle struct {
	X, Y                int
	AnimationStep       int
	Alive               bool
	animationFrameCount int
	StepDuration        int
	AnimationDuration   int
	Type                int
}
