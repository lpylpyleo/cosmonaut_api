CREATE TABLE IF NOT EXISTS post (
    id bigserial PRIMARY KEY NOT NULL,
    creator varchar(8) NOT NULL,
    title varchar(50) NULL,
    content text NULL,
    comment_count integer DEFAULT 0 NOT NULL,
    like_count integer DEFAULT 0 NOT NULL,
    is_public bool DEFAULT TRUE NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    deleted_at timestamp NULL DEFAULT NULL,
    FOREIGN KEY (creator) REFERENCES public."user" (id)
);

COMMENT ON COLUMN post.id IS '主键ID';

COMMENT ON COLUMN post.creator IS '用户uuid';

COMMENT ON COLUMN post.title IS '标题';

COMMENT ON COLUMN post.content IS '正文';

COMMENT ON COLUMN post.comment_count IS '评论数';

COMMENT ON COLUMN post.like_count IS '点赞数';

ALTER TABLE post OWNER TO postgres;

