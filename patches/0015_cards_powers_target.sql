ALTER TABLE cards_powers RENAME TO _cards_powers;
CREATE TABLE cards_powers (
	cardid INTEGER,
	id INTEGER,
	target TEXT,
	usesturn INTEGER,
	script TEXT
);
INSERT INTO cards_powers (cardid, id, target, usesturn, script)
  SELECT cardid, id, "", usesturn, script
  FROM _cards_powers;
DROP TABLE _cards_powers;
