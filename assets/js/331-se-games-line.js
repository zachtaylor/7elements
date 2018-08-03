SE.widget.control('se-games-line', function(data) {
	this.life = data.life;
	this.timer = data.timer;
	this.opponents = data.opponents;
	this.gameid = data.gameid;
	$('button', this).click(function() {
		window.location.href = '/play/?gameid='+data.gameid;
	});

	var me = this;
	this.tick = function() {
		if (me.timer>0) {
			me.timer = me.timer-1;
			window.setTimeout(function() {
				me.tick();
			}, 1000);
		} else {
			me.timer = "DONE"
		}
	};
	this.tick();
});
