SE.widget.control('se-element', function() {
	SE.widget.controlProperty(this, 'src', true)
	Object.defineProperty(this, 'elementid', {
		get: function() {
			return this.getAttribute('elementid');
		},
		set: function(val) {
			this.setAttribute('elementid', val);
			this.src = '/img/icon/element-'+val+'.png';
		},
		enumerable: true,
		configurable: true
	});
});