CREATE DATABASE hotkey;

/* \c hotkey */
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    prompt text NOT NULL UNIQUE,
    answer text NOT NULL
);

CREATE DATABASE hotkey_test;

/* \c hotkey_test */
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    prompt text NOT NULL UNIQUE,
    answer text NOT NULL
);


