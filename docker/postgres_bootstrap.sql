create schema source;
SET search_path TO source;

create table orders(
    id serial primary key,
    customer_id integer not null,
    total numeric(10, 2) not null,
    created_at timestamp not null default now(),
    modified_at timestamp not null default now()
);

insert into orders (customer_id, total) 
values (1, 100.00),
       (2, 45.00),
       (3, 125.00),
       (4, 66.00),
       (5, 423.00);