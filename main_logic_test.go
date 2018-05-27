// TestInsides
package main

import (
	"fmt"
	"testing"
	"time"
	"tools_sdl/interfaces"
	"tools_sdl/objects"
	"tools_sdl/tools"
)

const ITERATIONS int = 100000000

func TestInsideOutsideBounds(t *testing.T) {
	tri1 := objects.NewTriangle("t1", -50, -50, 0, 50, 50, -50, 400, 300, tools.GetColour("Coral Blue"), true)
	isInsideBounds(t, tri1, -50, 50)
	isInsideBounds(t, tri1, 0, 50)
	isInsideBounds(t, tri1, 50, -50)
	isInsideBounds(t, tri1, 50, 50)

	isOutsideBounds(t, tri1, -51, -51)
	isOutsideBounds(t, tri1, 0, 51)
	isOutsideBounds(t, tri1, 51, -51)
	isOutsideBounds(t, tri1, 51, 51)
}

func isOutsideBounds(t *testing.T, shape interfaces.Drawable, x float64, y float64) {
	if shape.PointInsideBounds(400+x, 300+y) {
		t.Errorf("Test failed x:%f, y:%f are inside", 400+x, 300+y)
	}
}

func isInsideBounds(t *testing.T, shape interfaces.Drawable, x float64, y float64) {
	if !shape.PointInsideBounds(400+x, 300+y) {
		t.Errorf("Test failed x:%f, y:%f are outside", 400+x, 300+y)
	}
}

func TestSumsMin(t *testing.T) {
	if tools.Min(0, 9, 10) != 0 {
		t.Errorf("Min failed. Should return 0")
	}
	if tools.Min(9, 0, 10) != 0 {
		t.Errorf("Min failed. Should return 0")
	}
	if tools.Min(9, 10, 0) != 0 {
		t.Errorf("Min failed. Should return 0")
	}
	if tools.Min(9, -10, 0) != -10 {
		t.Errorf("Min failed. Should return -10")
	}
}

func TestSumsMax(t *testing.T) {
	if tools.Max(0, 9, 10) != 10 {
		t.Errorf("Max failed. Should return 10")
	}
	if tools.Max(9, 10, 0) != 10 {
		t.Errorf("Max failed. Should return 10")
	}
	if tools.Max(10, 9, 0) != 10 {
		t.Errorf("Max failed. Should return 10")
	}
	if tools.Max(9, -10, 0) != 9 {
		t.Errorf("Max failed. Should return 10")
	}
}

func TestPointInsideBounds(t *testing.T) {
	shape := objects.NewTriangle("t1", -50, -50, 0, 51, 50, -50, 400, 300, tools.GetColour("Coral Blue"), true)
	timeTemp := time.Now().UnixNano()

	for i := 0; i < ITERATIONS; i++ {
		shape.PointInsideBounds(400, 300)
	}

	time := time.Now().UnixNano() - timeTemp
	fmt.Printf("NS: PointInsideBounds seconds: %f. (7.470014) %f nano seconds each", float64(time)/NANO_PER_SECOND, float64(time)/float64(ITERATIONS))
}

func TestPointInside(t *testing.T) {
	shape := objects.NewTriangle("t1", -50, -50, 0, 51, 50, -50, 400, 300, tools.GetColour("Coral Blue"), true)
	timeTemp := time.Now().UnixNano()

	for i := 0; i < ITERATIONS; i++ {
		shape.PointInside(400, 300)
	}

	time := time.Now().UnixNano() - timeTemp
	fmt.Printf("NS: PointInside seconds: %f. (12.167762) %f nano seconds each", float64(time)/NANO_PER_SECOND, float64(time)/float64(ITERATIONS))
}
