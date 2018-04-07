SE.widget.control('se-gameseat', function() {
	var me = this;
	var ready = Promise.Fake();

	SE.widget.new('se-card-spinner').then(function(spinner) {
		$('.se-gameseat-active', me).append(spinner);
		me.spinner = spinner;
		ready.resolve();
	});

	me.elements = {};

	me.addElement = function(elementid, active) {
		SE.widget.new('se-element').then(function(element) {
			me.elements[elementid] = me.elements[elementid] || [];
			me.elements[elementid].push(element);
			element.elementid = elementid;
			$('.se-gameseat-elements', me).append(element);
			$(element).css({
				opacity: active ? 1 : 0.2,
			});
		});
	};

	me.resetElements = function(data) {
		$.each(data, function(elementid, set) {
			me.elements[elementid] = me.elements[elementid] || [];
			for (var i=0; i<set.length; i++) {
				if (i < me.elements[elementid].length) $(me.elements[elementid][i]).css({opacity: set[i] ? 1 : 0.2});
				else me.addElement(elementid, set[i]);
			}
		});
	};

	me.addActiveCard = function(card) {
		$(card).detach().off('mouseover').off('mouseout');
		$(card).css({top:'0px'});
		ready.then(function() {
			me.spinner.append(card);
		});
	};

	me.resetActiveCards = function(data) {
		ready.then(function() {
			$(me.spinner)[0].empty();
			$.each(data, function(i, carddata) {
				vii.gamecard.set(carddata).then(function(gc) {
					me.addActiveCard(gc);
					gc.update(carddata);
				});
			});
		});
	};

	me.update = function(data) {
		me.username = data.username;
		me.life = data.life;
		me.hand = data.hand;
		me.deck = data.deck;
		me.spent = data.spent;
		me.resetElements(data.elements);
		me.resetActiveCards(data.active);
	};

	me.reactivate = function() {
		ready.then(function() {
			$.each(me.spinner.li.children, function(i, card) {
				card.awake = true;
				card.showClear();
			});
		});
		$.each($('.se-element', me), function(i, element) {
			$(element).css({opacity:'1'});
		});
	};

	me.updatetimer = function() {
		if (me.timer > 0) me.timer--;
		else me.timer = 0;
		window.setTimeout(function() {me.updatetimer();}, 1000);
	};
	me.updatetimer();
});
