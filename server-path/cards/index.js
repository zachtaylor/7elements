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
	}).then(function() {

		SE.req.getmycards().then(function(data) {
			console.log('completed getmycards', data);
			$.each(data, function(cardid, count) {
				var card = cardsContainer.children[cardid-1];
				if (card) {
					var hoverMessage = $('<span style="color:white;position:absolute;right:0px;top:36px;background:rgba(0,0,0,0.5);">You have '+count+'</span>');
					$(card).append(hoverMessage);
				} else {
					console.error("have copies of missing card", cardid, count);
				}
			});
		});
	});
});
