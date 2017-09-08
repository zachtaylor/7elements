SE.widget.control('se-game-active', function(data) {
	SE.widget.controlProperty(this, 'deckname');
	SE.widget.controlProperty(this, 'timer');
	SE.widget.controlProperty(this, 'life');
	SE.widget.controlProperty(this, 'opponents');

	this.deckname = data.deckname;
	this.life = data.life;
	this.timer = data.timer;
	this.opponents = data.opponents;
	$('a', this)[0].href = '/games/play/?gameid='+data.gameid;
});
