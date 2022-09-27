-- name: GetStudentByNumber :one
SELECT `id`, `name`, `ruby`, `student_number`, `date_of_birth`, `address`, `expiration_date`, `created_at`, `updated_at` FROM `students` WHERE `student_number`=?;

-- name: GetStudentByID :one
SELECT `id`, `name`, `ruby`, `student_number`, `date_of_birth`, `address`, `expiration_date`, `created_at`, `updated_at` FROM `students` WHERE `id`=?;

-- name: SetStudent :execresult
INSERT INTO `students` (`name`, `ruby`, `student_number`, `date_of_birth`, `address`, `expiration_date`) VALUES(?, ?, ?, ?, ?, ?);

-- name: GetAllStudents :many
SELECT `id`, `name`, `ruby`, `student_number`, `date_of_birth`, `address`, `expiration_date`, `created_at`, `updated_at` FROM `students` ORDER BY `student_number` ASC;
