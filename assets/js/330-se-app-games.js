SE.widget.control('se-app-games', function() {
	var me = this;

	me.update = function(data) {
		if (!data) return console.warn('se-app-games update failed');
		$(me.$gameslist).empty();
		$.each(data.games, function(gameid, gamedata) {
			SE.widget.new('se-games-line', gamedata).then(function(html) {
				$(me.$gameslist).append(html);
			});
		});
	};

	vii.ping().then(function(data) {
		me.update(data);
	});

	SE.event.on('/match', function() {
		console.log('se-app-games watch /match');
		vii.ping().then(function(data) {
			me.update(data);
		});
	});

	me.footer = $('<div></div>')[0];
	var btn = $('<button class="vii vii-blue">Click here to search for a new game</button>')[0];
	$(btn).click(function() {
		window.location.href='/#search';
	});
	$(me.footer).append(btn)
});