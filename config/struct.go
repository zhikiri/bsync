package config

import (
	"errors"
	"fmt"
	"os"
)

// Setup app configuration parameters
type Setup struct {
	SourcePath string `json:"source_path"`
	AWS        struct {
		AccessKey string `json:"access_key"`
		SecretKey string `json:"secret_key"`
		Bucket    string `json:"bucket"`
		Region    string `json:"region"`
	} `json:"aws"`
}

func (s *Setup) setAccessKey(v string) error {
	s.AWS.AccessKey = v
	return nil
}

func (s *Setup) setSecretKey(v string) error {
	s.AWS.SecretKey = v
	return nil
}

func (s *Setup) setBucket(v string) error {
	s.AWS.Bucket = v
	return nil
}

func (s *Setup) setRegion(v string) error {
	s.AWS.Region = v
	return nil
}

func (s *Setup) setSourcePath(v string) error {
	_, err := os.Stat(v)
	if os.IsNotExist(err) {
		return errors.New("blog path doesn't exist")
	}
	s.SourcePath = v
	return nil
}

func (s *Setup) save() error {
	fmt.Println("Saving your config...")
	return nil
}