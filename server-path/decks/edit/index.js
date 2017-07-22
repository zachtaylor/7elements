$(function() {
	var cards = false;
	var mycards = false;
	var mydecks = false;
	var activedeck = 1;
	var elementsFilter = {
		1:true,2:true,3:true,4:true,5:true,6:true,7:true
	};

	var newSECard = function(cardid, deck) {
		var card = $('<se-card></se-card>').draggable()[0];
		card.cardid = cardid;
		card.setAttribute("deck", deck);
		return card;
	};

	SE.api.get('cards', function(data) {
		cards = data;
		SE.event.fire('decks-edit-filter');
	});

	SE.api.get('mydecks', function(data) {
		mydecks = data;
		SE.event.fire('decks-edit-filter');
	});

	SE.api.get('mycards', function(data) {
		mycards = data;
		SE.event.fire('decks-edit-filter');
	});

	$('#deck-selector a').click(function() {
		activedeck = parseInt(this.innerHTML);

		if (!activedeck) {
			activedeck = 1;
		}

		$('#deck-selector-currect a').html(this.innerHTML);
		SE.event.fire('decks-edit-filter')
	});

	$('#deck-name').on('blur', function() {
		var val = $(this).val();
		if (mydecks[activedeck].name != val) {
			mydecks[activedeck].name = val;
			$('#deck-settings .input-group').addClass('has-warning');
			$('#save-button').removeClass('btn-success').removeClass('btn-danger').addClass('btn-warning');
		}
	});

	$('button[filter]').click(function(event) {
		var id = parseInt(this.id.charAt(this.id.length-1));

		if (parseInt(this.getAttribute('filter'))) {
			elementsFilter[id] = false;
			this.setAttribute('filter', '0');
		} else {
			elementsFilter[id] = true;
			this.setAttribute('filter', '1');
		}

		SE.event.fire('decks-edit-filter');
	});

	$('#search-button-left').click(function() {
		var card = $('#search-bar').find(':first-child')[0];
		$('#search-bar').append(newSECard(card.cardid, "false"));
		$(card).remove();
	});
	$('#search-button-right').click(function() {
		var card = $('#search-bar').find(':last-child')[0];
		$('#search-bar').prepend(newSECard(card.cardid, "false"));
		$(card).remove();
	});
	$('#deck-button-left').click(function() {
		var card = $('#deck-bar').find(':first-child')[0];
		$('#deck-bar').append(newSECard(card.cardid, "true"));
		$(card).remove();
	});
	$('#deck-button-right').click(function() {
		var card = $('#deck-bar').find(':last-child')[0];
		$('#deck-bar').prepend(newSECard(card.cardid, "true"));
		$(card).remove();
	});

	$('#save-button').click(function(event) {
		var saveButton = this;
		SE.api.post('mydecks/'+activedeck, JSON.stringify(mydecks[activedeck])).then(function(data) {
			SE.api.cache.mydecks = false;
			SE.api.get('mydecks');
			$(saveButton).removeClass('btn-error').removeClass('btn-warning').addClass('btn-success');
			$('#deck-bar').removeClass('has-warning');
			$('#deck-settings .input-group').removeClass('has-warning');
		}, function() {
			$(saveButton).removeClass('btn-warning').removeClass('btn-success',).addClass('btn-danger');
		});
	});

	$('#search-bar').droppable({
		accept: 'se-card[deck="true"]',
		drop: function(event, ui) {
			var card = ui.draggable[0];
			var cardid = card.cardid;
			mydecks[activedeck].cards[cardid]--;

			$(card).remove();
			$('#save-button').removeClass('btn-success').removeClass('btn-danger').addClass('btn-warning');
			$(this).addClass('has-warning').prepend(newSECard(cardid, "false"));
		}
	});

	$('#deck-bar').droppable({
		accept: 'se-card[deck="false"]',
		drop: function(event, ui) {
			var card = ui.draggable[0];
			var cardid = card.cardid;

			if (mydecks[activedeck].cards[cardid]) {
				mydecks[activedeck].cards[cardid]++;
			} else {
				mydecks[activedeck].cards[cardid] = 1;
			}

			$(card).remove();
			$('#save-button').removeClass('btn-success').removeClass('btn-danger').addClass('btn-warning');
			$(this).addClass('has-warning').prepend(newSECard(cardid, "true"));
		}
	});

	SE.event.on('decks-edit-filter', function() {
		if (!cards) return;
		if (!mycards) return;
		if (!mydecks) return;

		var searchbar = $('#search-bar').empty()[0];
		var deckbar = $('#deck-bar').empty()[0];

		$.each(mycards, function(cardid, count) {
			var card = cards[cardid];

			for (var i=0;i<count;i++) {
				if (i<mydecks[activedeck].cards[cardid]) {
					$(deckbar).append(newSECard(cardid, "true"));
					continue;
				}

				$.each(card.elementcosts, function(_,elcostobj) {
					if (elementsFilter[elcostobj.element]) {
						for (var i=0;i<count;i++) {
							$(searchbar).append(newSECard(cardid, "false"));
							return false;
						}
					}
				});
			}
		});

		if (mydecks[activedeck].name) {
			$('#deck-name').val(mydecks[activedeck].name)
		} else {
			$('#deck-name').val(null);
		}

		var deckCount = 0;
		$.each(mydecks[activedeck].cards, function(cardid, amount) {
			deckCount += amount;
		})
		$('#deck-count').html(''+deckCount);
	});

});
