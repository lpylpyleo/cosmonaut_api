create table if not exists poster
(
    id         bigserial primary key not null,
    poster     varchar(8)            not null,
    title      varchar(50)           null,
    content    text                  null,
    deleted    bool                       default false not null,
    created_at timestamp                  default now(),
    updated_at timestamp                  default now(),
    deleted_at timestamp             null default null,
    foreign key (poster) references public."user" (id)
);

comment on column poster.id is '主键ID';
comment on column poster.poster is '用户uuid';
comment on column poster.title is '标题';
comment on column poster.content is '正文';

alter table poster
    owner to postgres;