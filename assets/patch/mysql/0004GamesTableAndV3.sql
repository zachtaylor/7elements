-- bugfix

DELETE FROM accounts_decks_items WHERE amount=0;

-- database changes

ALTER TABLE cards DROP COLUMN image;
ALTER TABLE accounts_decks DROP COLUMN wins;
ALTER TABLE decks DROP COLUMN cover;
ALTER TABLE decks DROP COLUMN level;
ALTER TABLE decks ADD cover INT;
ALTER TABLE packs DROP COLUMN image;

CREATE TABLE games (
	time INT,
	winner VARCHAR(255),
	winner_deck VARCHAR(100),
	winner_skill INT,
	loser VARCHAR(255),
	loser_deck VARCHAR(100),
	loser_skill INT
) ENGINE=InnoDB;

-- gift

UPDATE accounts SET coins = coins + 21;

-- patch

DELETE FROM cards_powers_costs WHERE cardid=2;
UPDATE cards_powers
	SET text="Deal 1 Damage to Target Player or Being", target="player-being", usesturn=0, useskill=1
	WHERE cardid=2;

UPDATE cards_powers
	SET text="Target Player or other Being gains Life equal to my Life, then remove me from the Present", target="player-other-being", usesturn=0
	WHERE cardid=6;

UPDATE cards_powers
	SET useskill=0
	WHERE cardid=7;
INSERT INTO cards_powers_costs (cardid, powerid, element, count)
	VALUES (7, 1, 7, 1);

UPDATE cards_powers
	SET text="Deal 2 Damage to Target Player or Being", target="player-being"
	WHERE cardid=9;

UPDATE cards_powers
	SET text="Remove Target Being from the Present. Remove Target's Card from the Past, and add it to its owner's Hand"
	WHERE cardid=12;

UPDATE cards_powers
	SET text="Target Being gains 3 Life"
	WHERE cardid=13;

UPDATE cards_powers
	SET text="Target Being in your Past, add a Copy of Target's Card to your Hand"
	WHERE cardid=14;

UPDATE cards_powers
	SET text="Remove Target Card from the Past", target="past"
	WHERE cardid=15;

UPDATE cards_powers
	SET text="Deal 1 Damage to each enemy Player"
	WHERE cardid=16;

UPDATE cards_powers
	SET text="Add 3 Beings to your Present, each with 2 Attack and 2 Life"
	WHERE cardid=17;

UPDATE cards_element_costs
	SET count=2 WHERE cardid=18;

UPDATE cards_powers
	SET text="Target Being gains 1 Life"
	WHERE cardid=20 AND id=2;

UPDATE cards_powers_costs
	SET count=1 WHERE cardid=21;

-- cards

INSERT INTO cards (id, type, name, text)
	VALUES (22, 3, "Handrails", " ... precautions ... ");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (22, 1, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (22, 1, "", "being", 1, 0, "handrails", "Target Being becomes Awake");

INSERT INTO cards (id, type, name, text)
	VALUES (23, 1, "Fireball", "Now, burn!");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (23, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (23, 2, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (23, 1, "play", "being", 0, 0, "fireball", "Deal 3 Damage to Target Being");

INSERT INTO cards (id, type, name, text)
	VALUES (24, 2, "Golem", "From the clay");
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(24, 3, 3);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (24, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (24, 3, 1);

INSERT INTO cards (id, type, name, text)
	VALUES (25, 2, "Panther", "purrr...");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (25, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (25, 4, 1);
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(25, 0, 5);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (25, 1, "", "being", 1, 0, "panther", "Enter Combat with Target Being");

INSERT INTO cards (id, type, name, text)
	VALUES (26, 1, "Counter", "Nope!");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (26, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (26, 5, 2);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (26, 1, "play", "play", 0, 0, "counter", "Stop any Play, and put that Card into the Past");

INSERT INTO cards (id, type, name, text)
	VALUES (27, 3, "Cloning Pool", "");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (27, 6, 3);
INSERT INTO cards_powers_costs (cardid, powerid, element, count)
	VALUES (27, 1, 6, 3);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (27, 1, "", "my-being", 1, 0, "cloning-pool", "Target Being you own gains 1 Life. Create a new Being with Target's Card");

INSERT INTO cards (id, type, name, text)
	VALUES (28, 1, "Painful Memories", "");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (28, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (28, 7, 2);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (28, 1, "", "being", 0, 0, "painful-memories", "Each enemy Player loses 1 Life for every Being Card in your Past");

-- decks 

DELETE FROM decks WHERE id=2;
DELETE FROM decks WHERE id=3;
DELETE FROM decks WHERE id=4;
DELETE FROM decks WHERE id=5;
DELETE FROM decks WHERE id=6;
DELETE FROM decks WHERE id=7;

UPDATE decks SET cover=0 WHERE id=1;
INSERT INTO decks_items (deckid, cardid, amount)
  VALUES (1, 22, 1), (1, 23, 1), (1, 24, 1), (1, 25, 1), (1, 26, 1), (1, 27, 1), (1, 28, 1);

INSERT INTO decks (id, name, cover)
	VALUES (2, "Macro Burn", 8);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (2, 1, 3), (2, 2, 3), (2, 8, 3), (2, 9, 3), (2, 14, 3), (2, 22, 3), (2, 23, 3);

INSERT INTO decks (id, name, cover)
	VALUES (3, "Ketchup 'n' Mustard", 23);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (3, 2, 3), (3, 3, 3), (3, 9, 3), (3, 10, 3), (3, 16, 3), (3, 23, 3), (3, 24, 3);

INSERT INTO decks (id, name, cover)
	VALUES (4, "Face-Off", 17);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (4, 3, 3), (4, 10, 3), (4, 11, 3), (4, 14, 3), (4, 17, 3), (4, 23, 3), (4, 24, 3);

INSERT INTO decks (id, name, cover)
	VALUES (5, "To Combat!", 25);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (5, 4, 3), (5, 5, 3), (5, 10, 3), (5, 11, 3), (5, 13, 3), (5, 18, 3), (5, 25, 3);

INSERT INTO decks (id, name, cover)
	VALUES (6, "A Control Deck", 12);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (6, 5, 3), (6, 6, 3), (6, 12, 3), (6, 13, 3), (6, 19, 3), (6, 20, 3), (6, 26, 3);

INSERT INTO decks (id, name, cover)
	VALUES (7, "Zerg", 27);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (7, 6, 3), (7, 7, 3), (7, 13, 3), (7, 14, 3), (7, 20, 3), (7, 24, 3), (7, 27, 3);

INSERT INTO decks (id, name, cover)
	VALUES (8, "The Past Remembers", 14);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (8, 1, 3), (8, 7, 3), (8, 8, 3), (8, 14, 3), (8, 21, 3), (8, 24, 3), (8, 28, 3);

-- packs

UPDATE packs
	SET name="Items Pack(1)"
	WHERE id=7;
UPDATE packs
	SET name="Items Pack(5)"
	WHERE id=8;

INSERT INTO packs (id, name, size, cost)
	VALUES (9, "Support Pack(1)", 1, 3);
INSERT INTO packs_cards (packid, cardid, weight)
  VALUES (9, 22, 1), (9, 23, 1), (9, 24, 1), (9, 25, 1), (9, 26, 1), (9, 27, 1), (9, 28, 1);
INSERT INTO packs (id, name, size, cost)
	VALUES (10, "Support Pack(5)", 5, 7);
INSERT INTO packs_cards (packid, cardid, weight)
  VALUES (10, 22, 1), (10, 23, 1), (10, 24, 1), (10, 25, 1), (10, 26, 1), (10, 27, 1), (10, 28, 1);
