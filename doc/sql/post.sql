create table if not exists post
(
    id         bigserial primary key not null,
    creator    varchar(8)           not null,
    title      varchar(50)           null,
    content    text                  null,
    is_public  bool                       default true not null,
    created_at timestamp                  default now(),
    updated_at timestamp                  default now(),
    deleted_at timestamp             null default null,
    foreign key (creator) references public."user" (id)
);

comment on column post.id is '主键ID';
comment on column post.creator is '用户uuid';
comment on column post.title is '标题';
comment on column post.content is '正文';

alter table post owner to postgres;