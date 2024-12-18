package quadtree

/**
contentAt retourne le contenu de la node dans le quadtree, à la position (x, y).

Entrées:
- x, y: les coordonnées du point dont on veut connaître le contenu.

Sorties:
- le contenu de la node à la position (x, y).

Si la node est une feuille, le contenu est directement renvoyé.
Sinon, on regarde dans quelle partie de la node se trouve le point (x, y) (en haut, en bas, à gauche, à droite)
et on appelle la fonction de manière récursive sur la sous-node.
**/
func (n *node) contentAt(x, y int) int {

	// Si la node est nulle ou si le point qu'on cherche est en dehors de la node
	// (on calcule à partir des coordonnées de la node + sa taille)
	if n == nil || (x < n.topLeftX || x >= n.topLeftX+n.width || y < n.topLeftY || y >= n.topLeftY+n.height)  {
		return -1
	}
	// Si c’est une feuille, on renvoie directement le contenu
	if n.isLeaf {
		return n.content
	}

	// On calcule les coordonnées du point auquel on sépare la node
	splitX := n.topLeftX + n.width / 2
	splitY := n.topLeftY + n.height / 2

	// On cherche la partie dans lequel est le point qu'on cherche
	if y < splitY {
		// Si le point à trouver est dans la partie haute
		if x < splitX {
			// Partie Gauche
			return n.topLeftNode.contentAt(x, y)
		} else {
			// Partie droite
			return n.topRightNode.contentAt(x, y)
		}
	} else {
		// Si le point à trouver est dans la partie basse du quadtree
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
/**

Entrées:
- topLeftX, topLeftY: les coordonnées du point le plus en haut à gauche du terrain à remplir (par rapport au content global)
- contentHolder: le tableau à remplir avec le contenu du quadtree q, correspond au terrain qui sera affiché.

Sorties:
	Aucune, contentHolder est modifié en place.

**/
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {


	// On itère sur chaque pixel du terrain contentHolder
	for y := 0; y < len(contentHolder); y++ {
		for x := 0; x < len(contentHolder[y]); x++ {

			// On calcule la coordonnée du pixel par rapport au content global 
			globalX := topLeftX + x
			globalY := topLeftY + y

			// Si le point est en dehors du quadtree, on met -1 (donc un point hors du terrain, l'utilisateur ne peut pas y aller)
			if q.root == nil || (globalX < q.root.topLeftX || globalX >= q.root.topLeftX+q.root.width || globalY < q.root.topLeftY || globalY >= q.root.topLeftY+q.root.height)  {
				contentHolder[y][x] = -1 
			}else {
				// Sinon, on définit le contenu du pixel en recherchant sa valeur dans le quadtree
				contentHolder[y][x] = q.root.contentAt(globalX, globalY)
			}


		}
	}
}
