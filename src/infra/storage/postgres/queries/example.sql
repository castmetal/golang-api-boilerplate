-- name: GetExampleByID :one
select * from example
where deleted_at is null
    and id = $1;

-- name: ListAllExamples :many
select * from example
where deleted_at is null
order by created_at desc
LIMIT $1
OFFSET $2;


-- name: CreateExample :exec
insert into example (
    "id",
    "name",
    "created_at",
    "updated_at",
    "deleted_at"
) values ($1, $2, $3, $4, $5);

-- name: GetExampleByName :one
select * from example
where 
    name = $1;
