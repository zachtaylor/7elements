package player

func (cache *Cache) waitPlayer(player *T) {
	<-player.Session.Done()
	cache.Delete(player.Account.Username)
}
