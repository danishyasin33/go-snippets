package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// ec2Address returns the private IP address of the EC2 instance using the
// Instance Metadata API.
func Ec2Address() (string, error) {
	TOKEN := getInstanceMetaDataToken()
	if TOKEN == "" {
		return "", fmt.Errorf("failed to retrieve instance metadata token")
	}

	req, err := http.NewRequest("GET", "http://169.254.169.254/latest/meta-data/local-ipv4", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return "", err
	}

	req.Header.Set("X-aws-ec2-metadata-token", TOKEN)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to download user-data: %s", err)
	}

	defer res.Body.Close()

	ip, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}
	return string(ip), nil
}

func getInstanceMetaDataToken() string {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://169.254.169.254/latest/api/token", nil)
	if err != nil {
		fmt.Println("client: could not create request:", err)
		return ""
	}
	req.Header.Set("X-aws-ec2-metadata-token-ttl-seconds", "21600")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("failed to retrieve token:", err)
		return ""
	}
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read response body:", err)
		return ""
	}

	fmt.Println("Token retrieved successfully, token:", string(bodyText))

	return string(bodyText)
}
