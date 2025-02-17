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

CREATE TABLE public.personal_reminders (
	id serial primary key,
	reminder varchar NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	modified_at timestamp NOT NULL DEFAULT now(),
	reminder_date date NULL
);

INSERT INTO public.personal_reminders (reminder, created_at, modified_at, reminder_date)
VALUES 
    ('Take out the garbage', now(), now(), current_date),
    ('Text her back', now(), now(), current_date);