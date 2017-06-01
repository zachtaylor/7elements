CREATE TABLE cards(
	id INTEGER PRIMARY KEY,
	type INTEGER,
	image TEXT);
CREATE TABLE cards_element_costs(
	cardid INTEGER,
	element INTEGER,
	count INTEGER);
CREATE TABLE cards_text (
	cardid INTEGER,
	language TEXT,
	name TEXT,
	description TEXT,
	flavor TEXT);
