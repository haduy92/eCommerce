create database ecommerce;

\c ecommerce;

create table if not exists person (
  id          uuid primary key,
  name        varchar not null,
  email       varchar unique not null,
  password    varchar not null,
  role        varchar not null,
  created_at  TIMESTAMP,
  updated_at  TIMESTAMP,
  deleted_at  TIMESTAMP
);

create extension if not exists "uuid-ossp";

insert into person (id, name, email, password, role, created_at)
values
(uuid_generate_v4(), 'John Smith', 'john.smith@gmail.com', '0sa9dd0jaasd98', 'admin', now() at time zone 'utc'),
(uuid_generate_v4(), 'Lina Matha', 'lina.matha@gmail.com', '0sa9dd0jaasd98', 'admin', now() at time zone 'utc'),
(uuid_generate_v4(), 'Ori Jimenez','aliquam.arcu.aliquam@protonmail.ca', '0sa9dd0jaasd98', 'user', now() at time zone 'utc'),
(uuid_generate_v4(), 'Byron Richardson','suspendisse.dui@hotmail.net', '0sa9dd0jaasd98', 'user', now() at time zone 'utc'),
(uuid_generate_v4(), 'Christopher Wiggins','et.rutrum.non@icloud.com', '0sa9dd0jaasd98', 'admin', now() at time zone 'utc'),
(uuid_generate_v4(), 'Forrest Potts','vitae.mauris@hotmail.com', '0sa9dd0jaasd98', 'user', now() at time zone 'utc'),
(uuid_generate_v4(), 'Chase Mercer','neque.sed@icloud.org', '0sa9dd0jaasd98', 'admin', now() at time zone 'utc');