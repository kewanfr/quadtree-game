package floor

const GRASS = 0
const DIRT = 1
const BRICK = 2
const WOOD = 3
const WATER = 4

const DIRT_GRASS_BORDER_LEFT = 5
const DIRT_GRASS_BORDER_TOP = 6
const DIRT_GRASS_BORDER_BOT = 7
const DIRT_GRASS_BORDER_RIGHT = 8

const BRICK_GRASS_BORDER_LEFT = 21
const BRICK_GRASS_BORDER_TOP = 22
const BRICK_GRASS_BORDER_BOT = 23
const BRICK_GRASS_BORDER_RIGHT = 24

const DIRT_GRASS_BORDER_ANGLE_LEFT = 9
const DIRT_GRASS_BORDER_ANGLE_RIGHT = 10
const DIRT_GRASS_BORDER_ANGLE_LEFT_BOT = 11
const DIRT_GRASS_BORDER_ANGLE_RIGHT_BOT = 12

const BRICK_GRASS_BORDER_ANGLE_LEFT = 25
const BRICK_GRASS_BORDER_ANGLE_RIGHT = 26
const BRICK_GRASS_BORDER_ANGLE_LEFT_BOT = 27
const BRICK_GRASS_BORDER_ANGLE_RIGHT_BOT = 28

const PATH_LEFT = 26
const PATH_H = 27
const PATH_V = 36
const PATH_RIGHT = 28

func SmoothTerrain(content [][]int) [][]int {
	width := len(content[0])
	height := len(content)

	for y := range content {
		for x := range content[y] {
			if content[y][x] != -1 {
				if x < width-1 && y < height-1 {

					// on évite les outs of bounds dans l'array
					left := -1
					if x-1 >= 0 {
						left = content[y][x-1]
					}
					right := -1
					if x+1 < width {
						right = content[y][x+1]
					}
					top := -1
					if y+1 < height {
						top = content[y+1][x]
					}
					bot := -1
					if y-1 >= 0 {
						bot = content[y-1][x]
					}

					switch content[y][x] {

					case DIRT:
						{

						}

					case BRICK:
						{
							if left == GRASS && right == BRICK && (top == BRICK || top == BRICK_GRASS_BORDER_LEFT || top == -1) && (bot == BRICK || bot == BRICK_GRASS_BORDER_LEFT || bot == -1) {
								content[y][x] = BRICK_GRASS_BORDER_LEFT
							}

							if left == BRICK && right == GRASS && (top == BRICK || top == BRICK_GRASS_BORDER_RIGHT || top == -1) && (bot == BRICK || bot == BRICK_GRASS_BORDER_RIGHT || bot == -1) {
								content[y][x] = BRICK_GRASS_BORDER_RIGHT
							}

							if (left == BRICK || left == BRICK_GRASS_BORDER_BOT || left == -1) && (right == BRICK || right == BRICK_GRASS_BORDER_BOT || right == -1) && top == GRASS && bot == BRICK {
								content[y][x] = BRICK_GRASS_BORDER_BOT
							}
							if (left == BRICK || left == BRICK_GRASS_BORDER_TOP || left == -1) && (right == BRICK || right == BRICK_GRASS_BORDER_TOP || right == -1) && top == BRICK && bot == GRASS {
								content[y][x] = BRICK_GRASS_BORDER_TOP
							}

							if left == GRASS && (right == BRICK || right == BRICK_GRASS_BORDER_TOP) && top == GRASS && bot == BRICK {
								content[y][x] = BRICK_GRASS_BORDER_ANGLE_LEFT
							}
						}
					}

				}

			}
		}
	}

	return content
}
