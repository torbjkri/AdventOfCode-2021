package dec11

type Octopus struct {
	energy_ int
}

func CreateOctopus(e int) *Octopus {
	o := new(Octopus)
	o.energy_ = e
	return o
}

func (o *Octopus) CheckFlash() bool {

	if o.energy_ > 9 {
		o.energy_ = 0
		return true
	}
	return false
}

func (o *Octopus) HasFlashed() bool {
	return o.energy_ == 0
}

func (o *Octopus) Update() {
	o.energy_ += 1
}

func (o *Octopus) Pump() {
	if o.energy_ > 0 {
		o.energy_ += 1
	}
}
