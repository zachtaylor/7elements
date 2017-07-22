$(function() {
	var openButton = $('#packs-control-buttons-open');
	var opensound = new Audio('/mp3/openpack.mp3');

	openButton.on('click', function(e) {
		SE.api.cache.openpack = false;
		openButton[0].innerHTML = '...';

		SE.api.get('openpack').then(function(data) {
			opensound.play();
			openButton[0].innerHTML = 'Open Pack';

			var row = $('<div class="" style="white-space: nowrap;"></div>');
			var lbutton = $('<button class="btn btn-lg">&lt;</button>');
			row.append(lbutton);
			var rbutton = $('<button class="btn btn-lg">&gt;</button>');
			row.append(rbutton);
			row.append($('<span>Pack received: ' + data.register + '</span><br/>'));
			var cardpile = $('<div></div>');
			row.append(cardpile);

			$.each(data.cards, function(i, cardid) {
				var card = card = $('<se-card></se-card>')[0];
				card.cardid = cardid;
				cardpile.append(card);
			});

			lbutton.click(function(e) {
				var card = cardpile.find(':first-child')[0];
				var newcard = $('<se-card></se-card>')[0];
				newcard.cardid = card.cardid
				cardpile.append(newcard);
				$(card).remove();
			});

			rbutton.click(function(e) {
				var card = cardpile.find(':last-child')[0];
				var newcard = $('<se-card></se-card>')[0];
				newcard.cardid = card.cardid
				cardpile.prepend(newcard);
				$(card).remove();
			});

			$('#packs-opened-cards').append(row);
			$('#packs-opened-cards').append('<br/><br/>');

			SE.api.cache.myaccount = false;
			SE.api.get('myaccount');

		}, function() {
			console.error('/api/openpack.json: call failed', arguments);
		});
	})
});
