package mexif

import (
	"github.com/msvens/mexif/json"
	"time"
)

type ExifCompact struct {
	Title    string   `json:"title,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
	Software string   `json:"software,omitempty"`
	Rating   uint     `json:"rating,omitempty"`

	CameraMake              string  `json:"cameraMake,omitempty"`
	CameraModel             string  `json:"cameraModel,omitempty"`
	LensInfo                string  `json:"lensInfo,omitempty"`
	LensModel               string  `json:"lensModel,omitempty"`
	LensMake                string  `json:"lensMake,omitempty"`
	FocalLength             string  `json:"focalLength,omitempty"`
	FocalLengthIn35mmFormat string  `json:"focalLengthIn35mmFormat,omitempty"`
	MaxApertureValue        float32 `json:"maxApertureValue,omitempty"`
	Flash                   string  `json:"flash,omitempty"`

	ExposureTime         string  `json:"exposureTime,omitempty"`
	ExposureCompensation float32 `json:"exposureCompensationn,omitempty"`
	ExposureProgram      string  `json:"exposoureProgram,omitempty"`
	FNumber              float32 `json:"fNumber,omitempty"`
	ISO                  uint    `json:"ISO,omitempty"`
	ColorSpace           string  `json:"colorSpace,omitempty"`
	XResolution          uint    `json:"xResolution,omitempty"`
	YResolution          uint    `json:"yResolution,omitempty"`
	ImageWidth           uint    `json:"imageWidth,omitempty"`
	ImageHeight          uint    `json:"imageHeight,omitempty"`

	OriginalDate time.Time `json:"originalDate,omitempty"`
	ModifyDate   time.Time `json:"modifyDate,omitempty"`

	GPSLatitude  float64 `json:"gpsLatitude,omitempty"`
	GPSLongitude float64 `json:"gpsLongitude,omitempty"`
	City         string  `json:"city,omitempty"`
	Country      string  `json:"country,omitempty"`
	State        string  `json:"state,omitempty"`
}

func NewExifCompact(data *ExifData) *ExifCompact {
	ec := ExifCompact{}

	_ = json.ScanString("Title", data.Image, &ec.Title)

	//Keywords can either come as an array or strong
	kw := data.Other["Keywords"]
	if kw != nil {
		t := json.TypeOf(kw)
		if t == json.JString {
			ec.Keywords = append(ec.Keywords, kw.(string))
		} else if t == json.JArr {
			for _, v := range kw.([]interface{}) {
				ec.Keywords = append(ec.Keywords, v.(string))
			}
		}
	}

	_ = json.ScanString("Software", data.Image, &ec.Software)
	_ = json.ScanUInt("Rating", data.Image, &ec.Rating)

	_ = json.ScanString("Make", data.Camera, &ec.CameraMake)
	_ = json.ScanString("Model", data.Camera, &ec.CameraModel)

	_ = json.ScanString("LensInfo", data.Image, &ec.LensInfo)
	_ = json.ScanString("LensModel", data.Image, &ec.LensModel)
	_ = json.ScanString("LensMake", data.Image, &ec.LensMake)

	_ = json.ScanString("FocalLength", data.Camera, &ec.FocalLength)
	_ = json.ScanString("FocalLengthIn35mmFormat", data.Camera, &ec.FocalLengthIn35mmFormat)
	_ = json.ScanFloat32("MaxApertureValue", data.Camera, &ec.MaxApertureValue)
	_ = json.ScanString("Flash", data.Camera, &ec.Flash)

	_ = json.ScanString("ExposureTime", data.Image, &ec.ExposureTime)
	_ = json.ScanFloat32("ExposureCopmensation", data.Image, &ec.ExposureCompensation)
	_ = json.ScanString("ExposureProgram", data.Camera, &ec.ExposureProgram)
	_ = json.ScanFloat32("FNumber", data.Image, &ec.FNumber)
	_ = json.ScanUInt("ISO", data.Image, &ec.ISO)
	_ = json.ScanString("ColorSpace", data.Image, &ec.ColorSpace)
	_ = json.ScanUInt("XResolution", data.Image, &ec.XResolution)
	_ = json.ScanUInt("YResolution", data.Image, &ec.YResolution)
	_ = json.ScanUInt("ImageWidth", data.Image, &ec.ImageWidth)
	_ = json.ScanUInt("ImageHeight", data.Image, &ec.ImageHeight)

	_ = json.ScanDateTime("DateTimeOriginal", "OffsetTimeOriginal", data.Time, &ec.OriginalDate)
	_ = json.ScanDateTime("ModifyDate", "OffsetTime", data.Time, &ec.ModifyDate)

	_ = json.ScanFloat64("GPSLatitude", data.Location, &ec.GPSLatitude)
	_ = json.ScanFloat64("GPSLongitude", data.Location, &ec.GPSLongitude)
	_ = json.ScanString("City", data.Location, &ec.City)
	_ = json.ScanString("Country", data.Location, &ec.Country)
	_ = json.ScanString("State", data.Location, &ec.State)

	return &ec
}
