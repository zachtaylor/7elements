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
		$(me).css({opacity:1});
		$('.se-gc-eye', me)[0].src = '/img/icon/awake.png';
	};
	me.showAttack = function() {
		$(me).css({opacity:1}).css({'border-color':'maroon'});
		$('.se-gc-eye', me)[0].src = '/img/icon/attack.32px.png';
	};
	me.showClear = function() {
		$(me).css({'border-color':'black'});
		if (me.awake) me.showAwake();
		else me.showAsleep();
	};
	me.showAsleep = function() {
		$(me).css({opacity:0.5});
		$('.se-gc-eye', me)[0].src = '/img/icon/asleep.png';
	};

	me.update(data);
});
