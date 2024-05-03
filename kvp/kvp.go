package kvp

type KeyTypes interface {
	string | float32 | float64 | int | int8 | int16 | int32 | int64
}

type KeyValuePair[K KeyTypes, V any] struct {
	key   K
	value V
}

func (kvp *KeyValuePair[K, V]) Key() K {
	return kvp.key
}

func (kvp *KeyValuePair[K, V]) Value() V {
	return kvp.value
}

func NewKVP[K KeyTypes, V any](key K, value V) KeyValuePair[K, V] {
	return KeyValuePair[K, V]{
		key:   key,
		value: value,
	}
}
