package lib

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	stdWidth  = 100
	stdHeight = 40
	maxSkew   = 2
)

const (
	fontWidth  = 5
	fontHeight = 8
	blackChar  = 1
)

var font = [][]byte{
	{ // 0
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
	},
	{ // 1
		0, 0, 1, 0, 0,
		0, 1, 1, 0, 0,
		1, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		1, 1, 1, 1, 1,
	},
	{ // 2
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 1, 1,
		0, 1, 1, 0, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 1,
	},
	{ // 3
		1, 1, 1, 1, 0,
		0, 0, 0, 0, 1,
		0, 0, 0, 1, 0,
		0, 1, 1, 1, 0,
		0, 0, 0, 1, 0,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
	},
	{ // 4
		1, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 1, 1, 1, 1,
		0, 0, 0, 1, 0,
		0, 0, 0, 1, 0,
		0, 0, 0, 1, 0,
	},
	{ // 5
		1, 1, 1, 1, 1,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
	},
	{ // 6
		0, 0, 1, 1, 1,
		0, 1, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
	},
	{ // 7
		1, 1, 1, 1, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 1, 0,
		0, 0, 1, 0, 0,
		0, 1, 0, 0, 0,
		0, 1, 0, 0, 0,
		0, 1, 0, 0, 0,
	},
	{ // 8
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
	},
	{ // 9
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 1, 0, 0, 1,
		0, 1, 1, 1, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
	},
}

type Captcha struct {
	*image.NRGBA
	color   *color.NRGBA
	width   int //a digit width
	height  int //a digit height
	dotsize int
}

func init() {
	rand.Seed(int64(time.Second))
}

func NewImage(digits []int, width, height int) *Captcha {
	captcha := new(Captcha)
	r := image.Rect(captcha.width, captcha.height, stdWidth, stdHeight)
	captcha.NRGBA = image.NewNRGBA(r)
	captcha.color = &color.NRGBA{
		uint8(rand.Intn(129)),
		uint8(rand.Intn(129)),
		uint8(rand.Intn(129)),
		0xFF,
	}
	// Draw background (10 random circles of random brightness)
	captcha.calculateSizes(width, height, len(digits))
	captcha.fillWithCircles(10, captcha.dotsize)
	maxx := width - (captcha.width+captcha.dotsize)*len(digits) - captcha.dotsize
	maxy := height - captcha.height - captcha.dotsize*2
	x := rnd(captcha.dotsize*2, maxx)
	y := rnd(captcha.dotsize*2, maxy)
	// Draw digits.
	for _, n := range digits {
		captcha.drawDigit(font[n], x, y)
		x += captcha.width + captcha.dotsize
	}
	// Draw strike-through line.
	captcha.strikeThrough()
	return captcha
}

func (captcha *Captcha) WriteTo(w io.Writer) (int64, error) {
	return 0, png.Encode(w, captcha)
}

func (captcha *Captcha) calculateSizes(width, height, ncount int) {
	// Goal: fit all digits inside the image.
	var border int
	if width > height {
		border = height / 5
	} else {
		border = width / 5
	}
	// Convert everything to floats for calculations.
	w := float64(width - border*2)  //268
	h := float64(height - border*2) //48
	// fw takes into account 1-dot spacing between digits.
	fw := float64(fontWidth) + 1 //6
	fh := float64(fontHeight)    //8
	nc := float64(ncount)        //7
	// Calculate the width of a single digit taking into account only the
	// width of the image.
	nw := w / nc //38
	// Calculate the height of a digit from this width.
	nh := nw * fh / fw //51
	// Digit too high?
	if nh > h {
		// Fit digits based on height.
		nh = h //nh = 44
		nw = fw / fh * nh
	}
	// Calculate dot size.
	captcha.dotsize = int(nh / fh)
	// Save everything, making the actual width smaller by 1 dot to account
	// for spacing between digits.
	captcha.width = int(nw)
	captcha.height = int(nh) - captcha.dotsize
}

func (captcha *Captcha) fillWithCircles(n, maxradius int) {
	color := captcha.color
	maxx := captcha.Bounds().Max.X
	maxy := captcha.Bounds().Max.Y
	for i := 0; i < n; i++ {
		setRandomBrightness(color, 255)
		r := rnd(1, maxradius)
		captcha.drawCircle(color, rnd(r, maxx-r), rnd(r, maxy-r), r)
	}
}

func (captcha *Captcha) drawHorizLine(color color.Color, fromX, toX, y int) {
	for x := fromX; x <= toX; x++ {
		captcha.Set(x, y, color)
	}
}

func (captcha *Captcha) drawCircle(color color.Color, x, y, radius int) {
	f := 1 - radius
	dfx := 1
	dfy := -2 * radius
	xx := 0
	yy := radius
	captcha.Set(x, y+radius, color)
	captcha.Set(x, y-radius, color)
	captcha.drawHorizLine(color, x-radius, x+radius, y)
	for xx < yy {
		if f >= 0 {
			yy--
			dfy += 2
			f += dfy
		}
		xx++
		dfx += 2
		f += dfx
		captcha.drawHorizLine(color, x-xx, x+xx, y+yy)
		captcha.drawHorizLine(color, x-xx, x+xx, y-yy)
		captcha.drawHorizLine(color, x-yy, x+yy, y+xx)
		captcha.drawHorizLine(color, x-yy, x+yy, y-xx)
	}
}

func (captcha *Captcha) strikeThrough() {
	r := 0
	maxx := captcha.Bounds().Max.X
	maxy := captcha.Bounds().Max.Y
	y := rnd(maxy/3, maxy-maxy/3)
	for x := 0; x < maxx; x += r {
		r = rnd(1, captcha.dotsize/3)
		y += rnd(-captcha.dotsize/2, captcha.dotsize/2)
		if y <= 0 || y >= maxy {
			y = rnd(maxy/3, maxy-maxy/3)
		}
		captcha.drawCircle(captcha.color, x, y, r)
	}
}

func (captcha *Captcha) drawDigit(digit []byte, x, y int) {
	skf := rand.Float64() * float64(rnd(-maxSkew, maxSkew))
	xs := float64(x)
	minr := captcha.dotsize / 2                   // minumum radius
	maxr := captcha.dotsize/2 + captcha.dotsize/4 // maximum radius
	y += rnd(-minr, minr)
	for yy := 0; yy < fontHeight; yy++ {
		for xx := 0; xx < fontWidth; xx++ {
			if digit[yy*fontWidth+xx] != blackChar {
				continue
			}
			// Introduce random variations.
			or := rnd(minr, maxr)
			ox := x + (xx * captcha.dotsize) + rnd(0, or/2)
			oy := y + (yy * captcha.dotsize) + rnd(0, or/2)
			captcha.drawCircle(captcha.color, ox, oy, or)
		}
		xs += skf
		x = int(xs)
	}
}

func setRandomBrightness(c *color.NRGBA, max uint8) {
	minc := min3(c.R, c.G, c.B)
	maxc := max3(c.R, c.G, c.B)
	if maxc > max {
		return
	}
	n := rand.Intn(int(max-maxc)) - int(minc)
	c.R = uint8(int(c.R) + n)
	c.G = uint8(int(c.G) + n)
	c.B = uint8(int(c.B) + n)
}

func min3(x, y, z uint8) (o uint8) {
	o = x
	if y < o {
		o = y
	}
	if z < o {
		o = z
	}
	return
}

func max3(x, y, z uint8) (o uint8) {
	o = x
	if y > o {
		o = y
	}
	if z > o {
		o = z
	}
	return
}

// rnd returns a random number in range [from, to].
func rnd(from, to int) int {
	// fmt.Println(to + 1 - from)
	return rand.Intn(to+1-from) + from
}

func pic(w http.ResponseWriter, req *http.Request) {
	code := "1234"
	testData := make([]int, 4)
	for i := 0; i < len(code); i++ {
		iInt, _ := strconv.Atoi(code[i : i+1])
		testData[i] = iInt
	}
	w.Header().Set("Content-Type", "image/png")
	NewImage(testData, 100, 40).WriteTo(w)
}
