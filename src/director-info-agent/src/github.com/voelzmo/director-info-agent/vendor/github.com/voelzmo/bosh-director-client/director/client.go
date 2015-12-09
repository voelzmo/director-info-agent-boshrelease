package director

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func NewClient(rootCAPath string) *http.Client {
	pemData, err := ioutil.ReadFile(rootCAPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Parsing the rootCA cert at path '%s' failed: %s", rootCAPath, err))
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(pemData)
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	return client
}

type Client interface {
	RequestAndParseJSON(string, string, map[string]string, []byte, interface{}) (interface{}, error)
	RequestAndParseString(string, string, map[string]string, []byte) (string, error)
}

type authorizingClient struct {
	host       string
	rootCAPath string
	auth       string
}

func (c *authorizingClient) RequestAndParseString(method string, path string, headers map[string]string, bodyContent []byte) (string, error) {
	resp, err := c.request(method, path, headers, bodyContent)
	if err != nil {
		log.Fatalf("Error contacting director: %s", err)
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)

	return string(responseBody), nil
}

func (c *authorizingClient) RequestAndParseJSON(method string, path string, headers map[string]string, bodyContent []byte, jsonObject interface{}) (interface{}, error) {

	resp, err := c.request(method, path, headers, bodyContent)
	if err != nil {
		log.Fatalf("Error contacting director: %s", err)
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseBody, &jsonObject)

	return jsonObject, nil
}

func (c *authorizingClient) request(method string, path string, headers map[string]string, bodyContent []byte) (*http.Response, error) {
	var body io.Reader
	if bodyContent != nil {
		body = bytes.NewReader(bodyContent)
	}

	req, _ := http.NewRequest(method, fmt.Sprintf("%s%s", c.host, path), body)

	if c.auth != "" {
		req.Header.Add("Authorization", c.auth)
	}

	for header, value := range headers {
		req.Header.Add(header, value)
	}

	directorClient := NewClient(c.rootCAPath)
	resp, err := directorClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error contacting director: %s", err)
	}
	return resp, nil
}

func GetClient(host string, rootCAPath string, auth string) Client {
	return &authorizingClient{host, rootCAPath, auth}
}
