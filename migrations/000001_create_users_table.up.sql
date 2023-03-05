CREATE TABLE users (
    id serial primary key,
    username varchar(255) not null unique,
    bio text,
    last_online timestamp,
    password_hash varchar(512) not null,
    public_key varchar(256) not null unique
);

CREATE TABLE tokens (
    id serial primary key,
    user_id int references users(id),
    token varchar(128) not null unique,
    expires_at date not null
);

CREATE TABLE conversations (
    id serial primary key,
    title varchar(255) not null
);

CREATE TABLE conversation_members (
    conv_id int references conversations(id),
    user_username varchar(255) references users(username),
    unique (conv_id, user_username)
);

CREATE TABLE messages (
    msg text not null,
    sender_id int references users(id),
    conv_id int references conversations(id),
    send_date date,
    send_time time
);