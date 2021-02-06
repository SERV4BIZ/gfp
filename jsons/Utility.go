package jsons

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/SERV4BIZ/gfp/files"
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
	if err != nil {
		return nil, err
	}

	datalist := new(JSONArray).Factory()
	for _, value := range data {
		datalist.GetObjectData().Put(JSONDataValidate(value))
	}
	return datalist, nil
}

// JSONArrayString is convert json array string to JSONArray object
func JSONArrayString(buffer string) (*JSONArray, error) {
	return JSONArrayFromString(buffer)
}

// ArrayString is convert json array string to JSONArray object
func ArrayString(buffer string) (*JSONArray, error) {
	return JSONArrayFromString(buffer)
}

// JSONArrayParse is parse from slice array to json array object
func JSONArrayParse(v interface{}) (*JSONArray, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return JSONArrayFromString(string(b))
}

// ArrayParse is parse from slice array to json array object
func ArrayParse(v interface{}) (*JSONArray, error) {
	return JSONArrayParse(v)
}

// JSONObjectFromString is convert json object string to JSONObject object
func JSONObjectFromString(buffer string) (*JSONObject, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(buffer), &data)
	if err != nil {
		return nil, err
	}

	dataobj := new(JSONObject).Factory()
	for key, value := range data {
		dataobj.GetObjectData().Put(key, JSONDataValidate(value))
	}
	return dataobj, nil
}

// JSONObjectString is convert json object string to JSONObject object
func JSONObjectString(buffer string) (*JSONObject, error) {
	return JSONObjectFromString(buffer)
}

// ObjectString is convert json object string to JSONObject object
func ObjectString(buffer string) (*JSONObject, error) {
	return JSONObjectFromString(buffer)
}

// JSONObjectParse is parse from mapkey object to json object
func JSONObjectParse(v interface{}) (*JSONObject, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return JSONObjectFromString(string(b))
}

// ObjectParse is parse from mapkey object to json object
func ObjectParse(v interface{}) (*JSONObject, error) {
	return JSONObjectParse(v)
}

// JSONArrayFromFile is load JSONArray object from file
func JSONArrayFromFile(pathfile string) (*JSONArray, error) {
	if files.ExistFile(pathfile) {
		bbyte, err := files.ReadFile(pathfile)
		if err != nil {
			return nil, err
		}
		buffer := string(bbyte)
		return JSONArrayFromString(buffer)
	}
	nobj := new(JSONArray).Factory()
	return nobj, nil
}

// JSONArrayFile is load JSONArray object from file
func JSONArrayFile(pathfile string) (*JSONArray, error) {
	return JSONArrayFromFile(pathfile)
}

// ArrayFile is load JSONArray object from file
func ArrayFile(pathfile string) (*JSONArray, error) {
	return JSONArrayFromFile(pathfile)
}

// JSONObjectFromFile is load JSONObject object from file
func JSONObjectFromFile(pathfile string) (*JSONObject, error) {
	if files.ExistFile(pathfile) {
		bbyte, err := files.ReadFile(pathfile)
		if err != nil {
			return nil, err
		}
		buffer := string(bbyte)
		return JSONObjectFromString(buffer)
	}
	nobj := new(JSONObject).Factory()
	return nobj, nil
}

// JSONObjectFile is load JSONObject object from file
func JSONObjectFile(pathfile string) (*JSONObject, error) {
	return JSONObjectFromFile(pathfile)
}

// ObjectFile is load JSONObject object from file
func ObjectFile(pathfile string) (*JSONObject, error) {
	return JSONObjectFromFile(pathfile)
}
