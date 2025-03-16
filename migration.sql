create database orders;

create table clients(
	id int auto_increment primary key,
	name varchar(50),
	email varchar(100)
);

create table products(
	id int auto_increment primary key,
	name varchar(50),
	price decimal(10,2)
);

create table orders(
	id int auto_increment primary key,
	date_order date,
	total_value decimal(10,2),
	client_id int,
	foreign key (client_id) references clients(id)
);


create table order_product(
	id int auto_increment primary key,
	product_id int,
	order_id int,
	quantity int,
	unitary_price decimal(10,2),
	total_price decimal(10,2),
	foreign key (product_id) references products(id),
	foreign key (order_id) references orders(id)
);