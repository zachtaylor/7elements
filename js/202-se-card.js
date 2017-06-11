customElements.define('se-card', class extends HTMLElement {
	static get observedAttributes() {
		return ['cardid'];
	}

	constructor() {
		super();
	}

	getShadow() {
		return SE.shadow.attach(this, 'se-card');
	}

	connectedCallback() {
		this.getShadow();
		SE.req.getcards(this.updater());
	}

	get cardid() {
		return this.getAttribute('cardid');
	}

	set cardid(val) {
		this.setAttribute('cardid', val);
	}

	attributeChangedCallback(attr, oldval, newval) {
		if (attr == 'cardid') {
			SE.req.getcards().then(this.updater());
		}
	}

	update(data) {
		data = data[this.cardid];
		this.getShadow().then(function(shadow) {
			shadow.querySelector('#name').innerHTML = data.name;
			shadow.querySelector('#text').innerHTML = data.description;
			shadow.querySelector('#flavor').innerHTML = data.flavor;

			$(shadow.querySelector('#costs')).empty();
			$.each(data.elementcosts, function(i, elementcost) {
				for (var i=0; i<elementcost.count; i++) {
					var symbol = $('<se-symbol icon="element-'+elementcost.element+'"></se-symbol>')[0];
					shadow.querySelector('#costs').append(symbol);
				}
			});
		});
	}

	updater() {
		var card = this;
		return function(data) {
			card.update(data);
		};
	}
});
