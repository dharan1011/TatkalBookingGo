CREATE DATABASE RAILWAYS;
CREATE TABLE users(
    uid serial primary key,
    first_name varchar(200) not null,
    last_name varchar(200) not null,
    age integer
);
CREATE TABLE trains(
    train_number integer primary key,
    src varchar(5) not null,
    dest varchar(5) not null,
    capacity integer not null
);
CREATE TABLE bookings(
    booking_id serial primary key,
    train_number integer references trains(train_number) ON DELETE CASCADE,
    seat_number varchar(5) not null,
    uid integer references users(uid) ON DELETE CASCADE
);