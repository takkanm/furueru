package main

import (
	"image"
	"image/color/palette"
	"image/gif"
	"image/png"
	"math/rand"
	"os"
)

func writeGifFile(outFileName string, g gif.GIF) {
	outFile, _ := os.Create(outFileName)
	defer outFile.Close()

	g.Delay = make([]int, len(g.Image))
	g.LoopCount = 0

	gif.EncodeAll(outFile, &g)
}

func slideImage(in image.Image, slideX int, slideY int) *image.Paletted {
	margin := 2
	inBounds := in.Bounds()
	rect := image.Rect(inBounds.Min.X, inBounds.Min.Y, inBounds.Max.X+margin, inBounds.Max.Y+margin)
	pl := image.NewPaletted(rect, palette.WebSafe)

	for x := 0; x < inBounds.Max.X; x++ {
		for y := 0; y < inBounds.Max.Y; y++ {
			pl.Set(x+slideX, y+slideY, in.At(x, y))
		}
	}
	return pl
}

func calsSlideVolume() (int, int) {
	x := (rand.Int() % 3)
	y := (rand.Int() % 3)
	return x, y
}

func generateAnimeGif(inFileName string) {
	inFile, _ := os.Open(inFileName)
	defer inFile.Close()

	pngImage, _ := png.Decode(inFile)
	moveTimes := 100

	outGif := gif.GIF{
		Image: []*image.Paletted{},
	}

	for s := 0; s < moveTimes; s++ {
		palet := slideImage(pngImage, 1, 1)
		outGif.Image = append(outGif.Image, palet)

		slideX, slideY := calsSlideVolume()
		palet = slideImage(pngImage, slideX, slideY)
		outGif.Image = append(outGif.Image, palet)
	}

	writeGifFile(inFileName+".gif", outGif)
}

func main() {
	generateAnimeGif(os.Args[1])
}
