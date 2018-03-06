SE.widget.control('se-app-packs', function() {
	var me = this;

	$('button', me).click(function() {
		SE.api.get('openpack').then(function(data) {
			SE.api.cache.openpack = null;
			SE.websocket.send('ping');
			vii.sound.play('openpack');

			SE.widget.new('se-card-spinner').then(function(spinner) {
				$.each(data.cards, function(i, cardid) {
					SE.widget.new('se-card', cardid).then(function(card) {
						card.cardid = cardid;
						spinner.append(card);
					});
				});
				var div = $('<div></div>')[0];
				$(div).append('<span class="elemen7s-font-label">New Pack!</span>');
				$(div).append(spinner);
				$(div).append('<br/><br/>');
				$('[handle="contents"]', me).append(div);
				window.scrollTo(0,document.body.scrollHeight)
			});
		}, function() {
			console.error('/api/openpack.json: call failed', arguments);
		});
	});
});
