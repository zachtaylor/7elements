SE.widget.control('se-app-home', function() {
	var me = this;
	vii.ping().then(function(data) {
		me.online = data.online;
		me.updatetime = (new Date()).toTimeString();
	});
});
