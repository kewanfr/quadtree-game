package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/camera"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/character"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/flooroverlay"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/particles"
)

// Game est le type permettant de représenter les données du jeu.
// Aucun champs n'est exporté pour le moment.
//
// Les champs non exportés sont :
//   - camera : la représentation de la caméra
//   - floor : la représentation du terrain
//   - character : la représentation du personnage

type Portal struct {
	X, Y int
}

type Game struct {
	camera         camera.Camera
	floor          floor.Floor
	character      character.Character
	Portals        []Portal
	justTeleported bool
	particles      []particles.Particle
	tileOverlays   []flooroverlay.TileOverlay

	message       string // Message à afficher
	messageFrames int    // Nombre de frames pendant lesquelles le message est affiché
	CurrentState  int    // Menu/State à afficher (0 : Jeu, 1 : Menu Principal)
}
