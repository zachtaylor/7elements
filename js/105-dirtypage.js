SE.dirtypage = {
	on: function() {
		window.onbeforeunload = function() {
			return true;
		};
	},
	off: function() {
		window.onbeforeunload = null;
	}
};
