package db

const userTable = `
create table if not exists users
(
	id serial primary key,
	username varchar(30) not null unique,
	password varchar(256) not null,
	first_name varchar(30) not null,
	last_name varchar(30) not null,
	phone_number integer not null,
	balance	integer default 0,
	status text not null
);
`

const bookTable = `
create table if not exists books
(
	id serial,
	title text not null,
	author text not null,
	genre text not null,
	price int not null,
	description text
);
`

const TransactionTable = `
create table if not exists transactions
(
	id serial,
	book_id int,
	user_id int,
	amount int default 0,
	qty int not null
);
`
