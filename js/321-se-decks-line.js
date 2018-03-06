SE.widget.control('se-decks-line', function(data) {
	var me = this;
	if (data.name) me.name = data.name;
	me.deckid = data.id;
	me.wins = data.wins;
	me.cards = {};

	SE.widget.new('se-card-spinner').then(function(spinner) {
		$(me.spinner).replaceWith(spinner);
		me.spinner = spinner;

		$.each(data.cards, function(cardid, copies) {
			if (me.cards[cardid]) ;
			SE.widget.new('se-card', cardid).then(function(card) {
				me.cards[cardid] = card;
				me.spinner.append(card);
			});
		});
	});

	$('button', me).click(function() {
		vii.cookie('editid', me.deckid);
		window.location.href='/#edit';
	});
});
