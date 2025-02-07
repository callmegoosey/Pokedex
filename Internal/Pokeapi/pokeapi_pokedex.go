package pokeapi

// ===========================================
//
//	Returns the pokemon's data
//	streamlined version
//
// ===========================================

type PokedexData struct {
	Name   string
	Height int
	Weight int
	Stats  map[string]int
	Types  []string
}
