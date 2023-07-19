CREATE TABLE posts_tags (
    postId BIGINT UNSIGNED NOT NULL,
    tagId BIGINT UNSIGNED NOT NULL,
    createTime TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (postId, tagId),
    FOREIGN KEY (postId) REFERENCES posts(postId),
    FOREIGN KEY (tagId) REFERENCES tags(tagId)
);
