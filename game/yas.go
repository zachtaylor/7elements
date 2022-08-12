package game

import "taylz.io/yas"

type Cache = yas.SyncMap[*G]

func NewCache() *Cache { return yas.NewSyncMap[*G]() }
