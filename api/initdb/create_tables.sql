create table tweets
(
  id serial primary key,
  name text not null,
  description text not null,
  url text not null,
  description text not null,
  img text not null,
  num integer not null default 0,
);
