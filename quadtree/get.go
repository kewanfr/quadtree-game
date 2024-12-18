package quadtree

/**
GetContent sur un noeud remplit le tableau contentHolder (le terrain représenté par le quadtree)
à partir du contenu d'un noeud. Si le noeud est une feuille, on remplit le tableau avec la valeur,
sinon on appelle de façon récursive GetContent sur les 4 sous-noeuds.

Entrées:
- topLeftX, topLeftY: les coordonnées du point le plus en haut à gauche du terrain à remplir (par rapport au content global)
- contentHolder: le tableau à remplir avec le contenu du noeud

Sorties:
	Aucune, le tableau contentHolder est modifié en place.
**/
func (n *node) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
    if n == nil {
        return
    }

    // Si le noeud est une feuille
    if n.isLeaf {
		// On remplit le tableau contentHolder avec la valeur du noeud
		// En vérifiant que les coordonnées du noeud sont bien dans le terrain à remplir


		// Reduction: on itère sur la taille du noeud pour remplir les valeurs
		// (Quand le noeud a une taille supérieure à 1, il y a pluiseurs cases à remplir)
		for y := 0; y < n.height; y++ {
			for x := 0; x < n.width; x++ {

				// On vérifie que les coordonnées du noeud sont bien dans le terrain à remplir
				if n.topLeftY+y >= topLeftY && n.topLeftY+y < topLeftY+len(contentHolder) &&
					n.topLeftX+x >= topLeftX && n.topLeftX+x < topLeftX+len(contentHolder[0]) {
					contentHolder[n.topLeftY+y-topLeftY][n.topLeftX+x-topLeftX] = n.content
				}
			}
		}

		
        return
    }

    // Si le noeud n'est pas une feuille
    if n.topLeftNode != nil {
        n.topLeftNode.GetContent(topLeftX, topLeftY, contentHolder)
    }
    if n.topRightNode != nil {
        n.topRightNode.GetContent(topLeftX, topLeftY, contentHolder)
    }
    if n.bottomLeftNode != nil {
        n.bottomLeftNode.GetContent(topLeftX, topLeftY, contentHolder)
    }
    if n.bottomRightNode != nil {
        n.bottomRightNode.GetContent(topLeftX, topLeftY, contentHolder)
    }
}

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
/**

GetContent sur un quadtree remplit le tableau contentHolder (le terrain représenté par le quadtree) à partir du contenu du quadtree.
Si le quadtree est vide, on ne fait rien.
Sinon, on appelle la méthode GetContent sur le noeud racine.

Entrées:
- topLeftX, topLeftY: les coordonnées du point le plus en haut à gauche du terrain à remplir (par rapport au content global)
- contentHolder: le tableau à remplir avec le contenu du quadtree q, correspond au terrain qui sera affiché.

Sorties:
	Aucune, contentHolder est modifié en place.

**/
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {

	// On vérifie si le quadtree est vide
	if q.root == nil {
		return
	}


	// On remplit le quadtree avec -1
	// pour signifier que le terrain est vide
	for i := 0; i < len(contentHolder); i++ {
		for j := 0; j < len(contentHolder[i]); j++ {
			contentHolder[i][j] = -1
		}
	}


	// On appelle la méthode GetContent du noeud racine
	// avec les coordonnées du point le plus en haut à gauche du terrain à remplir
	// et le tableau à remplir
	q.root.GetContent(topLeftX, topLeftY, contentHolder)

}
