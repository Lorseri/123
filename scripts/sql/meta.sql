-- create database
CREATE DATABASE if not exists milvus_meta CHARACTER SET utf8mb4;

/*
 create tables script

 Notices:
    id, tenant_id, is_deleted, created_at, updated_at are 5 common columns for all collections.
 */

-- collections
CREATE TABLE if not exists milvus_meta.collections (
    id     BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    collection_id BIGINT NOT NULL,
    collection_name VARCHAR(128),
    description VARCHAR(2048) DEFAULT NULL,
    auto_id BOOL DEFAULT FALSE,
    shards_num INT,
    start_position VARCHAR(2048),
    consistency_level INT,
    ts BIGINT UNSIGNED DEFAULT 0,
    is_deleted BOOL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id),
    INDEX idx_collection_id_ts (collection_id, ts, is_deleted)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- collection aliases
CREATE TABLE if not exists milvus_meta.collection_aliases (
    id     BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    collection_id BIGINT NOT NULL,
    collection_alias VARCHAR(128),
    ts BIGINT UNSIGNED DEFAULT 0,
    is_deleted BOOL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id),
    INDEX idx_tenant_id_collection_id_ts_is_deleted (tenant_id, collection_id, ts, is_deleted)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- channels
CREATE TABLE if not exists milvus_meta.collection_channels (
    id     BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    collection_id BIGINT NOT NULL,
    virtual_channel_name VARCHAR(256) NOT NULL,
    physical_channel_name VARCHAR(256) NOT NULL,
    removed BOOL DEFAULT FALSE,
    ts BIGINT UNSIGNED DEFAULT 0,
    is_deleted BOOL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id),
    INDEX idx_collection_id_virtual_channel_name (collection_id, virtual_channel_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- fields
CREATE TABLE if not exists milvus_meta.field_schemas (
    id     BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    field_id BIGINT NOT NULL,
    field_name VARCHAR(128) NOT NULL,
    is_primary_key BOOL NOT NULL,
    description VARCHAR(2048) DEFAULT NULL,
    data_type INT UNSIGNED NOT NULL,
    type_params VARCHAR(2048),
    index_params VARCHAR(2048),
    auto_id BOOL NOT NULL,
    collection_id     BIGINT NOT NULL,
    ts BIGINT UNSIGNED DEFAULT 0,
    is_deleted BOOL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id),
    UNIQUE (collection_id, field_name, ts),
    INDEX idx_tenant_id_field_id_is_deleted (tenant_id, field_id, is_deleted)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- partitions
CREATE TABLE if not exists milvus_meta.`partitions` (
    id BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    partition_id     BIGINT NOT NULL,
    partition_name     VARCHAR(128),
    partition_created_timestamp bigint unsigned,
    collection_id     BIGINT NOT NULL,
    ts BIGINT UNSIGNED DEFAULT 0,
    is_deleted BOOL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id),
    UNIQUE (collection_id, partition_name, ts),
    INDEX idx_tenant_id_partition_id_is_deleted (tenant_id, partition_id, is_deleted)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- indexes
CREATE TABLE if not exists milvus_meta.`indexes` (
    id     BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    field_id BIGINT NOT NULL,
    collection_id BIGINT NOT NULL,
    index_id BIGINT NOT NULL,
    index_name VARCHAR(128),
    index_params VARCHAR(2048),
    is_deleted BOOL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id),
    UNIQUE (tenant_id, index_id),
    INDEX idx_collection_id_index_id (collection_id, index_id),
    INDEX idx_tenant_id_index_name_is_deleted (tenant_id, index_name, is_deleted)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- index file paths
CREATE TABLE if not exists milvus_meta.index_file_paths (
    id     BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    index_build_id BIGINT NOT NULL,
    index_file_path VARCHAR(256),
    is_deleted BOOL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- segments
CREATE TABLE if not exists milvus_meta.segments (
    id     BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    segment_id BIGINT NOT NULL,
    collection_id BIGINT NOT NULL,
    partition_id BIGINT NOT NULL,
    num_rows BIGINT NOT NULL,
    max_row_num INT COMMENT 'estimate max rows',
    dm_channel VARCHAR(128) NOT NULL,
    dml_position VARCHAR(2048) COMMENT 'checkpoint',
    start_position VARCHAR(2048),
    compaction_from VARCHAR(2048) COMMENT 'old segment IDs',
    created_by_compaction BOOL,
    segment_state TINYINT UNSIGNED NOT NULL,
    last_expire_time bigint unsigned COMMENT 'segment assignment expiration time',
    dropped_at bigint unsigned,
    is_deleted BOOL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- segment indexes
CREATE TABLE if not exists milvus_meta.segment_indexes (
    id     BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    collection_id BIGINT NOT NULL,
    partition_id BIGINT NOT NULL,
    segment_id BIGINT NOT NULL,
    field_id BIGINT NOT NULL,
    index_id BIGINT NOT NULL,
    index_build_id BIGINT,
    enable_index BOOL NOT NULL,
    index_file_paths VARCHAR(4096),
    index_size BIGINT UNSIGNED,
    `version` INT UNSIGNED,
    is_deleted BOOL DEFAULT FALSE COMMENT 'as mark_deleted',
    recycled BOOL DEFAULT FALSE COMMENT 'binlog files truly deleted',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id),
    INDEX idx_collection_id_segment_id (collection_id, segment_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- binlog files info
CREATE TABLE if not exists milvus_meta.binlogs (
    id     BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    field_id BIGINT NOT NULL,
    segment_id BIGINT NOT NULL,
    collection_id BIGINT NOT NULL,
    log_type SMALLINT UNSIGNED NOT NULL COMMENT 'binlog、stats binlog、delta binlog',
    num_entries BIGINT,
    timestamp_from BIGINT UNSIGNED,
    timestamp_to BIGINT UNSIGNED,
    log_path VARCHAR(256) NOT NULL,
    log_size BIGINT NOT NULL,
    is_deleted BOOL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id),
    INDEX idx_segment_id_log_type (segment_id, log_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- users
CREATE TABLE if not exists milvus_meta.credential_users (
    id     BIGINT NOT NULL AUTO_INCREMENT,
    tenant_id VARCHAR(128) DEFAULT NULL,
    username VARCHAR(128) NOT NULL,
    encrypted_password VARCHAR(256) NOT NULL,
    is_super BOOL NOT NULL DEFAULT false,
    is_deleted BOOL NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,
    PRIMARY KEY (id),
    UNIQUE (tenant_id, username, is_deleted)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
