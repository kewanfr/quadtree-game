package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Layout détermine la taille de l'image sur laquelle Ebitengine
// affiche les images du jeu en fonction de la taille de la fenêtre
// dans laquelle l'affichage a lieu. Pour le moment, cette taille
// ne dépend pas de celle de la fenêtre.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth = configuration.Global.ScreenWidth
	screenHeight = configuration.Global.ScreenHeight

	if g.CurrentState == 1 {
		// Dans le cas où on est dans le menu principal, on ajuste la taille de l'écran à celle de l'image du menu
		screenWidth = assets.TitleImage.Bounds().Dx()
		screenHeight = assets.TitleImage.Bounds().Dy()
	} else if configuration.Global.DebugMode {
		screenWidth += configuration.Global.NumTileForDebug * configuration.Global.TileSize
		screenHeight += configuration.Global.TileSize
	}
	return
}
