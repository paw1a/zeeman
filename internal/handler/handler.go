package handler

import (
	"bytes"
	"fmt"
	"github.com/paw1a/interferometer/internal/interferometer"
	"image/png"
	"net/http"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	params := interferometer.Params{
		Resolution:             500,
		PictureSize:            5,
		WaveLength:             532,
		FocalDistance:          100,
		GlassesDistance:        2,
		PathDifference:         10,
		RefractionFactor:       1.375,
		ReflectionFactor:       0.7,
		IncidentLightIntensity: 10,
	}
	img := interferometer.CreateImage(params)

	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, img)
	if err != nil {
		fmt.Printf("can't encode image pixels")
		return
	}
	data := buffer.Bytes()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)

	return
}
