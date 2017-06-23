$(function() {
	$('#packs-control-buttons-open').on('click', function() {
		SE.req.openpack().then(function(data) {
			var row = $('<div class="" style="white-space: nowrap;"></div>');
			row.append($('<button class="btn btn-lg">&lt;</button>'));
			row.append($('<button class="btn btn-lg">&rt;</button>'));
			row.append($('<span>Pack received: ' + data.register + '</span><br/>'));

			$.each(data.cards, function(i, cardid) {
				var card = card = $('<se-card></se-card>')[0];
				card.cardid = cardid;

				row.append(card);
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
