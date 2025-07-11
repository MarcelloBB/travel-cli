package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"travel-cli/internal/utils"
)

func Get(url string, verbose bool, headers string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	if headers != "" {
		if verbose {
			fmt.Println("--- Headers ---")
		}

		parsedHeaders := utils.ParseHeaders(headers)
		for key, value := range parsedHeaders {
			req.Header.Set(key, value)
			if verbose {
				fmt.Printf("* %s: %s\n", key, value)
			}
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if verbose {
		fmt.Println("\n--- Response ---")
		fmt.Println("* Status code:", resp.Status)
	}

	return string(body), nil
}
