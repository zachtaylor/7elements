$(function() {
	var cards = {};

	SE.req.getcards().then(function(data) {
		$.each(data, function(cardid, card) {
			cards[cardid] = card;
		});
		console.log(cards);
	}, function(error) {
		console.error(error);
	});

	$('#packs-control-buttons-open').on('click', function() {
		SE.req.openpack().then(function(data) {
			$('#packs-opened-cards').append($('<span>Pack received: ' + data.register + '</span><br/>'));
			$.each(data.cards, function(i, cardid) {
				$('#packs-opened-cards').append($('<span>received card: '+cards[cardid].name+'</span><br/>'));
			});
			$('#packs-opened-cards').append('<br/><br/>');

			SE.req.cache.myaccount = false;
			SE.req.getmyaccount();
		}, function() {
			console.error('/api/openpack.json: call failed');
			console.error(arguments);
		})
	})
});
