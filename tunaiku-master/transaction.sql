create table transaction(
	billing_id SERIAL primary key,
	name varchar(125),
	tax_code smallint,
	price float
);

CREATE TABLE users
(
  user_id serial,
  date timestamp not null,
  ktp varchar(100) not null,
  birth_date timestamp not null,
  gender varchar(20) not null,
  name varchar(100) not null,
  amount bigint not null,
  period int not  null,
   PRIMARY KEY(user_id)
);