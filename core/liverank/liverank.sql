-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- 主机： mysql
-- 生成日期： 2022-07-25 06:45:31
-- 服务器版本： 5.7.38
-- PHP 版本： 8.0.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `data`
--

-- --------------------------------------------------------

--
-- 表的结构 `bili_ranks`
--

CREATE TABLE `bili_ranks` (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT,
                              `timestamp` bigint(20) DEFAULT '0',
                              `buid` bigint(20) DEFAULT '0',
                              `room_id` bigint(20) DEFAULT '0',
                              `rank` bigint(20) DEFAULT '0',
                              `gift_value` bigint(20) DEFAULT '0',
                              `is_concern` tinyint(1) DEFAULT '0',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- 转储表的索引
--

--
-- 表的索引 `bili_ranks`
--
ALTER TABLE `bili_ranks`
    ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `bili_ranks`
--
ALTER TABLE `bili_ranks`
    MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
