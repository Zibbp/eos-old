package kv

var kv *KVStore

type KVStore struct {
	db map[string]string
}

func init() {
	kv = &KVStore{
		db: make(map[string]string),
	}
}

func DB() *KVStore {
	return kv
}

func (k *KVStore) Get(key string) (string, bool) {
	v, ok := k.db[key]
	return v, ok
}

func (k *KVStore) Set(key, value string) {
	k.db[key] = value
}
