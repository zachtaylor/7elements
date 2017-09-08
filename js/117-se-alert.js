SE.widget.control('se-alert', function() {
	var me = this;

	SE.widget.controlProperty(me, 'message');
	
	me.setMode = function(mode) {
		$(me).addClass('alert-'+mode);
	};

	me.autoDismissSeconds = function(s) {
		window.setTimeout(function() {
			$(me).slideUp();
			window.setTimeout(function() {
				$(me).remove();
			}, 1000);
		}, s*1000);
	};
});
