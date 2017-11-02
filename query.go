package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func Query(s yaml.MapSlice, path string) (interface{}, error) {
	pathChain := ParsePathString(path)
	return searchRecursive(s, pathChain)
}

func searchRecursive(current interface{}, pathChain []interface{}) (interface{}, error) {
	if len(pathChain) == 0 {
		switch v := current.(type) {
		case yaml.MapSlice:
			return nil, fmt.Errorf("Not exists.")
		case []interface{}:
			return nil, fmt.Errorf("Not exists.")
		default:
			return v, nil
		}
	}

	switch v := current.(type) {
	case yaml.MapSlice:
		if key, ok := pathChain[0].(string); ok {
			for _, item := range v {
				if item.Key == key {
					return searchRecursive(item.Value, pathChain[1:])
				}
			}
		}
	case []interface{}:
		if index, ok := pathChain[0].(int); ok {
			if len(v) > index {
				return searchRecursive(v[index], pathChain[1:])
			}
		}
	}
	return nil, fmt.Errorf("Not exists.")
}
