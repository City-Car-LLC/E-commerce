alter table Categories
    add column parent_id ulid references Categories(id) on delete cascade;