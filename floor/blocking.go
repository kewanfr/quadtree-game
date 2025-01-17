package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)


func (f Floor) IsBlockingBlock(block int) (isBlocking bool) {
	for _, v := range configuration.Global.BlockingBlocks {
		if v == block {
			isBlocking = true
		}
	}

	return isBlocking
}

// func RandomFloor() (f Floor) {
// 	randInt := 

// 	return f
// }
// func (f Floor) GenerateFloor() {
// 	f.content = make([][]int, configuration.Global.NumTileY)
// 	for y := 0; y < len(f.content); y++ {
// 		f.content[y] = make([]int, configuration.Global.NumTileX)
// 	}
// }

// func (f Floor) IsBlockingPos(posX, posY int) (isBlocking bool) {
// 	return posY < 0 || posX < 0 || posY >= configuration.Global.NumTileY || posX >= configuration.Global.NumTileX || f.IsBlockingBlock(f.content[posY][posX])
// }

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {
	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1
	blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1
	blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1
	blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1

	// if blocking[0] {
	// 	randInt := rand.Intn(4)
	// 	f.content[relativeYPos-1][relativeXPos] = randInt
	// 	// posOnFullX := relativeXPos + camXPos - configuration.Global.ScreenCenterTileX
	// 	// posOnFullY := relativeYPos - 1 + camYPos - configuration.Global.ScreenCenterTileY
	// 	// f.quadtreeContent.SetOrAddContent(posOnFullX, posOnFullY, randInt)

	// 	// f.quadtreeContent = quadtree.MakeFromArray(f.fullContent)


		
	// }

	if configuration.Global.ExtBlockingBlocks {

		// 0: top, 1: right, 2: bottom, 3: left
		blocking[0] = blocking[0] || f.IsBlockingBlock(f.content[relativeYPos-1][relativeXPos])
		blocking[1] = blocking[1] || f.IsBlockingBlock(f.content[relativeYPos][relativeXPos+1])
		blocking[2] = blocking[2] || f.IsBlockingBlock(f.content[relativeYPos+1][relativeXPos])
		blocking[3] = blocking[3] || f.IsBlockingBlock(f.content[relativeYPos][relativeXPos-1])
	}

	// blocking[0] = f.IsBlockingPos(relativeXPos, relativeYPos-1)
	// blocking[1] = f.IsBlockingPos(relativeXPos+1, relativeYPos)
	// blocking[2] = f.IsBlockingPos(relativeXPos, relativeYPos+1)
	// blocking[3] = f.IsBlockingPos(relativeXPos-1, relativeYPos)

	return blocking
}
