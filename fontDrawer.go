package fontDrawer

import (
	"image"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Config struct {
	FontPath string
	Size     float64
	Width    int
	Height   int
}

type Drawer struct {
	Drawer  *font.Drawer
	font    *truetype.Font
	img     *image.RGBA
	content []byte
}

// NewDrawer Create new *Drawer struct.
func NewDrawer(c *Config) (*Drawer, error) {
	readFont, _ := ioutil.ReadFile(c.FontPath)
	f, err := truetype.Parse(readFont)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(f, &truetype.Options{
		Size:    c.Size,
		Hinting: font.HintingVertical,
	})

	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))

	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: face,
	}

	return &Drawer{
		font:    f,
		img:     img,
		Drawer:  drawer,
		content: []byte{},
	}, nil
}

// Draw Drawing content on an image.
func (d *Drawer) Draw() *image.RGBA {
	d.Drawer.DrawBytes(d.content)
	return d.img
}

// Bounds Return the Drawer.BoundBytes.
func (d *Drawer) Bounds() (fixed.Rectangle26_6, fixed.Int26_6) {
	return d.Drawer.BoundBytes(d.content)
}

// Measure Return the Drawer.MeasureBytes.
func (d *Drawer) Measure() fixed.Int26_6 {
	return d.Drawer.MeasureBytes(d.content)
}

// SetContent Append to the content.
func (d *Drawer) SetContent(b []byte) {
	d.content = append(d.content, b...)
}

// CenterX Return the computed center from the content.
func (d *Drawer) CenterX() fixed.Int26_6 {
	return (fixed.I(d.img.Bounds().Max.X) - d.Measure()) / fixed.I(2)
}

// ChageFontOptions Changing Size and Hinting of the font.
func (d *Drawer) ChageFontOptions(size float64, hinting *font.Hinting) {
	d.Drawer.Face = truetype.NewFace(d.font, &truetype.Options{
		Size:    size,
		Hinting: *hinting,
	})
}

// ChageFaceColor Change the face color.
func (d *Drawer) ChageFaceColor(uni *image.Uniform) {
	d.Drawer.Src = uni
}

// SetPosition Set the font start position.
func (d *Drawer) SetPosition(x, y fixed.Int26_6) {
	d.Drawer.Dot.X = x
	d.Drawer.Dot.Y = y
}

// ClearContent clear the content.
func (d *Drawer) ClearContent() {
	d.content = []byte{}
}

// ClearImg Clear Only the image.
func (d *Drawer) ClearImg() {
	maxW := d.img.Bounds().Max.X
	maxH := d.img.Bounds().Max.Y

	d.img = image.NewRGBA(image.Rect(0, 0, maxW, maxH))
}

// ClearAll Clear the content and image.
func (d *Drawer) ClearAll() {
	d.ClearContent()
	d.ClearImg()
}
