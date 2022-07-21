-- bili_ranks definition

CREATE TABLE `bili_ranks` (`id` integer,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`buid` integer DEFAULT 0,`room_id` integer DEFAULT 0,`rank` integer DEFAULT 0,`gift_value` integer DEFAULT 0,`is_concern` numeric DEFAULT false,PRIMARY KEY (`id`));

CREATE INDEX `idx_bili_ranks_deleted_at` ON `bili_ranks`(`deleted_at`);