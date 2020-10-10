package log

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	//Errors file
	path := "./log"

	errFile, err := os.OpenFile(path+"/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open errors_file.", err)
	}
	//Info file
	infoFile, err := os.OpenFile(path+"/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open info_file.：", err)
	}

	Info = log.New(io.MultiWriter(os.Stderr, infoFile), "Info", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
	//Warning.Println("Warning")
	//Error.Println("Error")
}