package collection

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// MapKey is data mapping Key - Value
type MapKey struct {
	length int
	data   map[string]interface{}
}

// MapKeyFactory is global function create a new MapKey object
func MapKeyFactory() *MapKey {
	return new(MapKey).Factory()
}

// MapKeyNew is global function create a new MapKey object
func MapKeyNew() *MapKey {
	return MapKeyFactory()
}

// Factory is create a new MapKey object
func (me *MapKey) Factory() *MapKey {
	me.length = 0
	me.data = make(map[string]interface{})
	return me
}

// New is create a new MapKey object
func (me *MapKey) New() *MapKey {
	return me.Factory()
}

// SetData is set data in KeyMap object
func (me *MapKey) SetData(data map[string]interface{}) *MapKey {
	me.length = len(data)
	me.data = data
	return me
}

// GetData is get data in KeyMap object
func (me *MapKey) GetData() map[string]interface{} {
	return me.data
}

// Length is size or length or count item of MapKey
func (me *MapKey) Length() int {
	return me.length
}

// Size is same Length
func (me *MapKey) Size() int {
	return me.Length()
}

// Clear is delete all item in MapKay
func (me *MapKey) Clear() *MapKey {
	me.length = 0
	me.data = make(map[string]interface{})
	return me
}

// Clean is same Clear function
func (me *MapKey) Clean() *MapKey {
	return me.Clear()
}

// ToMap is convert MapKey to data map
func (me *MapKey) ToMap() map[string]interface{} {
	return me.data
}

// ToString is convert MapKey object to string
func (me *MapKey) ToString() (string, error) {
	buffer, err := json.Marshal(&me.data)
	return string(buffer), err
}

// GetKeys is get all keys in MapKey object
func (me *MapKey) GetKeys() []string {
	keys := make([]string, me.Length())
	var index int
	for k := range me.data {
		keys[index] = k
		index++
	}
	return keys
}

// ContainsKey is check key has in MapKey object
func (me *MapKey) ContainsKey(key string) bool {
	_, ok := me.data[key]
	return ok
}

// CheckKey is same ContainsKey function
func (me *MapKey) CheckKey(key string) bool {
	return me.ContainsKey(key)
}

// HasKey is same ContainsKey function
func (me *MapKey) HasKey(key string) bool {
	return me.ContainsKey(key)
}

// Get is get data from key in MapKey object
func (me *MapKey) Get(key string) interface{} {
	val, ok := me.data[key]
	if ok {
		return val
	}
	return nil
}

// Put is add new Key - Value to MapKey object
func (me *MapKey) Put(key string, value interface{}) *MapKey {
	ok := me.ContainsKey(key)
	if ok {
		me.data[key] = value
	} else {
		me.length++
		me.data[key] = value
	}
	return me
}

// Remove is delete item in MapKey from keyname
func (me *MapKey) Remove(key string) *MapKey {
	ok := me.ContainsKey(key)
	if ok {
		me.length--
		delete(me.data, key)
	}
	return me
}

// Delete is same Remove function
func (me *MapKey) Delete(key string) *MapKey {
	return me.Remove(key)
}

// GetType is get data type in MapKey object from keyname
func (me *MapKey) GetType(key string) string {
	tname := reflect.TypeOf(me.Get(key))
	if tname != nil {
		return fmt.Sprint(tname)
	}
	return "nil"
}
