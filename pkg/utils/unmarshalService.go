package utils

import (
	"encoding/json"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
)

func GetUnMarshalledHeartData(w http.ResponseWriter, r *http.Request, userHeartRate string) *models.HeartRate {
	var heartRate models.HeartRate
	err := json.Unmarshal([]byte(userHeartRate), &heartRate)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return &heartRate
}

func GetUnMarshalledHeartArr(w http.ResponseWriter, r *http.Request, userHeartRate string) []*models.HeartRateData {
	var heartRate []*models.HeartRateData
	err := json.Unmarshal([]byte(userHeartRate), &heartRate)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return heartRate
}

func GetUnMarshalledBloodPressure(w http.ResponseWriter, r *http.Request, userBloodPressure string) *models.BloodPressure {
	var bloodPressure models.BloodPressure
	err := json.Unmarshal([]byte(userBloodPressure), &bloodPressure)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return &bloodPressure
}

func GetUnMarshalledBloodPressureArr(w http.ResponseWriter, r *http.Request, userBloodPressure string) []*models.BloodPressureData {
	var bloodPressure []*models.BloodPressureData
	err := json.Unmarshal([]byte(userBloodPressure), &bloodPressure)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return bloodPressure
}

func GetUnMarshalledBodyTemperature(w http.ResponseWriter, r *http.Request, userBodyTemperature string) *models.BodyTemperature {
	var bodyTemperature models.BodyTemperature
	err := json.Unmarshal([]byte(userBodyTemperature), &bodyTemperature)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return &bodyTemperature
}

func GetUnMarshalledBodyTemperatureArr(w http.ResponseWriter, r *http.Request, userBodyTemperature string) []*models.BodyTemperatureData {
	var bodyTemperature []*models.BodyTemperatureData
	err := json.Unmarshal([]byte(userBodyTemperature), &bodyTemperature)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return bodyTemperature
}
func GetUnMarshalledGlucose(w http.ResponseWriter, r *http.Request, userGlucose string) *models.Glucose {
	var glucose models.Glucose
	err := json.Unmarshal([]byte(userGlucose), &glucose)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return &glucose
}

func GetUnMarshalledGlucoseArr(w http.ResponseWriter, r *http.Request, userGlucose string) []*models.GlucoseData {
	var glucose []*models.GlucoseData
	err := json.Unmarshal([]byte(userGlucose), &glucose)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return glucose
}
func GetUnMarshalledOxygenSaturation(w http.ResponseWriter, r *http.Request, userOxygenSaturation string) *models.OxygenSaturation {
	var oxygenSaturation models.OxygenSaturation
	err := json.Unmarshal([]byte(userOxygenSaturation), &oxygenSaturation)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return &oxygenSaturation
}

func GetUnMarshalledOxygenSaturationArr(w http.ResponseWriter, r *http.Request, userOxygenSaturation string) []*models.OxygenSaturationData {
	var oxygenSaturation []*models.OxygenSaturationData
	err := json.Unmarshal([]byte(userOxygenSaturation), &oxygenSaturation)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return oxygenSaturation
}

func GetUnMarshalledIntervals(w http.ResponseWriter, r *http.Request, interval string) *models.Intervals {
	var intervals models.Intervals
	err := json.Unmarshal([]byte(interval), &intervals)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return &intervals
}

func GetUnMarshalledMedicalRecord(w http.ResponseWriter, r *http.Request, record string) []*models.Record {
	var records []*models.Record
	err := json.Unmarshal([]byte(record), &records)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return records
}
