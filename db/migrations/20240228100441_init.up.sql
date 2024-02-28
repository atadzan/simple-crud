CREATE TABLE IF NOT EXISTS authors (
    id serial primary key,
    username varchar(255) not null,
    password_hash varchar(255) not null,
    created_at timestamp(0) default now()
);

CREATE TABLE IF NOT EXISTS genres (
   id serial primary key,
   title varchar(255) not null
);

CREATE TABLE IF NOT EXISTS books (
   id serial primary key,
   author_id integer references authors(id),
   genre_id integer references genres(id),
   title varchar(255) not null,
   description text not null,
   image_path varchar(255) not null,
   created_at timestamp(0) default now() not null,
   updated_at timestamp(0) default now() not null
);
