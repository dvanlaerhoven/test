package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	s "strings"
)

type device struct {
	Type     string
	artnr    string
	serialnr string
	firmware string
	hardware string
}

func GeneralData() device {

	var device device

	url := "https://192.168.10.11/wbm/GeneralData.html"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Cookie", "toggle_menu_administration=true; toggle_menu_security=true; toggle_menu_infos=true; toggle_menu_diagnostics=true; toggle_menu_config=true; sid=t.NO4O3ix0mISLKT")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	// Creating subsets and extracting values

	subset := s.SplitAfter(string(body), "</tr>")

	// get device Type
	subsubset := s.SplitAfter(subset[4], "</td>")
	value := s.TrimSpace(subsubset[1])
	fmt.Println(value[4:(len(value) - 5)])
	fmt.Println(device.Type)

	device.Type = GetDataGeneral(4, subset)
	device.artnr = GetDataGeneral(5, subset)
	device.serialnr = GetDataGeneral(6, subset)
	device.firmware = GetDataGeneral(7, subset)
	device.hardware = GetDataGeneral(8, subset)

	return device
}

func GetDataGeneral(x int, subset []string) string {

	subsubset := s.SplitAfter(subset[x], "</td>")
	value := s.TrimSpace(subsubset[1])
	fmt.Println(value[4:(len(value) - 5)])
	return value[4:(len(value) - 5)]
}

func main() {

	var device device

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // Disable TLS check due to bad certificate

	// ----------- Sending HTTP requests and parse Data

	device = GeneralData()

	println(device.Type)

	//
}
