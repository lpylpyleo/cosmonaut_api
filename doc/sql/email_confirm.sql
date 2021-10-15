create table if not exists "email_confirm"
(
    id           bigserial primary key not null unique,
    uid          varchar(8)            not null,
    email        varchar(50)           null,
    token       varchar(127)          null,
    confirmed_at timestamp             null,
    created_at   timestamp default now(),
    updated_at   timestamp default now(),
    foreign key (uid) references public."user" (id)
);

comment on column "email_confirm".id is '主键ID';
comment on column "email_confirm".uid is '用户uuid';
comment on column "email_confirm".email is '用户邮箱';
comment on column "email_confirm".token is '一次性token';
comment on column "email_confirm".confirmed_at is '确认时间';

alter table "email_confirm"
    owner to postgres;