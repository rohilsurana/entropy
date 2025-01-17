CREATE TABLE IF NOT EXISTS resources
(
    id                BIGSERIAL NOT NULL PRIMARY KEY,
    urn               TEXT      NOT NULL UNIQUE,
    kind              TEXT      NOT NULL,
    name              TEXT      NOT NULL,
    project           TEXT      NOT NULL,
    created_at        timestamp NOT NULL DEFAULT current_timestamp,
    updated_at        timestamp NOT NULL DEFAULT current_timestamp,
    spec_configs      bytea     NOT NULL,
    state_status      TEXT      NOT NULL,
    state_output      bytea     NOT NULL,
    state_module_data bytea     NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_resources_kind ON resources (kind);
CREATE INDEX IF NOT EXISTS idx_resources_name ON resources (name);
CREATE INDEX IF NOT EXISTS idx_resources_project ON resources (project);
CREATE INDEX IF NOT EXISTS idx_resources_state_status ON resources (state_status);

CREATE TABLE IF NOT EXISTS resource_dependencies
(
    resource_id    BIGINT NOT NULL REFERENCES resources (id),
    dependency_key TEXT   NOT NULL,
    depends_on     BIGINT NOT NULL REFERENCES resources (id),
    UNIQUE (resource_id, dependency_key)
);

CREATE TABLE IF NOT EXISTS resource_tags
(
    resource_id BIGINT NOT NULL REFERENCES resources (id),
    tag         TEXT   NOT NULL,
    UNIQUE (resource_id, tag)
);
CREATE INDEX IF NOT EXISTS idx_resource_tags_resource_id ON resource_tags (resource_id);
CREATE INDEX IF NOT EXISTS idx_resource_tags_tag ON resource_tags (tag);