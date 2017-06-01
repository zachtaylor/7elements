-- 1 drop creatures

INSERT INTO cards (id, type, image)
	VALUES (1, 3, "/img/card/1.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (1, 1, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (1, "en-US", "Time Dealer", "2/2<br/><span class='el-symbol-tap'></span>: Create a tapped land of any color<br/>:Time Dealer enters tapped", "");

INSERT INTO cards (id, type, image)
	VALUES (2, 3, "/img/card/2.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (2, 2, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (2, "en-US", "Novice Seer", "1/1<br/><span class='el-symbol-tap'></span>: Scry 1<br/>Sacrifice Novice Seer: Draw a card", "");

INSERT INTO cards (id, type, image)
	VALUES (3, 3, "/img/card/3.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (3, 3, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (3, "en-US", "Elven Child", "1/2<br/><span class='el-symbol-tap'></span>: Target creature gets +1/+0", "");

INSERT INTO cards (id, type, image)
	VALUES (4, 3, "/img/card/4.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (4, 4, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (4, "en-US", "Zealot", "2/2", "So where is the party?");

INSERT INTO cards (id, type, image)
	VALUES (5, 3, "/img/card/5.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (5, 5, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (5, "en-US", "Ifrit", "2/1<br/>:Whenever Ifrit dies, deal damage equal to its' attack to target creature or player", "");

INSERT INTO cards (id, type, image)
	VALUES (6, 3, "/img/card/6.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (6, 6, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (6, "en-US", "Pixie", "1/1<br/><span class='el-symbol-6'></span><span class='el-symbol-tap'></span>: Create a Pixie", "");

INSERT INTO cards (id, type, image)
	VALUES (7, 3, "/img/card/7.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (7, 7, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (7, "en-US", "Nightmare Ader", "1/1<br/>:Whenever Nightmare Ader dies, destroy target creature", "");

-- 1 drop instants

INSERT INTO cards (id, type, image)
	VALUES (8, 2, "/img/card/8.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (8, 1, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (8, "en-US", "Summon Land", "1. Choose a color<br/>2. Create a land (tapped) of chosen color<br/>3. Draw a card", "Prepare for tomorrow");

INSERT INTO cards (id, type, image)
	VALUES (9, 2, "/img/card/9.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (9, 2, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (9, "en-US", "Bend Will", "1. Target card<br/>2. Tap chosen card<br/>3. Draw a card", "It's a trap!");

INSERT INTO cards (id, type, image)
	VALUES (10, 2, "/img/card/10.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (10, 3, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (10, "en-US", "Invigorate", "1. Target creature card<br/>2. Chosen creature gets +2/+0<br/>3. Draw a card", "The best defense is a strong offense");

INSERT INTO cards (id, type, image)
	VALUES (11, 2, "img/card/11.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (11, 4, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (11, "en-US", "Energize", "1. Add <span class='mana-cost'>1</span><span class='el-symbol-4'></span> to your mana pool", "Now is the time!");

INSERT INTO cards (id, type, image)
	VALUES (12, 2, "/img/card/12.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (12, 5, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (12, "en-US", "Burn", "1. Target a creature card<br/>2. Deal 2 damage to chosen creature<br/>3. Draw a card", "Burn...");

INSERT INTO cards (id, type, image)
	VALUES (13, 2, "/img/card/13.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (13, 6, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (13, "en-US", "Glory", "1. Target a creature card<br/>2. Chosen creature gets +0/+2<br/>3. Draw a card", "Is that better?");

INSERT INTO cards (id, type, image)
	VALUES (14, 2, "/img/card/14.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (14, 7, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (14, "en-US", "Consume", "1. Target your creature<br/>2. Destroy chosen creature<br/>3. Gain 1 life and draw 1 card per hp of destroyed creature", "Never forget");
