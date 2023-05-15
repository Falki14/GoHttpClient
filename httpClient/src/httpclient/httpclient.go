package httpclient

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"time"
)

type HttpClient struct {
	url string
	data []string
	authToken string
	timeout int
}

func (hclient *HttpClient) Get(url string) string {
    resp, err := http.Get(url)

    if err != nil {
		fmt.Printf("follow error on get: %s\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func (hclient *HttpClient) Post(url string, data string, header map[string]string, output string) (string, int) {
	request, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		fmt.Println("Error creating request: ", err)
		os.Exit(1)
	}

	if hclient.authToken != "" {
		request.Header.Add("Authorization", "Bearer "+hclient.authToken)
	}

	if header != nil && len(header) > 0 {
		for headerKey, headerValue := range header {
			request.Header.Add(headerKey, headerValue)
		}
	}

	client := &http.Client{Timeout: 5 * time.Second}
	response, err := client.Do(request)

	if err != nil {
			fmt.Println("Error sending request: ", err)
			os.Exit(1)
		}

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	return string(body), response.StatusCode
}

func New(timeout ...int) *HttpClient {
	hclient := &HttpClient {}

	if timeout != nil && len(timeout) > 0 {
        hclient.timeout = timeout[0]
    }

	return hclient
}

func (h *HttpClient) SetToken(token string) {
	if token != "" {
		h.authToken = token
	}
}

func (hclient *HttpClient) GetUrl() {
	fmt.Print(hclient.url)
 }

func (hclient *HttpClient) GetToken() string {
	return hclient.authToken
}