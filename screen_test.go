package remotescreen

import (
	"image/color"
	"os"
	"testing"
)

func TestScreen_Line(t *testing.T) {

	f, err := os.Create("./test_data/line.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	s := NewScreen(400, 400)
	err = s.Clear().Line(100, 100, 200, 200, color.White).CaptureTo(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("success")

}

func TestScreen_Rect(t *testing.T) {

	f, err := os.Create("./test_data/rect.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	s := NewScreen(400, 400)
	err = s.Clear().Rect(100, 100, 200, 200, color.White).CaptureTo(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("success")

}

func TestScreen_RectFill(t *testing.T) {

	f, err := os.Create("./test_data/rect_fill.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	s := NewScreen(400, 400)
	err = s.ClearWithColor(color.White).RectFill(100, 100, 200, 200, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}).CaptureTo(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("success")

}

func TestScreen_DrawText(t *testing.T) {

	f, err := os.Create("./test_data/text.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	s := NewScreen(400, 400)
	err = s.ClearWithColor(color.White).DrawText("你好", 100, 100, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}).CaptureTo(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("success")

}
