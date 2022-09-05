package spacetime

/*
* GALAXY STRUCTURE
 */
type Galaxy struct {
	name string
	systems []*SolarSystem
}

func NewGalaxy() *Galaxy {
	galaxy := &Galaxy{}
	galaxy.newSystem()
	galaxy.name = "Milkdromeda"
	return galaxy
}

func (g *Galaxy) newSystem() {
	g.systems = append(g.systems, NewSolarSystem())
}

func (gal *Galaxy) GetSystem() *SolarSystem {
	if len(gal.systems) == 0 {
		return nil
	}

	return gal.systems[0]
}
