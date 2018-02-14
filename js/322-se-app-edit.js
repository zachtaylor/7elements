SE.widget.control('se-app-edit', function() {
	var me = this;

	// filtering
	me.filterEls = {1:true,2:true,3:true,4:true,5:true,6:true,7:true,};
	me.filterCosts = {1:true,2:true,3:true,7:true,};
	me.filterTypes = {spell:true,item:true,body:true,};
	SE.event.on('save-filter-el', function(element) {
		me.filterEls[element] = !me.filterEls[element];
		me.layout();
	});
	SE.event.on('save-filter-cost', function(cost) {
		me.filterCosts[cost] = !me.filterCosts[cost];
		me.layout();
	});
	SE.event.on('save-filter-type', function(type) {
		me.filterCosts[type] = !me.filterCosts[type];
		me.layout();
	});
	var filterShowElement = function(card) {
		for (var i=1;i<8;i++) {
			if (card.data.costs[i] && me.filterEls[i]) return true;
		}
		return false;
	};
	var filterShowCost = function(card) {
		var totalCost = 0;
		$.each(card.data.costs, function(elementid, count) {
			totalCost += count;
		});
		return me.filterCosts[totalCost]
	};
	var filterShowType = function(card) {
		if (me.filterTypes[card.data.type] !== undefined) return me.filterTypes[card.data.type];
		else console.warn('card type unrecognized', card.data.type);
	};
	me.layout = function() {
		if (!me.deckid) return;
		if (!me.data) return;

		for (var i=1;i<8;i++) {
			if (me.filterEls[i]) $(me['$filterel'+i]).addClass('active');
			else $(me['$filterel'+i]).removeClass('active');
		}
		for (var i=1;i<4;i++) {
			if (me.filterCosts[i]) $(me['$filtercost'+i]).addClass('active');
			else $(me['$filtercost'+i]).removeClass('active');
		}
		if (me.filterTypes.body) $(me.$filterbody).addClass('active');
		else $(me.$filterbody).removeClass('active');
		if (me.filterTypes.spell) $(me.$filterspell).addClass('active');
		else $(me.$filterspell).removeClass('active');
		if (me.filterTypes.item) $(me.$filteritem).addClass('active');
		else $(me.$filteritem).removeClass('active');

		for (; me.$cardsbank.children.length;) {
			$(me.$cardsbank.children[0]).detach();
		};
		$.each(me.cachebank, function(cardid, card) {
			if (filterShowElement(card) && filterShowCost(card) && filterShowType(card)) {
				me.AddToBankPane(cardid);
			}
		});

		for (; me.$cardsdeck.children.length;) {
			$(me.$cardsdeck.children[0]).detach();
		};
		$.each(me.data.decks[me.deckid].cards, function(cardid, copies) {
			me.AddToDeckPane(cardid, copies);
		});
	};
	// filtering

	// modding
	me.mod = {
		cards:{},
		name:''
	};
	// modding

	me.AddToBankPane = function(cardid, copies) {
		if (!me.cachebank) me.cachebank = {};
		if (me.cachebank[cardid]) {
			var card = me.cachebank[cardid];
			if (copies) card.copies=copies;
			card.banner='x'+card.copies;
			if (!card.parentElement) $(me.$cardsbank).append(card);
		} else SE.widget.new('se-card', cardid).then(function(card) {
			me.cachebank[cardid] = card;
			card.copies=copies;
			card.banner='x'+copies;
			$(card).click(function(e) {
				me.ClickBankCard(cardid);
			});
			$(me.$cardsbank).append(card);
		});
	};

	me.ClickBankCard = function(cardid) {
		vii.ping().then(function(data) {
			if (data.cards[cardid].copies > data.decks[me.deckid].cards[cardid]) {
				if (me.mod.cards[cardid]) me.mod.cards[cardid]++;
				else me.mod.cards[cardid] = 1;
				me.layout();
				SE.dirtypage.on();
			}
		});
	};

	me.AddToDeckPane = function(cardid, copies) {
		if (!me.cachedeck) me.cachedeck = {};
		if (me.cachedeck[cardid]) {
			var card = me.cachedeck[cardid];
			if(copies) card.copies=copies;
			if (me.mod.cards[cardid]) card.copies+=me.mod.cards[cardid];
			card.banner = 'x'+card.copies;
			if (!card.parentElement) $(me.$cardsdeck).append(card);
		} else SE.widget.new('se-card', cardid).then(function(card) {
			me.cachedeck[cardid] = card;
			card.copies = copies;
			card.banner = 'x'+me.cachedeck[cardid].copies;
			$(me.$cardsdeck).append(card);
		});
	};

	me.footer = $('<div></div>');
	me.savebutton = $('<button class="vii vii-gold" disabled>Save</button>');
	$(me.footer).append(me.savebutton);
	SE.event.on('dirtypage', function(on) {
		console.log('dirtypage');
		if (on) $(me.savebutton).removeClass('vii-green').addClass('vii-gold');
		else $(me.savebutton).removeClass('vii-gold').addClass('vii-green');
		$(me.savebutton).prop('disabled', !on);
	});

	vii.ping().then(function(data) {
		if (!data.username) return $(me).prepend('<h2>Edit Screen: <a href="/login/">Login</a> Required</h2>');
		me.data = data;
		$.each(data.cards, function(cardid, carddata) {
			if (carddata.copies) me.AddToBankPane(cardid, carddata.copies);
		});
		me.layout();
	});

	vii.cookie('editid', function(val) {
		var i = parseInt(val);
		if (!(i > 0)) vii.cookie('editid', 1);
		else {
			me.deckid = i;
			me.layout();
		}
	});
});
