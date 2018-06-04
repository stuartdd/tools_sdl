// Debug
package objects

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var debug_lines_printed int32 = 0

func Print(str string) {
	fmt.Printf("[%d] %s\n", debug_lines_printed, str)
	debug_lines_printed++
}

func Printf(format string, things interface{}) {
	fmt.Printf("[%d]"+format+"\n", debug_lines_printed, things)
	debug_lines_printed++
}

func PrintRect(title string, r *sdl.Rect) {
	fmt.Printf("[%d] %s Rect X:%d Y:%d W:%d H:%d\n", debug_lines_printed, title, r.X, r.Y, r.W, r.H)
	debug_lines_printed++
}
