SE.dirtypage = {
	on: function() {
		window.onbeforeunload = function() {
			return true;
		};
		SE.event.fire('dirtypage', true);
	},
	off: function() {
		window.onbeforeunload = null;
		SE.event.fire('dirtypage', false);
	},
	state: function() {
		return !!window.onbeforeunload;
	}
};
