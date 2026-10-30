// ManoMano Product API – Go example
// Usage: PILOTERR_API_KEY=your-key go run product.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	params := url.Values{}
	params.Set("query", "https://www.manomano.fr/p/pergola-retractable-298l-x-213l-x-222h-m-structure-metal-epoxy-anticorrosion-toile-polyester-haute-densite-180-g-m-incluse-gris-17786586?model_id=17784588")

	req, _ := http.NewRequest("GET", "https://api.piloterr.com/v2/manomano/product?"+params.Encode(), nil)
	req.Header.Set("x-api-key", os.Getenv("PILOTERR_API_KEY"))

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Piloterr error %d: %s\n", resp.StatusCode, body)
		os.Exit(1)
	}

	var pretty bytes.Buffer
	json.Indent(&pretty, body, "", "  ")
	fmt.Println(pretty.String())
}
