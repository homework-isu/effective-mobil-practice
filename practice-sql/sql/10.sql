DELETE FROM employees
WHERE id = (SELECT id FROM employees ORDER BY random() LIMIT 1);

DELETE FROM projects
WHERE id = (SELECT id FROM projects ORDER BY random() LIMIT 1);

DELETE FROM departments
WHERE id = (SELECT id FROM departments ORDER BY random() LIMIT 1);
