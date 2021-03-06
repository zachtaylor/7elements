SE.widget = {
	cache: {},
	controllers: {},
	_promise: function(name) {
		return new Promise(function(resolve, reject) {
			$.get('/html/'+name+'.html').done(resolve).fail(reject);
		});
	},
	_builder: function(name, data) {
		return function WidgetFactory() {
			var args = arguments;
			var widget = $(data)[0];
			$('[data-ctrl]', widget).each(function() {
				SE.widget.controlProperty(widget, this.getAttribute('data-ctrl'));
			});
			$('[handle]', widget).each(function() {
				widget[this.getAttribute('handle')] = this;
			});
			$.each(SE.widget.controllers[name], function(i, f) {
				f.apply(widget, args);
			});
			return widget;
		};
	},
	get: function(name) {
		if (!SE.widget.cache[name]) SE.widget.cache[name] = SE.widget._promise(name);
		return new Promise(function(resolve, reject) {
			SE.widget.cache[name].then(function(data) {
				resolve(SE.widget._builder(name, data));
			}, reject);
		});
	},
	new: function(name) {
		var args = Array.prototype.slice.call(arguments);
		args.shift();

		return new Promise(function(resolve, reject) {
			SE.widget.get(name).then(function(widgetFactory) {
				resolve(widgetFactory.apply(null, args));
			}, reject);
		});
	},
	provide: function(name, data) {
		SE.widget.cache[name] = Promise.resolve(data);
	},
	control: function(name, f) {
		SE.widget.controllers[name] = SE.widget.controllers[name] || [];
		SE.widget.controllers[name].push(f);
	},
	replace: function(name) {
		var script = document.currentScript;
		var args = Array.prototype.slice.call(arguments);
		return new Promise(function(resolve, reject) {
			SE.widget.new.apply(null, args).then(function(widget) {
				$(script).replaceWith(widget);
				resolve(widget);
			}, function() {
				$(script).replaceWith('<span>widget not found: '+name+'</span>');
				reject();
			});
		});
	},
	controlProperty: function(scope, property, hidden) {
		Object.defineProperty(scope, property, {
			get: function() {
				return scope.getAttribute(property)
			},
			set: function(val) {
				scope.setAttribute(property, val);
				var control = $('.data-control-'+property+',[data-ctrl="'+property+'"]', scope);
				if (hidden) {
				} else if (control.length) {
					control[0].innerHTML = val;
				} else {
					console.log("control property not found", '"'+property+'"', scope);
				}
			}
		});
	}
};
