create table if not exists "profile"
(
    id         bigserial primary key not null unique,
    uid        varchar(8)            not null,
    avatar     varchar(255)          null,
    nickname   varchar(20)           null,
    gender     varchar(1)            not null default 'f',
    motto      varchar(255)          null,
    created_at timestamp                  default now(),
    updated_at timestamp                  default now(),
    deleted_at timestamp             null default null,
    foreign key (uid) references public."user" (id)
);

comment on column "profile".id is '主键ID';
comment on column "profile".uid is '用户uuid';
comment on column "profile".avatar is '头像url';
comment on column "profile".nickname is '昵称';
comment on column "profile".gender is '性别 m/f';
comment on column "profile".motto is '签名';

alter table "profile" owner to postgres;

