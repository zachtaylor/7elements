SE.widget.control('se-app-decks', function() {
	var me = this;

	vii.ping().then(function(data) {
		if (!data.decks) $(me).prepend('<h2>Decks Screen: <a href="/login/">Login</a> Required</h2>');
		else $.each(data.decks, function(i, deck) {
			SE.widget.new('se-decks-line', deck).then(function(html) {
				$(me).append(html);
			});
		});
	});
});
