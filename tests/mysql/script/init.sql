DROP DATABASE IF EXISTS aggregator_storage;
CREATE DATABASE IF NOT EXISTS aggregator_storage;
USE aggregator_storage;

CREATE TABLE tasks (
    task_index INT UNSIGNED AUTO_INCREMENT,
    program_id VARCHAR(256) NOT NULL,
    input TEXT NOT NULL,
    input_hash VARCHAR(256) NOT NULL,
    batch_index INT NOT NULL,
    PRIMARY KEY (task_index),
    UNIQUE KEY(program_id, input_hash, batch_index)
);

CREATE TABLE task_batches (
    batch_index INT UNSIGNED NOT NULL,
    block_number INT UNSIGNED NOT NULL,
    quroum_number TEXT NOT NULL,
    quroum_threshold_precentage TEXT NOT NULL,
    PRIMARY KEY (batch_index)
);

CREATE TABLE task_responses (
    task_index INT UNSIGNED NOT NULL,
    operator_id VARCHAR(256) NOT NULL,
    output_hash TEXT NOT NULL,
    result_cid TEXT NOT NULL,
    UNIQUE KEY (task_index, operator_id)
);