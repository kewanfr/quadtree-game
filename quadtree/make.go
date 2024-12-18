package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	height := len(floorContent)
	if height == 0 {
		return Quadtree{}
	}
	width := len(floorContent[0])

	q.width = width
	q.height = height
	q.root = makeNode(floorContent, 0, 0, width, height)

	return q
}

func makeNode(floorContent [][]int, startX, startY, width, height int) *node {
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

	//Réduction et elimination des doublons
	if topLeftNode != nil && topRightNode != nil &&
		bottomLeftNode != nil && bottomRightNode != nil &&
		topLeftNode.isLeaf && topRightNode.isLeaf &&
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
