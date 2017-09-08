$(function() {
	var cardsContainer = $('#cards-container')[0];
	cardsContainer.children = {};

	SE.api.get('cards').then(function(cardsdata) {
		$(cardsContainer).empty();
		$.each(cardsdata, function(cardid, carddata) {
			SE.widget.new('se-card', cardid).then(function(card) {
				SE.api.get('navbar').then(function(navbardata) {
					if (navbardata.cards[cardid] > 0) card.banner = 'You have '+navbardata.cards[cardid];
				}, function(e) {
					if (e.status != 401) console.error(e.responseText);
				});
				$(cardsContainer).append(card);
			});
		});
	}, function(error) {
		console.error(error);
	});
});
