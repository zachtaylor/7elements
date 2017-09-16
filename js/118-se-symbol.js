SE.widget.control('se-symbol', function(name) {
	var me = this;

	this.update = function(name) {
		me.src = '/img/icon/'+name+'.png';
	};

	if (name) {
		this.update(name);
	}
});
