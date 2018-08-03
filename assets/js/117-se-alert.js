SE.widget.control('se-alert', function(data) {
	var me = this;

	$(me).addClass('alert-'+data.class);
	$('.se-alert-message', me)[0].innerHTML = '<strong>'+data.username+'</strong> : '+data.message;
	$(me).click(function() {
		me.timeout();
	});

	me.timeout = function() {
		$(me).slideUp();
		window.setTimeout(me.cleanup, 1000);
	};
	me.cleanup = function() {
		$(me).remove();
	};

	window.setTimeout(me.timeout, data.timeout || 7000);
});
