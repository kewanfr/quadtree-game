package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

/**
IsBlockingBlock renvoie si un block (à partir de sa valeur int) est un block bloquant ou non
Il sert à l'extension BlockBlocks "interdiction de marcher sur l'eau"
**/
func (f Floor) IsBlockingBlock(block int) (isBlocking bool) {
	// On itère sur les blocks bloquants
	for _, v := range configuration.Global.BlockingBlocks {
		if v == block {
			isBlocking = true
		}
	}

	return isBlocking
}

/**
	FindSpawn renvoie une coordonnée de bloc sur lequel le personnage peut apparaître, en fonction de sa position actuelle.
	La fonction est récursive, et teste les 4 positions adjacentes à la position actuelle.
	Renvoie true avec les coordonnées si une position est trouvée et false sinon
**/
func (f *Floor) FindSpawn(posX, posY int, camXPos, camYPos int, checked map[[2]int]bool) (bool, int, int) {
	// Vérif si la position sort du terrain
	if posY < 0 || posX < 0 || posX >= configuration.Global.NumTileX || posY >= configuration.Global.NumTileY {
		return false, 0, 0
	}

	// Si la position a déjà été vérifiée
	if checked[[2]int{posX, posY}]{
		return false, 0, 0
	}
	checked[[2]int{posX, posY}] = true

	relativeXPos := posX - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := posY - camYPos + configuration.Global.ScreenCenterTileY

	if relativeXPos < 0 || relativeYPos < 0 || relativeXPos >= len(f.content[0]) || relativeYPos >= len(f.content) {
		return false, 0, 0
	}

	// Verifier si la position est bloquée
	if f.IsBlockingBlock(f.content[relativeYPos][relativeXPos]) {

		// On essaye les 4 positions adjacentes

		positions := map[int][2]int{
			0: {0, -1}, // Haut
			1: {0, 1}, // Bas
			2: {-1, 0}, // Gauche
			3: {1, 0}, // Droite
		}

		for i := 0; i < len(positions); i++ {

			// On décale les coordonnées
			newX := posX + positions[i][0]
			newY := posY + positions[i][1]
			// On cherche un spawn à partir de la nouvelle position (de manière récursive)
			ok, newX, newY := f.FindSpawn(newX, newY, camXPos, camYPos, checked)
			if ok {
				return ok, newX, newY
			}
		}

		return false, 0, 0
	}

	return true, posX, posY
}


// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {
	

	// Position relative du personnage
	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	// On vérifie les bloquages en dehors du terrain
	blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1
	blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1
	blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1
	blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1

	// On vérifie les bloquages de blocs dans le terrain
	if configuration.Global.ExtBlockingBlocks {

		// Si les cases autour sont des blocs bloquants, on dit qu'ils sont bloquants, afin d'empêche de se déplacer dessus
		// 0: top, 1: right, 2: bottom, 3: left
		blocking[0] = blocking[0] || f.IsBlockingBlock(f.content[relativeYPos-1][relativeXPos])
		blocking[1] = blocking[1] || f.IsBlockingBlock(f.content[relativeYPos][relativeXPos+1])
		blocking[2] = blocking[2] || f.IsBlockingBlock(f.content[relativeYPos+1][relativeXPos])
		blocking[3] = blocking[3] || f.IsBlockingBlock(f.content[relativeYPos][relativeXPos-1])
	}

	return blocking
}
