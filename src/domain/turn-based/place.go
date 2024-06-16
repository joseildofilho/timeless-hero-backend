package turnbased

type Place struct {
	things []any
}

func NewPlace() *Place {
	return &Place {
		things: []any{},
	}
}

func (p *Place) AddToPlace(thing any) {
	p.things = append(p.things, thing)
}
