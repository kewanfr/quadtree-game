# project-quadtree

Projet de SAÉ Initiation au développement (R1.01) et de SA2 implémentation d'un besoin client (SAE1.01).

BERNIER Kéwan et GALLY Chianti

## Répartition des tâches

### Kéwan BERNIER

- [x] Fonction [readFloorFromFile](./floor/init.go#readFloorFromFile)
- [X] Fonction [GetContent](./quadtree/get.go)
- [X] Extension 6.1 Génération de terrain aléatoire (fichiers: [floor/generate.go](./floor/generate.go), [floor/init.go](./floor/init.go))
- [X] Extension 6.2 Enregistrement de Terrain (fichiers: [floor/save.go](./floor/save.go))
- [X] Extension 6.3 Animation des décors (implémentation pour les tuiles) (fichiers: [floor/draw.ho](./floor/draw.go))
- [X] Extension 6.4 Teleportation (fichiers: [game/teleportation.go](./game/teleportation.go), [game/update.go](./game/update.go), [game/draw.go](./game/draw.go))
- [X] Extension 6.5 Gestion des sols bloqués - "interdiction de marcher sur l'eau" (fichiers: [floor/blocking.go](./floor/blocking.go))

- [X] Extension Bonus (bug as a feature): ExtSpeedRun: speed x2 and bypass the blocking floor (fichiers: [game/update.go](./game/update.go))

### Chianti GALLY

- [x] Fonction [updateFromFileFloor](./floor/update.go#updateFromFileFloor)
- [x] Fonction [MakeFromArray](./quadtree/make.go)
- [X] Extension 6.10 Particules ([paquet particles](/particles), [game/update.go](./game/update.go), [game/draw.go](./game/draw.go), [character/update.go](./character/update.go))
- [X] Extension 6.12 Zoom (fichiers: ([game/update.go](./game/update.go))
- [X] Extension 6.3 Animation des décors (implémentation des surcouches pour des fleurs, buisson, etc.) (fichiers: [floor/init.go](./floor/init.go), [paquet flooroverlay](./flooroverlay))

- [X] Extension Bonus : Menu de démarrage (fichiers: [cmd/main.go](./cmd/main.go), [game/game.go](./game/game.go), [game/update.go](./game/update.go), [game/draw.go](./game/draw.go), [game/draw.go](./game/draw.go))