package compass

var (
	North = Vector{0, 1}
	East  = North.Rotate(90)
	South = North.Rotate(180)
	West  = North.Rotate(270)

	Zero = Vector{0, 0}
)

type Vector struct {
	X int
	Y int
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y}
}

func (v Vector) Scale(mul int) Vector {
	return Vector{v.X * mul, v.Y * mul}
}

func (v Vector) Rotate(angle int) Vector {
	switch mod(angle, 360) {
	case 90:
		return Vector{v.Y, -v.X}
	case 180:
		return Vector{-v.X, -v.Y}
	case 270:
		return Vector{-v.Y, v.X}
	}
	return v
}

func (v Vector) ManhattanDistance() int {
	return abs(v.X) + abs(v.Y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func mod(a, b int) int {
	return (a%b + b) % b
}
