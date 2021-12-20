package dec17

type Probe struct {
	x_     int
	y_     int
	vel_x_ int
	vel_y_ int
}

func (p *Probe) Update() {
	p.x_ += p.vel_x_
	p.y_ += p.vel_y_

	if p.vel_x_ < 0 {
		p.vel_x_ += 1
	} else if p.vel_x_ > 0 {
		p.vel_x_ -= 1
	}

	p.vel_y_ -= 1
}

func (p *Probe) ShouldStop(target Target) bool {
	if p.y_ < target.y_min_ && p.vel_y_ < 0 {
		return true
	}
	if p.x_ > target.x_max_ && p.vel_x_ > 0 {
		return true
	}

	return false
}
