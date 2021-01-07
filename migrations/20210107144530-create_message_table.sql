-- +migrate Up
CREATE TABLE IF NOT EXISTS messages (
    id      INTEGER      AUTO_INCREMENT NOT NULL,
    name    TEXT                        NOT NULL,
    message TEXT                        NOT NULL,
    PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE messages;
