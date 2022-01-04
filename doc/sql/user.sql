CREATE TABLE IF NOT EXISTS "user" (
    id varchar(8) NOT NULL UNIQUE PRIMARY KEY,
    account varchar(20) NULL,
    email varchar(50) NULL,
    phone varchar(20) NULL,
    role varchar(10) NULL,
    disabled bool NOT NULL DEFAULT FALSE,
    password varchar(64) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    deleted_at timestamp NULL DEFAULT NULL
);

COMMENT ON COLUMN "user".id IS '用户ID';

COMMENT ON COLUMN "user".account IS '用户账号';

COMMENT ON COLUMN "user".email IS '用户邮箱';

COMMENT ON COLUMN "user".password IS '用户密码';

ALTER TABLE "user" OWNER TO postgres;

