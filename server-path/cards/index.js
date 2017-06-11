$(function() {
	var cardsContainer = $('#cards-container')[0];
	cardsContainer.children = {};

	SE.req.getcards().then(function(data) {
		$.each(data, function(cardid, carddata) {
			var card = cardsContainer.children[cardid];

			if (!card) {
				card = $('<se-card></se-card>')[0];
				card.cardid = cardid;
				$(cardsContainer).append(card);
				cardsContainer.children[cardid] = card;
			}
		});
	}, function(error) {
		console.error(error);
	})
});
