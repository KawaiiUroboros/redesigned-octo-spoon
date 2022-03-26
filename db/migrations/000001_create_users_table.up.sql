BEGIN;
CREATE TABLE IF NOT EXISTS users(
    user_id serial primary key,
    user_name varchar (50) not null,
    user_surname varchar (50) not null,
    address varchar (300) not null,
    is_deleted boolean not null,
    end_date timestamp with time zone,
    notification_interval integer,
    last_confirmation timestamp with time zone
    );
COMMIT;