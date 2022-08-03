package mapstructure

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
)

func WeakMap2Struct(input interface{}, output interface{}, tags ...string) error {
	tag := "json"
	if len(tags) > 0 {
		tag = tags[0]
	}
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           output,
		TagName:          tag,
	}
	if len(tags) > 0 {
		config.TagName = tags[0]
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

func Map2Struct(input map[string]interface{}, output interface{}, tags ...string) error {
	tag := "json"
	if len(tags) > 0 {
		tag = tags[0]
	}
	config := &mapstructure.DecoderConfig{
		Result:  output,
		TagName: tag,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

func Struct2Map(input interface{}) map[string]interface{} {
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
