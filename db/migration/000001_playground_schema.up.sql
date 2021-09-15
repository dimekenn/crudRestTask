create table if not exists users(
    id bigserial primary key,
    firstname varchar(50) not null,
    lastname varchar(50) not null,
    email varchar(40),
    age int not null,
    created_at time default now()
    );