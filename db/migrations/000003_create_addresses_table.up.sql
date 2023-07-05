-- Создание таблицы "Addresses"
create table Addresses (
    id SERIAL PRIMARY KEY,
    shop_id ulid not null references Shops on delete restrict ,
    address VARCHAR(255),
    coordinates POINT
);
