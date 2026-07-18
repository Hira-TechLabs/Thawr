package application

import "time"

type Application struct {
	ID          string    `yaml:"id"`
	Name        string    `yaml:"name"`
	DisplayName string    `yaml:"displayName"`
	CreatedAt   time.Time `yaml:"createdAt"`
	Status      string    `yaml:"status"`
}
