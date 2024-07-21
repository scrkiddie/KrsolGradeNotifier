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

func GetGradeData(phpSessionID string, nameValue string) [][]string {
	alert := true
	for {
		data := url.Values{}
		data.Set("_object", nameValue)
		data.Set("_action", "onClick")
		req, err := http.NewRequest("POST", "http://krsol.unimma.ac.id", bytes.NewBufferString(data.Encode()))
		if err != nil {
			if alert {
				helper.LogError("Cek jaringan anda :", err)
				alert = false
			}
			continue
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Cookie", "PHPSESSID="+phpSessionID)
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			if alert {
				helper.LogError("Cek jaringan anda :", err)
				alert = false
			}
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			if alert {
				helper.LogError("Cek jaringan anda :", err)
				alert = false
			}
			continue
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
		if err != nil {
			helper.LogFatal("Hubungi developer :", err)
		}

		table12 := doc.Find("table").Eq(11)

		var data2 [][]string
		table12.Find("tr").Each(func(i int, row *goquery.Selection) {
			if i == 0 {
				return
			}
			var rowData []string
			row.Find("td").Each(func(j int, col *goquery.Selection) {
				if j >= 3 && j <= 6 && j != 4 {
					rowData = append(rowData, col.Text())
				}
			})
			data2 = append(data2, rowData)
		})
		return data2
	}

}
