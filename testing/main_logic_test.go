// TestInsides
package testing

import (
	"fmt"
	"testing"
	"time"
	"tools_sdl/interfaces"
	"tools_sdl/objects"
	"tools_sdl/structs"
	"tools_sdl/utils"
)

const ITERATIONS int = 100000000
const TEST_NANO_PER_SECOND float64 = 1000000000

func TestInsideOutsideBounds(t *testing.T) {
	world := &structs.World{Renderer: nil, X: 0, Y: 0}
	tri1 := objects.NewTriangle(world, -50, -50, 0, 50, 50, -50, 400, 300, utils.GetColour("Coral Blue"), true, true)
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
	if objects.Min3(0, 9, 10) != 0 {
		t.Errorf("Min failed. Should return 0")
	}
	if objects.Min3(9, 0, 10) != 0 {
		t.Errorf("Min failed. Should return 0")
	}
	if objects.Min3(9, 10, 0) != 0 {
		t.Errorf("Min failed. Should return 0")
	}
	if objects.Min3(9, -10, 0) != -10 {
		t.Errorf("Min failed. Should return -10")
	}
}

func TestSumsMax(t *testing.T) {
	if objects.Max3(0, 9, 10) != 10 {
		t.Errorf("Max failed. Should return 10")
	}
	if objects.Max3(9, 10, 0) != 10 {
		t.Errorf("Max failed. Should return 10")
	}
	if objects.Max3(10, 9, 0) != 10 {
		t.Errorf("Max failed. Should return 10")
	}
	if objects.Max3(9, -10, 0) != 9 {
		t.Errorf("Max failed. Should return 10")
	}
}

func TestPointInsideBounds(t *testing.T) {
	world := &structs.World{Renderer: nil, X: 0, Y: 0}
	shape := objects.NewTriangle(world, -50, -50, 0, 51, 50, -50, 400, 300, utils.GetColour("Coral Blue"), true, true)
	timeTemp := time.Now().UnixNano()

	for i := 0; i < ITERATIONS; i++ {
		shape.PointInsideBounds(400, 300)
	}

	time := time.Now().UnixNano() - timeTemp
	fmt.Printf("NS: PointInsideBounds seconds: %f. (7.470014) %f nano seconds each", float64(time)/TEST_NANO_PER_SECOND, float64(time)/float64(ITERATIONS))
}

func TestPointInside(t *testing.T) {
	world := &structs.World{Renderer: nil, X: 0, Y: 0}
	shape := objects.NewTriangle(world, -50, -50, 0, 51, 50, -50, 400, 300, utils.GetColour("Coral Blue"), true, true)
	timeTemp := time.Now().UnixNano()

	for i := 0; i < ITERATIONS; i++ {
		shape.PointInside(400, 300)
	}

	time := time.Now().UnixNano() - timeTemp
	fmt.Printf("NS: PointInside seconds: %f. (12.167762) %f nano seconds each", float64(time)/TEST_NANO_PER_SECOND, float64(time)/float64(ITERATIONS))
}
