-- name: ResetTaskRun :exec
UPDATE job
SET last_run = NOW()
FROM (
        SELECT last_run,
            task_name
        FROM job AS o
        WHERE o.task_name = $1 
  		FOR UPDATE
    ) AS old_values
WHERE job.task_name = $1
RETURNING job.last_run AS this_run,
    old_values.last_run AS last_run;
