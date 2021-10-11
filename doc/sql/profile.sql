create table "profile"
(
    id bigserial primary key,
    uid uuid null,
    avatar varchar(127) null,
    nickname varchar(20) null,
    create_at timestamptz,
    update_at timestamptz,
    foreign key (uid) references public."user"(id)
);

comment on column "profile".id is '主键ID';
comment on column "profile".uid is '用户uuid';
comment on column "profile".avatar is '头像url';
comment on column "profile".avatar is '昵称';

alter table "profile" owner to postgres;

