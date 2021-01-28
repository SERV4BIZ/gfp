package collection

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ArrayItem is item of ArrayList
type ArrayItem struct {
	nextNode *ArrayItem
	data     interface{}
}

// ArrayList is link list model container of ArrayItem
type ArrayList struct {
	headNode *ArrayItem
	lastNode *ArrayItem
	length   int

	pointerNode  *ArrayItem
	pointerIndex int
}

// ArrayListFactory is global function create a new ArrayList object
func ArrayListFactory() *ArrayList {
	return new(ArrayList).Factory()
}

// ArrayListNew is global function create a new ArrayList object
func ArrayListNew() *ArrayList {
	return ArrayListFactory()
}

// Factory is create new ArrayList object
func (me *ArrayList) Factory() *ArrayList {
	me.headNode = nil
	me.lastNode = nil
	me.length = 0

	me.pointerNode = nil
	me.pointerIndex = -1
	return me
}

// New is create new ArrayList object
func (me *ArrayList) New() *ArrayList {
	return me.Factory()
}

// Length is size or length or count all of ArrayList
func (me *ArrayList) Length() int {
	return me.length
}

// Size is same Length function
func (me *ArrayList) Size() int {
	return me.Length()
}

// Begin is reset loop befor use Next function
func (me *ArrayList) Begin() {
	me.pointerNode = nil
	me.pointerIndex = -1
}

// Next is check has item in ArrayList and move pointer to next item
func (me *ArrayList) Next() bool {
	if me.pointerNode == nil {
		me.pointerIndex = 0
		me.pointerNode = me.headNode
	} else {
		me.pointerIndex++
		me.pointerNode = me.pointerNode.nextNode
	}
	return me.pointerNode != nil
}

// Fetch is Get data item from pointer current in ArrayList
func (me *ArrayList) Fetch() (index int, node *ArrayItem) {
	index = me.pointerIndex
	node = me.pointerNode
	return
}

// Append is add item to last of ArrayList
func (me *ArrayList) Append(data interface{}) *ArrayList {
	nitem := new(ArrayItem)
	nitem.nextNode = nil
	nitem.data = data

	if me.headNode == nil {
		me.headNode = nitem
		me.lastNode = nitem

		me.pointerNode = nitem
		me.pointerIndex = 0
	} else {
		me.lastNode.nextNode = nitem
		me.lastNode = nitem

		me.pointerNode = me.lastNode
		me.pointerIndex = me.length - 1
	}
	me.length++
	return me
}

// Prepend is add item to first of ArrayList
func (me *ArrayList) Prepend(data interface{}) *ArrayList {
	nitem := new(ArrayItem)
	nitem.nextNode = nil
	nitem.data = data

	if me.headNode == nil {
		me.headNode = nitem
		me.lastNode = nitem

		me.pointerNode = nitem
		me.pointerIndex = 0
	} else {
		nitem.nextNode = me.headNode
		me.headNode = nitem

		me.pointerNode = me.headNode
		me.pointerIndex = 0
	}
	me.length++
	return me
}

// Get is get data from index
func (me *ArrayList) Get(index int) interface{} {
	if index >= 0 && index < me.Length() {
		me.Begin()
		for me.Next() {
			i, node := me.Fetch()
			if i == index {
				return node.data
			}
		}
	}
	return nil
}

// Add is same Append function
func (me *ArrayList) Add(data interface{}) *ArrayList {
	return me.Append(data)
}

// Put is same Append function
func (me *ArrayList) Put(data interface{}) *ArrayList {
	return me.Append(data)
}

// Unshift is same Prepend function
func (me *ArrayList) Unshift(data interface{}) *ArrayList {
	return me.Prepend(data)
}

// Shift is fetch and return first item of ArrayList and remove it
func (me *ArrayList) Shift() interface{} {
	if me.Length() > 0 {
		value := me.Get(0)
		me.Remove(0)
		return value
	}
	return nil
}

// Pop is fetch and return last item of ArrayList and remove it
func (me *ArrayList) Pop() interface{} {
	if me.Length() > 0 {
		value := me.Get(me.Length() - 1)
		me.Remove(me.Length() - 1)
		return value
	}
	return nil
}

// GetType is check data type of ArrayList from index
func (me *ArrayList) GetType(index int) string {
	tname := reflect.TypeOf(me.Get(index))
	if tname != nil {
		return fmt.Sprint(tname)
	}
	return ""
}

// Clear is remove all of ArrayList
func (me *ArrayList) Clear() *ArrayList {
	me.headNode = nil
	me.lastNode = nil
	me.length = 0

	me.pointerNode = nil
	me.pointerIndex = 0
	return me
}

// Clean is same Clear function
func (me *ArrayList) Clean() *ArrayList {
	return me.Clear()
}

// ToArray is convert ArrayList to array
func (me *ArrayList) ToArray() []interface{} {
	results := make([]interface{}, me.Length())
	me.Begin()
	for me.Next() {
		i, node := me.Fetch()
		results[i] = node.data
	}
	return results
}

// ToString is convert ArrayList to string
func (me *ArrayList) ToString() (string, error) {
	data := me.ToArray()
	buffer, err := json.Marshal(&data)
	return string(buffer), err
}

// Remove is delete item of ArrayList from index
func (me *ArrayList) Remove(index int) *ArrayList {
	if index >= 0 && index < me.Length() {
		var backNode *ArrayItem
		me.Begin()
		for me.Next() {
			i, node := me.Fetch()
			if i == index {
				if backNode == nil {
					me.headNode = me.headNode.nextNode
					me.pointerIndex = i
					me.pointerNode = me.headNode

					if me.headNode == nil {
						me.lastNode = me.headNode
					}
				} else {
					backNode.nextNode = node.nextNode
					me.pointerIndex = i
					me.pointerNode = node.nextNode

					if node.nextNode == nil {
						me.lastNode = backNode
					}
				}
				me.length--
				break
			}
			backNode = node
		}

	}
	return me
}

// Delete is same Remove function
func (me *ArrayList) Delete(index int) *ArrayList {
	return me.Remove(index)
}
