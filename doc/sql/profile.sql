create table if not exists "profile"
(
    id         bigserial primary key not null unique,
    uid        varchar(8)            not null,
    avatar     varchar(127)          null,
    nickname   varchar(20)           null,
    created_at timestamp                  default now(),
    updated_at timestamp                  default now(),
    deleted_at timestamp             null default null,
    foreign key (uid) references public."user" (id)
);

comment on column "profile".id is '主键ID';
comment on column "profile".uid is '用户uuid';
comment on column "profile".avatar is '头像url';
comment on column "profile".nickname is '昵称';

alter table "profile"
    owner to postgres;

