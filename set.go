package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func Set(s yaml.MapSlice, path string, value interface{}) (yaml.MapSlice, error) {
	pathChain := ParsePathString(path)
	return set(s, pathChain, value)
}

func set(before yaml.MapSlice, pathChain []interface{}, value interface{}) (yaml.MapSlice, error) {
	result, err := setRecursive(before, pathChain, value)
	if err != nil {
		return nil, err
	}

	after, ok := result.(yaml.MapSlice)
	if !ok {
		return nil, fmt.Errorf("Cannot set: unexpected error.")
	}
	return after, nil
}

func setRecursive(current interface{}, pathChain []interface{}, value interface{}) (interface{}, error) {
	switch v := current.(type) {
	case yaml.MapSlice:
		if key, ok := pathChain[0].(string); ok {
			pathChain = pathChain[1:]
			for i, item := range v {
				if item.Key == key {
					if len(pathChain) == 0 {
						switch item.Value.(type) {
						case yaml.MapSlice:
							return nil, fmt.Errorf("Cannot set: path indicates hash.")
						case []interface{}:
							return nil, fmt.Errorf("Cannot set: path indicates array.")
						}
						v[i].Value = value
						return v, nil
					}

					var err error = nil
					v[i].Value, err = setRecursive(v[i].Value, pathChain, value)
					if err != nil {
						return nil, err
					}
					return v, nil
				}
			}
			item := yaml.MapItem{}
			item.Key = key
			if len(pathChain) == 0 {
				item.Value = value
				v = append(v, item)
				return v, nil
			}
			if len(pathChain) > 0 {
				if _, isInt := pathChain[0].(int); isInt {
					newArray, err := createArray(pathChain)
					if err != nil {
						return nil, err
					}
					item.Value = newArray
				} else {
					item.Value = make(yaml.MapSlice, 0)
				}
			}
			v = append(v, item)

			var err error = nil
			v[len(v)-1].Value, err = setRecursive(v[len(v)-1].Value, pathChain, value)
			if err != nil {
				return nil, err
			}
			return v, nil
		}
	case []interface{}:
		if index, ok := pathChain[0].(int); ok {
			pathChain = pathChain[1:]
			if len(v) > index {
				if len(pathChain) == 0 {
					switch v[index].(type) {
					case yaml.MapSlice:
						return nil, fmt.Errorf("Cannot set: path indicates hash.")
					case []interface{}:
						return nil, fmt.Errorf("Cannot set: path indicates array.")
					}
					v[index] = value
					return v, nil
				}
			} else if len(v) == index {
				if len(pathChain) == 0 {
					v = append(v, value)
					return v, nil
				} else if _, isInt := pathChain[0].(int); isInt {
					newArray, err := createArray(pathChain)
					if err != nil {
						return nil, err
					}
					v = append(v, newArray)
				} else {
					s := make(yaml.MapSlice, 0)
					v = append(v, s)
				}
			} else {
				return nil, fmt.Errorf("Cannot set: array index out of bounds.")
			}

			var err error = nil
			v[index], err = setRecursive(v[index], pathChain, value)
			if err != nil {
				return nil, err
			}
			return v, nil
		}
	}
	return nil, fmt.Errorf("Cannot set: route of path is neither hash nor array.")
}

func createArray(pathChain []interface{}) ([]interface{}, error) {
	if len(pathChain) == 0 {
		return nil, fmt.Errorf("Cannot set: array index out of bounds.")
	}

	newArray := make([]interface{}, 0)
	current := newArray
	for {
		if len(pathChain) == 0 {
			current = append(current, nil)
			break
		}

		index, ok := pathChain[0].(int)
		if !ok {
			s := make(yaml.MapSlice, 0)
			current = append(current, s)
			break
		}
		if index > 0 {
			return nil, fmt.Errorf("Cannot set: array index out of bounds.")
		}

		a := make([]interface{}, 0)
		current = append(current, a)
		current = a

		pathChain = pathChain[1:]
	}
	return newArray, nil
}
