ALTER TABLE user_models
ADD COLUMN IF NOT EXISTS "role" VARCHAR (32);

UPDATE user_models
SET "role" = 'USER'
WHERE "role" IS NULL;
