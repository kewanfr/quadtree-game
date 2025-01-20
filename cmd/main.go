package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/game"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
)

func main() {

	var configFileName string
	flag.StringVar(&configFileName, "config", "config.json", "select configuration file")

	flag.Parse()

	configuration.Load(configFileName)

	assets.Load()

	g := &game.Game{}

	//Menu principal activé selon la config
	if configuration.Global.ExtTitleScreen {
		g.CurrentState = 1
	}

	g.Init()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("Quadtree")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
