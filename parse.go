package main

import (
	"gopkg.in/yaml.v2"
)

func Parse(data []byte) (yaml.MapSlice, error) {
	s := make(yaml.MapSlice, 0)
	err := yaml.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return s, err
}
