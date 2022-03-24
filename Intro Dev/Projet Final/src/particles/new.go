package particles

import (
	"project-particles/config"
	"math/rand"
	"time"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {

	
	var MyParticle []Particle
	rand.Seed(time.Now().UTC().UnixNano())

	for i:= 0 ; i < config.General.InitNumParticles ; i++ {
	
		MyParticle = append(MyParticle, NewParticle())
	}
	
	return System{Content: MyParticle}
}

func NewParticle() Particle {
	
	var Random bool = config.General.RandomSpawn
	var particle Particle
	rand.Seed(time.Now().UTC().UnixNano())
	var posx , posy float64
	var vx,vy float64

	if Random {
		posx,posy = rand.Float64()*float64(config.General.WindowSizeX),rand.Float64()*float64(config.General.WindowSizeY)
	} else {
		posx,posy = float64(config.General.SpawnX),float64(config.General.SpawnY)
	}
	if config.General.SlowSpeed {
		vx,vy = (rand.Float64()*5) - 2.5 , (rand.Float64()*5) - 2.5
	}
	if config.General.NormalSpeed {
		vx,vy = (rand.Float64()*10) - 5 , (rand.Float64()*10) - 5
	}
	if config.General.HighSpeed {
		vx,vy = (rand.Float64()*15) - 7.5 , (rand.Float64()*15) - 7.5
	}
	
	
	particle = Particle{
		PositionX: posx,
		PositionY: posy,
		ScaleX: 1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1,
		VitesseX: vx,
		VitesseY: vy,
	}
	
	return particle
}