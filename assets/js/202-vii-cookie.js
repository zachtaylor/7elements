vii.cookie = function(name, val) {
	if (typeof val == "function") {
		SE.event.on('vii.cookie.'+name, val);
		return val(vii.cookie(name));
	} else if (val) {
		console.debug('vii.cookie:', name, val);
		document.cookie = name+'='+val+';Path="/";';
		SE.event.fire('vii.cookie', name, val);
		SE.event.fire('vii.cookie.'+name, val);
	} else {
		$.each(document.cookie.split(';'), function(i, s) {
			var cookie = s.trim().split('=');
			var k = cookie[0];
			var v = cookie[1];
			if (name == k) {
				val = v;
				return false;
			}
		});
	}
	return val;
};
