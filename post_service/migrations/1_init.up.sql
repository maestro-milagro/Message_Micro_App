CREATE TABLE IF NOT EXISTS users_posts
(
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    bucket TEXT NOT NULL,
    key TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS subscriptions
(
    uid    integer,
    sub_id integer,
    PRIMARY KEY (uid, sub_id)
);

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    pass_hash bytea NOT NULL,
);

CREATE TABLE IF NOT EXISTS posts
(
    post_id integer PRIMARY KEY,
    email TEXT NOT NULL,
);

