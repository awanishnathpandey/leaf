-- name: GetDashboardKPICount :one
SELECT
  (SELECT COUNT(*) FROM users) AS users,
  (SELECT COUNT(*) FROM roles) AS roles,
  (SELECT COUNT(*) FROM permissions) AS permissions,
  (SELECT COUNT(*) FROM groups) AS groups,
  (SELECT COUNT(*) FROM folders) AS folders,
  (SELECT COUNT(*) FROM files) AS files;