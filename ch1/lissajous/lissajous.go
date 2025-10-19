// Lissajous генерирует анимированный gif из случайных фигур Лиссажу
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xFF}, // Черный
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // Зеленый
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // Синий
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // Желтый
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, // Пурпурный
	color.RGBA{0x32, 0xCD, 0x32, 0xFF}, // LimeGreen
	color.RGBA{0xFF, 0x14, 0x93, 0xFF}, // DeepPink
	color.RGBA{0xFF, 0xD7, 0x00, 0xFF}, // Gold
}

const (
// whiteIndex = 0 // Первый цвет. палитры
// blackIndex = 1 // Следующий цвет палитры
)

func Lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Кол-во кадров анимации
		delay   = 8     // Задержка между кадрами (еденица - 10мс)
	)
	rng := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	freq := rng.Float64() * 3.0 // Относительная частота колебания y
	randIndex := rng.Intn(len(palette) + 1)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Игнорируем ошибки
}
