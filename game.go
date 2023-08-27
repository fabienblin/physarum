package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Game *GameInstance
)

type GameInstance struct {
	ParticleImage  *ebiten.Image
	Particles      []*Particle
	PheromoneImage *ebiten.Image
}

func initGame() *GameInstance {
	Game = &GameInstance{
		ParticleImage:  ebiten.NewImage(ImageWidth, ImageHeight),
		Particles:      ParticleFactory(NbParticles),
		PheromoneImage: ebiten.NewImage(ImageWidth, ImageHeight),
	}

	return Game
}

func (g *GameInstance) Update() error {
	return nil
}

func (g *GameInstance) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)

	// g.ParticleImage.Fill(color.Transparent)
	for _, p := range g.Particles {
		p.MoveParticle()
		// p.DrawParticle()
	}
	DecayPheromones()

	// fuse images
	// g.PheromoneImage.DrawImage(g.ParticleImage, nil)

	// Draw the updated pheromone grid
	screen.DrawImage(g.PheromoneImage, op)
}

func (g *GameInstance) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ImageWidth, ImageHeight
}
