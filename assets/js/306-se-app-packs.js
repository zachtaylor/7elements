SE.widget.control('se-app-packs', function() {
	var me = this;

	me.footer = $('<div><button class="vii">Click here to open a pack of cards</button></div>')[0];

	$('button', me.footer).click(function() {
		SE.api.get('openpack').then(function(data) {
			console.debug('opened pack', data);
			SE.api.cache.openpack = null;
			SE.websocket.send('ping');

			SE.widget.new('se-card-spinner').then(function(spinner) {
				SE.widget.get('se-card').then(function(cardFactory) {
					$.each(data.cards, function(i, cardid) {
						spinner.append(cardFactory(cardid));
					});
					var div = $('<div></div>')[0];
					$(div).append('<span">New Pack!</span>');
					$(div).append(spinner);
					$(div).append('<br/><br/>');
					$('[handle="content"]', me).prepend(div);
					console.debug('append pack html', div);
					window.scrollTo(0,document.body.scrollHeight)
				})
			});
		}, function() {
			console.error('/api/openpack.json: call failed', arguments);
		});
	});

	me.update = function(data) {
		if (!data) return console.warn('se-app-packs update failed');
		me.packs = data.packs;
	};

	vii.ping().then(function(data) {
		me.update(data);
	});
});
