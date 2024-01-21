DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    published_at BIGINT NOT NULL,
    link TEXT NOT NULL
);

INSERT INTO posts (id, title, content, published_at, link) VALUES (0, 'Статья', 'Содержание статьи', 0, 'https://');