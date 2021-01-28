package jsons

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/handler"
)

// JSONDataValidate is convert data interface validate
func JSONDataValidate(data interface{}) interface{} {
	tname := fmt.Sprint(reflect.TypeOf(data))
	tname = strings.ReplaceAll(tname, " ", "")
	if tname == "map[string]interface{}" {
		objval := JSONObjectFactory()
		for key, val := range data.(map[string]interface{}) {
			objval.GetObjectData().Put(key, JSONDataValidate(val))
		}
		return objval
	} else if tname == "[]interface{}" {
		objval := JSONArrayFactory()
		for _, val := range data.([]interface{}) {
			objval.GetObjectData().Put(JSONDataValidate(val))
		}
		return objval
	} else if tname == "float32" {
		a := data.(float32)
		if float64(a) == math.Trunc(float64(a)) {
			return int(a)
		}
	} else if tname == "float64" {
		a := data.(float64)
		if a == math.Trunc(a) {
			return int(a)
		}
	}
	return data
}

// JSONArrayFromString is convert json array string to JSONArray object
func JSONArrayFromString(buffer string) (*JSONArray, error) {
	var data []interface{}
	err := json.Unmarshal([]byte(buffer), &data)
	if handler.Error(err) {
		return nil, err
	}

	datalist := new(JSONArray).Factory()
	for _, value := range data {
		datalist.GetObjectData().Put(JSONDataValidate(value))
	}
	return datalist, nil
}

// JSONObjectFromString is convert json object string to JSONObject object
func JSONObjectFromString(buffer string) (*JSONObject, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(buffer), &data)
	if handler.Error(err) {
		return nil, err
	}

	dataobj := new(JSONObject).Factory()
	for key, value := range data {
		dataobj.GetObjectData().Put(key, JSONDataValidate(value))
	}
	return dataobj, nil
}

// JSONArrayFromFile is load JSONArray object from file
func JSONArrayFromFile(pathfile string) (*JSONArray, error) {
	if files.ExistFile(pathfile) {
		bbyte, err := files.ReadFile(pathfile)
		if !handler.Error(err) {
			buffer := string(bbyte)
			return JSONArrayFromString(buffer)
		}
		return nil, err
	}
	nobj := new(JSONArray).Factory()
	return nobj, nil
}

// JSONObjectFromFile is load JSONObject object from file
func JSONObjectFromFile(pathfile string) (*JSONObject, error) {
	if files.ExistFile(pathfile) {
		bbyte, err := files.ReadFile(pathfile)
		if !handler.Error(err) {
			buffer := string(bbyte)
			return JSONObjectFromString(buffer)
		}
		return nil, err
	}
	nobj := new(JSONObject).Factory()
	return nobj, nil
}
