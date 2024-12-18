package quadtree

import (
	"fmt"
)

func (n *node) contentAt(x, y int) int {

	// Si c’est une feuille, on renvoie directement le contenu
	if n.isLeaf {
		fmt.Println(x, y, n.content)
		return n.content
	}

	// On calcule les coordonnées du point
	splitX := n.topLeftX + n.width / 2
	splitY := n.topLeftY + n.height / 2

	if y < splitY {
		// Si on est dans la partie haute de la node
		if x < splitX {
			// Partie Gauche
			return n.topLeftNode.contentAt(x, y)
		} else {
			// Partie droite
			return n.topRightNode.contentAt(x, y)
		}
	} else {
		// Si on est dans la partie basse 
		if x < splitX { 
			// partie gauche
			return n.bottomLeftNode.contentAt(x, y)
		} else {
			//partie droite
			return n.bottomRightNode.contentAt(x, y)
		}
	}
}

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	// Pour chaque pixel du contentHolder que l'on veut remplir

	// contentHolder = make([][]int, configuration.Global.NumTileY)
	// for i := range contentHolder {
	// 	contentHolder[i] = make([]int, configuration.Global.NumTileX)
	// }

	// fmt.Println(len(contentHolder), len(contentHolder[0]), configuration.Global.NumTileY, configuration.Global.NumTileX)

	for y := 0; y < len(contentHolder); y++ {
		for x := 0; x < len(contentHolder[y]); x++ {

			// On calcule la coordonnée du pixel par rapport au content global
			globalX := topLeftX + x
			globalY := topLeftY + y

			if q.root == nil || (globalX < q.root.topLeftX || globalX >= q.root.topLeftX+q.root.width || globalY < q.root.topLeftY || globalY >= q.root.topLeftY+q.root.height)  {
				contentHolder[y][x] = -1 
			}else {
				contentHolder[y][x] = q.root.contentAt(globalX, globalY)
			}

			// On cherche le contenu dans le quadtree aux coordonnées globalX et globalY

		}
	}
}