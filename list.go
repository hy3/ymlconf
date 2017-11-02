package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Property struct {
	Path  string
	Value interface{}
}

type Properties []*Property

func List(s yaml.MapSlice) Properties {
	return listRecursive(nil, s, "")
}

func listRecursive(props Properties, slice yaml.MapSlice, basePath string) Properties {
	if props == nil {
		props = make(Properties, 0)
	}
	for _, item := range slice {
		path, _ := item.Key.(string)
		if basePath != "" {
			path = basePath + "." + path
		}

		switch value := item.Value.(type) {
		case yaml.MapSlice:
			props = listRecursive(props, value, path)
		case []interface{}:
			props = listArrayRecursive(props, value, path)
		default:
			p := new(Property)
			p.Path = path
			p.Value = value
			props = append(props, p)
		}
	}
	return props
}

func listArrayRecursive(props Properties, array []interface{}, basePath string) Properties {
	for i, v := range array {
		path := fmt.Sprintf("%s[%d]", basePath, i)
		switch value := v.(type) {
		case yaml.MapSlice:
			props = listRecursive(props, value, path)
		case []interface{}:
			props = listArrayRecursive(props, value, path)
		default:
			p := new(Property)
			p.Path = path
			p.Value = value
			props = append(props, p)
		}
	}
	return props
}
