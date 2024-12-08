package utils

type Coordinate struct {
	X, Y int
}

type CoordinateSet map[Coordinate]struct{}

func NewCoordinateSet() CoordinateSet {
	return make(CoordinateSet)
}

func (cs CoordinateSet) Add(coord Coordinate) {
	cs[coord] = struct{}{}
}

func (cs CoordinateSet) Contains(coord Coordinate) bool {
	_, exists := cs[coord]
	return exists
}

func (cs CoordinateSet) Size() int {
	return len(cs)
}

type State struct {
	X, Y      int
	Direction int
}

type StateSet map[State]struct{}

func NewStateSet() StateSet {
	return make(StateSet)
}

func (cs StateSet) Add(coord State) {
	cs[coord] = struct{}{}
}

func (cs StateSet) Contains(coord State) bool {
	_, exists := cs[coord]
	return exists
}

func (cs StateSet) Size() int {
	return len(cs)
}
