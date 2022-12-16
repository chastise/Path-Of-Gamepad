package images

import (
	_ "embed"
)

var (
	//go:embed buttonA.png
	ButtonA_png []byte

	//go:embed buttonB.png
	ButtonB_png []byte

	//go:embed buttonX.png
	ButtonX_png []byte

	//go:embed buttonY.png
	ButtonY_png []byte

	//go:embed crosshair.png
	Crosshair_png []byte
)
