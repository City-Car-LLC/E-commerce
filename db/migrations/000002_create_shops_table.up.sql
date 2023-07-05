CREATE TABLE Shops (
    id      ulid    not null primary key,
    name    text    not null,
    logo    text    not null,
    opening_time time,
    closing_time time,
    phone_number text not null
);
