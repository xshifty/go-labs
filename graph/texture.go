package graph

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type texture struct {
	id                   string
	x, y                 int32
	sw, sh               int32
	tw, th               int32
	handle, target, unit uint32
}

func CreateTextureFromFile(file, id string, tw, th int32) (*texture, error) {
	infile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %s", file, err)
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		return nil, fmt.Errorf("cannot decode image: %s", err)
	}

	b := img.Bounds()
	rgba := image.NewRGBA(b)
	t := texture{id, 0, 0, int32(b.Max.X), int32(b.Max.Y), tw, th, 0, gl.TEXTURE_2D, gl.TEXTURE0}

	gl.GenTextures(1, &t.handle)
	gl.ActiveTexture(t.unit)
	gl.BindTexture(t.target, t.handle)

	defer func(t *texture) {
		t.unit = 0
		gl.BindTexture(t.target, 0)
	}(&t)

	gl.TexParameteri(t.target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(t.target, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(t.target, 0, gl.SRGB_ALPHA, tw, th, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(t.handle)

	return &t, nil
}

func (t *texture) Id() string {
	return t.id
}

func (t *texture) Render(rend Renderer) error {
	return nil
}

func (t *texture) Destroy() error {
	return nil
}
