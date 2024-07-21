package reverse

import (
	"KrsolGradeNotifier/internal/helper"
	"net/http"
	"time"
)

func GetPhpSess() string {
	alert := true
	for {
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		reqGet, err := http.NewRequest("GET", "http://krsol.unimma.ac.id", nil)
		if err != nil {
			if alert {
				helper.LogError("Cek jaringan anda :", err)
				alert = false
			}
			continue
		}
		respGet, err := client.Do(reqGet)
		if err != nil {
			if alert {
				helper.LogError("Cek jaringan anda :", err)
				alert = false
			}
			continue
		}
		defer respGet.Body.Close()
		cookies := respGet.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == "PHPSESSID" {
				return cookie.Value
			}
		}
	}
}
