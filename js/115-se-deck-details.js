SE.widget.control('se-deck-details', function(data) {
	var me = this;

	SE.widget.controlProperty(this, 'deckname');
	SE.widget.controlProperty(this, 'wins');

	SE.widget.replace('se-card-spinner', $('.se-deck-details-spinner', me)[0]);

	this.append = function(card) {
		var spinner = $('.se-card-spinner', me)[0];
		if (spinner) spinner.append(card);
	};
	this.prepend = function(card) {
		var spinner = $('.se-card-spinner', me)[0];
		if (spinner) spinner.prepend(card);
	};
	this.empty = function() {
		var spinner = $('.se-card-spinner', me)[0];
		if (spinner) spinner.empty();
	};
	this.update = function(data) {
		me.deckname = data.name;
		me.wins = data.wins;

		me.empty();
		$.each(data.cards, function(cardid, cardcount) {
			SE.widget.new('se-card', cardid).then(function(card) {
				card.banner = cardcount + ' copies';
				me.append(card);
			});
		});
	};

	if (data) this.update(data);
});
