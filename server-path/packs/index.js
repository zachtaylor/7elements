$(function() {
	var cards = {};

	SE.cards.get().then(function(data) {
		$.each(data, function(cardid, card) {
			cards[cardid] = card;
		});
		console.log(cards);
	}, function(error) {
		console.error(error);
	});

	SE.myaccount.get().then(function(data) {
		$('#data-myaccount-username').html(' ('+data.username+') ');
		$('#data-myaccount-cards').html(data.cards);

		if (data.packs > 0) {
			$('#data-myaccount-packs').html(data.packs);
		};
	}, function(error) {
		if (error.responseText == 'session missing') {
			$('#nav-myaccount-link')[0].href = '/login/';
			$('#data-myaccount-username').html(' (login) ');
			return;
		};

		console.error(error.responseText);
	});

	$('#packs-control-buttons-open').on('click', function() {
		console.log('wats up')
		$.getJSON('/api/openpack.json').then(function(data) {
			$('#packs-opened-cards').append($('<span>Pack received: ' + data.register + '</span><br/>'));

			$.each(data.cards, function(i, cardid) {
				console.log('create card entry: ' + cardid);
				$('#packs-opened-cards').append($('<span>received card: '+cards[cardid].name+'</span><br/>'));
			});

			$('#packs-opened-cards').append('<br/><br/>');
		}, function() {
			console.error("/api/openpack.json: call failed");
			console.error(arguments);
		})
	})
});
