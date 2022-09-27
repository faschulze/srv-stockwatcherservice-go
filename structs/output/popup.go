package output

import "image/color"

// PopUpWindowContent struct
type PopUpWindowContent struct {
	Title   string
	Content string
	Color   color.Color
	Height  uint
	Width   uint
}
