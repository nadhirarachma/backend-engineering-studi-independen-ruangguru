SELECT reports.id, fullname, class, status, study, score
FROM students
INNER JOIN reports
ON reports.student_id = students.id
WHERE score < 70
ORDER BY score ASc;