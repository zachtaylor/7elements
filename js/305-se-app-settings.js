SE.widget.control('se-app-settings', function(app) {
	var me = this;

	me.keybindAppDrawer = function(key) {
		if (key == 13) {
			$(me.keybindingAppDrawer).val('Enter');
		} else if (key == 27) {
			$(me.keybindingAppDrawer).val('Escape');
		} else {
			$(me.keybindingAppDrawer).val('[disabled]');
		}
	};

	me.keybindChat = function(key) {
		if (key == 13) {
			$(me.keybindingChat).val('Enter');
		} else if (key == 27) {
			$(me.keybindingChat).val('Escape');
		} else {
			$(me.keybindingChat).val('[disabled]');
		}
	};

	app.keybind.on(function(key, action) {
		if (action == 'appdrawer') {
			me.keybindAppDrawer(key);
		} else if (action == 'chat') {
			me.keybindChat(key);
		} else {
			console.warn('unrecognized keybind action', key, action);
		}
	});

	$(me.keybindingAppDrawer).keyup(function(e) {
		e.stopPropagation();
		$(this).val('');
		app.keybind.set(e.which, 'appdrawer');
	});
	$(me.keybindingChat).keyup(function(e) {
		e.stopPropagation();
		$(this).val('');
		app.keybind.set(e.which, 'chat');
	});
});
