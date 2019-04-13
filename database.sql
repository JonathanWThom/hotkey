CREATE DATABASE hotkey;

CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    prompt text NOT NULL,
    answer text NOT NULL
);

