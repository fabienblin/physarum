package main

import (
	"image/color"
	"math"
	"math/rand"
)

var Particles []*Particle

type Particle struct {
	pos            Pos
	direction      float64
	sensorAngle    float64
	sensorDistance float64
	sensors        [3]Pos
	spiecies       color.RGBA
}

func ParticleFactory(nbParticles int) []*Particle {
	Particles := make([]*Particle, nbParticles)

	for i := 0; i < nbParticles; i++ {
		Particles[i] = &Particle{
			pos:            Pos{x: rand.Float64() * float64(ImageWidth), y: rand.Float64() * float64(ImageHeight)},
			direction:      rand.Float64() * math.Pi * 2,
			sensorAngle:    math.Pi * SensorAngle,
			sensorDistance: SensorDistance,
			spiecies:       color.RGBA{255, 255, 255, 255},
		}
		Particles[i].RepositionSensors()
	}

	return Particles
}

func (p *Particle) RepositionSensors() {
	// Calculate the direction vector of the front sensor based on the particle's direction
	frontSensorAngle := p.direction
	frontSensorX := ((p.pos.x) + p.sensorDistance*math.Cos((frontSensorAngle)))
	frontSensorY := ((p.pos.y) + p.sensorDistance*math.Sin((frontSensorAngle)))
	p.sensors[1] = Pos{x: frontSensorX, y: frontSensorY}

	// Calculate the angle of the left sensor
	leftSensorAngle := (p.direction) + p.sensorAngle
	leftSensorX := ((p.pos.x) + p.sensorDistance*math.Cos(leftSensorAngle))
	leftSensorY := ((p.pos.y) + p.sensorDistance*math.Sin(leftSensorAngle))
	p.sensors[0] = Pos{x: leftSensorX, y: leftSensorY}

	// Calculate the angle of the right sensor
	rightSensorAngle := (p.direction) - p.sensorAngle
	rightSensorX := ((p.pos.x) + p.sensorDistance*math.Cos(rightSensorAngle))
	rightSensorY := ((p.pos.y) + p.sensorDistance*math.Sin(rightSensorAngle))
	p.sensors[2] = Pos{x: rightSensorX, y: rightSensorY}
}

func (p *Particle) DrawParticle() {
	Game.ParticleImage.Set(
		int(p.pos.x),
		int(p.pos.y),
		p.spiecies,
	)
	// Game.ParticleImage.Set(
	// 	int(p.sensors[0].x),
	// 	int(p.sensors[0].y),
	// 	color.RGBA{255, 0, 0, 255},
	// )
	// Game.ParticleImage.Set(
	// 	int(p.sensors[1].x),
	// 	int(p.sensors[1].y),
	// 	color.RGBA{0, 255, 0, 255},
	// )
	// Game.ParticleImage.Set(
	// 	int(p.sensors[2].x),
	// 	int(p.sensors[2].y),
	// 	color.RGBA{0, 0, 255, 255},
	// )
}

func (p *Particle) MoveParticle() {
	p.SensePheromone()
	dx := (ParticleSpeed * (math.Cos(float64(p.direction))))
	dy := (ParticleSpeed * (math.Sin(float64(p.direction))))

	nextX := dx + float64(p.pos.x)
	nextY := dy + float64(p.pos.y)

	for nextX < 0 || nextY < 0 || nextX >= float64(ImageWidth) || nextY >= float64(ImageHeight) {
		p.TurnParticle()

		dx := (ParticleSpeed * (math.Cos(float64(p.direction))))
		dy := (ParticleSpeed * (math.Sin(float64(p.direction))))

		nextX = dx + float64(p.pos.x)
		nextY = dy + float64(p.pos.y)
	}
	addPheromoneAlongPath(p.pos, Pos{(nextX), (nextY)})
	
	p.pos.x = (nextX)
	p.pos.y = (nextY)
	
	// wiggle
	p.direction += rand.Float64()* ParticleWiggle
	p.RepositionSensors()
}

func (p *Particle) TurnParticle() {
	p.direction = rand.Float64() * math.Pi * 2
}

func (p *Particle) SensePheromone() {
	var maxPheromone uint32 = 0
	bestDirection := p.direction // Initialize with current direction

	// Calculate the maximum pheromone value and find the best direction
	for i, sensor := range p.sensors {
		if sensor.x >= 0 && sensor.x < float64(ImageWidth) && sensor.y >= 0 && sensor.y < float64(ImageHeight) {
			_, _, pheromone, _ := Game.PheromoneImage.At(int(sensor.x), int(sensor.y)).RGBA()
			if pheromone > maxPheromone {
				maxPheromone = pheromone
				var angleOffset float64 = 0
				if i == 0 {
					angleOffset = p.sensorAngle
				} else if i == 2 {
					angleOffset = -p.sensorAngle
				}
				bestDirection = p.direction + angleOffset
			}
		}
	}

	// Update the particle's direction to turn toward the most pheromone-concentrated direction
	if maxPheromone != 0 {
		p.direction = bestDirection
	}
}
