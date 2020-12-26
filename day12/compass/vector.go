package compass

var (
	North = Vector{0, 1}
	East  = Vector{1, 0}
	South = Vector{0, -1}
	West  = Vector{-1, 0}

	Zero = Vector{0, 0}
)

type Vector struct {
	X int
	Y int
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y}
}

func (v Vector) Multiply(mul int) Vector {
	return Vector{v.X * mul, v.Y * mul}
}

func (v Vector) Turn(angle int) Vector {
	if mod(angle, 90) != 0 {
		return v
	}
	dir := angle / 90
	var directions [4]Vector = [4]Vector{North, East, South, West}
	var directionLookup = map[Vector]int{
		directions[0]: 0,
		directions[1]: 1,
		directions[2]: 2,
		directions[3]: 3,
	}
	curDirectionIdx := directionLookup[v]
	return directions[mod(curDirectionIdx+dir, 4)]
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
