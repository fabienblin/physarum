package main

import (
	"image/color"
	"math"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blur"
	"github.com/hajimehoshi/ebiten/v2"
)

func addPheromoneAlongPath(pos1, pos2 Pos) {
	dx := pos2.x - pos1.x
	dy := pos2.y - pos1.y

	steps := int(math.Max(math.Abs(dx), math.Abs(dy)))

	if steps == 0 {
		// If the positions are the same, just add pheromone to the single pixel
		addPheromoneToPixel(int(pos1.x), int(pos1.y))
		return
	}

	xStep := dx / float64(steps)
	yStep := dy / float64(steps)

	for i := 0; i <= steps; i++ {
		x := pos1.x + float64(i)*xStep
		y := pos1.y + float64(i)*yStep
		addPheromoneToPixel(int(x), int(y))
	}
}

func addPheromoneToPixel(x, y int) {
	if x >= 0 && y >= 0 && x < ImageWidth && y < ImageHeight {
		Game.PheromoneImage.Set(x, y, color.RGBA{0, 0, 255 * PheromoneDeposit, 255})
	}
}

func DecayPheromones() {
	Game.PheromoneImage = ebiten.NewImageFromImage(adjust.Brightness(blur.Gaussian(Game.PheromoneImage, DiffusionRate), DecayRate))
}
