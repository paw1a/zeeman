package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/paw1a/interferometer/backend/internal/interferometer"
	"image/png"
	"net/http"
)

type GetImageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type GetImageRequest struct {
	Resolution             int                           `json:"resolution" validate:"required,min=1"`
	PictureSize            float64                       `json:"pictureSize" validate:"required,min=0"`
	WaveLength             float64                       `json:"waveLength" validate:"min=0"`
	FocalDistance          float64                       `json:"focalDistance" validate:"min=0"`
	GlassesDistance        float64                       `json:"glassesDistance" validate:"min=0"`
	PathDifference         float64                       `json:"pathDifference" validate:"min=0"`
	RefractionFactor       float64                       `json:"refractionFactor" validate:"required,min=1"`
	MagneticInduction      float64                       `json:"magneticInduction" validate:"omitempty,min=0"`
	MagneticFieldDirection interferometer.FieldDirection `json:"magneticFieldDirection" validate:"omitempty,eq=1|eq=2"`
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorResponse(w, errors.New("post method expected"), http.StatusMethodNotAllowed)
		return
	}

	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var request GetImageRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		errorResponse(w, fmt.Errorf("request decode error: %w", err), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		errorResponse(w, fmt.Errorf("validate error: %w", err), http.StatusBadRequest)
		return
	}

	img, err := interferometer.CreateImage(requestToParams(request))
	if err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, img)
	if err != nil {
		errorResponse(w, fmt.Errorf("can't encode image pixels"), http.StatusInternalServerError)
		return
	}
	data := buffer.Bytes()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}

func errorResponse(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	resp := GetImageResponse{
		Status:  statusCode,
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(resp)
}

func requestToParams(request GetImageRequest) interferometer.Params {
	return interferometer.Params{
		Resolution:             request.Resolution,
		PictureSize:            request.PictureSize,
		WaveLength:             request.WaveLength,
		FocalDistance:          request.FocalDistance,
		GlassesDistance:        request.GlassesDistance,
		PathDifference:         request.PathDifference,
		RefractionFactor:       request.RefractionFactor,
		MagneticInduction:      request.MagneticInduction,
		MagneticFieldDirection: request.MagneticFieldDirection,
	}
}
