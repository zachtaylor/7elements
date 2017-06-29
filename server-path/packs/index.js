$(function() {
	$('#packs-control-buttons-open').on('click', function() {
		SE.req.openpack().then(function(data) {
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
				var firstcard = cardpile.find(':first-child').remove();
				cardpile.append(firstcard);
			});

			rbutton.click(function(e) {
				var firstcard = cardpile.find(':last-child').remove();
				cardpile.prepend(firstcard);
			});

			$('#packs-opened-cards').append(row);
			$('#packs-opened-cards').append('<br/><br/>');

			SE.req.cache.myaccount = false;
			SE.req.getmyaccount();
		}, function() {
			console.error('/api/openpack.json: call failed', arguments);
		});
	})
});
