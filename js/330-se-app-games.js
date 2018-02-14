SE.widget.control('se-app-games', function() {
	var me = this;

	me.update = function() {
		if (!me.data) return console.warn('se-app-games update failed');
		$(me.$gameslist).empty();
		$.each(me.data.games, function(gameid, gamedata) {
			SE.widget.new('se-games-line', gamedata).then(function(html) {
				$(me.$gameslist).append(html);
			});
		});
	};

	vii.ping().then(function(data) {
		me.data = data;
		me.update();
	});

	SE.event.on('/match', function(data) {
		console.log('se-app-games watch /match');
		vii.ping().then(function() {
			me.data.games[data.id] = data;
			me.update();
		});
	});

	me.footer = $('<div></div>')[0];
	var btn = $('<button class="vii vii-blue">Click here to search for a new game</button>')[0];
	$(btn).click(function() {
		window.location.href='/#search';
	});
	$(me.footer).append(btn)
});