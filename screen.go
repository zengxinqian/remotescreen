package remotescreen

import (
	"image"
	"image/color"
	"image/png"
	"io"

	"gopkg.in/fogleman/gg.v1"
)

type Screen struct {
	//Screen Size
	Width  int
	Height int
	//Updated Rect
	updatedRect image.Rectangle
	//FrameBuffer
	fb      *image.RGBA
	drawCtx *gg.Context
}

func NewScreen(width, height int) *Screen {

	fb := image.NewRGBA(image.Rect(0, 0, width, height))
	return &Screen{
		Width:   width,
		Height:  height,
		fb:      fb,
		drawCtx: gg.NewContextForRGBA(fb),
	}

}

func (s *Screen) CaptureTo(w io.Writer) error {

	return png.Encode(w, s.fb)

}

func (s *Screen) Clear() *Screen {
	return s.ClearWithColor(color.Black)
}

func (s *Screen) ClearWithColor(c color.Color) *Screen {

	s.drawCtx.SetColor(c)
	s.drawCtx.Clear()
	return s

}

func (s *Screen) Line(x0, y0, x1, y1 int, c color.Color) *Screen {

	s.drawCtx.DrawLine(float64(x0), float64(y0), float64(x1), float64(y1))
	s.drawCtx.SetColor(c)
	s.drawCtx.Stroke()
	return s

}

func (s *Screen) Rect(x0, y0, x1, y1 int, c color.Color) *Screen {

	if x0 > x1 {
		x0, x1 = x1, x0
	}

	if y0 > y1 {
		y0, y1 = y1, y0
	}

	s.drawCtx.DrawRectangle(float64(x0), float64(y0), float64(x1-x0), float64(y1-y0))
	s.drawCtx.SetColor(c)
	s.drawCtx.Stroke()

	return s

}

func (s *Screen) RectFill(x0, y0, x1, y1 int, c color.Color) *Screen {

	if x0 > x1 {
		x0, x1 = x1, x0
	}

	if y0 > y1 {
		y0, y1 = y1, y0
	}

	s.drawCtx.DrawRectangle(float64(x0), float64(y0), float64(x1-x0), float64(y1-y0))
	s.drawCtx.SetColor(c)
	s.drawCtx.Fill()

	return s

}

func (s *Screen) DrawText(text string, x, y int, c color.Color) *Screen {

	if err := s.drawCtx.LoadFontFace("/Library/Fonts/Arial Unicode.ttf", 16); err != nil {
		panic(err)
	}

	s.drawCtx.SetColor(c)
	s.drawCtx.DrawStringAnchored(text, float64(x), float64(y), 0.5, 0.5)

	return s

}
