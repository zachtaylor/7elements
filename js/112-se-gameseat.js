SE.widget.control('se-gameseat', function() {
	var me = this;
	var spinnerReady = Promise.Fake();

	SE.widget.controlProperty(me, 'username');
	SE.widget.controlProperty(me, 'deckname');
	SE.widget.controlProperty(me, 'turnphase');
	SE.widget.controlProperty(me, 'life');
	SE.widget.controlProperty(me, 'deck');
	SE.widget.controlProperty(me, 'spent');
	SE.widget.controlProperty(me, 'timer');

	SE.widget.new('se-card-spinner').then(function(spinner) {
		$('.se-gameseat-active', me).append(spinner);
		spinnerReady.resolve(spinner);
	});

	me.addElement = function(elementid, active) {
		var elements = $('.se-gameseat-elements', me);
		SE.widget.get('se-element').then(function(elementFactory) {
			var element = elementFactory();
			element.elementid = elementid;
			elements.append(element);
			if (!active) {
				$(element).css({
					opacity: '0.33',
				});
			}
		});
	};

	me.update = function(data) {
		me.username = data.username;
		me.deckname = data.deckname;
		me.life = data.life;
		me.deck = data.deck;
		me.spent = data.spent.length;
		me.updateactive(data.active);
		me.updateelements(data.elements);
	};

	me.updateactive = function(data) {
		spinnerReady.then(function(spinner) {
			spinner.empty();
			$.each(data, function(i, carddata) {
				SE.widget.new('se-card', carddata.cardid).then(function(card) {
					card.gcid = carddata.gcid;
					if (!carddata.active) {
						$(card).css({opacity:'0.33'});
					}
					$(card).css({display:'none'});
					spinner.append(card);
					$(card).slideDown();
				});
			});
		});
	};

	me.reactivate = function() {
		spinnerReady.then(function(spinner) {
			$.each($('.se-card', spinner), function(i, card) {
				$(card).css({opacity:'1'});
			});
		});
		$.each($('.se-gameseat-elements', me), function(i, element) {
			$(element).css({opacity:'1'});
		});
	};

	me.updateelements = function(data) {
		$('.se-gameseat-elements', me).empty();
		$.each(data, function(elementid, set) {
			for (var i=0; i<set.length; i++) {
				me.addElement(elementid, set[i]);
			}
		});
	};

	me.updatetimer = function() {
		if (me.timer > 0) {
			me.timer--;
		}
		window.setTimeout(me.updatetimer, 1000);
	};

	me.updatetimer();
});
