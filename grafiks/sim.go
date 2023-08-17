package main

type Particle struct {
	position FloatVector
	velocity FloatVector
}

type Sim struct {
	particles []*Particle
}

func makeSim() *Sim {
	particles := make([]*Particle, 0)
	p1 := Particle{
		position: FloatVector{
			X: 5.0,
			Y: 5.0,
		},
		velocity: FloatVector{
			X: 0.5,
			Y: 0.5,
		},
	}

	particles = append(particles, &p1)

	sim := Sim{
		particles,
	}

	return &sim
}

func (sim *Sim) step(t float32) {
	if sim == nil {
		return
	}

	for i := range sim.particles {
		particle := sim.particles[i]
		particle.position.X += particle.velocity.X * t
		particle.position.Y += particle.velocity.Y * t
	}
}
