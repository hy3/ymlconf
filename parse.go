package main

import (
	"regexp"
	"strconv"
	"strings"

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

func ParsePathString(path string) []interface{} {
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
