package pkg

import (
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"net/url"
	"path"

	"github.com/koron/go-ssdp"
)

func SearchForDevice() *Device {
	res, err := ssdp.Search("roku:ecp", 1, "")
	if err != nil {
		panic(err)
	}

	if len(res) < 1 {
		panic(errors.New("no device found"))
	}

	device := &Device{Addr: res[0].Location}
	return device
}

func DeviceFromIP(ip net.IP, port string) *Device {
	return &Device{
		Addr: "http://" + ip.String() + port,
	}
}

type Device struct {
	Addr string
}

type MediaType string

const (
	MediaTypeNone MediaType = ""
	MediaTypeLive           = "live"
)

func (d *Device) LaunchApp(appID, contentID string, mediaType MediaType) {
	u, _ := url.Parse(d.Addr)
	u.Path = "/launch/" + appID
	q := u.Query()
	q.Set("contentId", contentID)
	if mediaType != MediaTypeNone {
		q.Set("mediaType", string(mediaType))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", u.String(), http.NoBody)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
}

type DeviceInfo struct {
	PowerMode string `xml:"power-mode"`
}

func (d *Device) GetDeviceInfo() (*DeviceInfo, error) {
	u, _ := url.Parse(d.Addr)
	u.Path = path.Join("query", "device-info")

	req, err := http.NewRequest("GET", u.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := xml.NewDecoder(res.Body)
	var deviceInfo DeviceInfo
	if err := decoder.Decode(&deviceInfo); err != nil {
		return nil, err
	}

	return &deviceInfo, nil
}

func (d *Device) Keypress(button string) {
	u, _ := url.Parse(d.Addr)
	u.Path = path.Join("keypress", button)

	req, err := http.NewRequest("POST", u.String(), http.NoBody)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
}
