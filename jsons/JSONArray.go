package jsons

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/SERV4BIZ/gfp/collection"
	"github.com/SERV4BIZ/gfp/filesystem"
	"github.com/SERV4BIZ/gfp/handler"
)

// JSONArray is data struct for JSONArray object
type JSONArray struct {
	datalist *collection.ArrayList
}

// JSONArrayFactory is global create a new JSONArrayFactory
func JSONArrayFactory() *JSONArray {
	return new(JSONArray).Factory()
}

// Factory is create a new JSONArray object
func (me *JSONArray) Factory() *JSONArray {
	nlist := new(collection.ArrayList).Factory()
	me.datalist = nlist
	return me
}

// GetObjectData is get raw data
func (me *JSONArray) GetObjectData() *collection.ArrayList {
	return me.datalist
}

// SetObjectData is set raw data
func (me *JSONArray) SetObjectData(list *collection.ArrayList) *JSONArray {
	me.datalist = list
	return me
}

// PutString is put string data item
func (me *JSONArray) PutString(value string) *JSONArray {
	me.datalist.Put(value)
	return me
}

// PutInt is put int data item
func (me *JSONArray) PutInt(value int) *JSONArray {
	me.datalist.Put(value)
	return me
}

// PutDouble is put double data item
func (me *JSONArray) PutDouble(value float64) *JSONArray {
	me.datalist.Put(value)
	return me
}

// PutFloat is same PutDouble function
func (me *JSONArray) PutFloat(value float64) *JSONArray {
	return me.PutDouble(value)
}

// PutBool is put boolean data item
func (me *JSONArray) PutBool(value bool) *JSONArray {
	me.datalist.Put(value)
	return me
}

// PutNull is put null data item
func (me *JSONArray) PutNull() *JSONArray {
	me.datalist.Put(nil)
	return me
}

// PutObject is put object data item
func (me *JSONArray) PutObject(value *JSONObject) *JSONArray {
	me.datalist.Put(value)
	return me
}

// PutArray is put array data item
func (me *JSONArray) PutArray(value *JSONArray) *JSONArray {
	me.datalist.Put(value)
	return me
}

// Clear is remove all item
func (me *JSONArray) Clear() *JSONArray {
	me.datalist.Clear()
	return me
}

// Clean is same Clear function
func (me *JSONArray) Clean() *JSONArray {
	return me.Clear()
}

// Length is get size or count item of JSONArray object
func (me *JSONArray) Length() int {
	return me.datalist.Length()
}

// Remove is delete item from index
func (me *JSONArray) Remove(index int) *JSONArray {
	me.datalist.Remove(index)
	return me
}

// Delete is same Remove function
func (me *JSONArray) Delete(index int) *JSONArray {
	return me.Remove(index)
}

// GetType is get type data item
func (me *JSONArray) GetType(index int) string {
	if index < 0 || index >= me.Length() {
		return ""
	}

	tname := me.GetObjectData().GetType(index)
	switch tname {
	case "string":
		return "string"
	case "int":
		return "int"
	case "float32":
		a := me.GetObjectData().Get(index).(float32)
		if float64(a) == math.Trunc(float64(a)) {
			return "int"
		}
		return "double"
	case "float64":
		a := me.GetObjectData().Get(index).(float64)
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

// GetString is get string item from index
func (me *JSONArray) GetString(index int) string {
	if index < 0 || index >= me.Length() {
		return ""
	}

	tname := me.GetType(index)
	switch tname {
	case "string":
		return string(me.GetObjectData().Get(index).(string))
	case "int":
		return string(fmt.Sprintf("%d", me.GetInt(index)))
	case "double":
		return string(fmt.Sprintf("%f", me.GetDouble(index)))
	case "bool":
		if me.GetBool(index) {
			return "true"
		}
		return "false"
	case "null":
		return "null"
	case "object":
		return me.GetObject(index).ToString()
	case "array":
		return me.GetArray(index).ToString()
	}
	return ""
}

// GetInt is get int item from index
func (me *JSONArray) GetInt(index int) int {
	if index < 0 || index >= me.Length() {
		return 0
	}

	tname := me.GetType(index)
	switch tname {
	case "string":
		val, _ := strconv.Atoi(me.GetString(index))
		return val
	case "int":
		tp := me.GetObjectData().GetType(index)
		if tp == "float64" {
			return int(me.GetObjectData().Get(index).(float64))
		}
		return int(me.GetObjectData().Get(index).(int))
	case "double":
		return int(me.GetDouble(index))
	case "bool":
		if me.GetBool(index) {
			return 1
		}
		return 0
	case "null":
		return 0
	case "object":
		return me.GetObject(index).Length()
	case "array":
		return me.GetArray(index).Length()
	}
	return 0
}

// GetDouble is get double item from index
func (me *JSONArray) GetDouble(index int) float64 {
	if index < 0 || index >= me.Length() {
		return 0.0
	}

	tname := me.GetType(index)
	switch tname {
	case "string":
		val, _ := strconv.ParseFloat(me.GetString(index), 64)
		return val
	case "int":
		return float64(me.GetInt(index))
	case "double":
		return float64(me.GetObjectData().Get(index).(float64))
	case "bool":
		if me.GetBool(index) {
			return 1.0
		}
		return 0.0
	case "null":
		return 0.0
	case "object":
		return float64(me.GetObject(index).Length())
	case "array":
		return float64(me.GetArray(index).Length())
	}
	return 0.0
}

// GetFloat is same GetDouble
func (me *JSONArray) GetFloat(index int) float64 {
	return me.GetDouble(index)
}

// GetBool is get string item from index
func (me *JSONArray) GetBool(index int) bool {
	if index < 0 || index >= me.Length() {
		return false
	}

	tname := me.GetType(index)
	switch tname {
	case "string":
		val := me.GetString(index)
		return strings.ToLower(val)[0] == "t"[0]
	case "int":
		return me.GetInt(index) > 0
	case "double":
		return me.GetDouble(index) > 0
	case "bool":
		return me.GetObjectData().Get(index).(bool)
	case "null":
		return false
	case "object":
		return me.GetObject(index).Length() > 0
	case "array":
		return me.GetArray(index).Length() > 0
	}
	return false
}

// GetNull is get null item from index
func (me *JSONArray) GetNull(index int) interface{} {
	if index < 0 || index >= me.Length() {
		return nil
	}
	return me.GetObjectData().Get(index)
}

// GetObject is get object item from index
func (me *JSONArray) GetObject(index int) *JSONObject {
	if index < 0 || index >= me.Length() {
		nobj := new(JSONObject).Factory()
		return nobj
	}

	tname := me.GetType(index)
	if tname == "object" {
		return me.GetObjectData().Get(index).(*JSONObject)
	}

	nobj := new(JSONObject).Factory()
	return nobj
}

// GetArray is get array item from index
func (me *JSONArray) GetArray(index int) *JSONArray {
	if index < 0 || index >= me.Length() {
		narr := new(JSONArray).Factory()
		return narr
	}

	tname := me.GetType(index)
	if tname == "array" {
		return me.GetObjectData().Get(index).(*JSONArray)
	}

	narr := new(JSONArray).Factory()
	return narr
}

// ToString is convert to json array string
func (me *JSONArray) ToString() string {
	if me.Length() == 0 {
		return "[]"
	}

	var buff bytes.Buffer
	buff.WriteString("[")
	for i := 0; i < me.Length(); i++ {
		tname := me.GetType(i)
		switch tname {
		case "string":
			val := me.GetString(i)
			str, _ := json.Marshal(val)
			buff.WriteString(string(str))
		case "int":
			val := me.GetInt(i)
			str, _ := json.Marshal(val)
			buff.WriteString(string(str))
		case "double":
			val := me.GetDouble(i)
			str, _ := json.Marshal(val)
			buff.WriteString(string(str))
		case "bool":
			val := me.GetObjectData().Get(i).(bool)
			str, _ := json.Marshal(val)
			buff.WriteString(string(str))
		case "null":
			str, _ := json.Marshal(nil)
			buff.WriteString(string(str))
		case "object":
			val := me.GetObject(i)
			str := val.ToString()
			buff.WriteString(string(str))
		case "array":
			val := me.GetArray(i)
			str := val.ToString()
			buff.WriteString(string(str))
		default:
			str, _ := json.Marshal(nil)
			buff.WriteString(string(str))
		}

		if (i + 1) < me.Length() {
			buff.WriteString(",")
		}
	}
	buff.WriteString("]")
	return buff.String()
}

// ToFile is write JSON array data to file
func (me *JSONArray) ToFile(pathfile string) (int, error) {
	buffer := me.ToString()
	return filesystem.WriteFile(pathfile, []byte(buffer))
}

// FromString is load json array data from string buffer
func (me *JSONArray) FromString(buffer string) (*JSONArray, error) {
	nobj, err := JSONArrayFromString(buffer)
	if handler.Error(err) {
		return nil, err
	}

	me.datalist = nobj.datalist
	nobj = nil
	return me, err
}

// ReadString is same FromString function
func (me *JSONArray) ReadString(buffer string) (*JSONArray, error) {
	return me.FromString(buffer)
}

// FromFile is load json array data from file
func (me *JSONArray) FromFile(pathfile string) (*JSONArray, error) {
	nobj, err := JSONArrayFromFile(pathfile)
	if handler.Error(err) {
		return nil, err
	}

	me.datalist = nobj.datalist
	nobj = nil
	return me, err
}

// ReadFile is same FromFile function
func (me *JSONArray) ReadFile(pathfile string) (*JSONArray, error) {
	return me.FromFile(pathfile)
}

// Fill is merge json array data from other JSONArray object
func (me *JSONArray) Fill(src *JSONArray) *JSONArray {
	for i := 0; i < src.Length(); i++ {
		switch src.GetType(i) {
		case "string":
			me.PutString(src.GetString(i))
		case "int":
			me.PutInt(src.GetInt(i))
		case "double":
			me.PutFloat(src.GetFloat(i))
		case "bool":
			me.PutBool(src.GetBool(i))
		case "null":
			me.PutNull()
		case "object":
			me.PutObject(src.GetObject(i))
		case "array":
			me.PutArray(src.GetArray(i))
		}
	}

	return me
}

// Equals is check JSONArray equal Other JSONArray
func (me *JSONArray) Equals(src *JSONArray) bool {
	if me.Length() != src.Length() {
		return false
	}

	for i := 0; i < me.Length(); i++ {
		switch me.GetType(i) {
		case "string":
			if strings.Compare(me.GetString(i), src.GetString(i)) != 0 {
				return false
			}
		case "int":
			if me.GetInt(i) != src.GetInt(i) {
				return false
			}
		case "double":
			if me.GetDouble(i) != src.GetDouble(i) {
				return false
			}
		case "bool":
			if me.GetBool(i) != src.GetBool(i) {
				return false
			}
		case "null":
			if me.GetBool(i) != src.GetBool(i) {
				return false
			}
		case "object":
			if !me.GetObject(i).Equals(src.GetObject(i)) {
				return false
			}
		case "array":
			if !me.GetArray(i).Equals(src.GetArray(i)) {
				return false
			}
		}
	}

	return true
}

// Copy is clone data myseft to a new JSONArray object
func (me *JSONArray) Copy() (*JSONArray, error) {
	return JSONArrayFromString(me.ToString())
}

// Clone is same Copy function
func (me *JSONArray) Clone() (*JSONArray, error) {
	return me.Copy()
}

// Merge is merge myseft and other JSONArray object to a new JSONArray object
func (me *JSONArray) Merge(src *JSONArray) *JSONArray {
	nobj := new(JSONArray).Factory()
	nobj.Fill(me)
	nobj.Fill(src)
	return nobj
}
