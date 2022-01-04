CREATE TABLE IF NOT EXISTS user_post_like (
    id bigserial PRIMARY KEY NOT NULL,
    uid varchar(8) NOT NULL,
    pid bigint NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    deleted_at timestamp NULL DEFAULT NULL,
    FOREIGN KEY (uid) REFERENCES public."user" (id),
    FOREIGN KEY (pid) REFERENCES public."post" (id)
);

COMMENT ON COLUMN user_post_like.id IS '主键ID';

COMMENT ON COLUMN user_post_like.uid IS '用户id';

COMMENT ON COLUMN user_post_like.pid IS 'post id';

ALTER TABLE user_post_like OWNER TO postgres;
