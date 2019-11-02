-- cards

INSERT INTO cards (id, type, name, text, image)
	VALUES (15, 3, "Banhammer", "shwing", "/img/card/15.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (15, 1, 3);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (15, 1, "", "being-item", 1, 0, "banhammer", "Remove Target Being or Item from the Present");

INSERT INTO cards (id, type, name, text, image)
	VALUES (16, 3, "Burning Rage", "", "/img/card/16.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (16, 2, 2);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (16, 1, "sunset", "self", 0, 0, "burning-rage", "Burning Rage does 1 damage to each enemy Player");

INSERT INTO cards (id, type, name, text, image)
	VALUES (17, 1, "Call the Banners", "Join me!", "/img/card/17.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (17, 3, 3);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (17, 1, "play", "self", 0, 0, "call-banners", "Add 3 Beings to your Present, each with 2 Attack and 2 Health");

INSERT INTO cards (id, type, name, text, image)
	VALUES (18, 3, "Symbiosis", "Mycelliated Spiranthes", "/img/card/18.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (18, 4, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (18, 1, "sunrise", "", 0, 0, "symbiosis", "Target Being gains 1 Attack");

INSERT INTO cards (id, type, name, text, image)
	VALUES (19, 3, "Crystal Ball", "I see...", "/img/card/19.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (19, 5, 2);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (19, 1, "", "self", 1, 0, "crystal-ball", "Look at the next card in your Future, then you can choose to Shuffle your Future");

INSERT INTO cards (id, type, name, text, image)
	VALUES (20, 3, "Font of Life", "", "/img/card/20.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (20, 6, 2);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (20, 1, "", "my-being", 1, 0, "font-life", "Target Being gains 2 Health");

INSERT INTO cards (id, type, name, text, image)
	VALUES (21, 3, "Necromancy 101", "", "/img/card/21.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (21, 7, 3);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (21, 1, "", "mypast-being", 1, 0, "intro-necromancy", "Add Target Being in your Past to your Present, set its' Health to 1, you lose 1 Life");

-- packs

DELETE FROM packs WHERE id=2;
DELETE FROM packs WHERE id=5;

INSERT INTO packs (id, name, size, cost, image)
	VALUES (7, "November2019 Cards(1)", 1, 3, "/img/pack/7.jpg");
INSERT INTO packs_cards (packid, cardid, weight)
  VALUES (7, 15, 1), (7, 16, 1), (7, 17, 1), (7, 18, 1), (7, 19, 1), (7, 20, 1), (7, 21, 1);
INSERT INTO packs (id, name, size, cost, image)
	VALUES (8, "November2019 Cards(7)", 5, 7, "/img/pack/8.jpg");
INSERT INTO packs_cards (packid, cardid, weight)
  VALUES (8, 15, 1), (8, 16, 1), (8, 17, 1), (8, 18, 1), (8, 19, 1), (8, 20, 1), (8, 21, 1);

-- decks

UPDATE decks_items SET amount=1 WHERE deckid=1;
INSERT INTO decks_items (deckid, cardid, amount)
  VALUES (1, 15, 1), (1, 16, 1), (1, 17, 1), (1, 18, 1), (1, 19, 1), (1, 20, 1), (1, 21, 1);

DELETE FROM decks WHERE id=2;
DELETE FROM decks WHERE id=3;

INSERT INTO decks (id, name, cover)
	VALUES (2, "WhiteGold", 8);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (2, 1, 3), (2, 3, 3), (2, 8, 3), (2, 9, 3), (2, 10, 3), (2, 15, 3), (2, 16, 3);

INSERT INTO decks (id, name, cover)
	VALUES (3, "RememberForever", 7);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (3, 3, 3), (3, 6, 3), (3, 7, 3), (3, 10, 3), (3, 14, 3), (3, 17, 3), (3, 21, 3);

INSERT INTO decks (id, name, cover)
	VALUES (4, "Grow", 11);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (4, 3, 3), (4, 4, 3), (4, 10, 3), (4, 11, 3), (4, 13, 3), (4, 17, 3), (4, 18, 3);

INSERT INTO decks (id, name, cover)
	VALUES (5, "ItemTech", 19);
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (5, 1, 3), (5, 5, 3), (5, 8, 3), (5, 15, 3), (5, 18, 3), (5, 19, 3), (5, 20, 3);
