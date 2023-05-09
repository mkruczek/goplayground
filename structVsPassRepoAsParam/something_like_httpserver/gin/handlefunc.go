package gin

import (
	"context"
	"math"
)

type HandlerFunc func(ctx context.Context)

type Gin struct {
	// Fields
}

func (g *Gin) GET(relativePath string, handlers ...HandlerFunc) {
	// Code
}

func (g *Gin) Run() {
	//advance have math operation
	var e float64 = 1.0
	var factorial float64 = 1.0
	for i := 1; i <= 10; i++ {
		factorial *= float64(i)
		e += math.Pow(23, float64(i)) / factorial
	}
}
