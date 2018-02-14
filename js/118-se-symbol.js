SE.widget.control('se-symbol', function(name) {
	var me = this;
	me.update = function(name) {
		me.src = '/img/icon/'+name+'.png';
	}
	if (name) {
		me.update(name);
	}
});
