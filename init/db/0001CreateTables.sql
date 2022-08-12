CREATE TABLE patch (
	patch INT) ENGINE=InnoDB;
INSERT INTO patch (patch) VALUES (0);
CREATE TABLE accounts (
	username VARCHAR(24) UNIQUE PRIMARY KEY NOT NULL,
	email VARCHAR(255),
	password TINYBLOB,
	skill INT,
	coins INT,
	register INT,
	lastlogin INT) ENGINE=InnoDB;
CREATE TABLE accounts_cards(
	username VARCHAR(24) NOT NULL,
	card INT,
	register INT,
	notes VARCHAR(255),
	CONSTRAINT FOREIGN KEY (username)	REFERENCES accounts(username)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB;
CREATE TABLE accounts_decks (
	username VARCHAR(24),
	id INT,
	name VARCHAR(255),
	wins INT,
	cover INT,
	register INT,
	CONSTRAINT FOREIGN KEY (username)	REFERENCES accounts(username)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB;
CREATE TABLE accounts_decks_items (
	username VARCHAR(24),
	id INT,
	cardid INT,
	amount INT,
	CONSTRAINT FOREIGN KEY (username)	REFERENCES accounts(username)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB;
CREATE TABLE cards (
	id INT PRIMARY KEY NOT NULL,
	name VARCHAR(255),
	text VARCHAR(255),
	type INT,
	image VARCHAR(255)
) ENGINE=InnoDB;
CREATE TABLE cards_bodies (
	cardid INT,
	attack INT,
	health INT,
	CONSTRAINT FOREIGN KEY (cardid)	REFERENCES cards(id)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB;
CREATE TABLE cards_element_costs(
	cardid INT,
	element INT,
	count INT,
	CONSTRAINT FOREIGN KEY (cardid)	REFERENCES cards(id)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB;
CREATE TABLE cards_powers (
	cardid INT,
	id INT,
	text VARCHAR(255),
	xtrigger VARCHAR(255),
	target VARCHAR(255),
	usesturn INT,
	useskill INT,
	script VARCHAR(255),
	CONSTRAINT FOREIGN KEY (cardid)	REFERENCES cards(id)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB;
CREATE TABLE cards_powers_costs (
	cardid INT,
	powerid INT,
	element INT,
	count INT,
	CONSTRAINT FOREIGN KEY (cardid)	REFERENCES cards(id)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB;
CREATE TABLE decks (
	id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	name VARCHAR(255),
	level INT,
	cover VARCHAR(255)) ENGINE=InnoDB;
CREATE TABLE decks_items (
	deckid INT,
	cardid INT,
	amount INT,
	CONSTRAINT FOREIGN KEY (deckid)	REFERENCES decks(id)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB;
CREATE TABLE httptrack (
	name VARCHAR(255),
	addr VARCHAR(255),
	heat INT,
	t INT
) ENGINE=InnoDB;
CREATE TABLE packs (
	id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	name VARCHAR(255),
	size INT,
	cost INT,
	image VARCHAR(255)
) ENGINE=InnoDB;
CREATE TABLE packs_cards (
	packid INT,
	cardid INT,
	weight INT,
	CONSTRAINT FOREIGN KEY (packid)	REFERENCES packs(id)
		ON DELETE CASCADE
		ON UPDATE RESTRICT,
	CONSTRAINT FOREIGN KEY (cardid)	REFERENCES cards(id)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB;

UPDATE patch SET patch=1;
