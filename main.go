package main

import (
	"fmt"
	i3 "github.com/johnae/go-i3"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	t, err := i3.GetTree()
	panicIf(err)
	a := getActiveWindow(t.Root.Nodes)

	fmt.Printf("%v,%v %vx%v\n",
		a.Rect.X,
		a.Rect.Y-a.DecoRect.Height,
		a.Rect.Width,
		a.Rect.Height+a.DecoRect.Height,
	)
}

func getActiveWindow(nodes []*i3.Node) *i3.Node {
	for _, n := range nodes {
		if len(n.Nodes) > 0 {
			active := getActiveWindow(n.Nodes)
			if active != nil {
				return active
			}
		}

		if n.Focused {
			return n
		}
	}

	return nil
}
