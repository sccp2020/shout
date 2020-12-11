CREATE TABLE users (
    id varchar(255) PRIMARY KEY,
    screen_id varchar(255) NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    hash varchar(255) NOT NULL,
    biography varchar(255) NOT NULL,
    profile_image_url varchar(255),
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp on UPDATE current_timestamp
);

CREATE TABLE shouts (
    id varchar(255) PRIMARY KEY,
    user_id varchar(255) NOT NULL,
    parent_id varchar(255),
    content varchar(255),
    created_at DATETIME
);

CREATE TABLE users_shouts (
    id varchar(255) PRIMARY KEY,
    user_id varchar(255) NOT NULL,
    shouts_id varchar(255) NOT NULL,
    is_reshout boolean NOT NULL,
    created_at DATETIME
);

CREATE TABLE likes (
    id varchar(255) PRIMARY KEY,
    user_id varchar(255) NOT NULL,
    shouts_id varchar(255) NOT NULL,
    created_at DATETIME
);

CREATE TABLE shouts_images (
    id varchar(255) PRIMARY KEY,
    shouts_id varchar(255) NOT NULL,
    image_url varchar(255) NOT NULL,
    created_at DATETIME
);

CREATE TABLE follows (
    followee varchar(255),
    follower varchar(255),
    is_mutual boolean,
    created_at DATETIME
);
