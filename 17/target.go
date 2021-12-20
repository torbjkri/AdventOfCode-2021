package dec17

type Target struct {
	x_min_ int
	x_max_ int
	y_min_ int
	y_max_ int
}

func (t *Target) IsInside(x, y int) bool {
	return (x >= t.x_min_ && x <= t.x_max_) && (y >= t.y_min_ && y <= t.y_max_)
}
