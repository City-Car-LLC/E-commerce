create table Orders(
    id ulid not null PRIMARY KEY,
    shop_id ulid not null references Shops on delete restrict,
    product_id ulid not null references Products on delete restrict,
    quantity int not null,
    status text not null
);
