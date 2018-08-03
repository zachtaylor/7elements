SE.widget.control('se-card-power', function(data) {
	var me = this;
	me.data = data;
	me.innerHTML = data.description;

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
