BEGIN;

DROP INDEX IF EXISTS idx_lightwell_advisory_notifications_repo_config_advisory_org;

ALTER TABLE lightwell_advisory_notifications
    DROP COLUMN IF EXISTS org_id;

CREATE UNIQUE INDEX IF NOT EXISTS idx_lightwell_advisory_notifications_repo_config_advisory
    ON lightwell_advisory_notifications (repository_configuration_uuid, advisory_id, package_name);

COMMIT;
