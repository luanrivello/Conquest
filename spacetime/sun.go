package spacetime

/*
* SOLAR SYSTEM STRUCTURE
 */
type Sun struct {
	name string
	size int
	//?megastrutcture &Megastructure
}

// CONSTRUCTOR
func NewSun() *Sun {
	return &Sun{name: "Sol", size: 1}
}

// GET
func (s *Sun) GetName() string {
	return s.name
}

func (s *Sun) GetSize() int {
	return s.size
}

