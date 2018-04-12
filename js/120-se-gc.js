SE.widget.control('se-gc', function(data) {
	var me = this;

	me.update = function(data) {
		me.data = data;
		me.awake = data.awake;
		me.gcid = data.gcid;
		me.name = data.name;
		me.username = data.username;
		me.powers = Object.keys(data.powers).length;

		$('.se-gc-art', me)[0].src = data.image;

		if (data.body) {
			me.attack = data.body.attack;
			me.health = data.body.health;
		} else {
			$('.se-gc-banner', me)[0].innerHTML = 'ITEM';
		}

		me.showClear();
	};

	$(me).on('click', function() {
		SE.event.fire('se-gc-click', me);
	});

	me.showAwake = function() {
		me.state = 'awake';
		$(me).css({opacity:1});
		$(me).css({'border-color':'black'});
		$('.se-gc-eye', me)[0].src = '/img/icon/awake.png';
	};
	me.showAttack = function() {
		me.state = 'attack';
		me.banner = 'attack';
		$(me).css({opacity:1}).css({'border-color':'maroon'});
		$('.se-gc-eye', me)[0].src = '/img/icon/attack.32px.png';
	};
	me.showAsleep = function() {
		me.state = 'asleep';
		$(me).css({opacity:0.5});
		$(me).css({'border-color':'black'});
		$('.se-gc-eye', me)[0].src = '/img/icon/asleep.png';
	};
	me.showTrigger = function() {
		$(me).css({opacity:1}).css({'border-color':'gold'});
		$('.se-gc-eye', me)[0].src = '/img/icon/star.32px.png';
	};
	me.showTarget = function() {
		me.state = 'target';
		$(me).css({opacity:1}).css({'border-color':'gold'});
		$('.se-gc-eye', me)[0].src = '/img/icon/use.32px.png';
	}
	me.showClear = function() {
		me.state = '';
		if (me.awake) me.showAwake();
		else me.showAsleep();
	};

	me.update(data);
});
