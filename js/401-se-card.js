SE.widget.control('se-card', function(cardid) {
	var me = this;
	me.cardid = cardid;

	var powersReady = Promise.Fake();

	me.showAwake = function() {
		$(me).css({opacity:1});
		powersReady.then(function() {
			$('.se-card-power', me).addClass('active').on('click', function() {
				console.log('card power click', this, me);
			});
		});
	};
	me.showAsleep = function() {
		$(me).css({opacity:0.33});
		$('.se-card-power', me).removeClass('active').off('click');
	};

	me.update = function(data) {
		me.data = data;
		me.name = data.name;
		me.flavor = data.flavor;
		$('.se-card-art', me)[0].src = data.image;

		if (data.body) {
			me.attack = data.body.attack;
			me.health = data.body.health;
			me.body = "true";
		} else {
			me.type = data.type;
			me.body = "false";
		}

		me.updateCosts(data.costs);
		me.updatePowers(data.powers);
	};
	me.updateCosts = function(data) {
		$('.se-card-costs', me).empty();
		$.each(data, function(elementid, cost) {
			for (var i=0; i<cost; i++) {
				SE.widget.new('se-symbol', 'element-'+elementid).then(function(symbol) {
					me.costs.append(symbol);
				});
			}
		});
	};
	me.updatePowers = function(data) {
		$(me.powers).empty();
		$.each(data, function(i, powerdata) {
			SE.widget.new('se-card-power', powerdata).then(function(power) {
				$(me.powers).append(power);
			});
		});
		SE.widget.get('se-card-power').then(function() {
			powersReady.resolve();
		});
	};

	if (cardid) vii.ping().then(function(data) {
		me.update(data.cards[cardid]);
	}, function(e) {
		console.error('cards error', e);
	});
	else console.error('card id missing', cardid);
});
