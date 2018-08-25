INSERT INTO cards (id, type, image)
	VALUES (8, 1, "/img/card/8.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (8, 1, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (8, "en-US", "New Element", "Create a new element, then draw a card", "Prepare for tomorrow");

INSERT INTO cards (id, type, image)
	VALUES (9, 1, "/img/card/9.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (9, 2, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (9, "en-US", "Bend Will", "Spend target item or body, then draw a card", "It's a trap!");

INSERT INTO cards (id, type, image)
	VALUES (10, 1, "/img/card/10.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (10, 3, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (10, "en-US", "Invigorate", "Target body gets +2 attack, then draw a card", "The best defense is a strong offense");

INSERT INTO cards (id, type, image)
	VALUES (11, 1, "img/card/11.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (11, 4, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (11, "en-US", "Energize", "Add <se-symbol icon='element-4\'></se-symbol><se-symbol icon='element-4'></se-symbol> to your available elements this turn", "Now is the time!");

INSERT INTO cards (id, type, image)
	VALUES (12, 1, "/img/card/12.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (12, 5, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (12, "en-US", "Burn", "Deal 2 damage to target body, then draw a card", "Burn...");

INSERT INTO cards (id, type, image)
	VALUES (13, 1, "/img/card/13.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (13, 6, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (13, "en-US", "Grace", "Target body gets +2 life, then draw a card", "Is that better?");

INSERT INTO cards (id, type, image)
	VALUES (14, 1, "/img/card/14.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (14, 7, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (14, "en-US", "Consume", "Select 1 of your bodies, then gain 'x' lives and draw 'x' cards, where 'x' is the body's life, then destroy that body", "Never forget");
