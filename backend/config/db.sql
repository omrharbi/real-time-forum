-- database: :memory:

CREATE TABLE user(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nickname text   not NULL,
    firstname text   not NULL,
    lastname text not NULL,
    Age INTEGER not NULL,
    gender text not NULL,
    email text not NULL UNIQUE,
    password text not NULL,
    expires DATETIME,
    createAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    status text not NULL,
    UUID text
);

CREATE TABLE post (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    card_id INTEGER,
    FOREIGN KEY (card_id) REFERENCES card(id)
);

CREATE TABLE category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name text not NULL
);


CREATE TABLE post_category (
    post_id INTEGER,
    category_id INTEGER,
    PRIMARY KEY (post_id, category_id),
    FOREIGN KEY (post_id) REFERENCES post(id),
    FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE card (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER ,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE comment (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    card_id INTEGER,
    target_id INTEGER,
    FOREIGN KEY (card_id) REFERENCES card(id),
    FOREIGN KEY (target_id) REFERENCES card(id)
);

CREATE TABLE likes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    card_id INTEGER NOT NULL,
    is_like BOOLEAN NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id), 
    FOREIGN KEY (card_id) REFERENCES card(id),
    -- Ensure is_like is always either TRUE or FALSE
    CONSTRAINT check_is_like CHECK (
        is_like IN (TRUE, FALSE)
    ),
    -- Prevent duplicate likes or dislikes by the same user on the same card
    CONSTRAINT unique_like_per_user_card UNIQUE 
    (user_id, card_id)
);


INSERT INTO category (name) VALUES ('General');
INSERT INTO category (name) VALUES ('Technology');
INSERT INTO category (name) VALUES ('Sports');
INSERT INTO category (name) VALUES ('Entertainment');
INSERT INTO category (name) VALUES ('Science');
INSERT INTO category (name) VALUES ('Food');
INSERT INTO category (name) VALUES ('Travel');
INSERT INTO category (name) VALUES ('Fashion');
INSERT INTO category (name) VALUES ('Art');
INSERT INTO category (name) VALUES ('Music');