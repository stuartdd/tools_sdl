// Collection
package objects

import (
	"tools_sdl/interfaces"
)

type Collection struct {
	Objects []interfaces.Drawable
	Enabled bool
}

func NewCollection() *Collection {
	return &Collection{Objects: make([]interfaces.Drawable, 0), Enabled: true}
}

func (p *Collection) Add(obj interfaces.Drawable) {
	p.Objects = append(p.Objects, obj)
}

func (p *Collection) UpdateDraw(currentTime float64) {
	if p.Enabled {
		for i := 0; i < len(p.Objects); i++ {
			(p.Objects[i]).Update(currentTime)
			(p.Objects[i]).Draw()
		}
	} else {
		p.Update(currentTime)
	}
}

func (p *Collection) PointInside(x float64, y float64) bool {
	if p.Enabled {
		for i := 0; i < len(p.Objects); i++ {
			if (p.Objects[i]).PointInside(x, y) {
				return true
			}
		}
	}
	return false
}

func (p *Collection) PointInsideBounds(x float64, y float64) bool {
	if p.Enabled {
		for i := 0; i < len(p.Objects); i++ {
			if (p.Objects[i]).PointInsideBounds(x, y) {
				return true
			}
		}
	}
	return false
}

func (p *Collection) Draw() {
	if p.Enabled {
		for i := 0; i < len(p.Objects); i++ {
			(p.Objects[i]).Draw()
		}
	}
}

func (p *Collection) Update(currentTime float64) {
	for i := 0; i < len(p.Objects); i++ {
		(p.Objects[i]).Update(currentTime)
	}
}
