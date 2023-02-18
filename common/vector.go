package common

type Vector struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

func (vector *Vector) Add(other Vector) *Vector {
	vector.X += other.X
	vector.Y += other.Y
	return vector
}

func (vector *Vector) Equals(other Vector) bool {
	return vector.X == other.X && vector.Y == other.Y
}

func NORTH() Vector {
	return Vector{0, 1}
}

func EAST() Vector {
	return Vector{1, 0}
}

func SOUTH() Vector {
	return Vector{0, -1}
}

func WEST() Vector {
	return Vector{-1, 0}
}
