CREATE TABLE IF NOT EXISTS api_infos (
    version VARCHAR(8) NOT NULL,
    release_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
