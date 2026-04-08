INSERT INTO posts (user_id, title, body)
SELECT u.id, 'Primo post', 'Contenuto primo post'
FROM users u
WHERE u.email = 'gabriele@mail.com';

INSERT INTO posts (user_id, title, body)
SELECT u.id, 'Secondo post', 'Contenuto secondo post'
FROM users u
WHERE u.email = 'elena@mail.com';