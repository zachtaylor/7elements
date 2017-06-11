INSERT INTO cards (id, type, image)
	VALUES (8, 2, "/img/card/8.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (8, 1, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (8, "en-US", "Summon Land", "Create a tapped land of any color", "Prepare for tomorrow");

INSERT INTO cards (id, type, image)
	VALUES (9, 2, "/img/card/9.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (9, 2, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (9, "en-US", "Bend Will", "Tap target permanent, then draw a card", "It's a trap!");

INSERT INTO cards (id, type, image)
	VALUES (10, 2, "/img/card/10.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (10, 3, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (10, "en-US", "Invigorate", "Target creature gets +2/+0, then draw a card", "The best defense is a strong offense");

INSERT INTO cards (id, type, image)
	VALUES (11, 2, "img/card/11.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (11, 4, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (11, "en-US", "Energize", "Add <se-symbol icon='element-4\'></se-symbol><se-symbol icon='element-4'></se-symbol> to your mana pool", "Now is the time!");

INSERT INTO cards (id, type, image)
	VALUES (12, 2, "/img/card/12.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (12, 5, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (12, "en-US", "Burn", "Deal 2 damage to target creature, then draw a card", "Burn...");

INSERT INTO cards (id, type, image)
	VALUES (13, 2, "/img/card/13.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (13, 6, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (13, "en-US", "Grace", "Target creature gets +0/+2, then draw a card", "Is that better?");

INSERT INTO cards (id, type, image)
	VALUES (14, 2, "/img/card/14.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (14, 7, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (14, "en-US", "Consume", "Sacrifice a creature, then gain 1 life and draw 1 card per hp of sacrificed creature", "Never forget");
