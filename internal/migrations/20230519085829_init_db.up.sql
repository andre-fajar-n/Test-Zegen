create table if not exists "users" (
    id serial primary key,
    username varchar(100) not null,
    password varchar(255) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    deleted_at timestamptz
);

create table if not exists "books" (
    id serial primary key,
    title varchar(255) not null,
    published_year integer not null,
    isbn varchar(255) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    deleted_at timestamptz
);

create table if not exists "authors" (
    id serial primary key,
    name varchar(255) not null,
    country varchar(255) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    deleted_at timestamptz
);

create table if not exists "book_author" (
    id serial primary key,
    book_id integer not null,
    author_id integer not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    deleted_at timestamptz,
    constraint fk_book foreign key(book_id) references books(id) on update restrict on delete restrict,
    constraint fk_author foreign key(author_id) references authors(id) on update restrict on delete restrict
);