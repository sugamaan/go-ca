CREATE
    DATABASE go_ca;

USE
    go_ca;

CREATE TABLE task
(
    task_id INT          NOT NULL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    reward  INT UNSIGNED
);

CREATE TABLE status_ms
(
    status_id INT          NOT NULL PRIMARY KEY,
    name      VARCHAR(255) NOT NULL
);

CREATE TABLE task_status
(
    task_id   INT NOT NULL,
    status_id INT NOT NULL,
    PRIMARY KEY (task_id, status_id),
    FOREIGN KEY (task_id) REFERENCES task (task_id),
    FOREIGN KEY (status_id) REFERENCES status_ms (status_id)
);

CREATE TABLE user
(
    user_id INT NOT NULL PRIMARY KEY
);

CREATE TABLE task_assignee
(
    user_id INT NOT NULL,
    task_id INT NOT NULL,
    PRIMARY KEY (user_id, task_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id),
    FOREIGN KEY (task_id) REFERENCES task (task_id)
);