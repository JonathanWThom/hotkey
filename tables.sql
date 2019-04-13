CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    prompt text NOT NULL,
    answer text NOT NULL
);

CREATE UNIQUE INDEX questions_pkey ON questions(id int4_ops);