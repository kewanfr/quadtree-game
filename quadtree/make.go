package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	// on vérifie que floorContent n'est pas vide, auquel cas ont return un quadtree vide
	if len(floorContent) == 0 {
		return q
	}

	height := len(floorContent)   // y
	width := len(floorContent[0]) // x

	//idem
	if height == 0 || width == 0 {
		return q
	}

	// ne devrait normalement pas arriver
	for _, row := range floorContent {
		if len(row) == 0 {
			panic("floorContent invalide : les lignes ne doivent pas être vides")
		}
	}

	q.width = width
	q.height = height

	// node parent
	q.root = makeNode(floorContent, 0, 0, width, height)

	return q
}

/*
*

	  	makeNode est une fonction récursive qui construit un nœud du quadtree.

		Elle prend en entrée un tableau représentant le contenu du terrain,
		les coordonnées du coin supérieur gauche de la zone à représenter,
		ainsi que sa largeur et sa hauteur.

		Elle s'arrête lorsqu'un node a pour dimension 1x1 contient une valeur à 0 dans ses dimensions.

		Paramètres :
		- floorContent : [][]int - Le contenu du terrain, représenté par un tableau 2D d'entiers.
		- startX : int - La coordonnée X du coin supérieur gauche de la zone à représenter.
		- startY : int - La coordonnée Y du coin supérieur gauche de la zone à représenter.
		- width : int - La largeur de la zone à représenter.
		- eight : int - La hauteur de la zone à représenter.

		Sortie :
		- *node : Un pointeur vers le nœud du quadtree créé.

*
*/
func makeNode(floorContent [][]int, startX, startY, width, height int) *node {
	// Arrêt de la fonction récursive si les dimensions sont 1x1 ou contiennent 0
	// Ont return une feuille comme node
	if width <= 0 || height <= 0 || (width == 1 && height == 1) {
		return &node{
			topLeftX: startX,
			topLeftY: startY,
			width:    width,
			height:   height,
			content:  floorContent[startY][startX],
			isLeaf:   true,
		}
	}

	// Création des sous nodes avec le découpage présenté dans le pdf

	var halfWidth int = width / 2
	var halfHeight int = height / 2

	var topLeftNode *node = makeNode(floorContent, startX, startY, halfWidth, halfHeight)
	var topRightNode *node = makeNode(floorContent, startX+halfWidth, startY, width-halfWidth, halfHeight)
	var bottomLeftNode *node = makeNode(floorContent, startX, startY+halfHeight, halfWidth, height-halfHeight)
	var bottomRightNode *node = makeNode(floorContent, startX+halfWidth, startY+halfHeight, width-halfWidth, height-halfHeight)

	// Réduction et elimination des doublons
	if topLeftNode.isLeaf && topRightNode.isLeaf &&
		bottomLeftNode.isLeaf && bottomRightNode.isLeaf &&
		topLeftNode.content == topRightNode.content &&
		topRightNode.content == bottomLeftNode.content &&
		bottomLeftNode.content == bottomRightNode.content {
		return &node{
			topLeftX: startX,
			topLeftY: startY,
			width:    width,
			height:   height,
			content:  topLeftNode.content,
			isLeaf:   true,
		}
	}

	// On retourne le nœud contenant les sous nœuds
	return &node{
		topLeftX:        startX,
		topLeftY:        startY,
		width:           width,
		height:          height,
		isLeaf:          false,
		topLeftNode:     topLeftNode,
		topRightNode:    topRightNode,
		bottomLeftNode:  bottomLeftNode,
		bottomRightNode: bottomRightNode,
	}
}
