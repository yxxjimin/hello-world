package types

// Non-exported fields can only be accessed or initialized within the same
// package.
type Vertex struct {
	PublicX     int
	PublicY     int
	privateName string
}
