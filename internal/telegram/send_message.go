package telegram

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func SendMessage(token, chatID, message string) {
	alert := true
	for {
		apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
		data := url.Values{}
		data.Set("chat_id", chatID)
		data.Set("text", message)

		resp, err := http.PostForm(apiURL, data)
		if err != nil {
			if alert {
				fmt.Println("[Error] Cek jaringan anda : ", err)
				alert = false
			}
			continue
		}

		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Println("[Error] Error response dari Telegram pastikan chat id dan bot token sudah benar :", resp.Status)
		}
		break
	}

}
