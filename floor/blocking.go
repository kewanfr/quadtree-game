package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

func isBlockingBlock(block int) (isBlocking bool) {
	for _, v := range configuration.Global.BlockingBlocks {
		if v == block {
			isBlocking = true
		}
	}

	return isBlocking
}

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {
	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	blocking[0] = relativeYPos <= 0 || (relativeYPos-1 >= 0 && f.content[relativeYPos-1][relativeXPos] == -1)
	blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || (relativeXPos+1 < len(f.content[relativeYPos]) && f.content[relativeYPos][relativeXPos+1] == -1)
	blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || (relativeYPos+1 < len(f.content) && f.content[relativeYPos+1][relativeXPos] == -1)
	blocking[3] = relativeXPos <= 0 || (relativeXPos-1 >= 0 && f.content[relativeYPos][relativeXPos-1] == -1)

	if configuration.Global.ExtBlockingBlocks {
		blocking[0] = blocking[0] || (relativeYPos-1 >= 0 && isBlockingBlock(f.content[relativeYPos-1][relativeXPos]))
		blocking[1] = blocking[1] || (relativeXPos+1 < len(f.content[relativeYPos]) && isBlockingBlock(f.content[relativeYPos][relativeXPos+1]))
		blocking[2] = blocking[2] || (relativeYPos+1 < len(f.content) && isBlockingBlock(f.content[relativeYPos+1][relativeXPos]))
		blocking[3] = blocking[3] || (relativeXPos-1 >= 0 && isBlockingBlock(f.content[relativeYPos][relativeXPos-1]))
	}

	return blocking
}
