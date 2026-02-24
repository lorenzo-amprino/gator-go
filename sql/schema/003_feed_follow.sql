-- +goose Up 
create table feed_follows (
    id UUID primary key,
    feed_id UUID not null,
    user_id UUID not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    unique (feed_id, user_id),
foreign key (feed_id) references feeds(id) on delete cascade,
foreign key (user_id) references users(id) on delete cascade);

-- +goose Down
drop table feed_follows;