ALTER TABLE accounts_decks RENAME TO _accounts_decks;
CREATE TABLE accounts_decks (
	username TEXT,
	id INTEGER,
	version TEXT,
	name TEXT,
	wins INTEGER,
	color TEXT,
	register INTEGER
);
INSERT INTO accounts_decks (username, id, version, name, wins, color, register)
  SELECT username, id, "", name, wins, color, register
  FROM _accounts_decks;
DROP TABLE _accounts_decks;
ALTER TABLE accounts_decks_items RENAME TO _accounts_decks_items;
CREATE TABLE accounts_decks_items (
	username TEXT,
	id INTEGER,
	version TEXT,
	cardid INTEGER,
	amount INTEGER
);
INSERT INTO accounts_decks_items (username, id, version, cardid, amount)
  SELECT username, id, "", cardid, amount
  FROM _accounts_decks_items;
DROP TABLE _accounts_decks_items;
