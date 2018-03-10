// vii.storage
vii.storage = {
	ready: false,
	detect: function() {
		if (vii.storage.ready === true) {
			return true;
		} else if (vii.storage.ready) {
			return false;
		} else {
			vii.storage.ready = vii.storage.test();
		}
	},
	get:function(k, v) {
		if (vii.storage.detect()) return localStorage.getItem(k) || v;
		else return false;
	},
	set: function(k, v) {
		if (vii.storage.detect()) {
			SE.event.fire('vii.storage.'+k, v);
			return localStorage.setItem(k, v);
		} else {
			console.warn('storage set failed', k);
		}
	},
	test: function() {
		try {
			var s = '_';
			localStorage.setItem(s, s);
			localStorage.removeItem(s);
			return true;
		} catch(e) {
			return e;
		}
	},
	with: function(k, f) {
		SE.go(function() {
			f(vii.storage.get(k));
			SE.event.on('vii.storage.'+k, f)
		});
	}
};
