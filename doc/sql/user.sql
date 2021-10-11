create table "user"
(
    id uuid default gen_random_uuid() constraint user_pkey primary key,
    email  varchar(20) null,
    phone  varchar(20) null,
    role  varchar(10) null,
    disabled  bool not null default false,
    password  varchar(20) not null,
    create_at timestamptz default now(),
    update_at timestamptz default now()
);

comment on column "user".id is '用户ID';
comment on column "user".email is '用户账号';
comment on column "user".password is '用户密码';
comment on column "user".create_at is '创建时间';
comment on column "user".update_at is '更新时间';

alter table "user" owner to postgres;

