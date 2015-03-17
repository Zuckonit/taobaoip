// Package ipinfo provides ...
package taobaoip

//assume your network is OK
import (
	"testing"
)

func TestIsIPV4(t *testing.T) {
	ip := "192.18.10.100s"
	if isIPV4(ip) {
		t.Errorf("%s is not a ipv4", ip)
	}

	ip = "192.168.1.1"
	if !isIPV4(ip) {
		t.Errorf("%s is a ipv4", ip)
	}
}

func TestURLOpen(t *testing.T) {
	req := Req{ip: "192.128.1.10s"}
	rb, err := req.URLOpen()
	if err == nil {
		t.Error("IP is invalid")
	}
	if rb != nil {
		_, err = rb.GetIPInfo()

		req = Req{ip: "127.0.0.1"}
		rb, err = req.URLOpen()
		info, err := rb.GetIPInfo()
		if err != nil {
			t.Error("get ip info failed")
		} else if info.CountryId != "IANA" {
			t.Error("its a LAN")
		}

		req = Req{ip: "192.1.1.1"}
		rb, err = req.URLOpen()
		info, err = rb.GetIPInfo()
		if err != nil {
			t.Error("get ip info failed")
		} else if info.CountryId != "US" {
			t.Error("its a US IP")
		}
	}

}
