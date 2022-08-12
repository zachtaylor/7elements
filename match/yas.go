package match

import "taylz.io/yas"

type Cache = yas.SyncMap[*Queue]

func NewCache() *Cache { return yas.NewSyncMap[*Queue]() }
