package jsons

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/SERV4BIZ/collection"
	"github.com/SERV4BIZ/files"
	"github.com/SERV4BIZ/handler"
)

// JSONObject is data struct JSONObject object
type JSONObject struct {
	datamap *collection.MapKey
}

// JSONObjectFactory is global create a new JSONObject object
func JSONObjectFactory() *JSONObject {
	return new(JSONObject).Factory()
}

// Factory is create a new JSONObject object
func (me *JSONObject) Factory() *JSONObject {
	mapkey := new(collection.MapKey).Factory()
	me.datamap = mapkey
	return me
}

// GetObjectData is get raw data
func (me *JSONObject) GetObjectData() *collection.MapKey {
	return me.datamap
}

// SetObjectData is set raw data
func (me *JSONObject) SetObjectData(mapkey *collection.MapKey) *JSONObject {
	me.datamap = mapkey
	return me
}

// ContainsKey is check has key
func (me *JSONObject) ContainsKey(key string) bool {
	return me.datamap.ContainsKey(key)
}

// CheckKey is same ContainsKey
func (me *JSONObject) CheckKey(key string) bool {
	return me.ContainsKey(key)
}

// HasKey is same ContainsKey
func (me *JSONObject) HasKey(key string) bool {
	return me.ContainsKey(key)
}

// PutString is put string data item
func (me *JSONObject) PutString(key string, value string) *JSONObject {
	me.datamap.Put(key, value)
	return me
}

// PutInt is put int data item
func (me *JSONObject) PutInt(key string, value int) *JSONObject {
	me.datamap.Put(key, value)
	return me
}

// PutDouble is put double data item
func (me *JSONObject) PutDouble(key string, value float64) *JSONObject {
	me.datamap.Put(key, value)
	return me
}

// PutFloat is same PutDouble function
func (me *JSONObject) PutFloat(key string, value float64) *JSONObject {
	return me.PutDouble(key, value)
}

// PutBool is put boolean data item
func (me *JSONObject) PutBool(key string, value bool) *JSONObject {
	me.datamap.Put(key, value)
	return me
}

// PutNull is put null data item
func (me *JSONObject) PutNull(key string) *JSONObject {
	me.datamap.Put(key, nil)
	return me
}

// PutObject is put object data item
func (me *JSONObject) PutObject(key string, value *JSONObject) *JSONObject {
	me.datamap.Put(key, value)
	return me
}

// PutArray is put array data item
func (me *JSONObject) PutArray(key string, value *JSONArray) *JSONObject {
	me.datamap.Put(key, value)
	return me
}

// Clear is remove all
func (me *JSONObject) Clear() *JSONObject {
	me.datamap.Clear()
	return me
}

// Clean is same Clear function
func (me *JSONObject) Clean() *JSONObject {
	return me.Clear()
}

// Length is get size or count key of JSONObject object
func (me *JSONObject) Length() int {
	return me.datamap.Length()
}

// Remove is delete data item from key
func (me *JSONObject) Remove(key string) *JSONObject {
	me.datamap.Remove(key)
	return me
}

// Delete is same Remove function
func (me *JSONObject) Delete(key string) *JSONObject {
	return me.Remove(key)
}

// GetType is get data type of key
func (me *JSONObject) GetType(key string) string {
	if !me.ContainsKey(key) {
		return ""
	}

	tname := me.GetObjectData().GetType(key)
	switch tname {
	case "string":
		return "string"
	case "int":
		return "int"
	case "float32":
		a := me.GetObjectData().Get(key).(float32)
		if float64(a) == math.Trunc(float64(a)) {
			return "int"
		}
		return "double"
	case "float64":
		a := me.GetObjectData().Get(key).(float64)
		if a == math.Trunc(a) {
			return "int"
		}
		return "double"
	case "bool":
		return "bool"
	case "nil":
		return "null"
	case "<nil>":
		return "null"
	default:
		if strings.Contains(tname, "JSONObject") {
			return "object"
		}
		if strings.Contains(tname, "JSONArray") {
			return "array"
		}
	}

	return ""
}

// GetString is get string data from key
func (me *JSONObject) GetString(key string) string {
	if !me.ContainsKey(key) {
		return ""
	}
	tname := me.GetType(key)
	switch tname {
	case "string":
		return string(me.GetObjectData().Get(key).(string))
	case "int":
		return string(fmt.Sprintf("%d", me.GetInt(key)))
	case "double":
		return string(fmt.Sprintf("%f", me.GetDouble(key)))
	case "bool":
		if me.GetBool(key) {
			return "true"
		}
		return "false"
	case "null":
		return "null"
	case "object":
		return me.GetObject(key).ToString()
	case "array":
		return me.GetArray(key).ToString()
	}
	return ""
}

// GetInt is get int data from key
func (me *JSONObject) GetInt(key string) int {
	if !me.ContainsKey(key) {
		return 0
	}

	tname := me.GetType(key)
	switch tname {
	case "string":
		val, _ := strconv.Atoi(me.GetString(key))
		return val
	case "int":
		tp := me.GetObjectData().GetType(key)
		if tp == "float64" {
			return int(me.GetObjectData().Get(key).(float64))
		}
		return int(me.GetObjectData().Get(key).(int))
	case "double":
		return int(me.GetDouble(key))
	case "bool":
		if me.GetBool(key) {
			return 1
		}
		return 0
	case "null":
		return 0
	case "object":
		return me.GetObject(key).Length()
	case "array":
		return me.GetArray(key).Length()
	}
	return 0
}

// GetDouble is get double data from key
func (me *JSONObject) GetDouble(key string) float64 {
	if !me.ContainsKey(key) {
		return 0.0
	}

	tname := me.GetType(key)
	switch tname {
	case "string":
		val, _ := strconv.ParseFloat(me.GetString(key), 64)
		return val
	case "int":
		return float64(me.GetInt(key))
	case "double":
		return float64(me.GetObjectData().Get(key).(float64))
	case "bool":
		if me.GetBool(key) {
			return 1.0
		}
		return 0.0
	case "null":
		return 0.0
	case "object":
		return float64(me.GetObject(key).Length())
	case "array":
		return float64(me.GetArray(key).Length())
	}
	return 0.0
}

// GetFloat is same GetDouble
func (me *JSONObject) GetFloat(key string) float64 {
	return me.GetDouble(key)
}

// GetBool is get boolean data from key
func (me *JSONObject) GetBool(key string) bool {
	if !me.ContainsKey(key) {
		return false
	}

	tname := me.GetType(key)
	switch tname {
	case "string":
		val := me.GetString(key)
		return strings.ToLower(val)[0] == "t"[0]
	case "int":
		return me.GetInt(key) > 0
	case "double":
		return me.GetDouble(key) > 0
	case "bool":
		return me.GetObjectData().Get(key).(bool)
	case "null":
		return false
	case "object":
		return me.GetObject(key).Length() > 0
	case "array":
		return me.GetArray(key).Length() > 0
	}
	return false
}

// GetNull is get null data from key
func (me *JSONObject) GetNull(key string) interface{} {
	if !me.ContainsKey(key) {
		return nil
	}
	return me.GetObjectData().Get(key)
}

// GetObject is get object data from key
func (me *JSONObject) GetObject(key string) *JSONObject {
	if !me.ContainsKey(key) {
		nobj := new(JSONObject).Factory()
		return nobj
	}

	tname := me.GetType(key)
	if tname == "object" {
		return me.GetObjectData().Get(key).(*JSONObject)
	}

	nobj := new(JSONObject).Factory()
	return nobj
}

// GetArray is get array data from key
func (me *JSONObject) GetArray(key string) *JSONArray {
	if !me.ContainsKey(key) {
		narr := new(JSONArray).Factory()
		return narr
	}

	tname := me.GetType(key)
	if tname == "array" {
		return me.GetObjectData().Get(key).(*JSONArray)
	}

	narr := new(JSONArray).Factory()
	return narr
}

// GetKeys is get all keys in JSONObject object
func (me *JSONObject) GetKeys() []string {
	return me.datamap.GetKeys()
}

// ToString is convert to json object string
func (me *JSONObject) ToString() string {
	if me.Length() == 0 {
		return "{}"
	}

	var buff bytes.Buffer
	buff.WriteString("{")
	keys := me.GetKeys()
	for i, key := range keys {
		strkey, _ := json.Marshal(key)
		tname := me.GetType(key)
		switch tname {
		case "string":
			val := me.GetString(key)
			strval, _ := json.Marshal(val)
			buff.WriteString(fmt.Sprintf("%s:%s", string(strkey), string(strval)))
		case "int":
			val := me.GetInt(key)
			strval, _ := json.Marshal(val)
			buff.WriteString(fmt.Sprintf("%s:%s", string(strkey), string(strval)))
		case "double":
			val := me.GetDouble(key)
			strval, _ := json.Marshal(val)
			buff.WriteString(fmt.Sprintf("%s:%s", string(strkey), string(strval)))
		case "bool":
			val := me.GetObjectData().Get(key).(bool)
			strval, _ := json.Marshal(val)
			buff.WriteString(fmt.Sprintf("%s:%s", string(strkey), string(strval)))
		case "null":
			strval, _ := json.Marshal(nil)
			buff.WriteString(fmt.Sprintf("%s:%s", string(strkey), string(strval)))
		case "object":
			val := me.GetObject(key)
			strval := val.ToString()
			buff.WriteString(fmt.Sprintf("%s:%s", string(strkey), string(strval)))
		case "array":
			val := me.GetArray(key)
			strval := val.ToString()
			buff.WriteString(fmt.Sprintf("%s:%s", string(strkey), string(strval)))
		default:
			strval, _ := json.Marshal(nil)
			buff.WriteString(fmt.Sprintf("%s:%s", string(strkey), string(strval)))
		}

		if (i + 1) < me.Length() {
			buff.WriteString(",")
		}
	}
	buff.WriteString("}")
	return buff.String()
}

// ToFile to write json object data to file
func (me *JSONObject) ToFile(pathfile string) (int, error) {
	return filesystem.WriteFile(pathfile, []byte(me.ToString()))
}

// FromString is load json object data from string
func (me *JSONObject) FromString(buffer string) (*JSONObject, error) {
	nobj, err := JSONObjectFromString(buffer)
	if handler.Error(err) {
		return nil, err
	}

	me.datamap = nobj.datamap
	nobj = nil
	return me, err
}

// ReadString is same FromString function
func (me *JSONObject) ReadString(buffer string) (*JSONObject, error) {
	return me.FromString(buffer)
}

// FromFile is load json object data from file
func (me *JSONObject) FromFile(pathfile string) (*JSONObject, error) {
	nobj, err := JSONObjectFromFile(pathfile)
	if handler.Error(err) {
		return nil, err
	}

	me.datamap = nobj.datamap
	nobj = nil
	return me, err
}

// ReadFile is same FromFile function
func (me *JSONObject) ReadFile(pathfile string) (*JSONObject, error) {
	return me.FromFile(pathfile)
}

// Fill is merge myseft with other JSONObject object
func (me *JSONObject) Fill(src *JSONObject) *JSONObject {
	keys := src.GetKeys()
	for _, key := range keys {
		switch src.GetType(key) {
		case "string":
			me.PutString(key, src.GetString(key))
		case "int":
			me.PutInt(key, src.GetInt(key))
		case "double":
			me.PutFloat(key, src.GetFloat(key))
		case "bool":
			me.PutBool(key, src.GetBool(key))
		case "null":
			me.PutNull(key)
		case "object":
			me.PutObject(key, src.GetObject(key))
		case "array":
			me.PutArray(key, src.GetArray(key))
		}
	}

	return me
}

// Equals is check JSONObject equal Other JSONObject
func (me *JSONObject) Equals(src *JSONObject) bool {
	if me.Length() != src.Length() {
		return false
	}

	keys := me.GetKeys()
	for _, key := range keys {
		switch me.GetType(key) {
		case "string":
			if strings.Compare(me.GetString(key), src.GetString(key)) != 0 {
				return false
			}
		case "int":
			if me.GetInt(key) != src.GetInt(key) {
				return false
			}
		case "double":
			if me.GetDouble(key) != src.GetDouble(key) {
				return false
			}
		case "bool":
			if me.GetBool(key) != src.GetBool(key) {
				return false
			}
		case "null":
			if me.GetBool(key) != src.GetBool(key) {
				return false
			}
		case "object":
			if !me.GetObject(key).Equals(src.GetObject(key)) {
				return false
			}
		case "array":
			if !me.GetArray(key).Equals(src.GetArray(key)) {
				return false
			}
		}
	}

	return true
}

// Copy is clone data myseft to a new JSONObject object
func (me *JSONObject) Copy() (*JSONObject, error) {
	return JSONObjectFromString(me.ToString())
}

// Clone is same Copy function
func (me *JSONObject) Clone() (*JSONObject, error) {
	return me.Copy()
}

// Merge is merge myseft and other JSONArray object to a new JSONArray object
func (me *JSONObject) Merge(src *JSONObject) *JSONObject {
	nobj := new(JSONObject).Factory()
	nobj.Fill(me)
	nobj.Fill(src)
	return nobj
}
