package account

import "taylz.io/yas"

type Cache = yas.SyncMap[*T]

func NewCache() *Cache { return yas.NewSyncMap[*T]() }
