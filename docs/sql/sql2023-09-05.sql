/*
 Navicat Premium Data Transfer

 Source Server         : bysystem
 Source Server Type    : MySQL
 Source Server Version : 50740
 Source Host           : 43.136.117.207:3306
 Source Schema         : bysystem

 Target Server Type    : MySQL
 Target Server Version : 50740
 File Encoding         : 65001

 Date: 05/09/2023 19:19:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=265 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (239, 'p', '888', '/ApiService/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (240, 'p', '888', '/ApiService/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (245, 'p', '888', '/ApiService/deleteApisByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (246, 'p', '888', '/ApiService/getAllApiGroups', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (243, 'p', '888', '/ApiService/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (244, 'p', '888', '/ApiService/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (242, 'p', '888', '/ApiService/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (241, 'p', '888', '/ApiService/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (247, 'p', '888', '/AuthorityService/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (248, 'p', '888', '/AuthorityService/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (249, 'p', '888', '/AuthorityService/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (251, 'p', '888', '/AuthorityService/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (252, 'p', '888', '/AuthorityService/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (250, 'p', '888', '/AuthorityService/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (254, 'p', '888', '/CasbinService/GetPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (253, 'p', '888', '/CasbinService/UpdateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (255, 'p', '888', '/MenuService/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (263, 'p', '888', '/MenuService/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (257, 'p', '888', '/MenuService/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (259, 'p', '888', '/MenuService/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (261, 'p', '888', '/MenuService/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (256, 'p', '888', '/MenuService/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (262, 'p', '888', '/MenuService/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (260, 'p', '888', '/MenuService/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (258, 'p', '888', '/MenuService/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (264, 'p', '888', '/ProductService/create', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (236, 'p', '888', '/UserService/ChangePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (229, 'p', '888', '/UserService/DeleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (234, 'p', '888', '/UserService/GetUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (231, 'p', '888', '/UserService/GetUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (228, 'p', '888', '/UserService/Login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (230, 'p', '888', '/UserService/Register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (238, 'p', '888', '/UserService/ResetPassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (233, 'p', '888', '/UserService/SetSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (235, 'p', '888', '/UserService/SetUserAuthorities', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (237, 'p', '888', '/UserService/SetUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (232, 'p', '888', '/UserService/SetUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (181, 'p', '9528', '/MenuService/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (180, 'p', '9528', '/MenuService/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (178, 'p', '9528', '/UserService/ChangePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (177, 'p', '9528', '/UserService/GetUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (174, 'p', '9528', '/UserService/GetUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (172, 'p', '9528', '/UserService/Login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (173, 'p', '9528', '/UserService/Register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (176, 'p', '9528', '/UserService/SetSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (179, 'p', '9528', '/UserService/SetUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (175, 'p', '9528', '/UserService/SetUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (191, 'p', '9999', '/MenuService/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (190, 'p', '9999', '/MenuService/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (188, 'p', '9999', '/UserService/ChangePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (187, 'p', '9999', '/UserService/GetUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (184, 'p', '9999', '/UserService/GetUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (182, 'p', '9999', '/UserService/Login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (183, 'p', '9999', '/UserService/Register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (186, 'p', '9999', '/UserService/SetSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (189, 'p', '9999', '/UserService/SetUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (185, 'p', '9999', '/UserService/SetUserInfo', 'PUT', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for p_product
-- ----------------------------
DROP TABLE IF EXISTS `p_product`;
CREATE TABLE `p_product` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `count` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_p_product_deleted_at` (`deleted_at`),
  KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of p_product
-- ----------------------------
BEGIN;
INSERT INTO `p_product` VALUES (1, '2023-09-05 13:36:53.128', '2023-09-05 13:36:53.128', NULL, '产品1', 100, 1);
INSERT INTO `p_product` VALUES (2, '2023-09-05 13:56:17.695', '2023-09-05 13:56:17.695', NULL, '产品2', 200, 2);
INSERT INTO `p_product` VALUES (3, '2023-09-05 13:56:40.948', '2023-09-05 13:56:40.948', NULL, '产品3', 300, 3);
INSERT INTO `p_product` VALUES (4, '2023-09-05 13:36:53.128', '2023-09-05 13:36:53.128', NULL, '产品4', 400, 4);
INSERT INTO `p_product` VALUES (5, '2023-09-05 13:56:17.695', '2023-09-05 13:56:17.695', NULL, '产品5', 500, 5);
INSERT INTO `p_product` VALUES (6, '2023-09-05 13:56:40.948', '2023-09-05 13:56:40.948', NULL, '产品6', 600, 6);
INSERT INTO `p_product` VALUES (7, '2023-09-05 17:40:33.487', '2023-09-05 17:40:33.487', NULL, '产品new', 100, 0);
INSERT INTO `p_product` VALUES (8, '2023-09-05 17:40:47.117', '2023-09-05 17:40:47.117', NULL, '产品new', 100, 0);
INSERT INTO `p_product` VALUES (9, '2023-09-05 17:46:58.637', '2023-09-05 17:46:58.637', NULL, '产品新增', 2000, 100);
INSERT INTO `p_product` VALUES (10, '2023-09-05 17:48:39.123', '2023-09-05 17:48:39.123', NULL, '产品新增', 2000, 100);
INSERT INTO `p_product` VALUES (11, '2023-09-05 17:48:46.103', '2023-09-05 17:48:46.103', NULL, '产品新增', 2000, 100);
INSERT INTO `p_product` VALUES (12, '2023-09-05 17:48:47.519', '2023-09-05 17:48:47.519', NULL, '产品新增', 2000, 100);
INSERT INTO `p_product` VALUES (13, '2023-09-05 17:48:48.915', '2023-09-05 17:48:48.915', NULL, '产品新增', 2000, 100);
INSERT INTO `p_product` VALUES (14, '2023-09-05 17:50:40.453', '2023-09-05 17:50:40.453', NULL, '产品新增', 2000, 100);
INSERT INTO `p_product` VALUES (15, '2023-09-05 17:51:10.868', '2023-09-05 17:51:10.868', NULL, '产品新增', 2000, 100);
COMMIT;

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
BEGIN;
INSERT INTO `sys_apis` VALUES (1, '2023-07-09 11:55:50.339', '2023-09-03 16:00:24.088', NULL, '/UserService/Login', '用户登录(必选)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (3, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/UserService/DeleteUser', '删除用户', '系统用户', 'DELETE');
INSERT INTO `sys_apis` VALUES (4, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/UserService/Register', '用户注册', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/UserService/GetUserList', '获取用户列表', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (6, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/UserService/SetUserInfo', '设置用户信息', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (7, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/UserService/SetSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (8, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/UserService/GetUserInfo', '获取自身信息(必选)', '系统用户', 'GET');
INSERT INTO `sys_apis` VALUES (9, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/UserService/SetUserAuthorities', '设置权限组', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (10, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/UserService/ChangePassword', '修改密码(建议选择)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/UserService/SetUserAuthority', '修改用户角色(必选)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/UserService/ResetPassword', '重置用户密码', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (13, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/ApiService/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (14, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/ApiService/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (15, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/ApiService/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/ApiService/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/ApiService/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/ApiService/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (19, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/ApiService/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` VALUES (20, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/AuthorityService/copyAuthority', '拷贝角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (21, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/AuthorityService/createAuthority', '创建角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (22, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/AuthorityService/deleteAuthority', '删除角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/AuthorityService/updateAuthority', '更新角色信息', '角色', 'PUT');
INSERT INTO `sys_apis` VALUES (24, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/AuthorityService/getAuthorityList', '获取角色列表', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (25, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/AuthorityService/setDataAuthority', '设置角色资源权限', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (26, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/CasbinService/UpdateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (27, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/CasbinService/GetPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (28, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/MenuService/addBaseMenu', '新增菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (29, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/MenuService/getMenu', '获取菜单树(必选)', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/MenuService/deleteBaseMenu', '删除菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/MenuService/updateBaseMenu', '更新菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (32, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/MenuService/getBaseMenuById', '根据id获取菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (33, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/MenuService/getMenuList', '分页获取基础menu列表', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (34, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/MenuService/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/MenuService/getMenuAuthority', '获取指定角色menu', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (36, '2023-07-09 11:55:50.339', '2023-07-09 11:55:50.339', NULL, '/MenuService/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (85, '2023-09-04 13:55:47.293', '2023-09-04 13:55:47.293', NULL, '/ApiService/getAllApiGroups', '获取api的全部分组', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (86, '2023-09-05 13:34:01.903', '2023-09-05 13:34:01.903', '2023-09-05 13:34:06.885', 'aa', 'aaa', 'aa', 'POST');
INSERT INTO `sys_apis` VALUES (87, '2023-09-05 13:34:38.960', '2023-09-05 13:34:38.960', NULL, '111', 'aaa', '111', 'POST');
INSERT INTO `sys_apis` VALUES (88, '2023-09-05 13:39:50.229', '2023-09-05 13:39:50.229', '2023-09-05 15:07:08.005', 'aaaa', 'asd', 'asds', 'POST');
INSERT INTO `sys_apis` VALUES (89, '2023-09-05 13:40:29.403', '2023-09-05 13:40:29.403', '2023-09-05 15:07:00.955', 'aaaa1', 'sgd', 'asdas', 'POST');
INSERT INTO `sys_apis` VALUES (90, '2023-09-05 17:39:04.774', '2023-09-05 17:39:04.774', NULL, '/ProductService/create', '数据库添加产品', '产品模块', 'POST');
COMMIT;

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  `authority_name` longtext COMMENT '角色名',
  `parent_id` longtext COMMENT '父角色ID',
  `default_router` varchar(191) DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `authority_id` (`authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
BEGIN;
INSERT INTO `sys_authorities` VALUES ('2023-07-09 11:55:51.229', '2023-09-03 16:22:47.105', NULL, '888', '管理员', '0', 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2023-07-09 11:55:51.229', '2023-09-03 18:53:28.575', NULL, '9528', '系统管理员', '0', 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2023-09-03 19:16:21.603', '2023-09-03 19:16:36.444', NULL, '9999', '普通学生', '0', 'dashboard');
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `sys_authority_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_authority_menus` VALUES (1, '888');
INSERT INTO `sys_authority_menus` VALUES (1, '9528');
INSERT INTO `sys_authority_menus` VALUES (1, '9999');
INSERT INTO `sys_authority_menus` VALUES (3, '888');
INSERT INTO `sys_authority_menus` VALUES (4, '888');
INSERT INTO `sys_authority_menus` VALUES (5, '888');
INSERT INTO `sys_authority_menus` VALUES (6, '888');
INSERT INTO `sys_authority_menus` VALUES (7, '888');
INSERT INTO `sys_authority_menus` VALUES (8, '888');
INSERT INTO `sys_authority_menus` VALUES (14, '888');
INSERT INTO `sys_authority_menus` VALUES (14, '9528');
INSERT INTO `sys_authority_menus` VALUES (14, '9999');
INSERT INTO `sys_authority_menus` VALUES (17, '888');
INSERT INTO `sys_authority_menus` VALUES (17, '9528');
INSERT INTO `sys_authority_menus` VALUES (17, '9999');
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL,
  `type` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` varchar(191) DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_base_menus` VALUES (1, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, 0, 0, '仪表盘', 'odometer', 0);
INSERT INTO `sys_base_menus` VALUES (2, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', '2023-09-03 19:19:49.653', 0, '0', 'about', 'about', 0, 'view/about/index.vue', 7, 0, 0, '关于我们', 'info-filled', 0);
INSERT INTO `sys_base_menus` VALUES (3, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, 0, 0, '超级管理员', 'user', 0);
INSERT INTO `sys_base_menus` VALUES (4, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, 0, 0, '角色管理', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (5, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, 1, 0, '菜单管理', 'tickets', 0);
INSERT INTO `sys_base_menus` VALUES (6, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, 1, 0, 'api管理', 'platform', 0);
INSERT INTO `sys_base_menus` VALUES (7, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, 0, 0, '用户管理', 'coordinate', 0);
INSERT INTO `sys_base_menus` VALUES (8, '2023-07-09 11:55:51.025', '2023-09-03 16:44:48.943', NULL, 0, '3', 'myselfinfo', 'myselfinfo', 0, 'view/superAdmin/myselfinfo/myselfinfo.vue', 4, 0, 0, '个人信息', 'medal', 0);
INSERT INTO `sys_base_menus` VALUES (14, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, 0, 0, '系统工具', 'tools', 0);
INSERT INTO `sys_base_menus` VALUES (17, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', NULL, 0, '14', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, 0, 0, '系统配置', 'operation', 0);
INSERT INTO `sys_base_menus` VALUES (18, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', '2023-09-03 12:44:29.505', 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, 0, 0, '字典管理', 'notebook', 0);
INSERT INTO `sys_base_menus` VALUES (19, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', '2023-09-03 12:44:33.094', 0, '3', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, 0, 0, '字典详情', 'order', 0);
INSERT INTO `sys_base_menus` VALUES (20, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', '2023-09-03 12:43:53.718', 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, 0, 0, '操作历史', 'pie-chart', 0);
INSERT INTO `sys_base_menus` VALUES (23, '2023-07-09 11:55:51.025', '2023-07-09 11:55:51.025', '2023-09-03 12:43:17.611', 0, '0', 'state', 'state', 0, 'view/system/state.vue', 6, 0, 0, '服务器状态', 'cloudy', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
BEGIN;
INSERT INTO `sys_data_authority_id` VALUES ('888', '888');
INSERT INTO `sys_data_authority_id` VALUES ('888', '9528');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '9528');
INSERT INTO `sys_data_authority_id` VALUES ('9999', '9528');
COMMIT;

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `sys_authority_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_authority` VALUES (1, '888');
INSERT INTO `sys_user_authority` VALUES (1, '9528');
INSERT INTO `sys_user_authority` VALUES (2, '888');
INSERT INTO `sys_user_authority` VALUES (3, '888');
INSERT INTO `sys_user_authority` VALUES (9, '888');
INSERT INTO `sys_user_authority` VALUES (9, '9528');
INSERT INTO `sys_user_authority` VALUES (10, '888');
COMMIT;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` longtext COMMENT '用户UUID',
  `username` longtext COMMENT '用户登录名',
  `password` longtext COMMENT '用户登录密码',
  `nick_name` varchar(191) DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` varchar(191) DEFAULT '888' COMMENT '用户角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` VALUES (1, '2023-07-09 11:55:50.575', '2023-09-04 13:40:22.480', NULL, '9b9b9ab7-1639-4c9f-8ebc-b0ab21f1f30c', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '用户12', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', '888');
INSERT INTO `sys_users` VALUES (2, '2023-07-09 11:55:50.575', '2023-07-11 16:15:46.680', NULL, 'c2d280d6-b222-4a33-8dde-4f948b8b3dc4', 'a303176530', 'e10adc3949ba59abbe56e057f20f883e', 'QMPlusUser', 'dark', 'https:///qmplusimg.henrongyi.top/1572075907logo.png', '#fff', '#1890ff', '888');
INSERT INTO `sys_users` VALUES (3, '2023-07-11 17:10:08.376', '2023-07-17 11:11:55.995', NULL, '27b808ac-365e-498b-8148-57b4b347d2ac', '125554566', 'e10adc3949ba59abbe56e057f20f883e', 'wdf', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', '888');
INSERT INTO `sys_users` VALUES (9, '2023-09-02 16:25:48.644', '2023-09-03 19:01:47.917', NULL, '511bcbae-20cd-411b-9bbd-871024b3b7a9', 'wdfwds', '4124bc0a9335c27f086f24ba207a4912', 'wdfwds', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', '888');
INSERT INTO `sys_users` VALUES (10, '2023-09-02 16:26:06.420', '2023-09-03 19:01:57.783', NULL, 'cb437f40-2fa6-4643-bedc-ade904d8b464', 'wdfwdsa', 'd06d94c9b591a22a356dfc8afccbc303', 'wdfwds', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', '888');
COMMIT;

-- ----------------------------
-- View structure for authority_menu
-- ----------------------------
DROP VIEW IF EXISTS `authority_menu`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`sort` AS `sort`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`default_menu` AS `default_menu`,`sys_base_menus`.`close_tab` AS `close_tab`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)));

SET FOREIGN_KEY_CHECKS = 1;
