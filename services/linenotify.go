package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func LineNotify(msg string) (err error) {
	url := "https://notify-api.line.me/api/notify"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("message=%s", msg))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("LINE_NOTIFY_TOKEN")))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	return
}
