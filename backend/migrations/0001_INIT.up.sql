CREATE TABLE Users(
    id UUID PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    hashed_password BLOB NOT NULL,
    created_at DATETIME NOT NULL
);

CREATE TABLE Sessions(
    id UUID PRIMARY KEY,
    user_id UUID,
    created_at DATETIME NOT NULL,

    FOREIGN KEY(user_id) REFERENCES Users(id)
);