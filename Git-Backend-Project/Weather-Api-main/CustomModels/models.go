package CustomModels

import (
	"encoding/binary"
	"errors"
	"math"
)

// ------------------------------------------------------+
type WeatherResponse struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
} //u can add more parameters check doc:www.visualcrossing.com/resources/documentation/weather-api
//------------------------------------------------------+

func (w WeatherResponse) MarshalBinary() ([]byte, error) {
	if w.Latitude < -90 || w.Latitude > 90 || w.Longitude < -180 || w.Longitude > 180 {
		return nil, errors.New("invalid latitude or longitude value")
	}

	data := make([]byte, 24)
	binary.LittleEndian.PutUint64(data[0:8], math.Float64bits(w.Latitude))
	binary.LittleEndian.PutUint64(data[8:16], math.Float64bits(w.Longitude))
	copy(data[16:], []byte(w.ResolvedAddress))

	return data, nil
}

func (w *WeatherResponse) UnmarshalBinary(data []byte) error {
	if len(data) != 24 {
		return errors.New("invalid data length")
	}

	w.Latitude = math.Float64frombits(binary.LittleEndian.Uint64(data[0:8]))
	w.Longitude = math.Float64frombits(binary.LittleEndian.Uint64(data[8:16]))
	w.ResolvedAddress = string(data[16:])

	return nil
}
