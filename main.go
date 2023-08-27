package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game
const ImageHeight int = 500
const ImageWidth int = 500

// Particle
var NbParticles int = 2000
var ParticleSpeed float64 = 2
var SensorDistance float64 = 30
var SensorAngle float64 = 0.2       // value should be [0, 0.5]
var ParticleWiggle float64 = 0.0000 // [0.0, 0.5]

// Pheromone
const DiffusionRate float64 = 1   // [-1.0, 1.0]
const DecayRate float64 = -0.05   // must be negative neer 0
var PheromoneDeposit float64 = .9 // [0, 1]

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
