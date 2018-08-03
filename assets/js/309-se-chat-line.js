SE.widget.control('se-chat-line', function(data) {
	var me = this;
	me.username = data.username;
	me.message = data.message;
});
