SE.widget.control('se-card', function(cardid) {
	var me = this;

	SE.widget.controlProperty(me, 'cardid', true);
	SE.widget.controlProperty(me, 'gcid', true);
	SE.widget.controlProperty(me, 'name');
	SE.widget.controlProperty(me, 'description');
	SE.widget.controlProperty(me, 'flavor');
	SE.widget.controlProperty(me, 'type');
	SE.widget.controlProperty(me, 'attack');
	SE.widget.controlProperty(me, 'health');
	SE.widget.controlProperty(me, 'body', true);
	SE.widget.controlProperty(me, 'banner');

	me.cardid = cardid;
	if (cardid) SE.api.get('cards').then(function(data) {
		me.update(data[cardid]);
	}, function(e) {
		console.error('cards error', e);
	});
	else console.error('card id missing', cardid);

	me.update = function(data) {
		me.name = data.name;
		me.description = data.description;
		me.flavor = data.flavor;

		if (data.body) {
			me.attack = data.attack;
			me.health = data.health;
			me.body = "true";
		} else {
			me.type = data.type;
			me.body = "false";
		}

		if (data.image) {
			$('.se-card-art', me)[0].src = data.image;
		}

		$('.se-card-costs', me).empty();
		$.each(data.costs, function(elementid, cost) {
			for (var i=0; i<cost; i++) {
				var symbol = $('<se-symbol icon="element-'+elementid+'"></se-symbol>')[0];
				$('.se-card-costs', me).append(symbol);
			}
		});
	}
});
