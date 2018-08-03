SE.widget.control('se-app-play', function() {
	var me = this;

	me.footer = $('<div></div')[0];

	me.reload = function() {
	};
	me.AddToHand = function(card) {
	};
	me.RemoveFromHand = function(card) {
	};

	me.gethand = Promise.Fake();
	SE.widget.new('se-card-spinner').then(function(spinner) {
		$(me.footer).append(spinner);
		me.gethand.resolve(spinner);
	});

	vii.ping().then(function(data) {
		console.log('se-app-play', data);
	});
});
