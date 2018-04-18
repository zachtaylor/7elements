SE.widget.control('se-app-notification', function(data) {
	var me = this;
	me.title = data.title;
	me.message = data.message;

	if (data.class == 'chat') {
		$('img', me)[0].src='/img/icon/chat.black.128px.png';
	} else if (data.class == 'match') {
		$('img', me)[0].src='/img/icon/attack.32px.png';
	} else if (data.class == 'error') {
		$('img', me)[0].src='/img/icon/x.20px.png';
	}

	$(me).click(function() {
		$(me).remove();
	});

	if (data.timeout) window.setTimeout(function() {
		$(me).remove();
	}, data.timeout);
});
