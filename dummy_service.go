package main

import (
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
)

var serviceConfig = &service.Config{
	Name:        "Dummy Service",
	DisplayName: "Dummy Service",
	Description: "Dummy Service Description, Trend Micro Inc.",
}

func main() {
	p := &program{}
	s, err := service.New(p, serviceConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) < 2 {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
		return
	}

	control := os.Args[1]
	srvName := os.Args[2]
	srvDisplayName := os.Args[3]

	if srvName == "" || srvDisplayName == "" {
		log.Fatal("Service Name or Service Display Name is empty.")
	}

	if control == "/install" {
		serviceInstall(s, srvName, srvDisplayName)
	}
	if control == "/uninstall" {
		serviceUninstall(s, srvName, srvDisplayName)
	}
}

func serviceInstall(s service.Service, srvName, srvDisplayName string) {
	serviceConfig.Name = srvName
	serviceConfig.DisplayName = srvDisplayName
	err := s.Install()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Install Successful")
	return
}

func serviceUninstall(s service.Service, srvName, srvDisplayName string) {
	serviceConfig.Name = srvName
	serviceConfig.DisplayName = srvDisplayName
	err := s.Uninstall()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uninstall Successful")
	return
}

type program struct{}

func (p *program) Start(s service.Service) error {
	log.Println("Start Service")
	go p.run()
	return nil
}
func (p *program) Stop(s service.Service) error {
	log.Println("Stop Service")
	return nil
}
func (p *program) run() {
	fmt.Println("Run Service")
}
