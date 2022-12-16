package overlay

import (
	"bytes"
	_ "embed"
	"image"

	//"image/color"
	_ "image/png"
	"log"

	"github.com/Steven-Ireland/path-of-gamepad/config"
	"github.com/Steven-Ireland/path-of-gamepad/images"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	overlayImgA     *ebiten.Image
	overlayImgB     *ebiten.Image
	overlayImgX     *ebiten.Image
	overlayImgY     *ebiten.Image
	overlayImgCross *ebiten.Image
)

func init() {
	initImage(images.ButtonA_png, overlayImgA)
	initImage(images.ButtonB_png, overlayImgB)
	initImage(images.ButtonX_png, overlayImgX)
	initImage(images.ButtonY_png, overlayImgY)
	initImage(images.Crosshair_png, overlayImgCross)
}

func initImage(someImage []byte, imageTarget *ebiten.Image) {
	img, _, err := image.Decode(bytes.NewReader(someImage))
	if err != nil {
		log.Fatal(err)
	}
	origOverlayImg := ebiten.NewImageFromImage(img)
	w, h := origOverlayImg.Size()
	imageTarget = ebiten.NewImage(w, h)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 1)
	imageTarget.DrawImage(origOverlayImg, op)
}

type Sprite struct {
	imageWidth  int
	imageHeight int
	x           int
	y           int
}

func loadSprite(img *ebiten.Image, xOffsetPercent float32, yOffsetPercent float32) *Sprite {
	w, h := img.Size()
	newSprite := &Sprite{
		imageWidth:  w,
		imageHeight: h,
		x:           int(xOffsetPercent * float32(config.ScreenWidth())),
		y:           int(yOffsetPercent * float32(config.ScreenHeight())),
	}

	return newSprite
}

func (s *Sprite) Update() {
	// Do literally nothing we love sprites but they don't move
}

type Sprites struct {
	sprites []*Sprite
	num     int
}

func (s *Sprites) Update() {
	for i := 0; i < s.num; i++ {
		s.sprites[i].Update()
	}
}

type Game struct {
	touchIDs []ebiten.TouchID
	sprites  Sprites
	op       ebiten.DrawImageOptions
	inited   bool
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	// numSpritesInOverlay := 5
	// g.sprites.sprites = make([]*Sprite, numSpritesInOverlay)
	// g.sprites.num = numSpritesInOverlay

	// g.sprites.sprites[0] = loadSprite(overlayImgCross, 0, 0)

	// g.sprites.sprites[1] = loadSprite(overlayImgX, 0.8215-0.02875*3, 0.97)
	// g.sprites.sprites[2] = loadSprite(overlayImgA, 0.8215-0.02875*2, 0.97)
	// g.sprites.sprites[3] = loadSprite(overlayImgB, 0.8215-0.02875*1, 0.97)
	// g.sprites.sprites[4] = loadSprite(overlayImgY, 0.8215, 0.97)

}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	//vector.DrawFilledCircle(screen, 400, 400, 100, color.RGBA{0x80, 0x00, 0x80, 0x80})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth(), config.ScreenHeight()
}

func StartOverlay() {
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowSize(config.ScreenWidth(), config.ScreenHeight())
	ebiten.SetWindowTitle("pog")
	ebiten.SetWindowResizable(false)
	ebiten.SetTPS(20)
	ebiten.SetScreenTransparent(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
