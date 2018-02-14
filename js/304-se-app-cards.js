SE.widget.control('se-app-cards', function() {
	var me = this;

	vii.ping().then(function(data) {
		$.each(data.cards, function(cardid, carddata) {
			SE.widget.new('se-card', cardid).then(function(card) {
				if (carddata.copies) card.banner = carddata.copies+' copies';
				$(card).css({'margin':'2px'});
				$(me).append(card);
			}, function(e) {
				if (e.status != 401) console.error(e.responseText);
			});
		});
	}, function(error) {
		console.error(error);
	});
});