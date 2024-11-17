CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);
CREATE TABLE todo_list (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL
);
CREATE TABLE users_lists (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    list_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_list (id) ON DELETE CASCADE
);
CREATE TABLE todo_item (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE TABLE list_items (
    id SERIAL PRIMARY KEY,
    list_id INTEGER,
    item_id INTEGER,
    FOREIGN KEY (list_id) REFERENCES todo_list (id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES todo_item (id) ON DELETE CASCADE
);