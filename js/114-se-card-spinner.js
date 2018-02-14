SE.widget.control('se-card-spinner', function() {
	var me = this;
	var li = $('li', this);
	me.li = li[0];

	this.append = function(card) {
		li.append(card);
		$(card).show('slide', {direction:'right'}, 500);
	};
	this.prepend = function(card) {
		li.prepend(card);
		$(card).show('slide', {direction:'left'}, 500);
	};
	this.empty = function() {
		li.empty();
	};

	$('button.se-card-spinner-left', me).click(function(e) {
		var card = li.find('.se-card:first-child, .se-gc:first-child')[0];
		$(card).hide('slide', {direction:'left'}, 500);
		window.setTimeout(function() {
			$(card).detach();
			me.append(card);
		}, 500);
	});

	$('button.se-card-spinner-right', me).click(function(e) {
		var card = li.find('.se-card:last-child, .se-gc:last-child')[0];
		$(card).hide('slide', {direction:'right'}, 500);
		window.setTimeout(function() {
			$(card).detach();
			me.prepend(card);
		}, 500);
	});

	$('button', me).mouseover(function() {
		$(this).css({opacity:1});
	}).mouseout(function() {
		$(this).css({opacity:0.1});
	}).css({opacity:0.1}); // css warm up
});