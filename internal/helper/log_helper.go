package helper

import "log"

func LogError(v ...interface{}) {
	log.SetPrefix("[Error] ")
	log.Println(v...)
}

func LogInfo(v ...interface{}) {
	log.SetPrefix("[Info] ")
	log.Println(v...)
}

func LogFatal(v ...interface{}) {
	log.SetPrefix("[Fatal] ")
	log.Fatalln(v...)
}
