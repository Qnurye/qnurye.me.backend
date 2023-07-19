CREATE TABLE tags (
    tagId SERIAL PRIMARY KEY,
    tagName VARCHAR(31) NOT NULL,
    parentId BIGINT UNSIGNED,
    createTime TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (parentId) REFERENCES tags(tagId)
);
