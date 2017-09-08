$(function() {
	var data = {};
	var deckid = 1;
	var elementsFilter = {
		1:true,2:true,3:true,4:true,5:true,6:true,7:true
	};
	var costsFilter = {
		1:true,2:true,3:true,7:true
	};

	SE.api.get('cards', function(cardsdata) {
		data.cards = cardsdata;
	}).then(function(cardsdata) {
		data.cards = cardsdata;
		SE.event.fire('decks-edit-filter')
	});

	SE.api.get('navbar', function(navbardata) {
		data.navbar = navbardata;
	}).then(function(navbardata) {
		data.navbar = navbardata;
		SE.event.fire('decks-edit-filter');
	});

	SE.event.on('save-protector', function() {
		$('#save-button').prop('disabled', false)[0].innerHTML = "Save Needed!";
		$('#save-button').removeClass('btn-success').removeClass('btn-danger').addClass('btn-warning');
		SE.dirtypage.on();
	});

	SE.widget.controlProperty($('#deck-settings')[0], 'deckcount');
	SE.widget.controlProperty($('#deck-settings')[0], 'deckcountmod');

	var makeCardClicker = function(card) {
		return function() {
			card.banner = 'not saved';
			SE.event.fire('save-protector');

			if (card.getAttribute('deck') == 'true') {
				data.navbar.decks[deckid].cards[card.cardid]--;
				card.setAttribute('deck', false);
				$(card).detach()
				$('#search-bar').prepend(card);
				if ($('#deck-settings')[0].deckcountmod) $('#deck-settings')[0].deckcountmod--;
			} else if (card.getAttribute('deck') == 'false') {
				if (data.navbar.decks[deckid].cards[card.cardid]) data.navbar.decks[deckid].cards[card.cardid]++;
				else data.navbar.decks[deckid].cards[card.cardid] = 1;
				card.setAttribute('deck', true);
				$(card).detach()

				$.each(data.cards[card.cardid].costs, function(elementid,count) {
					if (!elementsFilter[elementid]) return true;
					$('#deck-bar .se-card-spinner')[0].prepend(card);
					return false;
				});

				$('#deck-settings')[0].deckcountmod++;
			}
		}
	};

	$('#deck-selector-list a').click(function() {
		$('#save-button').click();
		var me = this;
		window.setTimeout(function() {
			deckid = parseInt(me.innerHTML);
			if (!deckid) deckid = 1;
			SE.event.fire('decks-edit-filter');
		}, 1000);
	});

	$('#deck-name').on('blur', function() {
		var val = $(this).val();
		SE.api.get('navbar').then(function(navbardata) {
			if (navbardata.decks[deckid].name != val) {
				navbardata.decks[deckid].name = val;
				$('#deck-settings .input-group').addClass('has-warning');
				SE.event.fire('save-protector');
			}
		});
	});

	$('button[filter-element]').click(function(event) {
		var id = parseInt(this.getAttribute('filter-element'));
		if (parseInt(this.getAttribute('filter-toggle'))) {
			elementsFilter[id] = false;
			this.setAttribute('filter-toggle', '0');
		} else {
			elementsFilter[id] = true;
			this.setAttribute('filter-toggle', '1');
		}
		SE.event.fire('decks-edit-filter');
	});

	$('button[filter-cost]').click(function(event) {
		var id = parseInt(this.getAttribute('filter-cost'));
		if (parseInt(this.getAttribute('filter-toggle'))) {
			costsFilter[id] = false;
			this.setAttribute('filter-toggle', '0');
		} else {
			costsFilter[id] = true;
			this.setAttribute('filter-toggle', '1');
		}
		SE.event.fire('decks-edit-filter');
	});

	$('#save-button').click(function(event) {
		var saveButton = this;
		saveButton.innerHTML = "Saving...";
		$(saveButton).prop('disabled', true);
		SE.api.post('decks/'+deckid, JSON.stringify(data.navbar.decks[deckid])).then(function() {
			saveButton.innerHTML = "Saved";
			SE.api.cache.navbar = false;
			SE.api.get('navbar').then();
			SE.dirtypage.off();
			$('#deck-settings')[0].deckcountmod = 0;
			$(saveButton).removeClass('btn-error').removeClass('btn-warning').addClass('btn-success');
			$('#deck-bar').removeClass('has-warning');
			$('#deck-settings .input-group').removeClass('has-warning');
			SE.event.fire('deck-selector-refresh');
		}, function() {
			$(saveButton).removeClass('btn-warning').removeClass('btn-success').addClass('btn-danger');
		});
	});

	SE.event.on('deck-selector-refresh', function() {
		SE.api.get('navbar').then(function(navbardata) {
			$('#deck-selector-currect a')[0].innerHTML = deckid+': '+data.navbar.decks[deckid].name;
			$.each(navbardata.decks, function(deckid, deck) {
				var deckli = $('#deck-selector-list a')[deckid-1];
				if (deckli) {
					deckli.innerHTML = deckid + ": " + deck.name;
				} else {
					console.error('deck id other than 1,2,3');
				}
			});
		});
	});

	SE.event.on('decks-edit-filter', function() {
		if (!data.cards) return;
		if (!data.navbar) return;
		console.log('start decks-edit-filter', data);
		$('#deck-bar .se-card-spinner')[0].empty();
		$('#collection-bar .se-card-spinner')[0].empty();

		$.each(data.navbar.cards, function(cardid, count) {
			for (var i=0;i<data.navbar.decks[deckid].cards[cardid];i++) {
				SE.widget.new('se-card', cardid).then(function(card) {
					$(card).click(makeCardClicker(card));
					card.setAttribute("deck", true);
					$('#deck-bar .se-card-spinner')[0].append(card);
				});
			}
			for (var i=data.navbar.decks[deckid].cards[cardid]||0;i<count;i++) {
				SE.widget.new('se-card', cardid).then(function(card) {
					$(card).click(makeCardClicker(card));
					card.setAttribute("deck", false);

					var cardCost = 0;
					$.each(data.cards[cardid].costs, function(_, count) {
						cardCost+=count;
					});
					if (!costsFilter[cardCost]) return true;

					$.each(data.cards[cardid].costs, function(elementid,count) {
						if (!elementsFilter[elementid]) return true;
						$('#collection-bar .se-card-spinner')[0].append(card);
						return false;
					});
				});
			}
		});

		SE.event.fire('deck-selector-refresh');

		if (data.navbar.decks[deckid].name) {
			$('#deck-name').val(data.navbar.decks[deckid].name)
		} else {
			$('#deck-name').val(null);
		}

		var deckCount = 0;
		$.each(data.navbar.decks[deckid].cards, function(cardid, amount) {
			deckCount += amount;
		});
		$('#deck-settings')[0].deckcount = deckCount;
	});
});
