package reverse

import (
	"KrsolGradeNotifier/internal/helper"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"time"
)

func GetGradePageButtonValue(account *helper.Config) (string, string) {
	alert := true
	for {
		phpSessionID := GetPhpSess()
		for i := 0; i < 2; i++ {
			data := url.Values{}
			data.Set("username", account.Username)
			data.Set("password", account.Password)
			data.Set("login", "Sign in")

			reqPost, err := http.NewRequest("POST", "http://krsol.unimma.ac.id", bytes.NewBufferString(data.Encode()))
			if err != nil {
				if alert {
					helper.LogError("Cek jaringan anda :", err)
					alert = false
				}
				break
			}
			reqPost.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			if phpSessionID != "" {
				reqPost.Header.Set("Cookie", "PHPSESSID="+phpSessionID)
			}
			client := &http.Client{
				Timeout: 5 * time.Second,
			}
			respPost, err := client.Do(reqPost)
			if err != nil {
				if alert {
					helper.LogError("Cek jaringan anda :", err)
					alert = false
				}
				break
			}
			defer respPost.Body.Close()
			if i == 1 {
				body, err := io.ReadAll(respPost.Body)
				if err != nil {
					if alert {
						helper.LogError("Cek jaringan anda :", err)
						alert = false
					}
					break
				}

				doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
				if err != nil {
					helper.LogFatal("Hubungi developer :", err)
				}
				var nameValue string
				doc.Find("a[name]").Each(func(index int, item *goquery.Selection) {
					if index == 4 {
						nameValue, _ = item.Attr("name")
					}
				})

				if nameValue != "" {
					return phpSessionID, nameValue
				} else {
					helper.LogFatal("Username atau password akun krsol anda salah")
				}
			}
		}
	}

}
