package organization

import "time"

type Organization struct {
	ID          string    `yaml:"id"`
	Name        string    `yaml:"name"`
	DisplayName string    `yaml:"displayName"`
	CreatedAt   time.Time `yaml:"createdAt"`
	Status      string    `yaml:"status"`
}
