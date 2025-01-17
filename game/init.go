package game

import (
	"fmt"
)

func (g *Game) RecursiveBlockingVerify(posX, posY int) {


	// block := g.floor.GetBlockContent(posX, posY);

	// fmt.Println(g.floor.IsBlockingBlock(block))

	blocking := g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y)

	fmt.Println(blocking)
	// 0: top, 1: right, 2: bottom, 3: left
	if blocking[0] {
		if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX, posY - 1)) {
			g.character.Y -= 1
			posY -= 1
			g.RecursiveBlockingVerify(posX, posY)
		} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX + 1, posY)) {
			g.character.X += 1
			posX += 1
			g.RecursiveBlockingVerify(posX, posY)
		} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX - 1, posY)) {
			g.character.X -= 1
			posX -= 1
			g.RecursiveBlockingVerify(posX, posY)
		}
		fmt.Println("top")
	} else if !blocking[0] {
		g.character.Y -= 1
		posY -= 1
	} else if blocking[1] {
		if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX + 1, posY)) {
			g.character.X += 1
			posX += 1
			g.RecursiveBlockingVerify(posX, posY)
		} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX, posY + 1)) {
			g.character.Y += 1
			posY += 1
			g.RecursiveBlockingVerify(posX, posY)
		} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX, posY - 1)) {
			g.character.Y -= 1
			posY -= 1
			g.RecursiveBlockingVerify(posX, posY)
		}
	} else if !blocking[1] {
		g.character.X += 1
		posX += 1
	} else if blocking[2] {
		if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX, posY + 1)) {
			g.character.Y += 1
			posY += 1
			g.RecursiveBlockingVerify(posX, posY)
		} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX + 1, posY)) {
			g.character.X += 1
			posX += 1
			g.RecursiveBlockingVerify(posX, posY)
		} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX - 1, posY)) {
			g.character.X -= 1
			posX -= 1
			g.RecursiveBlockingVerify(posX, posY)
		}
	}else if !blocking[2] {
		g.character.Y += 1
		posY += 1
	} else if blocking[3] {
		if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX - 1, posY)) {
			g.character.X -= 1
			posX -= 1
			g.RecursiveBlockingVerify(posX, posY)
		} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX, posY + 1)) {
			g.character.Y += 1
			posY += 1
			g.RecursiveBlockingVerify(posX, posY)
		} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX, posY - 1)) {
			g.character.Y -= 1
			posY -= 1
			g.RecursiveBlockingVerify(posX,
				posY)
		}
	}else if !blocking[3] {
		g.character.X -= 1
		posX -= 1
	}


	// if g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y)[0] {
	// 	// if !g.floor.IsBlockingPos(posX + 1, posY) {
	// 	// 	g.character.X += 1
	// 	// 	posX += 1
	// 	// 	g.RecursiveBlockingVerify(posX, posY)
	// 	// } else if !g.floor.IsBlockingPos(posX - 1, posY) {
	// 	// 	g.character.X -= 1
	// 	// 	posX -= 1
	// 	// 	g.RecursiveBlockingVerify(posX, posY)
	// 	// } else if !g.floor.IsBlockingPos(posX, posY + 1) {
	// 	// 	g.character.Y += 1
	// 	// 	posY += 1
	// 	// 	g.RecursiveBlockingVerify(posX, posY)
	// 	// } else if !g.floor.IsBlockingPos(posX, posY - 1) {
	// 	// 	g.character.Y -= 1
	// 	// 	posY -= 1
	// 	// 	g.RecursiveBlockingVerify(posX, posY)
	// 	// }
	// } else {
	// 	return
	// }

	// return


}

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.floor.Init()
	g.character.Init(g.floor.GetWidth(), g.floor.GetHeight())


	// posX := configuration.Global.ScreenCenterTileX - g.camera.X + g.character.X
	// posY := configuration.Global.ScreenCenterTileY - g.camera.Y + g.character.Y
	// g.RecursiveBlockingVerify(posX, posY);

	g.camera.Init(g.character.X, g.character.Y)


	// topLeftX := configuration.Global.ScreenCenterTileX - g.camera.X
	// topLeftY := configuration.Global.ScreenCenterTileY - g.camera.Y
	// posX := topLeftX + g.character.X
	// posY := topLeftY + g.character.Y

	// // fmt.Println()
	// block := g.floor.GetBlockContent(posX, posY)
	// bloque := g.floor.IsBlockingBlock(block)
	// end := false
	// fmt.Println(block, bloque, posX, posY)
	
	// for ;bloque && !end; {
	// 	fmt.Println("for")
	// 	if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX + 1, posY)) {
	// 		g.character.X += 1
	// 		posX += 1
	// 		bloque = false
	// 	} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX - 1, posY)) {
	// 		g.character.X -= 1
	// 		posX -= 1
	// 		bloque = false
	// 	} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX, posY + 1)) {
	// 		g.character.Y += 1
	// 		posY += 1
	// 		bloque = false
	// 	} else if !g.floor.IsBlockingBlock(g.floor.GetBlockContent(posX, posY - 1)) {
	// 		g.character.Y -= 1
	// 		posY -= 1
	// 		bloque = false
	// 	} else {
	// 		end = true
	// 	}
	// }

	// fmt.Println(posX, posY)

	// g.Update()
	// if g.floor.IsBlockingBlock(block) {

	// }
	// fmt.Println(g.floor.GetContent()[g.character.X][g.character.Y])

	// if g.floor.IsBlockingBlock(g.character.X)
}
