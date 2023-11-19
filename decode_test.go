package svg

import (
	"bytes"
	_ "embed"
	"image"
	"image/jpeg"
	"io/ioutil"
	"testing"
)

//go:embed testdata/test.svg
var testSVG []byte

//go:embed testdata/test.jpg
var testJPEG []byte

func TestDecode(t *testing.T) {
	img, _, err := image.Decode(bytes.NewReader(testSVG))
	if err != nil {
		t.Fatal(err)
	}

	err = jpeg.Encode(ioutil.Discard, img, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeConfig(t *testing.T) {
	cfg, _, err := image.DecodeConfig(bytes.NewReader(testSVG))
	if err != nil {
		t.Fatal(err)
	}

	w, h := 620, 472

	if cfg.Width != w {
		t.Errorf("width: got %d, want %d", cfg.Width, w)
	}

	if cfg.Height != h {
		t.Errorf("height: got %d, want %d", cfg.Height, h)
	}
}

func BenchmarkDecodeJPEG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := image.Decode(bytes.NewReader(testJPEG))
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDecodeSVG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Decode(bytes.NewReader(testSVG))
		if err != nil {
			b.Error(err)
		}
	}
}
