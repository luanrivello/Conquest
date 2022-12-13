package spacetime

/*
* SOLAR SYSTEM STRUCTURE
 */
type SolarSystem struct {
	name   string
	sun    *Sun
	planet *Planet
}

// CONSTRUCTOR
func NewSolarSystem() *SolarSystem {
	bornsun := NewSun()
	bornplanet := NewPlanet(16)

	return &SolarSystem{name: "Origin", sun: bornsun, planet: bornplanet}
}

// GET
func (s *SolarSystem) GetSun() *Sun {
	return s.sun
}

func (s *SolarSystem) GetPlanet() *Planet {
	return s.planet
}

func (s *SolarSystem) GetName() string {
	return s.name
}

// SET
func (s *SolarSystem) SetName(name string) {
	s.name = name
}
