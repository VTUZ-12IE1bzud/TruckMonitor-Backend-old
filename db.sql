-- Schema
DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA public;

CREATE TABLE role(
  id serial PRIMARY KEY,
	name VARCHAR UNIQUE NOT NULL);

CREATE TABLE account (
	id serial PRIMARY KEY,
	role_id INTEGER NOT NULL REFERENCES role,
	email VARCHAR UNIQUE NOT NULL,
	password VARCHAR (50) NOT NULL,
	surname VARCHAR UNIQUE NOT NULL,
	name VARCHAR UNIQUE NOT NULL,
	patronymic VARCHAR,
	date_of_birth TIMESTAMP NOT NULL,
	photo VARCHAR UNIQUE NOT NULL,
	phone VARCHAR (20) UNIQUE NOT NULL);


-- Default Data
--   Roles
INSERT INTO public.role(name) VALUES ('admin');
INSERT INTO public.role(name) VALUES ('client');
INSERT INTO public.role(name) VALUES ('driver');
INSERT INTO public.role(name) VALUES ('manager');

--  Users
INSERT INTO public.account
		(role_id, email, password, surname, name, patronymic, date_of_birth, photo, phone)
		VALUES (3, 'ivanov@truck.ru', 'qwerty12345', 'Иванов', 'Иван', 'Иванович', '1990-12-3','http://images.clipartpanda.com/driver-clipart-smiling-driver-in-car-holding-steering-wheel.jpg', '88000000000');
