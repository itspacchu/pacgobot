package miscutils

import (
	"fmt"
	"image"
	"net/http"
)

func GetAverageColor(imageURL string) int {
	resp, err := http.Get(imageURL)
	if err != nil {
		fmt.Printf("[WARN] No Image found from url")
		return 0xFFFFFF
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		fmt.Printf("[WARN] Image coudnt be loaded")
		return 0xFFFFFF
	}

	var totalRed, totalGreen, totalBlue, totalAlpha float64
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			r, g, b, a := img.At(x, y).RGBA()
			totalRed += float64(r)
			totalGreen += float64(g)
			totalBlue += float64(b)
			totalAlpha += float64(a)
		}
	}

	averageRed := (totalRed / 255.) / float64(img.Bounds().Dx()*img.Bounds().Dy())
	averageGreen := (totalGreen / 255.) / float64(img.Bounds().Dx()*img.Bounds().Dy())
	averageBlue := (totalBlue / 255.) / float64(img.Bounds().Dx()*img.Bounds().Dy())

	finalColor := (int(averageRed) << 16) | (int(averageGreen) << 8) | int(averageBlue)
	if finalColor > 16777215 {
		return 0xFF5555
	} else {
		return finalColor
	}
}
