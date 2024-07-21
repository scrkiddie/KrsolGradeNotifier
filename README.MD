# Krsol-Grade-Notifier
![image_2024-07-21_03-04-43](https://github.com/user-attachments/assets/42c21a65-66f6-4e1a-927e-999fd8a5901a)
Krsol-Grade-Notifier adalah aplikasi Go yang secara berkala memeriksa pembaruan nilai KRS di situs krsol.unimma.ac.id dan mengirimkan notifikasi melalui Telegram saat ada perubahan. Pengguna bisa mengatur jadwal pemeriksaan menggunakan format cron.
## Fitur

* Memeriksa pembaruan nilai KRS secara berkala.
* Mengirim notifikasi melalui Telegram jika ada perubahan.
* Penjadwalan yang bisa ditentukan oleh pengguna dengan format cron.

## Prasyarat

* Golang terinstal di sistem Anda.
* Token bot telegram dan id chat telegram untuk mengirim notifikasi.
* Akun krsol.unimma.ac.id yang aktif.

## Instalasi

1. Clone repository:
```shell
   git clone https://github.com/Hilmi-Raif/Krsol-Grade-Notifier.git
   cd Krsol-Grade-Notifier
```
2. Ubah nama file config.example.json menjadi config.json dan isi dengan informasi Anda:

```json
{
  "username" : "ISI_USERNAME_KRSOL",
  "password" : "ISI_PASSWORD_KRSOL",
  "telegramToken": "ISI_TELEGRAM_BOT_TOKEN",
  "chatID": "ISI_TELEGRAM_CHAT_ID"
}
```
3. Install dependensi:
```shell
go get -u -d ./...
```
4. Build aplikasi:
- Untuk Windows:
```shell
go build -o KrsolGradeNotifier.exe cmd/krsol_grade_notifier/main.go
```
- Untuk Unix (Linux dan macOS):
```shell
go build -o KrsolGradeNotifier cmd/krsol_grade_notifier/main.go
```
## Penggunaan
Jalankan aplikasi dengan jadwal cron yang diinginkan. Defaultnya adalah memeriksa setiap jam.

- Untuk Windows:
```shell
./KrsolGradeNotifier.exe -schedule "0 * * * *"
```
- Untuk Unix (Linux dan macOS):
```shell
./KrsolGradeNotifier -schedule "0 * * * *"
```
###  Command Line
- -schedule: String jadwal cron. Default adalah "0 * * * *" (setiap jam).
### Contoh Jadwal Cron
* Setiap setengah jam: */30 * * * *
* Setiap 2 jam: 0 */2 * * *
* Setiap hari pada tengah malam: 0 0 * * *

Untuk belajar lebih lanjut tentang format cron, kunjungi [crontab guru](https://crontab.guru/).
## Catatan
Agar aplikasi ini terus berjalan dan memeriksa pembaruan:
* Perangkat harus tetap menyala 24 jam tanpa henti.
* Perangkat harus memiliki koneksi internet yang stabil.

Pastikan perangkat Anda dalam kondisi yang baik dan memiliki sumber daya yang memadai untuk tetap berjalan dan terhubung ke internet.
