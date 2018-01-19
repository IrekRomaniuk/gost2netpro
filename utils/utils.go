package utils

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	//"fmt"
	"os"
)

//Netpro st2 webhook
type Netpro struct {
	Hosts string `json:"hosts"`
	User  string `json:"user"`
	Pass  string `json:"pass"`
	Cmd   string `json:"cmd"`
}

//Hosts creates output slice of targets
func (list *Netpro) ReadFile(hosts string) error {

	lines, err := readHosts(hosts)

	if err != nil {
		//fmt.Println("Error reading file", hosts)
		return errors.New("Error reading file")
	} else {
		*list = Netpro{Hosts: strings.Join(deletempty(lines), ",")}
	}
	return nil
}

//Response from Phantom
type Response struct {
	Cmd   string `json:"cmd"`
	Hosts string `json:"hosts"`
	User  string `json:"user"`
	Pass  string `json:"pass"`
}

//PostPage to st2
func PostPage(url, api string, data interface{}) (string, error) {
	var response Response
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(data)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("St2-Api-Key", api)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	err = json.Unmarshal(htmlData, &response)
	if err != nil {
		return "", err
	}
	return response.Cmd, nil
}

func readHosts(path string) ([]string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func deletempty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
