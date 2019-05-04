package server

import (
	"log"

	"github.com/gangozero/hackprague2019-be/generated/restapi/operations/data"
	"github.com/gangozero/hackprague2019-be/generated/restapi/operations/profile"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) GetProfiles(params profile.GetProfileParams) middleware.Responder {
	resp, err := s.gc.GetProfiles()
	if err != nil {
		log.Printf("[GetProfiles] Error: %s", err.Error())
		return profile.NewGetProfileDefault(500)
	}
	return profile.NewGetProfileOK().WithPayload(resp)
}

func (s *Server) GetGradesByID(params data.GetGradeListParams) middleware.Responder {
	resp, err := s.gc.GetGradesByID(params.ProfileID)
	if err != nil {
		log.Printf("[GetGradesByID] Error: %s", err.Error())
		return data.NewGetGradeListDefault(500)
	}
	return data.NewGetGradeListOK().WithPayload(resp)
}

func (s *Server) GetGradesByIDAndUser(params data.GetUserGradeListParams) middleware.Responder {
	resp, err := s.gc.GetGradesByIDAndUser(params.ProfileID, params.UserID)
	if err != nil {
		log.Printf("[GetGradesByIDAndUser] Error: %s", err.Error())
		return data.NewGetUserGradeListDefault(500)
	}
	return data.NewGetUserGradeListOK().WithPayload(resp)
}

func (s *Server) PostGrade(params data.PostNewGradeParams) middleware.Responder {
	err := s.gc.PostGrade(params.ProfileID, params.Data)
	if err != nil {
		log.Printf("[PostGrade] Error: %s", err.Error())
		return data.NewPostNewGradeDefault(500)
	}
	return data.NewPostNewGradeAccepted()
}
