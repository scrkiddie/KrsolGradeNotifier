package main

import (
	"KrsolGradeNotifier/internal/helper"
	"KrsolGradeNotifier/internal/reverse"
	"KrsolGradeNotifier/internal/telegram"
	"flag"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"strings"
)

func checkAndNotify() {
	helper.LogInfo("Memulai pengecekan pada situs krsol")
	config, err := helper.LoadConfig("config.json")
	if err != nil {
		helper.LogFatal("Error membaca file config.json pastikan sudah sesuai instruksi :", err)
	}
	phpSessionId, nameValue := reverse.GetGradePageButtonValue(config)
	gradeData := reverse.GetGradeData(phpSessionId, nameValue)
	filename := "data.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		helper.LogInfo("Memulai proses inisialisasi data awal karena file data.txt tidak ditemukan")
		helper.WriteDataToFile(filename, gradeData)
		helper.LogInfo("Data baru sudah ditulis ke file data.txt silahkan cek untuk melihat data awal")
		helper.LogInfo("Notifikasi akan dikirimkan jika ada perubahan pada pengecekan berikutnya")
	} else {
		existingData, err := helper.ReadDataFromFile(filename)
		if err != nil {
			helper.LogError("[Error] Error membaca data yang ada :", err)
			return
		}
		changes := helper.DetectAndLogChanges(existingData, gradeData)
		if len(changes) > 0 {
			for _, change := range changes {
				helper.LogInfo(change)
			}
			combinedMessage := strings.Join(changes, "\n")
			telegram.SendMessage(config.TelegramToken, config.ChatID, combinedMessage)
		} else {
			helper.LogInfo("Tidak ada perubahan terdeteksi")
		}
		helper.WriteDataToFile(filename, gradeData)
	}
}

func main() {
	schedule := flag.String("schedule", "0 * * * *", "Cron schedule string")
	flag.Parse()
	if _, err := cron.ParseStandard(*schedule); err != nil {
		helper.LogFatal("Format cron salah :", err)
	}
	c := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(log.New(log.Writer(), "[Cron] ", log.LstdFlags))))
	_, err := c.AddFunc(*schedule, checkAndNotify)
	if err != nil {
		helper.LogFatal("Error menambahkan fungsi ke cron :", err)
	}
	c.Start()
	helper.LogInfo("Terjadwal dengan format cron :", *schedule)
	select {}
}
