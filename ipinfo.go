// Package taobaoip provides ...
package taobaoip

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
)

const (
	BaseURL = "http://ip.taobao.com/service/getIpInfo.php?ip="
)

var (
	RegIPv4        = regexp.MustCompile("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$")
	ErrInvalidIPV4 = errors.New("Invalid IPV4.")
	ErrRequestFail = errors.New("Request failed, check your network.")
	ErrEmptyResp   = errors.New("Empty response")
	ErrUnkown      = errors.New("Unkown error")
	ErrInvalidJson = errors.New("Response is not json")
)

type IPInfo struct {
	Country   string `json:country`
	CountryId string `json:country_id`
	Area      string `json:area`
	AreaId    string `json:area_id`
	Region    string `json:region`
	RegionId  string `json:region_id`
	City      string `json:city`
	CityId    string `json:city_id`
	County    string `json:county`
	CountyId  string `json:county_id`
	Isp       string `isp`
	IspId     string `isp_id`
	Ip        string `ip`
}

type ValidData struct {
	Data IPInfo `json:data`
}

type ResponseInfo struct {
	Code int
	Data interface{}
}

type ResponseBody struct {
	Buf  []byte
	Info ResponseInfo
}

type Req struct {
	ip string
}

func isIPV4(ip string) bool {
	if m := RegIPv4.MatchString(ip); !m {
		return false
	}
	return true
}

func (r *Req) URLOpen() (*ResponseBody, error) {
	if !isIPV4(r.ip) {
		return nil, ErrInvalidIPV4
	}
	url := BaseURL + r.ip
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, ErrRequestFail
	}

	var buf []byte
	var rbinfo ResponseInfo
	if buf, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, ErrRequestFail
	}

	if err = json.Unmarshal(buf, &rbinfo); err != nil {
		return nil, ErrInvalidJson
	}
	return &ResponseBody{Buf: buf, Info: rbinfo}, nil
}

func (rb *ResponseBody) GetIPInfo() (*IPInfo, error) {
	rbinfo := rb.Info
	if rbinfo.Code == 1 {
		return nil, errors.New(rbinfo.Data.(string))
	} else if rbinfo.Code == 0 {
		var vd ValidData
		json.Unmarshal(rb.Buf, &vd)
		return &vd.Data, nil
	} else {
		return nil, ErrUnkown
	}
}

func (rb *ResponseBody) Format() string {
	info, err := rb.GetIPInfo()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}

	s := reflect.ValueOf(info).Elem()
	t := s.Type()
	sfmt := ""
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		sfmt += fmt.Sprintf("%s: %v\r\n", t.Field(i).Name, f.Interface())
	}
	return sfmt
}

func (rb *ResponseBody) Print() {
	sfmt := rb.Format()
	fmt.Println(sfmt)
}
