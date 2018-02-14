SE.widget.control('se-card-power', function(power) {
	this.power = power;
	this.innerHTML = power.description;

	var me = this;
	$(me).click(function() {
		if (!me.gameid) {
			console.warn('se-card-power: gameid not set');
		} else {
			SE.websocket.send('play', {
				gameid: me.gameid
			});
		}
	});
});
