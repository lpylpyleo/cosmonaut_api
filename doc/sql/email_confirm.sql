CREATE TABLE IF NOT EXISTS "email_confirm" (
    id bigserial PRIMARY KEY NOT NULL UNIQUE,
    uid varchar(8) NOT NULL,
    email varchar(50) NULL,
    token varchar(127) NULL,
    confirmed_at timestamp NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    FOREIGN KEY (uid) REFERENCES public."user" (id)
);

COMMENT ON COLUMN "email_confirm".id IS '主键ID';

COMMENT ON COLUMN "email_confirm".uid IS '用户uuid';

COMMENT ON COLUMN "email_confirm".email IS '用户邮箱';

COMMENT ON COLUMN "email_confirm".token IS '一次性token';

COMMENT ON COLUMN "email_confirm".confirmed_at IS '确认时间';

ALTER TABLE "email_confirm" OWNER TO postgres;

