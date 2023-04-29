-- taskテーブルに10件のテストデータを挿入
INSERT INTO task (task_id, name, reward)
VALUES (1, 'task1', 100),
       (2, 'task2', 200),
       (3, 'task3', 300),
       (4, 'task4', 400),
       (5, 'task5', 500),
       (6, 'task6', 600),
       (7, 'task7', 700),
       (8, 'task8', 800),
       (9, 'task9', 900),
       (10, 'task10', 1000);

-- status_msテーブルに10件のテストデータを挿入
INSERT INTO status_ms (status_id, name)
VALUES (1, '未完了'),
       (2, '作業中'),
       (3, '確認中'),
       (4, '完了');

-- task_statusテーブルに10件のテストデータを挿入
INSERT INTO task_status (task_id, status_id)
VALUES (1, 1),
       (2, 2),
       (3, 3),
       (4, 4),
       (5, 1),
       (6, 2),
       (7, 3),
       (8, 4),
       (9, 1),
       (10, 2);

-- userテーブルに10件のテストデータを挿入
INSERT INTO user (user_id)
VALUES (1),
       (2),
       (3),
       (4),
       (5),
       (6),
       (7),
       (8),
       (9),
       (10);

-- task_assigneeテーブルに10件のテストデータを挿入
INSERT INTO task_assignee (user_id, task_id)
VALUES (1, 1),
       (2, 2),
       (3, 3),
       (4, 4),
       (5, 5),
       (6, 6),
       (7, 7),
       (8, 8),
       (9, 9),
       (10, 10);