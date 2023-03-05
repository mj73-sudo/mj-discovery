package main

import (
	"sync"
	"time"
)

type Service struct {
	Name           string    `json:"name"`
	InstanceId     string    `json:"instanceId"`
	Ip             string    `json:"ip"`
	Port           string    `json:"port"`
	RegisterDate   time.Time `json:"registerDate"`
	HealthCheckUrl string    `json:"healthCheckUrl"`
}

var services = make(map[string]*[]Service, 0)

var lock = &sync.Mutex{}

func (r *Service) SaveNewService() {
	lock.Lock()
	value, exist := services[r.Name]
	if exist {
		existValue := false
		for _, service := range *value {
			if service.Ip == r.Ip && service.Port == r.Port {
				existValue = true
				break
			}
		}
		if !existValue {
			*value = append(*value, *r)
		}
	} else {
		serviceNewList := make([]Service, 0)
		serviceNewList = append(serviceNewList, *r)
		services[r.Name] = &serviceNewList
	}
	lock.Unlock()
}

func (r *Service) DeleteService() {
	lock.Lock()
	delete(services, r.Name)
	lock.Unlock()
}
