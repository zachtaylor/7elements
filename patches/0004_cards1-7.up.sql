INSERT INTO cards (id, type, image) VALUES (1, 2, '/img/card/1.jpg');
INSERT INTO cards_element_costs (cardid, element, count) VALUES (1, 1, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor) VALUES (1, "en-US", "Summon Land", "1. Choose a color\n2. Create a land (tapped) of chosen color", "Prepare for tomorrow");

INSERT INTO cards (id, type, image) VALUES (2, 2, '/img/card/2.jpg');
INSERT INTO cards_element_costs (cardid, element, count) VALUES (2, 2, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor) VALUES (2, "en-US", "Bend Will", "1. Target card\n2. Tap chosen card\n3. Draw a card", "It's a trap!");

INSERT INTO cards (id, type, image) VALUES (3, 2, '/img/card/3.jpg');
INSERT INTO cards_element_costs (cardid, element, count) VALUES (3, 3, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor) VALUES (3, "en-US", "Invigorate", "1. Target creature card\n2. Chosen creature gets +2/+0\n3. Draw a card", "The best defense is a strong offense");

INSERT INTO cards (id, type, image) VALUES (4, 3, '/img/card/4.jpg');
INSERT INTO cards_element_costs (cardid, element, count) VALUES (4, 4, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor) VALUES (4, "en-US", "Zealot", "2/2", "So where is the party?");

INSERT INTO cards (id, type, image) VALUES (5, 2, '/img/card/5.jpg');
INSERT INTO cards_element_costs (cardid, element, count) VALUES (5, 5, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor) VALUES (5, "en-US", "Burn", "1. Target a creature card\n2. Deal 2 damage to chosen creature\n3. Draw a card", "Burn...");

INSERT INTO cards (id, type, image) VALUES (6, 2, '/img/card/6.jpg');
INSERT INTO cards_element_costs (cardid, element, count) VALUES (6, 6, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor) VALUES (6, "en-US", "Glory", "1. Target a creature card\n2. Chosen creature gets +0/+2\n3. Draw a card", "Is that better?");

INSERT INTO cards (id, type, image) VALUES (7, 2, '/img/card/7.jpg');
INSERT INTO cards_element_costs (cardid, element, count) VALUES (7, 7, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor) VALUES (7, "en-US", "Consume", "1. Target your creature\n2. Destroy chosen creature\n3. Gain 1 life and draw 1 card per hp of destroyed creature", "Never forget");
