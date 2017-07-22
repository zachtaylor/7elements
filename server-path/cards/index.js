$(function() {
	var cardsContainer = $('#cards-container')[0];
	cardsContainer.children = {};

	SE.api.get('cards').then(function(cards) {
		$.each(cards, function(cardid, carddata) {
			var card = cardsContainer.children[cardid];

			if (!card) {
				card = $('<se-card></se-card>')[0];
				card.cardid = cardid;
				$(cardsContainer).append(card);
				cardsContainer.children[cardid] = card;
			}
		});

		SE.api.get('mycards').then(function(mycards) {
			$.each(mycards, function(cardid, count) {
				var card = cardsContainer.children[cardid-1];
				if (card) {
					$(card).append($('<span style="color:white;position:absolute;right:0px;top:36px;background:rgba(0,0,0,0.5);">You have '+count+'</span>'));
				} else {
					console.error("have copies of missing card", cardid, count);
				}
			});
		});
	}, function(error) {
		console.error(error);
	})
});
