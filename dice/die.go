package dice

// Side represents the contents one of a die's faces.
type Side interface {
	// String returns the string representation of the side.
	String() string

	// Value returns the integer representation of the side.
	//
	// Value is used in sum calculations.
	Value() int
}

type Die interface {
	Name() string
	Roll() Side
	Sides() []Side
	Up() *int
	SetUp(*int)
}

/*
// NewDie creates an instance of DieType d
func NewDie(d DieType) *Die {
	return &Die{Type: d}
}

func (d *Die) Name() string {
	return d.Type.Name
}

// Roll sets d.up to a random side.
func (d *Die) Roll() Side {
	up := rand.Intn(len(d.Sides()))
	d.Up = &up
	return d.Type.Sides[*d.Up]
}

func (r *Die) Side() Side {
	if r.Up == nil {
		return nil
	}
	return r.Sides()[*r.Up]
}

func (d *Die) Sides() []Side {
	return d.Type.Sides
}

func (r *Die) String() string {
	return r.Side().String()
}

func (d *Die) Value() int {
	if d.Up == nil {
		return 0
	}
	return d.Type.Sides[*d.Up].Value()
}
*/
