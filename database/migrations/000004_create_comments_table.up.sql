CREATE TABLE comments (
    commentId SERIAL PRIMARY KEY,
    parentId BIGINT UNSIGNED,
    userId BIGINT UNSIGNED NOT NULL,
    postId BIGINT UNSIGNED NOT NULL,
    content TEXT NOT NULL,
    updateTime TIMESTAMP ON UPDATE NOW(),
    createTime TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (parentId) REFERENCES comments(commentId),
    FOREIGN KEY (postId) REFERENCES posts(postId)
);
