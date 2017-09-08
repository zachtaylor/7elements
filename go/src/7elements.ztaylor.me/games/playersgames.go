package games

var playersActiveGames = make(map[string]map[int]bool)

func GetActiveGames(username string) map[int]bool {
	if playersActiveGames[username] == nil {
		playersActiveGames[username] = make(map[int]bool)
	}
	return playersActiveGames[username]
}
