package quadtree

func (n *node) contentAt(x, y int) int {
	/*
	A partir d'un node, renvoie le contenu aux coordonnées (x, y)
	*/
	// Si la node est nil, il n'y a pas de sous-quadtree, donc on renvoie -1
	if n == nil {
		return -1
	}
    
	// On vérifie si (x, y) est en dehors des bornes (+ la taille de la node )
	if x < n.topLeftX || x >= n.topLeftX+n.width ||
	   y < n.topLeftY || y >= n.topLeftY+n.height {
		return -1
	}
    
	// Si c’est une feuille, on renvoie directement le contenu
	if n.isLeaf {
		return n.content
	}
    
	// SI c’est un noeud interne : on cherche (x, y)
	// On calcule la taille de la moitié de la node
	halfW := n.width / 2
	halfH := n.height / 2

	// On calcule les coordonnées du point
	splitX := n.topLeftX + halfW
	splitY := n.topLeftY + halfH

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
	for y := 0; y < len(contentHolder); y++ {
		for x := 0; x < len(contentHolder[y]); x++ {

			// On calcule la coordonnée du pixel par rapport au content global
			globalX := topLeftX + x
			globalY := topLeftY + y

			// On cherche le contenu dans le quadtree aux coordonnées globalX et globalY
			contentHolder[y][x] = q.root.contentAt(globalX, globalY)

		}
	}
}