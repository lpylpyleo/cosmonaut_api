create table if not exists "user"
(
    id varchar(8) not null unique primary key,
    account  varchar(20) null,
    email  varchar(50) null,
    phone  varchar(20) null,
    role  varchar(10) null,
    disabled  bool not null default false,
    password  varchar(64) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp null default null
);

comment on column "user".id is '用户ID';
comment on column "user".account is '用户账号';
comment on column "user".email is '用户邮箱';
comment on column "user".password is '用户密码';


alter table "user" owner to postgres;

