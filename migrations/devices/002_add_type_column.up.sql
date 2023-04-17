ALTER TABLE IF EXISTS gotor.devices
    ADD COLUMN device_type VARCHAR(255) NOT NULL DEFAULT 'unknown';