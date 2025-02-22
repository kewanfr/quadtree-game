package assets

import (
	"bytes"
	"image"
	"log"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed floor.png
var floorBytes []byte

// FloorImage contient une version compatible avec Ebitengine de l'image
// qui contient les différents éléments qui peuvent s'afficher au sol
// (herbe, sable, etc).
// Dans la version du projet qui vous est fournie, ces éléments sont des
// carrés de 16 pixels de côté. Vous pourrez changer cela si vous le voulez.
var FloorImage *ebiten.Image

//go:embed character.png
var characterBytes []byte

// CharacterImage contient une version compatible avec Ebitengine de
// l'image qui contient les différentes étapes de l'animation du
// personnage.
// Dans la version du projet qui vous est fournie, ce personnage tient
// dans un carré de 16 pixels de côté. Vous pourrez changer cela si vous
// le voulez.
var CharacterImage *ebiten.Image

//go:embed particles/dust.png
var dustBytes []byte

//go:embed particles/dust_grass.png
var dustGrassBytes []byte

//go:embed particles/dust_sand.png
var dustSandBytes []byte

//go:embed particles/dust_wood.png
var dustWoodBytes []byte

var DustImage *ebiten.Image
var DustGrassImage *ebiten.Image
var DustSandImage *ebiten.Image
var DustWoodImage *ebiten.Image

//go:embed teleporter.png
var teleporterBytes []byte

var TeleporterImage *ebiten.Image

//go:embed teleporter_end.png
var teleporter_endBytes []byte

var Teleporter_endImage *ebiten.Image

//go:embed water.png
var waterBytes []byte

var WaterImage *ebiten.Image

//go:embed title.png
var titleBytes []byte

var TitleImage *ebiten.Image

// OVERLAY

//go:embed overlay/flower.png
var flowerBytes []byte

var FlowerImage *ebiten.Image

//go:embed overlay/herbe.png
var buissonBytes []byte

var BuissonImage *ebiten.Image

// Load est la fonction en charge de transformer, à l'exécution du programme,
// les images du jeu en structures de données compatibles avec Ebitengine.
// Ces structures de données sont stockées dans les variables définies ci-dessus.
func Load() {
	decoded, _, err := image.Decode(bytes.NewReader(floorBytes))
	if err != nil {
		log.Fatal(err)
	}
	FloorImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(characterBytes))
	if err != nil {
		log.Fatal(err)
	}
	CharacterImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(teleporterBytes))
	if err != nil {
		log.Fatal(err)
	}
	TeleporterImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(waterBytes))
	if err != nil {
		log.Fatal(err)
	}
	WaterImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(teleporter_endBytes))
	if err != nil {
		log.Fatal(err)
	}
	Teleporter_endImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(dustBytes))
	if err != nil {
		log.Fatal(err)
	}
	DustImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(dustGrassBytes))
	if err != nil {
		log.Fatal(err)
	}
	DustGrassImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(dustWoodBytes))
	if err != nil {
		log.Fatal(err)
	}
	DustWoodImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(dustSandBytes))
	if err != nil {
		log.Fatal(err)
	}
	DustSandImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(flowerBytes))
	if err != nil {
		log.Fatal(err)
	}
	FlowerImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(buissonBytes))
	if err != nil {
		log.Fatal(err)
	}
	BuissonImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(titleBytes))
	if err != nil {
		log.Fatal(err)
	}
	TitleImage = ebiten.NewImageFromImage(decoded)

}
