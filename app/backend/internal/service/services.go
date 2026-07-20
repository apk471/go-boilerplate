package service

import (
	job "github.com/ayush-amin/go-boilerplate/internal/lib/jobs"
	"github.com/ayush-amin/go-boilerplate/internal/repository"
	"github.com/ayush-amin/go-boilerplate/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}