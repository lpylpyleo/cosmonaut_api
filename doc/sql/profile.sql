CREATE TABLE IF NOT EXISTS "profile" (
    id bigserial PRIMARY KEY NOT NULL UNIQUE,
    uid varchar(8) NOT NULL,
    avatar varchar(255) NULL,
    nickname varchar(20) NULL,
    gender char(1) NOT NULL DEFAULT 'f',
    motto varchar(255) NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    deleted_at timestamp NULL DEFAULT NULL,
    FOREIGN KEY (uid) REFERENCES public."user" (id)
);

COMMENT ON COLUMN "profile".id IS '主键ID';

COMMENT ON COLUMN "profile".uid IS '用户uuid';

COMMENT ON COLUMN "profile".avatar IS '头像url';

COMMENT ON COLUMN "profile".nickname IS '昵称';

COMMENT ON COLUMN "profile".gender IS '性别 m/f';

COMMENT ON COLUMN "profile".motto IS '签名';

ALTER TABLE "profile" OWNER TO postgres;

