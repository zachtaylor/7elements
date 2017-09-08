CREATE TABLE accounts (
	username TEXT PRIMARY KEY,
	email TEXT,
	password TEXT,
	skill INTEGER,
	coins INTEGER,
	packs INTEGER,
	language TEXT,
	register INTEGER,
	lastlogin INTEGER);
CREATE TABLE accounts_cards(
	username TEXT,
	card INTEGER,
	register INTEGER,
	notes TEXT);
CREATE TABLE accounts_decks (
	username TEXT,
	name TEXT,
	id INTEGER,
	wins INTEGER,
	register INTEGER);
CREATE TABLE accounts_decks_items (
	username TEXT,
	id INTEGER,
	cardid INTEGER,
	amount INTEGER);
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
CREATE TABLE cards_bodies (
	cardid INTEGER,
	attack INTEGER,
	health INTEGER);
