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
		console.warn("deprecated use of <se-card> type", this);
		SE.api.get('cards', this.updater());
	}

	get cardid() {
		return this.getAttribute('cardid');
	}

	set cardid(val) {
		this.setAttribute('cardid', val);
	}

	attributeChangedCallback(attr, oldval, newval) {
		if (attr == 'cardid') {
			SE.api.get('cards').then(this.updater());
		}
	}

	update(data) {
		data = data[this.cardid];
		this.getShadow().then(function(shadow) {
			shadow.querySelector('#name').innerHTML = data.name;
			shadow.querySelector('#text').innerHTML = data.description;
			shadow.querySelector('#flavor').innerHTML = data.flavor;
			shadow.querySelector('#type').innerHTML = data.type;

			if (data.body) {
				shadow.querySelector('.se-card-attack').innerHTML = data.attack;
				shadow.querySelector('.se-card-health').innerHTML = data.health;
				shadow.querySelector('.se-card-body-divider').style.display = 'initial';
			}

			$(shadow.querySelector('#costs')).empty();
			$.each(data.costs, function(elementid, cost) {
				for (var i=0; i<cost; i++) {
					var symbol = $('<se-symbol icon="element-'+elementid+'"></se-symbol>')[0];
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
