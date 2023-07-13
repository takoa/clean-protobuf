create table if not exists features(
    id         uuid    default gen_random_uuid(),
    latitude   integer not null,
    longitude  integer not null,
    name       text    not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key(id)
);

create table if not exists messages(
    id         uuid DEFAULT gen_random_uuid(),
    feature_id uuid not null,
    body       text not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key(id),
    constraint messages_feature_fkey foreign key(feature_id) references features(id)
);
