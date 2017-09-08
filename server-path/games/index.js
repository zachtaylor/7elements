$(function() {
	SE.api.get('navbar').then(function(data) {
		$.each(data.games, function(gameid, gamedata) {
			console.log('games', gameid, gamedata);
			SE.widget.new('se-game-active', gamedata).then(function(gameActive) {
				$('#games-inprogress').append(gameActive);
			})
		});
	})
});
