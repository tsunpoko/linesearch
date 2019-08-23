create table chats
(
  id serial primary key,
  name text not null,
  description text not null,
  url text not null,
  img text not null,
  num integer not null default 0
);
