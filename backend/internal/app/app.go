package app

import (
	"github.com/paw1a/interferometer/backend/internal/handler"
	"log"
	"net/http"
)

func Run() {
	//params := interferometer.Params{
	//	Resolution:             500,
	//	PictureSize:            5,
	//	WaveLength:             532,
	//	FocalDistance:          100,
	//	GlassesDistance:        2,
	//	PathDifference:         10,
	//	RefractionFactor:       1.375,
	//	ReflectionFactor:       0.7,
	//	IncidentLightIntensity: 10,
	//}
	//img := interferometer.CreateImage(params)
	//imageFile, err := os.Create("image.png")
	//if err != nil {
	//	fmt.Printf("can't create file")
	//	return
	//}
	//
	//err = png.Encode(imageFile, img)
	//if err != nil {
	//	fmt.Printf("can't encode image pixels")
	//	return
	//}
	http.HandleFunc("/interferometer", handler.GetImage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
