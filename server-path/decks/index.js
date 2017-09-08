$(function() {
	SE.api.get('navbar').then(function(navbardata) {
		$.each(navbardata.decks, function(deckid, deck) {
			SE.widget.new('se-deck-details', deck).then(function(deckDetails) {
				$('#decks-details').append(deckDetails);
			});
		});
	});
});