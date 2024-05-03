package dictionary

import (
	"github.com/robertjshirts/data-structures/kvp"
)

type dictionary[K kvp.KeyTypes, V any] struct {
	elements []kvp.KeyValuePair[K, V]
}

func NewDict[K kvp.KeyTypes, V any]() *dictionary[K, V] {
	// Please note that this is a SLOW dictionary
	// We don't got no O(1) access times here.
	// It's O(n). (bc i do not know how hashmap :( sad)
	return &dictionary[K, V]{}
}

func (d *dictionary[K, V]) Add(key K, value V) {
	// Look for duplicate key, and replace if exists
	for i, element := range d.elements {
		if element.Key() == key {
			// Reassign element
			d.elements[i] = kvp.NewKVP(key, value)
			return
		}
	}
	// If we didn't find a matching key, append kvp
	d.elements = append(d.elements, kvp.NewKVP(key, value))
}

func (d *dictionary[K, V]) AddKVP(kvp kvp.KeyValuePair[K, V]) {
	// Look for duplicate key, and replace if exists
	for i, element := range d.elements {
		if element.Key() == kvp.Key() {
			// Reassign element
			d.elements[i] = kvp
			return
		}
	}
	// If we didn't find a matching key, append kvp
	d.elements = append(d.elements, kvp)
}

func (d *dictionary[K, V]) Get(key K) *V {
	// Just a wrapper around the GetKVP func
	kvp := d.GetKVP(key)
	if kvp == nil {
		return nil
	}
	value := kvp.Value()
	return &value
}

func (d *dictionary[K, V]) GetKVP(key K) *kvp.KeyValuePair[K, V] {
	// Iterate through slice of elements
	for _, element := range d.elements {
		if element.Key() == key {
			// Return element if found
			return &element
		}
	}
	// Else return nil
	return nil
}

func (d *dictionary[K, V]) Remove(key K) bool {
	// Iterate through the elements
	for i, element := range d.elements {
		if element.Key() == key {
			// If we find it, append all elements besides it
			// We use ... to unwrap the slice into individual
			// variables (for the append function)
			d.elements = append(d.elements[:i], d.elements[i+1:]...)
			return true
		}
	}
	// If we don't return true during that loop, then there
	// Isn't an element to remove
	return false
}

func (d *dictionary[K, V]) GetKVPs() []kvp.KeyValuePair[K, V] {
	// Return a copy, not the reference to the array
	return d.elements
}
