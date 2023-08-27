package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game
const ImageHeight int = 500
const ImageWidth int = 500

// Particle
var NbParticles int = 1000
var ParticleSpeed float64 = 10
var SensorDistance float64 = 20
var SensorAngle float64 = 0.1                   // value should be [0, 0.5]
var ParticleWiggle float64 = math.Pi * 2 * 0.05 // PI * 2 * [0.0, 2.0]

// Pheromone
const DiffusionRate float64 = 1 // [-1.0, 1.0]
const DecayRate float64 = -0.1  // must be negative neer 0
const PheromoneDeposit = 1

var PheromoneColor = color.RGBA{0, 0, 255, 255}

func init() {
	Game = initGame()
	ebiten.SetFullscreen(true)
	ebiten.SetTPS(10)
}

// https://cargocollective.com/sagejenson/physarum
func main() {
	if err := ebiten.RunGame(Game); err != nil {
		panic(err)
	}
}

type Pos struct {
	x, y float64
}
