$(function() {
	var openButton = $('#packs-control-buttons-open');
	var opensound = new Audio('/mp3/openpack.mp3');

	function openpack(data) {
		$('#se-navbar')[0].packs--;
		opensound.play();
		openButton[0].innerHTML = 'Open Pack';

		SE.widget.new('se-card-spinner').then(function(spinner) {
			$.each(data.cards, function(i, cardid) {
				SE.widget.new('se-card', cardid).then(function(card) {
					card.cardid = cardid;
					spinner.append(card);
				});
			});
			$('#packs-spinners').append(spinner);
			$('#packs-spinners').append('<br/><br/>');
		});
	};

	openButton.on('click', function(e) {
		SE.api.cache.openpack = false;
		openButton[0].innerHTML = '...';

		SE.api.get('openpack').then(openpack, function() {
			console.error('/api/openpack.json: call failed', arguments);
		});
	})
});
