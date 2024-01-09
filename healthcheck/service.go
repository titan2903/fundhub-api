package healthcheck

import "time"

type Service interface {
	HealthcheckService() (Healthcheck, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) HealthcheckService() (Healthcheck, error) {
	check := Healthcheck{
		ServiceName: "Golang Fundhub Services v3",
		Status:      "OK",
		Description: "Everything is running well",
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
	}
	return check, nil
}
