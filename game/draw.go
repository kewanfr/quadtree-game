package game

import (
	"fmt"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"image/color"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"golang.org/x/image/font/basicfont"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Draw permet d'afficher à l'écran tous les éléments du jeu
// (le sol, le personnage, les éventuelles informations de debug).
// Il faut faire attention à l'ordre d'affichage pour éviter d'avoir
// des éléments qui en cachent d'autres.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.CurrentState == 0 {

		g.floor.Draw(screen)

		if configuration.Global.ExtTeleportation {
			g.DrawTeleport(screen, g.camera.X, g.camera.Y)
		}

		if configuration.Global.ExtParticles {
			for _, particle := range g.particles {
				particle.Draw(screen, g.camera.X, g.camera.Y)
			}
		}

		if configuration.Global.ExtFloorAnimation {
			for _, overlay := range g.tileOverlays {
				// On ne dessine que les overlays qui sont dans la zone de la caméra
				if overlay.Y >= g.camera.Y-configuration.Global.ScreenCenterTileY && overlay.Y <= g.camera.Y+configuration.Global.ScreenCenterTileY && overlay.X >= g.camera.X-configuration.Global.ScreenCenterTileX && overlay.X <= g.camera.X+configuration.Global.ScreenCenterTileX {
					overlay.Draw(screen, g.camera.X, g.camera.Y)
				}
			}
		}

		g.character.Draw(screen, g.camera.X, g.camera.Y)

		if configuration.Global.DebugMode {
			g.drawDebug(screen)
		}

		if g.messageFrames > 0 {
			text.Draw(screen, g.message, basicfont.Face7x13, 20, screen.Bounds().Dy()-20, color.White)
		}
	} else if g.CurrentState == 1 {
		g.drawTitleScreen(screen)
	}
}

func (g Game) drawTitleScreen(screen *ebiten.Image) {
	/*startButton := "Appuyez sur Espace!"
	keys := []string{
		"Fleches - Deplacement",
		"T - Placer un portail",
		") - Dezoomer    = - Zoomer",
		"F5 - Sauvegarder la carte",
		"D - DebugMode",
	}*/

	//Image de fond
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(screen.Bounds().Dx())/float64(assets.TitleImage.Bounds().Dx()), float64(screen.Bounds().Dy())/float64(assets.TitleImage.Bounds().Dy()))
	screen.DrawImage(assets.TitleImage, op)

	/*
		// Invite pour lancer le jeu
		text.Draw(screen, startButton, basicfont.Face7x13, screen.Bounds().Dx()/3-len(startButton)*2, screen.Bounds().Dy()/3, color.RGBA{255, 0, 0, 255})

		// Liste des touches
		for i, key := range keys {
			text.Draw(screen, key, basicfont.Face7x13, screen.Bounds().Dx()/3-len(key)*2, screen.Bounds().Dy()/3+20*(i+1), color.Black)
		}

	*/
}

// drawDebug se charge d'afficher les informations de debug si
// l'utilisateur le demande (positions absolues du personnage
// et de la caméra, grille avec les coordonnées, etc).
func (g Game) drawDebug(screen *ebiten.Image) {

	gridColor := color.NRGBA{R: 255, G: 255, B: 255, A: 63}
	gridHoverColor := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	gridLineSize := 2
	cameraColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	cameraLineSize := 1

	mouseX, mouseY := ebiten.CursorPosition()

	xMaxPos := configuration.Global.ScreenWidth
	yMaxPos := configuration.Global.ScreenHeight

	for x := 0; x < configuration.Global.NumTileX; x++ {
		xGeneralPos := x*configuration.Global.TileSize + configuration.Global.TileSize/2
		xPos := float32(xGeneralPos)

		lineColor := gridColor
		if mouseX+1 >= xGeneralPos && mouseX+1 <= xGeneralPos+gridLineSize {
			lineColor = gridHoverColor
		}

		vector.StrokeLine(screen, xPos, 0, xPos, float32(yMaxPos), float32(gridLineSize), lineColor, false)

		xPrintValue := g.camera.X + x - configuration.Global.ScreenCenterTileX
		xPrint := fmt.Sprint(xPrintValue)
		if len(xPrint) <= (2*configuration.Global.TileSize)/16 || (xPrintValue > 0 && xPrintValue%2 == 0) || (xPrintValue < 0 && (-xPrintValue)%2 == 0) {
			xTextPos := xGeneralPos - 3*len(xPrint) - 1
			ebitenutil.DebugPrintAt(screen, xPrint, xTextPos, yMaxPos)
		}
	}

	for y := 0; y < configuration.Global.NumTileY; y++ {
		yGeneralPos := y*configuration.Global.TileSize + configuration.Global.TileSize/2
		yPos := float32(yGeneralPos)

		lineColor := gridColor
		if mouseY+1 >= yGeneralPos && mouseY+1 <= yGeneralPos+gridLineSize {
			lineColor = gridHoverColor
		}

		vector.StrokeLine(screen, 0, yPos, float32(xMaxPos), yPos, float32(gridLineSize), lineColor, false)

		yPrint := fmt.Sprint(g.camera.Y + y - configuration.Global.ScreenCenterTileY)
		xTextPos := xMaxPos + 1
		yTextPos := yGeneralPos - 8
		ebitenutil.DebugPrintAt(screen, yPrint, xTextPos, yTextPos)
	}

	vector.StrokeRect(screen, float32(configuration.Global.ScreenCenterTileX*configuration.Global.TileSize), float32(configuration.Global.ScreenCenterTileY*configuration.Global.TileSize), float32(configuration.Global.TileSize+1), float32(configuration.Global.TileSize+1), float32(cameraLineSize), cameraColor, false)

	ySpace := 16
	ebitenutil.DebugPrintAt(screen, "Camera:", xMaxPos+2*configuration.Global.TileSize, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("(", g.camera.X, ",", g.camera.Y, ")"), xMaxPos+2*configuration.Global.TileSize+configuration.Global.TileSize/2, ySpace)

	ebitenutil.DebugPrintAt(screen, "Character:", xMaxPos+2*configuration.Global.TileSize, 3*ySpace)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("(", g.character.X, ",", g.character.Y, ")"), xMaxPos+2*configuration.Global.TileSize+configuration.Global.TileSize/2, 4*ySpace)

	ebitenutil.DebugPrintAt(screen, "Zoom:", xMaxPos+2*configuration.Global.TileSize, 5*ySpace)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("(", configuration.Global.NumTileX, "/", configuration.Global.MaxZoom, ")"), xMaxPos+2*configuration.Global.TileSize+configuration.Global.TileSize/2, 6*ySpace)
}
