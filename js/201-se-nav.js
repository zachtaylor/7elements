customElements.define('se-nav', class extends HTMLElement {
	constructor() {
		super();
		SE.event.on('data.myaccount', this.updater())
	}

	connectedCallback() {
		var updater = this.updater();

		SE.shadow.attach(this, 'se-nav').then(function(shadow) {
			var pathname = window.location.pathname;

			if (pathname == '/') {
				shadow.querySelector('.navbar-brand').classList.add('active');
			} else if (pathname == '/cards/') {
				shadow.querySelector('#cards-link').classList.add('active');
			} else if (pathname == '/packs/') {
				shadow.querySelector('#packs-link').classList.add('active');
			} else if (pathname == '/login/') {
				shadow.querySelector('#myaccount-link').classList.add('active');
			} else if (pathname.substr(0, 7) == '/decks/') {
				shadow.querySelector('#decks-link').classList.add('active');
			}

			SE.req.getmyaccount().then(updater, function() {
				$(shadow.querySelector('#decks-link')).hide();
				$(shadow.querySelector('#login-link')).show();
			});
		});
	}

	update(data) {
		if (!this.shadowRoot) return;
		this.shadowRoot.querySelector('#data-cards').innerHTML = data.cards;
		this.shadowRoot.querySelector('#data-packs').innerHTML = data.packs;
		this.shadowRoot.querySelector('#data-username').innerHTML = '('+data.username+')';
	}

	updater() {
		var nav = this;
		return function(data) {
			nav.update(data);
		};
	}
});
