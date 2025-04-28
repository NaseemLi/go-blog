-- MySQL dump 10.13  Distrib 5.7.44, for Linux (x86_64)
--
-- Host: localhost    Database: blog
-- ------------------------------------------------------
-- Server version	5.7.44-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

use blog;
--
-- Table structure for table `article_digg_models`
--

DROP TABLE IF EXISTS `article_digg_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `article_digg_models` (
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `article_id` bigint(20) unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`user_id`,`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_digg_models`
--

LOCK TABLES `article_digg_models` WRITE;
/*!40000 ALTER TABLE `article_digg_models` DISABLE KEYS */;
INSERT INTO `article_digg_models` VALUES (1,7,'2025-04-20 18:10:51.539',1,NULL);
/*!40000 ALTER TABLE `article_digg_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article_models`
--

DROP TABLE IF EXISTS `article_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `article_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL,
  `abstract` varchar(256) DEFAULT NULL,
  `content` longtext,
  `category_id` bigint(20) unsigned DEFAULT NULL,
  `tag_list` longtext,
  `cover` varchar(256) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `look_count` bigint(20) DEFAULT NULL,
  `digg_count` bigint(20) DEFAULT NULL,
  `comment_count` bigint(20) DEFAULT NULL,
  `collect_count` bigint(20) DEFAULT NULL,
  `open_comment` tinyint(1) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_models`
--

LOCK TABLES `article_models` WRITE;
/*!40000 ALTER TABLE `article_models` DISABLE KEYS */;
INSERT INTO `article_models` VALUES (7,'2025-04-15 21:41:14.870','2025-04-15 21:41:14.870','5.go环境搭建','hello 你好\n','hello 你好',2,'后端','',1,0,0,0,0,1,3),(8,'2025-04-15 22:10:02.270','2025-04-15 22:11:41.015','.go环境搭建','你好','你好',2,'go,java','xxx',1,0,0,0,0,1,3),(9,'2025-04-16 22:10:02.270','2025-04-16 22:11:41.015','文章搜索测试','我在这','here',2,'vue3','xxx',1,1,2,2,3,3,3),(10,'2025-04-23 23:04:02.000','2025-04-23 23:04:02.000','如何用 Go 构建微服务','介绍使用 Go 构建微服务架构的基本方法','使用 Gin + gRPC 构建高并发服务',1,'go,grpc,microservice','cover1.jpg',1,123,12,3,5,1,3),(11,'2025-04-23 23:04:02.000','2025-04-23 23:04:02.000','深入理解 Goroutine 调度机制','Goroutine 是 Go 的核心并发特性','Go 使用 M:N 的调度模型实现轻量线程',2,'go,goroutine,scheduler','cover2.jpg',2,88,9,1,2,1,3),(12,'2025-04-23 23:04:02.000','2025-04-23 23:04:02.000','Go 与 Java 协程对比分析','从开发者角度对比两种并发模型','Java 使用线程池，Go 使用 goroutine',3,'go,java,concurrency','cover3.jpg',1,144,18,4,6,1,3),(13,'2025-04-23 23:04:02.000','2025-04-23 23:04:02.000','构建一个博客系统的后端 API','一步步教你用 Go 写博客后端','包括文章、评论、用户模块',1,'go,restful,backend','cover4.jpg',3,256,21,8,11,1,3),(14,'2025-04-23 23:04:02.000','2025-04-23 23:04:02.000','Go 语言中的错误处理哲学','探索 Go 的 error 设计哲学','Go 不抛异常，而是显式返回 error',2,'go,error,design','cover5.jpg',2,72,5,0,1,1,3),(15,'2025-04-25 14:43:08.051','2025-04-25 14:43:08.051','四月二十四日','这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的','这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文这是我的正文',NULL,'go,后端','',1,0,0,0,0,1,3),(17,'2025-04-25 14:45:48.105','2025-04-25 14:45:48.105','四月二十四日新','四月二十四日笔记\n\n这是我的正文，用来测试 Markdown 渲染。\n\n一、今日目标\n\n\n学习 Golang 接口定义\n阅读 Gin 框架文档\n写一个 POST 请求的测试用例\n\n\n二、代码片段\n\np','# 四月二十四日笔记\n\n这是我的正文，用来测试 **Markdown 渲染**。\n\n## 一、今日目标\n\n- 学习 Golang 接口定义\n- 阅读 Gin 框架文档\n- 写一个 POST 请求的测试用例\n\n## 二、代码片段\n\n```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, Markdown in Go!\")\n}\n```\n\n## 三、注意事项\n\n> 在接口设计中，要特别注意请求体的结构，字段不要命名歧义。\n\n## 四、总结\n\n写 Markdown 的时候：\n\n1. 注意标题层级要清晰\n2. 内容语义要准确\n3. 保持格式统一，比如代码块、引用、列表等\n\n---\n\n以上就是今天的记录！\n',NULL,'go,后端','',1,0,0,0,0,1,3),(21,'2025-04-25 15:33:01.685','2025-04-25 15:33:49.045','四月二十四日最新','你好111','你好111',3,'go,后端','xxx',1,0,0,0,0,1,3),(23,'2025-04-25 16:51:40.929','2025-04-25 16:51:49.582','接口设计实战技巧','接口设计实战技巧分享\n\n在实际开发中，优雅的接口设计是后端开发的核心能力。\n本文总结了我在构建 API 时的一些思考和实战经验，供大家参考。\n\n\n\n? 核心原则\n\n1. 一致性优先\n路径风格、请求方式','# 接口设计实战技巧分享\n\n在实际开发中，优雅的接口设计是后端开发的核心能力。\n本文总结了我在构建 API 时的一些思考和实战经验，供大家参考。\n\n---\n\n## ? 核心原则\n\n**1. 一致性优先**  \n路径风格、请求方式、字段命名都要保持统一。\n\n**2. 显式比隐式好**  \n不要让客户端猜测字段意义，能表达清楚就不要省字。\n\n**3. 面向资源设计**  \n以资源为中心，而不是动作。例如：\n\n```\n✅ /articles/123\n❌ /getArticleById\n```\n\n---\n\n## ? 接口结构示例\n\n```json\n{\n  \"code\": 0,\n  \"message\": \"OK\",\n  \"data\": {\n    \"id\": 123,\n    \"title\": \"Golang 实战入门\",\n    \"author\": \"Naseem\",\n    \"tags\": [\"Go\", \"REST\"]\n  }\n}\n```\n\n---\n\n## ? 常见坑位\n\n> 字段歧义会导致客户端开发效率极低，例如 `status` 字段到底是发布状态、审核状态还是登录状态？\n\n建议使用更具体的命名，比如：`publishStatus`, `auditStatus`, `loginStatus`。\n\n---\n\n## ? 工具推荐\n\n- Postman / Hoppscotch：接口测试神器\n- Swagger / Redoc：自动生成 API 文档\n- go-playground/validator：字段验证利器\n\n---\n\n## ✅ 结语\n\n> 优秀的接口设计 = 技术审美 + 沟通能力 + 实践积累。\n\n保持学习，保持输出！?\n',NULL,'接口设计,后端,go','',1,0,0,0,0,1,3),(24,'2025-04-25 16:54:47.867','2025-04-25 16:54:47.867','接口设计实战技巧12312312','四月二十四日笔记这是我的正文，用来测试 Markdown 渲染。## 一、今日目标- 学习 Golang 接口定义- 阅读 Gin 框架文档- 写一个 POST 请求的测试用例## 二、代码片段”`g','# 四月二十四日笔记这是我的正文，用来测试 **Markdown 渲染**。## 一、今日目标- 学习 Golang 接口定义- 阅读 Gin 框架文档- 写一个 POST 请求的测试用例## 二、代码片段```gopackage mainimport Hello, Markdown in Go!',NULL,'接口设计,后端,go','',1,0,0,0,0,1,3);
/*!40000 ALTER TABLE `article_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `banner_models`
--

DROP TABLE IF EXISTS `banner_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `banner_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `show` tinyint(1) DEFAULT NULL,
  `cover` longtext,
  `href` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `banner_models`
--

LOCK TABLES `banner_models` WRITE;
/*!40000 ALTER TABLE `banner_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `banner_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category_models`
--

DROP TABLE IF EXISTS `category_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category_models`
--

LOCK TABLES `category_models` WRITE;
/*!40000 ALTER TABLE `category_models` DISABLE KEYS */;
INSERT INTO `category_models` VALUES (3,'2025-04-17 19:36:47.018','2025-04-17 21:13:37.954','新的分类名称',1),(4,'2025-04-17 19:36:56.076','2025-04-17 19:36:56.076','golang',1),(5,'2025-04-17 19:37:00.950','2025-04-17 19:37:00.950','java',1);
/*!40000 ALTER TABLE `category_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `chat_models`
--

DROP TABLE IF EXISTS `chat_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chat_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `send_user_id` bigint(20) unsigned DEFAULT NULL,
  `rev_user_id` bigint(20) unsigned DEFAULT NULL,
  `msg_type` tinyint(4) DEFAULT NULL,
  `msg` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chat_models`
--

LOCK TABLES `chat_models` WRITE;
/*!40000 ALTER TABLE `chat_models` DISABLE KEYS */;
INSERT INTO `chat_models` VALUES (13,'2025-04-23 16:21:23.544','2025-04-23 16:21:23.544',1,2,1,'{\"textMsg\":{\"content\":\"nihao\"}}'),(14,'2025-04-23 16:39:03.985','2025-04-23 16:39:03.985',2,1,1,'{\"textMsg\":{\"content\":\"metoo\"}}'),(15,'2025-04-23 16:39:07.022','2025-04-23 16:39:07.022',2,1,1,'{\"textMsg\":{\"content\":\"metoo\"}}'),(16,'2025-04-23 17:02:43.579','2025-04-23 17:02:43.579',2,1,1,'{\"textMsg\":{\"content\":\"metoo\"}}'),(17,'2025-04-23 17:02:44.593','2025-04-23 17:02:44.593',2,1,1,'{\"textMsg\":{\"content\":\"metoo\"}}');
/*!40000 ALTER TABLE `chat_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `collect_models`
--

DROP TABLE IF EXISTS `collect_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `collect_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL,
  `abstract` varchar(256) DEFAULT NULL,
  `cover` varchar(256) DEFAULT NULL,
  `article_count` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `is_default` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `collect_models`
--

LOCK TABLES `collect_models` WRITE;
/*!40000 ALTER TABLE `collect_models` DISABLE KEYS */;
INSERT INTO `collect_models` VALUES (2,NULL,NULL,'test','学习',NULL,NULL,2,0),(3,'2025-04-17 21:09:31.808','2025-04-19 00:04:34.036','study1','学习2','',NULL,1,0),(4,'2025-04-19 00:00:45.635','2025-04-20 18:38:28.983','默认收藏夹','','',NULL,1,1);
/*!40000 ALTER TABLE `collect_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment_digg_models`
--

DROP TABLE IF EXISTS `comment_digg_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment_digg_models` (
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `comment_id` bigint(20) unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`user_id`,`comment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment_digg_models`
--

LOCK TABLES `comment_digg_models` WRITE;
/*!40000 ALTER TABLE `comment_digg_models` DISABLE KEYS */;
INSERT INTO `comment_digg_models` VALUES (1,16,'2025-04-18 22:46:56.329',1,NULL),(1,8,'2025-04-20 17:55:40.813',2,NULL),(1,7,'2025-04-20 18:19:32.866',3,NULL),(1,9,'2025-04-20 20:37:24.011',4,'2025-04-20 20:37:24.011'),(1,10,'2025-04-20 20:38:50.486',5,'2025-04-20 20:38:50.486');
/*!40000 ALTER TABLE `comment_digg_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment_models`
--

DROP TABLE IF EXISTS `comment_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `content` varchar(256) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `article_id` bigint(20) unsigned DEFAULT NULL,
  `parent_id` bigint(20) unsigned DEFAULT NULL,
  `root_parent_id` bigint(20) unsigned DEFAULT NULL,
  `digg_count` bigint(20) DEFAULT NULL,
  `apply_count` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment_models`
--

LOCK TABLES `comment_models` WRITE;
/*!40000 ALTER TABLE `comment_models` DISABLE KEYS */;
INSERT INTO `comment_models` VALUES (1,'2025-04-17 23:12:06.700','2025-04-17 23:12:06.700','我是第1条评论',1,7,NULL,NULL,0,NULL),(2,'2025-04-17 23:20:13.957','2025-04-17 23:20:13.957','我是1的子评论',1,7,1,1,0,NULL),(3,'2025-04-17 23:21:16.619','2025-04-17 23:21:16.619','我是2的子评论',1,7,2,1,0,NULL),(7,'2025-04-18 18:44:37.496','2025-04-18 18:44:37.496','你好',1,8,NULL,NULL,0,NULL),(8,'2025-04-18 18:45:17.289','2025-04-18 18:45:17.289','你好2',1,8,1,1,0,NULL),(9,'2025-04-18 18:45:25.102','2025-04-18 18:45:25.102','你好3',1,8,2,1,0,NULL),(10,'2025-04-18 18:48:05.896','2025-04-18 18:48:05.896','你好4',1,8,3,1,0,NULL),(14,'2025-04-18 19:31:43.276','2025-04-18 19:31:43.276','你好8',1,8,2,1,0,NULL),(15,'2025-04-18 20:47:11.659','2025-04-18 20:47:11.659','你好8',1,8,2,1,0,0),(16,'2025-04-18 20:51:11.222','2025-04-18 20:51:11.222','你好8',1,8,8,1,0,0),(17,'2025-04-20 16:57:32.938','2025-04-20 16:57:32.938','根评论',1,8,NULL,NULL,0,0),(18,'2025-04-20 17:39:00.173','2025-04-20 17:39:00.173','18的子评论',1,8,17,17,0,0),(19,'2025-04-20 17:42:06.359','2025-04-20 17:42:06.359','18的子评论',1,8,17,17,0,0),(20,'2025-04-20 17:42:33.704','2025-04-20 17:42:33.704','18的子评论',1,8,17,17,0,0);
/*!40000 ALTER TABLE `comment_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `global_notification_models`
--

DROP TABLE IF EXISTS `global_notification_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `global_notification_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL,
  `icon` varchar(256) DEFAULT NULL,
  `content` varchar(64) DEFAULT NULL,
  `href` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `global_notification_models`
--

LOCK TABLES `global_notification_models` WRITE;
/*!40000 ALTER TABLE `global_notification_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `global_notification_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `image_models`
--

DROP TABLE IF EXISTS `image_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `image_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `filename` varchar(64) DEFAULT NULL,
  `path` varchar(256) DEFAULT NULL,
  `size` bigint(20) DEFAULT NULL,
  `hash` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `image_models`
--

LOCK TABLES `image_models` WRITE;
/*!40000 ALTER TABLE `image_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `image_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_models`
--

DROP TABLE IF EXISTS `log_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `log_type` tinyint(4) DEFAULT NULL,
  `title` varchar(64) DEFAULT NULL,
  `content` longtext,
  `level` tinyint(4) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `ip` varchar(32) DEFAULT NULL,
  `addr` varchar(64) DEFAULT NULL,
  `is_read` tinyint(1) DEFAULT NULL,
  `login_status` tinyint(1) DEFAULT NULL,
  `username` varchar(32) DEFAULT NULL,
  `pwd` varchar(32) DEFAULT NULL,
  `login_type` tinyint(4) DEFAULT NULL,
  `service_name` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_models`
--

LOCK TABLES `log_models` WRITE;
/*!40000 ALTER TABLE `log_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `log_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message_models`
--

DROP TABLE IF EXISTS `message_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `message_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL,
  `rev_user_id` bigint(20) unsigned DEFAULT NULL,
  `action_user_id` bigint(20) unsigned DEFAULT NULL,
  `action_user_nickname` longtext,
  `action_user_avatar` longtext,
  `title` longtext,
  `content` longtext,
  `article_id` bigint(20) unsigned DEFAULT NULL,
  `article_title` longtext,
  `comment_id` bigint(20) unsigned DEFAULT NULL,
  `linktitle` longtext,
  `link_href` longtext,
  `is_read` tinyint(1) DEFAULT NULL,
  `link_title` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message_models`
--

LOCK TABLES `message_models` WRITE;
/*!40000 ALTER TABLE `message_models` DISABLE KEYS */;
INSERT INTO `message_models` VALUES (1,'2025-04-20 16:57:33.720','2025-04-20 16:57:33.720',1,1,1,'admin','','','根评论',8,'.go环境搭建',17,'','',0,NULL),(5,'2025-04-20 17:42:34.479','2025-04-20 17:42:34.479',1,1,1,'admin','','','18的子评论',8,'.go环境搭建',20,'','',0,NULL),(6,'2025-04-20 17:42:34.810','2025-04-20 17:42:34.810',1,17,1,'admin','','','18的子评论',8,'.go环境搭建',20,'','',0,NULL),(8,'2025-04-20 18:19:34.160','2025-04-20 18:19:34.160',5,1,1,'admin','','','',8,'.go环境搭建',0,'','',0,NULL),(9,'2025-04-20 18:38:29.570','2025-04-20 18:38:29.570',7,1,1,'admin','','','',7,'5.go环境搭建',0,'','',0,NULL),(10,'2025-04-20 20:37:25.932','2025-04-20 20:37:25.932',5,1,1,'admin','','','你好3',8,'.go环境搭建',0,NULL,'',0,''),(11,'2025-04-20 20:38:51.736','2025-04-20 23:43:54.746',5,1,1,'admin','','','你好4',8,'.go环境搭建',0,NULL,'',1,''),(12,'2025-04-25 15:16:23.992','2025-04-25 15:16:23.992',9,1,0,'','','管理员删除了你的文章','四月二十四日最新内容不符合社区规范',0,'',0,NULL,'',0,''),(13,'2025-04-25 16:51:51.414','2025-04-25 16:51:51.414',9,1,0,'','','管理员审核了你的文章','审核成功',0,'',0,NULL,'/article/23',0,'接口设计实战技巧');
/*!40000 ALTER TABLE `message_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `models`
--

DROP TABLE IF EXISTS `models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `models`
--

LOCK TABLES `models` WRITE;
/*!40000 ALTER TABLE `models` DISABLE KEYS */;
/*!40000 ALTER TABLE `models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `text_models`
--

DROP TABLE IF EXISTS `text_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `text_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `article_id` bigint(20) unsigned DEFAULT NULL,
  `head` longtext,
  `body` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `text_models`
--

LOCK TABLES `text_models` WRITE;
/*!40000 ALTER TABLE `text_models` DISABLE KEYS */;
INSERT INTO `text_models` VALUES (25,'2025-04-25 15:33:49.989','2025-04-25 15:33:49.989',21,'四月二十四日最新','你好111'),(40,'2025-04-25 16:54:48.067','2025-04-25 16:54:48.067',24,'接口设计实战技巧12312312','啊啊啥的撒的'),(41,'2025-04-25 16:54:48.067','2025-04-25 16:54:48.067',24,'四月二十四日笔记这是我的正文，用来测试 **Markdown 渲染**。## 一、今日目标- 学习 Golang 接口定义- 阅读 Gin 框架文档- 写一个 POST 请求的测试用例## 二、代码片段```gopackage mainimport Hello, Markdown in Go!','是大叔大婶大叔大婶大舍大得发发');
/*!40000 ALTER TABLE `text_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_article_collect_models`
--

DROP TABLE IF EXISTS `user_article_collect_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_article_collect_models` (
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `article_id` bigint(20) unsigned DEFAULT NULL,
  `collect_id` bigint(20) unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`user_id`,`article_id`,`collect_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_article_collect_models`
--

LOCK TABLES `user_article_collect_models` WRITE;
/*!40000 ALTER TABLE `user_article_collect_models` DISABLE KEYS */;
INSERT INTO `user_article_collect_models` VALUES (1,7,1,'2025-04-19 00:04:33.524',4,'2025-04-19 00:04:33.524'),(1,7,4,'2025-04-20 18:38:28.498',5,'2025-04-20 18:38:28.498');
/*!40000 ALTER TABLE `user_article_collect_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_article_look_history_models`
--

DROP TABLE IF EXISTS `user_article_look_history_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_article_look_history_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `article_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_article_look_history_models`
--

LOCK TABLES `user_article_look_history_models` WRITE;
/*!40000 ALTER TABLE `user_article_look_history_models` DISABLE KEYS */;
INSERT INTO `user_article_look_history_models` VALUES (1,NULL,NULL,NULL,NULL),(2,NULL,NULL,NULL,NULL),(3,NULL,NULL,NULL,NULL),(4,'2025-04-23 23:26:55.815','2025-04-23 23:26:55.815',1,10),(5,'2025-04-23 23:27:37.517','2025-04-23 23:27:37.517',1,11);
/*!40000 ALTER TABLE `user_article_look_history_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_chat_action_models`
--

DROP TABLE IF EXISTS `user_chat_action_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_chat_action_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `chat_id` bigint(20) unsigned DEFAULT NULL,
  `is_read` tinyint(1) DEFAULT NULL,
  `is_delete` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_chat_action_models`
--

LOCK TABLES `user_chat_action_models` WRITE;
/*!40000 ALTER TABLE `user_chat_action_models` DISABLE KEYS */;
INSERT INTO `user_chat_action_models` VALUES (1,NULL,NULL,2,1,1,0),(2,'2025-04-22 16:57:41.033','2025-04-22 18:51:36.626',1,1,0,0),(3,'2025-04-22 16:59:31.577','2025-04-22 19:29:28.472',1,2,1,0),(4,'2025-04-22 17:29:53.626','2025-04-22 18:56:03.389',1,3,0,0),(5,'2025-04-22 17:29:53.626','2025-04-22 18:56:03.389',1,4,0,0),(6,'2025-04-22 17:46:26.860','2025-04-22 18:56:03.389',1,5,0,0),(7,'2025-04-22 17:46:26.860','2025-04-22 18:56:03.389',1,6,0,0);
/*!40000 ALTER TABLE `user_chat_action_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_conf_models`
--

DROP TABLE IF EXISTS `user_conf_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_conf_models` (
  `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `like_tags` longtext,
  `update_username_date` datetime(3) DEFAULT NULL,
  `open_collect` tinyint(1) DEFAULT NULL,
  `open_follow` tinyint(1) DEFAULT NULL,
  `open_fans` tinyint(1) DEFAULT NULL,
  `home_style_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `uni_user_conf_models_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_conf_models`
--

LOCK TABLES `user_conf_models` WRITE;
/*!40000 ALTER TABLE `user_conf_models` DISABLE KEYS */;
INSERT INTO `user_conf_models` VALUES (1,NULL,NULL,1,1,1,1),(2,NULL,NULL,1,1,1,1),(3,NULL,NULL,1,1,1,1);
/*!40000 ALTER TABLE `user_conf_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_focus_models`
--

DROP TABLE IF EXISTS `user_focus_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_focus_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `focus_user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_focus_models`
--

LOCK TABLES `user_focus_models` WRITE;
/*!40000 ALTER TABLE `user_focus_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_focus_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_global_notification_models`
--

DROP TABLE IF EXISTS `user_global_notification_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_global_notification_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `notification_id` bigint(20) unsigned DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `is_read` tinyint(1) DEFAULT NULL,
  `is_delete` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_global_notification_models`
--

LOCK TABLES `user_global_notification_models` WRITE;
/*!40000 ALTER TABLE `user_global_notification_models` DISABLE KEYS */;
INSERT INTO `user_global_notification_models` VALUES (1,'2025-04-21 11:30:26.985','2025-04-21 11:32:08.023',1,1,1,1);
/*!40000 ALTER TABLE `user_global_notification_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login_models`
--

DROP TABLE IF EXISTS `user_login_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_login_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `ip` varchar(32) DEFAULT NULL,
  `addr` varchar(64) DEFAULT NULL,
  `ua` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login_models`
--

LOCK TABLES `user_login_models` WRITE;
/*!40000 ALTER TABLE `user_login_models` DISABLE KEYS */;
INSERT INTO `user_login_models` VALUES (1,'2025-04-15 16:49:05.221','2025-04-15 16:49:05.221',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(2,'2025-04-15 20:31:10.046','2025-04-15 20:31:10.046',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(3,'2025-04-15 22:07:08.020','2025-04-15 22:07:08.020',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(4,'2025-04-15 22:08:57.913','2025-04-15 22:08:57.913',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(5,'2025-04-16 14:07:48.474','2025-04-16 14:07:48.474',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(6,'2025-04-16 17:36:13.915','2025-04-16 17:36:13.915',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(7,'2025-04-16 20:50:41.579','2025-04-16 20:50:41.579',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(8,'2025-04-17 12:07:53.339','2025-04-17 12:07:53.339',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(9,'2025-04-17 18:36:23.638','2025-04-17 18:36:23.638',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(10,'2025-04-17 23:12:02.332','2025-04-17 23:12:02.332',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(11,'2025-04-17 23:20:10.436','2025-04-17 23:20:10.436',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(12,'2025-04-18 17:31:02.256','2025-04-18 17:31:02.256',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(13,'2025-04-18 20:47:08.059','2025-04-18 20:47:08.059',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(14,'2025-04-18 23:58:53.089','2025-04-18 23:58:53.089',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(15,'2025-04-20 16:57:29.820','2025-04-20 16:57:29.820',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(16,'2025-04-20 19:51:54.383','2025-04-20 19:51:54.383',3,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(17,'2025-04-20 20:30:19.873','2025-04-20 20:30:19.873',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(18,'2025-04-20 23:42:33.186','2025-04-20 23:42:33.186',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(19,'2025-04-21 11:16:38.218','2025-04-21 11:16:38.218',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(20,'2025-04-21 14:22:03.809','2025-04-21 14:22:03.809',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(21,'2025-04-21 17:43:05.346','2025-04-21 17:43:05.346',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(22,'2025-04-21 21:44:33.751','2025-04-21 21:44:33.751',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(23,'2025-04-22 14:07:07.506','2025-04-22 14:07:07.506',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(24,'2025-04-22 17:25:55.473','2025-04-22 17:25:55.473',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(25,'2025-04-22 22:29:55.465','2025-04-22 22:29:55.465',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(26,'2025-04-23 14:44:33.613','2025-04-23 14:44:33.613',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(27,'2025-04-23 21:23:59.872','2025-04-23 21:23:59.872',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(28,'2025-04-25 14:42:53.765','2025-04-25 14:42:53.765',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(29,'2025-04-25 23:33:22.706','2025-04-25 23:33:22.706',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)'),(30,'2025-04-26 13:52:10.119','2025-04-26 13:52:10.119',1,'127.0.0.1','','Apifox/1.0.0 (https://apifox.com)');
/*!40000 ALTER TABLE `user_login_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_message_conf_models`
--

DROP TABLE IF EXISTS `user_message_conf_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_message_conf_models` (
  `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `open_comment_message` tinyint(1) DEFAULT NULL,
  `open_digg_message` tinyint(1) DEFAULT NULL,
  `open_private_chat` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `uni_user_message_conf_models_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_message_conf_models`
--

LOCK TABLES `user_message_conf_models` WRITE;
/*!40000 ALTER TABLE `user_message_conf_models` DISABLE KEYS */;
INSERT INTO `user_message_conf_models` VALUES (1,1,1,1),(2,1,1,1),(3,1,1,1);
/*!40000 ALTER TABLE `user_message_conf_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_models`
--

DROP TABLE IF EXISTS `user_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `username` varchar(32) DEFAULT NULL,
  `nickname` varchar(32) DEFAULT NULL,
  `avatar` varchar(256) DEFAULT NULL,
  `abstract` varchar(256) DEFAULT NULL,
  `register_source` tinyint(4) DEFAULT NULL,
  `code_age` bigint(20) DEFAULT NULL,
  `password` varchar(64) DEFAULT NULL,
  `email` varchar(256) DEFAULT NULL,
  `open_id` varchar(64) DEFAULT NULL,
  `role` tinyint(4) DEFAULT NULL,
  `ip` longtext,
  `addr` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_models`
--

LOCK TABLES `user_models` WRITE;
/*!40000 ALTER TABLE `user_models` DISABLE KEYS */;
INSERT INTO `user_models` VALUES (1,'2025-04-14 19:34:33.376','2025-04-14 19:34:33.376','naseemli','user1','','',3,0,'$2a$10$tXTc3yk7mKFAJcAYP3SC9eDWxjkwBHv6Nb2qqSFWYNqbQUwqTMYIG','','',1,'',''),(2,'2025-04-15 16:47:41.234','2025-04-15 16:47:41.234','testuser','user2','','',3,0,'$2a$10$Jhij44fhYrWHckLZfF9uM.UJmqmBtkVpcRGozTCFD.1ur4V0/YQ4O','','',1,'',''),(3,'2025-04-20 19:51:36.740','2025-04-20 19:51:36.740','lizeyang','user3','','',3,0,'$2a$10$uTs1DijNQEXcvGzj8D87/./5G99MP9lXNAm9uM1ptLh.E1kl6xRCS','','',2,'',''),(4,NULL,NULL,'testfocus','user4',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `user_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_top_article_models`
--

DROP TABLE IF EXISTS `user_top_article_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_top_article_models` (
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `article_id` bigint(20) unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  UNIQUE KEY `idx_name` (`user_id`,`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_top_article_models`
--

LOCK TABLES `user_top_article_models` WRITE;
/*!40000 ALTER TABLE `user_top_article_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_top_article_models` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-04-26 16:51:57
