package main

import (
	"github.com/Dimadetected/notesAPI/internal/handler"
	"github.com/Dimadetected/notesAPI/internal/server"
	"github.com/sirupsen/logrus"
)

func main()  {
	router := handler.RouteInit()

	err := server.Start("8080",router)
	if err != nil{
		logrus.Error(err)
	}
}
