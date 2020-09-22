package main

import "github.com/xshifty/go-labs/graph"

func main() {
	g, err := graph.New("Sample GO", 800, 600, false)
	if err != nil {
		panic(err)
	}

	t, err := graph.CreateTextureFromFile("./assets/full_moon.png", "full-moon", 640, 480)
	if err != nil {
		panic(err)
	}

	if err := g.AppendFigure(t); err != nil {
		panic(err)
	}

	if err := g.Run(); err != nil {
		panic(err)
	}
}
