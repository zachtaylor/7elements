customElements.define('se-symbol', class extends HTMLElement {
	static get observedAttributes() {
		return ['icon'];
	}

	constructor() {
		super();
	}

	getShadow() {
		return SE.shadow.attach(this, 'se-symbol')
	}

	connectedCallback() {
		this.getShadow()
		console.warn("deprecated use of <se-symbol> type", this);
	}

	get icon() {
		return this.getAttribute('icon')
	}

	set icon(val) {
		this.setAttribute('icon', val);
	}

	attributeChangedCallback(attr, oldval, newval) {
		if (attr == 'icon') {
			$(this).css({
				content: "url('/img/icon/"+newval+".png')"
			});
		}
	}
});

customElements.define('se-symbol-tap', class extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback() {
		SE.shadow.attach(this, 'se-symbol-tap')
	}
});
