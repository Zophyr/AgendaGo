package service
import (
	"log"
	//"os"
	"io"
)

var (
	logfile io.Writer
	logger *log.Logger
)

func InitLogger(){
	/*
	filename := "agenda.log"
	logfile,_ = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	logger = log.New(logfile,"[Info]",log.Llongfile)
	logger.Println("--start log--")
	*/
	
}


func logln(message string){
	logger.Println(message)
}

