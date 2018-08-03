SE.widget.control('se-app-edit', function() {
	var me = this;

	// caching
	me.cache = {
		bank:{},
		deck:{},
	};
	// caching

	// filtering
	me.filter = {
		el:{1:true,2:true,3:true,4:true,5:true,6:true,7:true},
		cost: {1:true,2:true,3:true,7:true},
		type:{spell:true,item:true,body:true}
	}
	me.filter.testElement = function(carddata) {
		for (var i=1;i<8;i++) {
			if (carddata.costs[i] && me.filter.el[i]) return true;
		}
		return false;
	};
	me.filter.testCost = function(carddata) {
		var totalCost = 0;
		$.each(carddata.costs, function(elementid, count) {
			totalCost += count;
		});
		return me.filter.cost[totalCost]
	};
	me.filter.testType = function(carddata) {
		if (me.filter.type[carddata.type] !== undefined) return me.filter.type[carddata.type];
		else console.warn('card type unrecognized', carddata.type);
	};
	SE.event.on('save-filter-el', function(element) {
		me.filter.el[element] = !me.filter.el[element];
		me.layout();
	});
	SE.event.on('save-filter-cost', function(cost) {
		me.filter.cost[cost] = !me.filter.cost[cost];
		me.layout();
	});
	SE.event.on('save-filter-type', function(type) {
		me.filter.type[type] = !me.filter.type[type];
		me.layout();
	});
	// filtering

	// modding
	me.mod = {
		cards:{},
		name:''
	};
	me.mod.reset = function() {
		me.mod.cards = {};
		me.mod.name = me.data.decks[me.deckid].name;
	};
	me.mod.addCard = function(cardid) {
		var copiesInDeck = me.data.decks[me.deckid].cards[cardid] || 0;
		var copiesMod = me.mod.cards[cardid] || 0;
		me.deckcountmod = '' + (parseInt(me.deckcountmod || 0) + 1);
		if ((copiesInDeck + copiesMod) >= me.data.cards[cardid].copies) {
			return console.warn('mod.addCard: count exceeded', me.data.cards[cardid].copies);
		}
		if (me.mod.cards[cardid]) {
			me.mod.cards[cardid] = copiesMod+1;
		} else {
			me.mod.cards[cardid] = 1;
		}
		me.showDeckCard(cardid).then(me.mod.updateCardBanner);
	};
	me.mod.removeCard = function(cardid) {
		var copiesInDeck = me.data.decks[me.deckid].cards[cardid] || 0;
		var copiesMod = me.mod.cards[cardid] || 0;
		me.mod.cards[cardid] = copiesMod - 1;
		me.deckcountmod = '' + (parseInt(me.deckcountmod || 0) - 1);
		if ((copiesInDeck + copiesMod) > 1) {
			me.mod.updateCardBanner(me.cache.deck[cardid]);
		}	else {
			$(me.cache.deck[cardid]).detach();
		}
	};
	me.mod.updateCardBanner = function(card) {
		if (!card.copies) card.copies = 0;
		card.banner = 'x'+card.copies;
		if (me.mod.cards[card.cardid] > 0) card.banner += '+'+me.mod.cards[card.cardid];
		else if (me.mod.cards[card.cardid] < 0) card.banner += me.mod.cards[card.cardid];
	};
	// modding

	me.getCacheCard = function(cachename, cardid) {
		return new Promise(function(resolve, reject) {
			if (me.cache[cachename][cardid]) {
				resolve(me.cache[cachename][cardid]);
			} else SE.widget.new('se-card', cardid).then(function(card) {
				me.cache[cachename][cardid] = card;
				resolve(card);
			});
		});
	};
	me.showBankCard = function(cardid, copies) {
		return new Promise(function(resolve, reject) {
			me.getCacheCard('bank', cardid).then(function(card) {
				if (!card.parentElement) {
					$(card).off('click');
					$(card).click(function(e) {
						SE.dirtypage.on();
						me.mod.addCard(cardid);
					});
					$(me.$cardsbank).append(card);
				}
				if (copies) {
					card.copies = copies;
					card.banner = 'x'+copies;
				}
				resolve(card);
			});
		});
	};
	me.showDeckCard = function(cardid, copies) {
		if (!me.showDeckCard[cardid]) me.showDeckCard[cardid] = new Promise(function(resolve, reject) {
			me.getCacheCard('deck', cardid).then(function(card) {
				$(card).click(function(e) {
					SE.dirtypage.on();
					me.mod.removeCard(cardid);
				});
				resolve(card);
			});
		});
		return new Promise(function(resolve, reject) {
			me.showDeckCard[cardid].then(function(card) {
				if (!card.parentElement) {
					$(me.$cardsdeck).append(card);
				}
				if (copies) {
					card.copies = copies;
					card.banner = 'x'+copies;
				}
				resolve(card);
			});
		});
	};

	me.layout = function() {
		if (!me.data || !me.data.decks) return;
		if (!me.deckid) return;

		me.mod.reset();


		// update filter activity indicators
		for (var i=1;i<8;i++) {
			if (me.filter.el[i]) $(me['$filterel'+i]).addClass('active');
			else $(me['$filterel'+i]).removeClass('active');
		}
		for (var i=1;i<4;i++) {
			if (me.filter.cost[i]) $(me['$filtercost'+i]).addClass('active');
			else $(me['$filtercost'+i]).removeClass('active');
		}
		if (me.filter.type.body) $(me.$filterbody).addClass('active');
		else $(me.$filterbody).removeClass('active');
		if (me.filter.type.spell) $(me.$filterspell).addClass('active');
		else $(me.$filterspell).removeClass('active');
		if (me.filter.type.item) $(me.$filteritem).addClass('active');
		else $(me.$filteritem).removeClass('active');
		// update filter activity indicators

		// bank pane
		for (; me.$cardsbank.children.length;) {
			$(me.$cardsbank.children[0]).detach();
		};
		var showEmptyBankWarning = true;
		var cardSetCount = Object.keys(me.data.cards).length;
		for (var cardid=1;cardid<=cardSetCount;cardid++) {
			var cdat = me.data.cards[cardid];
			if (!cdat.copies) continue;
			cardCount+= cdat.copies;
			showEmptyBankWarning = false;
			if (me.filter.testElement(cdat) && me.filter.testCost(cdat) && me.filter.testType(cdat)) {
				me.showBankCard(cardid, cdat.copies);
			}
		};
		if (showEmptyBankWarning) {
			$(me.$cardsbank).append('<h2 style="color:red;">Uh-oh! It appear you have no cards...</h2>');
		}
		// bank pane

		// deck pane
		for (; me.$cardsdeck.children.length;) {
			$(me.$cardsdeck.children[0]).detach();
		}
		var showEmptyDeckWarning = true;
		var cardCount = 0;
		$.each(me.data.decks[me.deckid].cards, function(cardid, copies) {
			showEmptyDeckWarning = false;
			cardCount += copies;
			me.showDeckCard(cardid, copies);
		});
		if (showEmptyDeckWarning) {
			$(me.$cardsdeck).append('<h2 style="color:red;">Uh-oh! It appears this deck is empty...</h2>');
		} else {
			me.deckcount = ''+cardCount;
		}
		// deck pane

		// deck rename
		$(me.nameinput).val(me.data.decks[me.deckid].name);
	};

	me.footer = $('<div></div>');
	me.nameinput = $('<input autosize class="vii"></input>')[0];
	$(me.nameinput).keyup(function(e) {
		me.mod.name = $(me.nameinput).val();
		$(me.nameinput).attr('placeholder', me.data.decks[me.deckid].name);
		if (me.mod.name != me.data.decks[me.deckid].name) {
			SE.dirtypage.on();
		}
	});
	$(me.footer).append(me.nameinput);
	me.savebutton = $('<button class="vii vii-green" autosize disabled style="padding:7px;position:absolute;">Saved</button>')[0];
	$(me.savebutton).click(function() {
		var dat = {
			name: me.mod.name || me.data.decks[me.deckid].name,
			cards: {}
		};
		$.each(me.data.decks[me.deckid].cards, function(cardid, copies) {
			dat.cards[cardid] = copies;
		});
		$.each(me.mod.cards, function(cardid, copies) {
			dat.cards[cardid] = dat.cards[cardid] || 0;
			dat.cards[cardid] += copies;
		});
		SE.api.post('decks/'+me.deckid, JSON.stringify(dat)).then(function() {
			SE.dirtypage.off();
			SE.websocket.send('ping');
		}, function(err) {
			console.error('error saving deck', err);
		});
	});
	$(me.footer).append(me.savebutton);

	SE.event.on('dirtypage', function(on) {
		$(me.savebutton).prop('disabled', !on);
		if (on) {
			$(me.savebutton).removeClass('vii-green').addClass('vii-gold');
			me.savebutton.innerHTML = 'Save';
		} else {
			$(me.savebutton).removeClass('vii-gold').addClass('vii-green');
			me.savebutton.innerHTML = 'Saved';
		}
	});

	vii.ping().then(function(data) {
		if (!data.username) return $(me).prepend('<h2>Edit Screen: <a href="/login/">Login</a> Required</h2>');
		me.data = data;
		me.layout();
	});

	vii.storage.with('editid', function(v) {
		if (!v) return vii.storage.set('editid', 1);
		me.deckid = v;
		me.layout();
	});
});
