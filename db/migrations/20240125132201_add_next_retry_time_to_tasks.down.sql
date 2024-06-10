BEGIN;

DROP VIEW IF EXISTS ready_tasks;

ALTER TABLE tasks DROP COLUMN next_retry_time;

CREATE OR REPLACE VIEW ready_tasks AS
SELECT *
FROM tasks
WHERE started_at IS NULL
  AND (status != 'canceled' OR status is null)
  AND id NOT IN (
    SELECT task_id
    FROM task_dependencies JOIN tasks ON dependency_id = id
    WHERE finished_at IS NULL
)
ORDER BY queued_at ASC;

COMMIT;
