package shape

import (
	"github.com/AkshachRd/ood-2023/labs/01/shapes/gfx"
)

type IShapeBehavior interface {
	Move(dx, dy float64)
	GetShapeInfo() string
	GetShapeName() string
	Draw(canvas gfx.ICanvas)
}
