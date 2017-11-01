package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

func Query(s yaml.MapSlice, path string) (interface{}, error) {
	pathChain := parsePathString(path)
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

func parsePathString(path string) []interface{} {
	pathChain := make([]interface{}, 0)
	dotSplited := strings.Split(path, ".")
	for _, keyStr := range dotSplited {
		key, indexes := parseKeyString(keyStr)
		pathChain = append(pathChain, key)
		if indexes != nil {
			pathChain = append(pathChain, indexes...)
		}
	}
	return pathChain
}

func parseKeyString(key string) (string, []interface{}) {
	r := regexp.MustCompile(`^([^\[]+)((?:\[(?:\d+)\])+)$`)
	keyAndIndexes := r.FindStringSubmatch(key)
	if keyAndIndexes == nil {
		return key, nil
	}

	r = regexp.MustCompile(`\d+`)
	indexStrings := r.FindAllString(keyAndIndexes[2], -1)

	indexes := make([]interface{}, len(indexStrings))
	for i, indexString := range indexStrings {
		indexes[i], _ = strconv.Atoi(indexString)
	}

	return keyAndIndexes[1], indexes
}
