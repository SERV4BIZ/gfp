package collection

import (
	"encoding/json"
)

// KeySet is map key string
type KeySet struct {
	length int
	data   map[string]bool
}

// KeySetFactory is global function create a new KeySet object
func KeySetFactory() *KeySet {
	return new(KeySet).Factory()
}

// KeySetNew is global function create a new KeySet object
func KeySetNew() *KeySet {
	return KeySetFactory()
}

// Factory is create a new KeySet
func (me *KeySet) Factory() *KeySet {
	me.length = 0
	me.data = make(map[string]bool)
	return me
}

// New is create a new KeySet
func (me *KeySet) New() *KeySet {
	return me.Factory()
}

// SetData is set data map in KeySet
func (me *KeySet) SetData(data map[string]bool) *KeySet {
	me.length = len(data)
	me.data = data
	return me
}

// GetData is get data map in KeySet
func (me *KeySet) GetData() map[string]bool {
	return me.data
}

// Length is Get size or length or count of Item in KeySet
func (me *KeySet) Length() int {
	return me.length
}

// Size is same Length
func (me *KeySet) Size() int {
	return me.Length()
}

// Clear is remove all item in KeySet
func (me *KeySet) Clear() *KeySet {
	me.length = 0
	me.data = make(map[string]bool)
	return me
}

// Clean is same Clear function
func (me *KeySet) Clean() *KeySet {
	return me.Clear()
}

// ToArray is convert KeySet to array string
func (me *KeySet) ToArray() []string {
	result := make([]string, me.Length())
	var index int
	for key := range me.data {
		result[index] = key
		index++
	}
	return result
}

// ToString is convert KeySet to String
func (me *KeySet) ToString() (string, error) {
	data := me.ToArray()
	buffer, err := json.Marshal(&data)
	return string(buffer), err
}

// Add is add key item to KeySet
func (me *KeySet) Add(key string) *KeySet {
	ok := me.Contains(key)
	if !ok {
		me.length++
		me.data[key] = true
	}
	return me
}

// Put is same Add function
func (me *KeySet) Put(key string) *KeySet {
	return me.Add(key)
}

// Remove is delete key item in KeySet
func (me *KeySet) Remove(key string) *KeySet {
	ok := me.Contains(key)
	if ok {
		delete(me.data, key)
		me.length--
	}
	return me
}

// Delete is same Remove function
func (me *KeySet) Delete(key string) *KeySet {
	return me.Remove(key)
}

// Contains is check key item has in KeySet
func (me *KeySet) Contains(key string) bool {
	_, ok := me.data[key]
	return ok
}

// Check is same Contains function
func (me *KeySet) Check(key string) bool {
	return me.Contains(key)
}

// Has is same Contains function
func (me *KeySet) Has(key string) bool {
	return me.Contains(key)
}
