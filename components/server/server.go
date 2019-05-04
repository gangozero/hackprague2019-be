package server

import (
	"github.com/gangozero/hackprague2019-be/generated/restapi/operations/profile"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) GetProfiles(params profile.GetProfileParams) middleware.Responder {
	resp, err := s.gc.GetProfiles()
	if err != nil {
		return profile.NewGetProfileDefault(500)
	}
	return profile.NewGetProfileOK().WithPayload(resp)
}
