package mexif

import (
	"github.com/msvens/mexif/json"
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
	Audio    json.JSONObject `json:"audio,omitempty"`
	Author   json.JSONObject `json:"author,omitempty"`
	Camera   json.JSONObject `json:"camera,omitempty"`
	Device   json.JSONObject `json:"device,omitempty"`
	Document json.JSONObject `json:"document,omitempty"`
	ExifTool json.JSONObject `json:"exiftool,omitempty"`
	Image    json.JSONObject `json:"image,omitempty"`
	Location json.JSONObject `json:"location,omitempty"`
	Other    json.JSONObject `json:"other,omitempty"`
	Preview  json.JSONObject `json:"preview,omitempty"`
	Printing json.JSONObject `json:"printing,omitempty"`
	Time     json.JSONObject `json:"time,omitempty"`
	Unknown  json.JSONObject `json:"unknown,omitempty"`
	Video    json.JSONObject `json:"video,omitempty"`
}

func NewExifData(root json.JSONObject) *ExifData {
	ret := ExifData{}
	_ = json.ScanObject(Audio, root, &ret.Audio)
	_ = json.ScanObject(Author, root, &ret.Author)
	_ = json.ScanObject(Camera, root, &ret.Camera)
	_ = json.ScanObject(Device, root, &ret.Device)
	_ = json.ScanObject(Document, root, &ret.Document)
	_ = json.ScanObject(ExifTool, root, &ret.ExifTool)
	_ = json.ScanObject(Image, root, &ret.Image)

	_ = json.ScanObject(Location, root, &ret.Location)
	_ = json.ScanObject(Other, root, &ret.Other)
	_ = json.ScanObject(Preview, root, &ret.Preview)
	_ = json.ScanObject(Printing, root, &ret.Printing)
	_ = json.ScanObject(Time, root, &ret.Time)
	_ = json.ScanObject(Unknown, root, &ret.Unknown)
	_ = json.ScanObject(Video, root, &ret.Video)
	return &ret
}
