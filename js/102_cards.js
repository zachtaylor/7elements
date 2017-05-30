SE.cards = {
	get: function() {
		return new Promise(function(resolve, reject) {
			$.getJSON('/api/cards.json').done(resolve).fail(reject);
		});
	}
};
