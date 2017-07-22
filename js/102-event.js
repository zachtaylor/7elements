SE.event = {
	cache: {},
	on: function(event, f) {
		SE.event.cache[event] = SE.event.cache[event] || [];
		SE.event.cache[event].push(f);
		return SE.event.cache[event].length - 1;
	},
	off: function(event, id) {
		SE.event.cache[event][id] = false;
	},
	fire: function() {
		var args = Array.prototype.slice.call(arguments);
		var event = args.shift();
		var eventList = SE.event.cache[event];
		if (!eventList) return;
		for (var i = 0; i < eventList.length; i++) {
			var f = eventList[i];
			if (f) f.apply(null, args);
		}
	}
};
