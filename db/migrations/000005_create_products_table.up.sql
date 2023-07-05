create table Products (
    id ulid not null primary key,
    shop_id ulid not null references Shops on delete restrict,
    category_id ulid not null references Categories on delete restrict,
    name  text not null,
    image text not null,
    description text not null,
    price decimal not null,
    status text not null
);


