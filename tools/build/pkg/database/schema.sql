-- This file generated by protoc-ddl.
-- Generated by MySQL Generator (v0.1)

SET foreign_key_checks=0;

DROP TABLE IF EXISTS `source_repository`;
CREATE TABLE `source_repository` (
	`id` INTEGER NOT NULL AUTO_INCREMENT,
	`url` VARCHAR(255) NOT NULL,
	`clone_url` VARCHAR(255) NOT NULL,
	`name` VARCHAR(100) NOT NULL,
	`private` TINYINT(1) NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NULL,
	PRIMARY KEY(`id`)
) Engine=InnoDB;

DROP TABLE IF EXISTS `job`;
CREATE TABLE `job` (
	`id` INTEGER NOT NULL AUTO_INCREMENT,
	`repository_id` INTEGER NOT NULL,
	`command` VARCHAR(20) NOT NULL,
	`target` VARCHAR(255) NOT NULL,
	`active` TINYINT(1) NOT NULL,
	`all_revision` TINYINT(1) NOT NULL,
	`github_status` TINYINT(1) NOT NULL,
	`cpu_limit` VARCHAR(255) NOT NULL,
	`memory_limit` VARCHAR(255) NOT NULL,
	`exclusive` TINYINT(1) NOT NULL,
	`sync` TINYINT(1) NOT NULL,
	`config_name` VARCHAR(255) NOT NULL,
	`bazel_version` VARCHAR(255) NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NULL,
	PRIMARY KEY(`id`)
) Engine=InnoDB;

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
	`id` INTEGER NOT NULL AUTO_INCREMENT,
	`job_id` INTEGER NOT NULL,
	`revision` VARCHAR(255) NOT NULL,
	`success` TINYINT(1) NOT NULL,
	`log_file` VARCHAR(255) NOT NULL,
	`command` VARCHAR(255) NOT NULL,
	`target` VARCHAR(255) NOT NULL,
	`via` VARCHAR(255) NOT NULL,
	`config_name` VARCHAR(255) NOT NULL,
	`start_at` DATETIME NULL,
	`finished_at` DATETIME NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NULL,
	PRIMARY KEY(`id`)
) Engine=InnoDB;

DROP TABLE IF EXISTS `trusted_user`;
CREATE TABLE `trusted_user` (
	`id` INTEGER NOT NULL AUTO_INCREMENT,
	`github_id` BIGINT NOT NULL,
	`username` VARCHAR(255) NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NULL,
	PRIMARY KEY(`id`)
) Engine=InnoDB;

DROP TABLE IF EXISTS `permit_pull_request`;
CREATE TABLE `permit_pull_request` (
	`id` INTEGER NOT NULL AUTO_INCREMENT,
	`repository` VARCHAR(255) NOT NULL,
	`number` INTEGER NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NULL,
	PRIMARY KEY(`id`)
) Engine=InnoDB;

SET foreign_key_checks=1;
