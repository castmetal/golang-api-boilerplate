create table if not exists "example" (
    "id" uuid not null primary key,
    "name" text not null,
    "created_at" timestamptz not null default now(),
    "updated_at" timestamptz not null default now(),
    "deleted_at" timestamptz
);

create index if not exists "example_name_idx" on "example" (
    "name"
);

create index if not exists "example_deleted_at_idx" on "example" (
    "deleted_at"
);


INSERT INTO example
(id, "name", created_at, updated_at, deleted_at)
VALUES('204ef11c-7471-429e-80c1-851e79664bb4'::uuid, 'bifarma', '2021-12-22 11:04:17.346', '2021-12-22 11:04:17.346', NULL);
