
-- +goose Up
CREATE TABLE addresses (
	address_id serial,
	street varchar(1000),
	city varchar(1000),
	region varchar(100),
	zip varchar(10),
	country char(2),
	PRIMARY KEY (address_id)
);

CREATE TABLE users (
	user_id serial,
	first_name varchar(1000),
	last_name varchar(1000),
	address_id int,
	PRIMARY KEY (user_id),
	FOREIGN KEY (address_id) REFERENCES addresses (address_id)
);

INSERT INTO addresses VALUES
	(1, '123 Fake Street', 'Toronto', 'Ontario', 'M1V2XC', 'CA'),
	(2, '4920 5th Avenue', 'New York', 'New York', '10010', 'US');

INSERT INTO users VALUES
	(1, 'Tharsan', 'Bhuvanendran', 1),
	(2, 'Shawnette', 'Bhuvanendran', 1),
	(3, 'Jerry', 'Lo', 2);

-- +goose Down
DROP TABLE addresses;
DROP TABLE users;