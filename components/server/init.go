package server

import (
	"log"

	"github.com/gangozero/hackprague2019-be/components/gcapi"
)

type Server struct {
	gc *gcapi.Client
}

func NewServer() *Server {
    //TODO: change to env variable
	projectID := "hackprague2019-239615"

	gcClient, err := gcapi.NewClient(projectID)
	if err != nil {
		log.Fatalf("Cannot create Google Cloud client: %s", err.Error())
	}
	return &Server{
		gc: gcClient,
	}
}

func (s *Server) Stop() {
    s.gc.Close()
}