package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PagerDuty/go-pagerduty"
)

type Service struct {
	Endpoint    string
	ServiceKey  string
	IncidentKey string
}

func (service *Service) runCheck() {
	resp, err := http.Get(service.Endpoint)
	if err != nil {
		log.Printf("Could not reach service: %s", err)
		if service.IncidentKey != "" {
			return
		}

		service.IncidentKey, err = service.sendAlert(err)
		if err != nil {
			log.Printf("Could not trigger event: %s", err)
		}
		return
	}

	log.Printf("GET %s %d", resp.Request.URL, resp.StatusCode)

	if service.IncidentKey != "" {
		err := service.resolveAlert()
		if err != nil {
			log.Printf("Could not resolve event: %s", err)
			return
		}
	}
}

func (service *Service) sendAlert(err error) (string, error) {
	event := pagerduty.Event{
		Type:        "trigger",
		ServiceKey:  service.ServiceKey,
		Description: fmt.Sprintf("Healthcheck on %s failed", service.Endpoint),
		Details:     fmt.Sprintf("Could not reach service: %s", err),
	}

	resp, err := pagerduty.CreateEvent(event)
	if err != nil {
		return "", err
	}

	log.Printf("Triggered incident %s", resp.IncidentKey)
	return resp.IncidentKey, nil
}

func (service *Service) resolveAlert() error {
	event := pagerduty.Event{
		Type:        "resolve",
		ServiceKey:  service.ServiceKey,
		IncidentKey: service.IncidentKey,
		Description: "Resolved",
	}

	_, err := pagerduty.CreateEvent(event)
	if err != nil {
		return err
	}

	log.Printf("Resolved incident %s", service.IncidentKey)
	service.IncidentKey = ""
	return nil
}
