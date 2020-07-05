package fldb

import (
	bolt "go.etcd.io/bbolt"
)

type Cache struct {
	db      *bolt.DB
	module  []byte
	prepath []byte
}

func Open(module []byte, prepath []byte) *Cache {
	db, err := bolt.Open("test.db", 0666, nil)
	if err != nil {
		return nil
	}
	return &Cache{
		db:      db,
		module:  module,
		prepath: prepath,
	}
}

//func (cache *Cache) Put(info *rsync.FileInfo) error {
//	key := []byte(cache.prepath + info.Path)
//	value, err := proto.Marshal(&FInfo{
//		Size:  info.Size,
//		Mtime: info.Mtime,
//		Mode:  info.Mode,
//	})
//	if err != nil {
//		return err
//	}
//	return cache.module.Put(key, value)
//}
//
//func (cache *Cache) Get(key []byte) *FInfo {
//	value := cache.module.Get(key)
//	if value != nil {
//		info := &FInfo{}
//		err := proto.Unmarshal(value, info)
//		if err == nil {
//			return info
//		}
//	}
//	return nil
//}
//
//func (cache *Cache) PutAll(list *rsync.FileList) error {
//	for _, info := range *list {
//		err := cache.Put(&info)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}


