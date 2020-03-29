package mexif

import (
	"github.com/msvens/mexif/json"
	"time"
)

const Audio = "Audio"
const Author = "Author"
const Camera = "Camera"
const Device = "Device"
const Document = "Document"
const ExifTool = "ExifTool"
const Image = "Image"
const Location = "Location"
const Other = "Other"
const Preview = "Preview"
const Printing = "Printinng"
const Time = "Time"
const Unknown = "Unknow"
const Video = "Video"



type ExifData struct {
	Audio json.JSONObject `json:"audio,omitempty"`
	Author json.JSONObject `json:"author,omitempty"`
	Camera json.JSONObject `json:"camera,omitempty"`
	Device json.JSONObject `json:"device,omitempty"`
	Document json.JSONObject `json:"document,omitempty"`
	ExifTool json.JSONObject `json:"exiftool,omitempty"`
	Image json.JSONObject `json:"image,omitempty"`
	Location json.JSONObject `json:"location,omitempty"`
	Other json.JSONObject `json:"other,omitempty"`
	Preview json.JSONObject `json:"preview,omitempty"`
	Printing json.JSONObject `json:"printing,omitempty"`
	Time json.JSONObject `json:"time,omitempty"`
	Unknown json.JSONObject `json:"unknown,omitempty"`
	Video json.JSONObject `json:"video,omitempty"`
}

func NewExifData(root json.JSONObject) *ExifData{
	ret := ExifData{}
	_ = getAndSet(Audio, root, &ret.Audio)
	_ = getAndSet(Author, root, &ret.Author)
	_ = getAndSet(Camera, root, &ret.Camera)
	_ = getAndSet(Device, root, &ret.Device)
	_ = getAndSet(Document, root, &ret.Document)
	_ = getAndSet(ExifTool, root, &ret.ExifTool)
	_ = getAndSet(Image, root,  &ret.Image)

	_ = getAndSet(Location, root, &ret.Location)
	_ = getAndSet(Other, root, &ret.Other)
	_ = getAndSet(Preview, root, &ret.Preview)
	_ = getAndSet(Printing, root, &ret.Printing)
	_ = getAndSet(Time, root, &ret.Time)
	_ = getAndSet(Unknown, root, &ret.Unknown)
	_ = getAndSet(Video, root, &ret.Video)
	return &ret
}

//Common Fields
func (d *ExifData) FNumber() (float32, error) {
	return json.GetFloat32("FNumber", d.Image)
}

func (d *ExifData) ISO() (uint, error) {
	return json.GetUInt("ISO", d.Image)
}

func (d *ExifData) ShutterSpeedValue() (string, error) {
	return json.GetString("ShutterSpeedValue", d.Image)
}

func (d *ExifData) OriginalDate() (time.Time, error) {
	return json.GetDateTime("DateTimeOriginal", "OffsetTimeOriginal", d.Time)
}

func (d *ExifData) ModifyDate() (time.Time, error) {
	return json.GetDateTime("ModifyDate", "OffsetTime", d.Time)
}


func getAndSet(group string, root json.JSONObject, val *json.JSONObject) error {
	o, err := json.GetObject(group, root)
	if err != nil {
		return err
	}
	*val = o
	return nil
}

