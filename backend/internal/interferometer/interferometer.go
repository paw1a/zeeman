package interferometer

import (
	"errors"
	"image"
	"image/color"
	"math"
)

var (
	errNegativeWaveLength = errors.New("magnetic induction is invalid")
)

type FieldDirection int

const (
	Disabled      FieldDirection = 0
	Parallel                     = 1
	Perpendicular                = 2
)

const (
	mm            = 1e-3
	nm            = 1e-9
	lightVelocity = 3 * 1e8
)

type Params struct {
	Resolution             int
	PictureSize            float64
	WaveLength             float64
	FocalDistance          float64
	GlassesDistance        float64
	PathDifference         float64
	RefractionFactor       float64
	MagneticInduction      float64
	MagneticFieldDirection FieldDirection
}

func CreateImage(params Params) (*image.RGBA, error) {
	normalizeParams(&params)

	intensity := make([][]float64, params.Resolution)
	for i := 0; i < params.Resolution; i++ {
		intensity[i] = make([]float64, params.Resolution)
	}

	paramsArray, err := applyMagneticField(params)
	if err != nil {
		return nil, err
	}

	for _, params := range paramsArray {
		draw(params, intensity)
	}

	return imageFromIntensity(intensity), nil
}

func applyMagneticField(params Params) ([]Params, error) {
	paramsArray := []Params{params}

	frequency := lightVelocity / params.WaveLength
	angleFrequency := 2 * math.Pi * frequency
	larmorFrequency := (1.6 * math.Pow(10, -19) * params.MagneticInduction) /
		(2 * 9 * math.Pow(10, -31))

	if params.MagneticFieldDirection != Disabled {
		frequency = (angleFrequency + larmorFrequency) / (2 * math.Pi)
		params.WaveLength = lightVelocity / frequency

		paramsArray = append(paramsArray, params)
	}

	if params.MagneticFieldDirection == Perpendicular {
		frequency = (angleFrequency - larmorFrequency) / (2 * math.Pi)
		params.WaveLength = lightVelocity / frequency
		if params.WaveLength < 0 {
			return nil, errNegativeWaveLength
		}

		paramsArray = append(paramsArray, params)
	}

	return paramsArray, nil
}

func draw(params Params, intensity [][]float64) {
	waveK := 2 * math.Pi / params.WaveLength
	fineness := 4 * params.RefractionFactor / math.Pow(1-params.RefractionFactor, 2)

	step := params.PictureSize / float64(params.Resolution) / mm
	for i := 0; i < params.Resolution; i++ {
		rayX := float64(i) * step

		for j := 0; j < params.Resolution; j++ {
			rayY := float64(j) * step
			x := rayX*mm - params.PictureSize/2
			y := rayY*mm - params.PictureSize/2

			radius := math.Sqrt(x*x + y*y)
			theta := radius / params.FocalDistance

			var delta float64
			delta = 2 * waveK * params.RefractionFactor *
				params.GlassesDistance * math.Cos(theta)
			lightIntensity := 1 / (1 + fineness*math.Pow(math.Sin(delta/2), 2))

			intensity[i][j] += lightIntensity
		}
	}
}

func imageFromIntensity(intensity [][]float64) *image.RGBA {
	min, max := matrixMinMax(intensity)

	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: len(intensity), Y: len(intensity[0])}
	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	for i := 0; i < len(intensity); i++ {
		for j := 0; j < len(intensity[0]); j++ {
			intensity[i][j] = (intensity[i][j] - min) / (max - min) * 255
			img.Set(i, j, color.RGBA{
				R: 0,
				G: uint8(intensity[i][j]),
				B: 0,
				A: 255,
			})
		}
	}

	return img
}

func matrixMinMax(matrix [][]float64) (float64, float64) {
	var min, max float64
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			min = math.Min(min, matrix[i][j])
			max = math.Max(max, matrix[i][j])
		}
	}

	return min, max
}

func normalizeParams(params *Params) {
	params.PictureSize *= mm
	params.WaveLength *= nm
	params.FocalDistance *= mm
	params.GlassesDistance *= mm
	params.PathDifference *= nm
}
