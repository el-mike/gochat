ALTER TABLE user_models
ADD "role" VARCHAR (255);

UPDATE user_models
SET "role" = 'User';
