package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisations car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.floor.Init(&g.tileOverlays)

	g.character.Init(g.floor.GetWidthContent(), g.floor.GetHeightContent())
	g.camera.Init(g.character.X, g.character.Y)
	g.floor.Update(g.character.X, g.character.Y)

	if configuration.Global.ExtBlockingBlocks {

		// Pour faire fonctionner l'extension, il faut que le terrain soit correctement initialisé

		// On cherche si on peut trouver une position de spawn non blocante pour le personnage
		ok, posX, posY := g.floor.FindSpawn(g.character.X, g.character.Y, g.camera.X, g.camera.Y, make(map[[2]int]bool, 0))
		if ok {
			g.character.X = posX
			g.character.Y = posY
		}

	}

	// Par défaut, on considère que le joueur vient d'être téléporté
	// Afin d'éviter d'être téléporté en posant un portail
	g.justTeleported = true

}
