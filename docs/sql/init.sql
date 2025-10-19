/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80041 (8.0.41)
 Source Host           : localhost:3306
 Source Schema         : goat

 Target Server Type    : MySQL
 Target Server Version : 80041 (8.0.41)
 File Encoding         : 65001

 Date: 19/10/2025 01:12:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `config_id` bigint NOT NULL AUTO_INCREMENT COMMENT '参数id',
  `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'N' COMMENT '系统内置：Y-是；N-否',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`config_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='参数配置表';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (1, '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', 'admin', '2025-10-06 02:44:02', '', NULL, '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', 'admin', '2025-10-06 02:44:02', '', NULL, '初始化密码 123456');
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (3, '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', 'admin', '2025-10-06 02:44:02', '', NULL, '深色主题theme-dark，浅色主题theme-light');
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (4, '账号自助-验证码开关', 'sys.account.captchaEnabled', 'true', 'Y', 'admin', '2025-10-06 02:44:02', '', NULL, '是否开启验证码功能（true开启，false关闭）');
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (5, '账号自助-是否开启用户注册功能', 'sys.account.registerUser', 'false', 'Y', 'admin', '2025-10-06 02:44:02', '', NULL, '是否开启注册用户功能（true开启，false关闭）');
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` bigint NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int NOT NULL DEFAULT '0' COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '状态：0-正常；1-停用',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (100, 0, ',0', '总公司', 0, 'hugo', '18243088680', 'zhang8680@outlook.com', '0', 'admin', '2025-10-06 02:44:02', 'admin', '2025-10-19 00:53:06', NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (101, 100, ',0,100', '深圳总公司', 1, 'hugo', '18243088680', 'zhang8680@outlook.com', '0', 'admin', '2025-10-06 02:44:02', 'admin', '2025-10-19 00:53:34', NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (102, 100, ',0,100', '长沙分公司', 2, 'hugo', '18243088680', 'zhang8680@outlook.com', '0', 'admin', '2025-10-06 02:44:02', 'admin', '2025-10-19 00:53:55', '2025-10-19 00:54:17');
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (103, 101, ',0,100,101', '研发部门', 1, 'hugo', '18243088680', 'zhang8680@outlook.com', '0', 'admin', '2025-10-06 02:44:02', 'admin', '2025-10-19 00:55:01', NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (104, 101, '0,100,101', '市场部门', 2, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '2025-10-19 00:54:23');
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (105, 101, ',0,100,101', '测试部门', 3, 'hugo', '18243088680', 'zhang8680@outlook.com', '0', 'admin', '2025-10-06 02:44:02', 'admin', '2025-10-19 00:55:15', NULL);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (106, 101, '0,100,101', '财务部门', 4, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '2025-10-19 00:54:34');
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (107, 101, '0,100,101', '运维部门', 5, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '2025-10-19 00:54:28');
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (108, 102, '0,100,102', '市场部门', 1, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '2025-10-19 00:54:09');
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (109, 102, '0,100,102', '财务部门', 2, '若依', '15888888888', 'ry@qq.com', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '2025-10-19 00:54:15');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `dict_code` bigint NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int NOT NULL DEFAULT '0' COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '表格回显样式',
  `is_default` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'N' COMMENT '是否默认：Y-是；N-否',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '状态：0-正常；1-停用',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典数据表';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (1, 1, '男', '0', 'sys_user_sex', '', '', 'Y', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '性别男');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (2, 2, '女', '1', 'sys_user_sex', '', '', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '性别女');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (3, 3, '未知', '2', 'sys_user_sex', '', '', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '性别未知');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (4, 1, '显示', '0', 'sys_show_hide', '', 'primary', 'Y', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '显示菜单');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (5, 2, '隐藏', '1', 'sys_show_hide', '', 'danger', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '隐藏菜单');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (6, 1, '正常', '0', 'sys_normal_disable', '', 'primary', 'Y', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '正常状态');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (7, 2, '停用', '1', 'sys_normal_disable', '', 'danger', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '停用状态');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (8, 1, '正常', '0', 'sys_job_status', '', 'primary', 'Y', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '正常状态');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (9, 2, '暂停', '1', 'sys_job_status', '', 'danger', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '停用状态');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (10, 1, '默认', 'DEFAULT', 'sys_job_group', '', '', 'Y', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '默认分组');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (11, 2, '系统', 'SYSTEM', 'sys_job_group', '', '', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '系统分组');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (12, 1, '是', 'Y', 'sys_yes_no', '', 'primary', 'Y', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '系统默认是');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (13, 2, '否', 'N', 'sys_yes_no', '', 'danger', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '系统默认否');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (14, 1, '通知', '1', 'sys_notice_type', '', 'warning', 'Y', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '通知');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (15, 2, '公告', '2', 'sys_notice_type', '', 'success', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '公告');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (16, 1, '正常', '0', 'sys_notice_status', '', 'primary', 'Y', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '正常状态');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (17, 2, '关闭', '1', 'sys_notice_status', '', 'danger', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '关闭状态');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (18, 99, '其他', '0', 'sys_oper_type', '', 'info', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '其他操作');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (19, 1, '新增', '1', 'sys_oper_type', '', 'info', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '新增操作');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (20, 2, '修改', '2', 'sys_oper_type', '', 'info', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '修改操作');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (21, 3, '删除', '3', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '删除操作');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (22, 4, '授权', '4', 'sys_oper_type', '', 'primary', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '授权操作');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (23, 5, '导出', '5', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '导出操作');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (24, 6, '导入', '6', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '导入操作');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (25, 7, '强退', '7', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '强退操作');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (26, 8, '生成代码', '8', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '生成操作');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (27, 9, '清空数据', '9', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '清空操作');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (28, 1, '成功', '0', 'sys_common_status', '', 'primary', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '正常状态');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (29, 2, '失败', '1', 'sys_common_status', '', 'danger', 'N', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '停用状态');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `dict_id` bigint NOT NULL AUTO_INCREMENT COMMENT '字典id',
  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '字典类型',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '状态：0-正常；1-停用',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE KEY `dict_type` (`dict_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典类型表';

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (1, '用户性别', 'sys_user_sex', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '用户性别列表');
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (2, '菜单状态', 'sys_show_hide', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '菜单状态列表');
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (3, '系统开关', 'sys_normal_disable', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '系统开关列表');
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (4, '任务状态', 'sys_job_status', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '任务状态列表');
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (5, '任务分组', 'sys_job_group', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '任务分组列表');
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (6, '系统是否', 'sys_yes_no', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '系统是否列表');
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (7, '通知类型', 'sys_notice_type', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '通知类型列表');
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (8, '通知状态', 'sys_notice_status', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '通知状态列表');
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (9, '操作类型', 'sys_oper_type', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '操作类型列表');
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (10, '系统状态', 'sys_common_status', '0', 'admin', '2025-10-06 02:44:02', '', NULL, '登录状态列表');
INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES (11, 'test', 'test', '0', 'admin', '2025-10-17 03:37:58', '', '2025-10-17 03:37:58', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log` (
  `info_id` bigint NOT NULL AUTO_INCREMENT COMMENT '访问id',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `ipaddr` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '登录ip地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '状态：0-成功；1-失败',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime NOT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`) USING BTREE,
  KEY `idx_sys_logininfor_s` (`status`) USING BTREE,
  KEY `idx_sys_logininfor_lt` (`login_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统访问记录';

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
BEGIN;
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (1, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '1', '验证码错误', '2025-10-06 03:35:28');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (2, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-06 03:35:34');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (3, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '1', '密码错误次数超过限制，请0分钟后重试', '2025-10-16 02:16:29');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (4, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '1', '密码错误次数超过限制，请0分钟后重试', '2025-10-16 02:16:40');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (5, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '1', '验证码错误', '2025-10-16 02:16:57');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (6, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '1', '密码错误次数超过限制，请0分钟后重试', '2025-10-16 02:17:08');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (7, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '1', '密码错误次数超过限制，请0分钟后重试', '2025-10-16 02:17:43');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (8, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '1', '密码错误次数超过限制，请0分钟后重试', '2025-10-16 13:23:36');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (9, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '1', '密码错误次数超过限制，请0分钟后重试', '2025-10-16 13:27:20');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (10, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 13:32:38');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (11, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 13:33:08');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (12, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 13:33:23');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (13, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 13:33:42');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (14, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 13:37:50');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (15, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 13:51:42');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (16, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 13:55:07');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (17, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 13:57:20');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (18, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 13:59:20');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (19, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 13:59:48');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (20, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 14:02:02');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (21, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 14:02:53');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (22, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 14:04:25');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (23, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 14:08:41');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (24, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 14:44:15');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (25, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 14:56:07');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (26, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 14:57:48');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (27, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 15:00:06');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (28, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 15:01:34');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (29, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 15:06:41');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (30, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 15:07:23');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (31, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '1', '验证码错误', '2025-10-16 15:09:00');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (32, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 15:09:06');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (33, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 15:10:56');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (34, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 15:45:35');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (35, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-16 23:12:35');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (36, 'admin', '::1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-17 00:34:49');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (37, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-17 03:31:42');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (38, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '1', '验证码错误', '2025-10-17 15:23:03');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (39, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-17 22:39:25');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (40, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-17 22:46:22');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (41, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-17 23:41:07');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (42, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-17 23:45:07');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (43, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-17 23:50:45');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (44, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 00:16:17');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (45, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 00:22:35');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (46, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 00:27:47');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (47, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 00:30:35');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (48, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 00:36:29');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (49, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 00:40:19');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (50, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 00:41:58');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (51, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 00:43:30');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (52, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 00:45:59');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (53, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 00:49:03');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (54, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 01:33:00');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (55, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 01:43:36');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (56, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 01:48:28');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (57, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 02:19:00');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (58, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 02:30:04');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (59, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 15:04:41');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (60, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 22:38:37');
INSERT INTO `sys_login_log` (`info_id`, `user_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (61, 'admin', '127.0.0.1', '内网地址', 'Chrome', 'macOS', '0', '成功', '2025-10-18 23:01:59');
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单id',
  `menu_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单名称',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '父菜单id',
  `order_num` int NOT NULL DEFAULT '0' COMMENT '显示顺序',
  `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '组件路径',
  `query` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '路由参数',
  `route_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '路由名称',
  `is_frame` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否为外链：0-是；1-否',
  `is_cache` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否缓存：0-缓存；1-不缓存',
  `menu_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '菜单类型：M-目录；C-菜单；F-按钮',
  `visible` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '菜单状态：0-显示；1-隐藏',
  `perms` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '权限标识',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '#' COMMENT '菜单图标',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '状态：0-正常；1-停用',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1046 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单权限表';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1, '系统管理', 0, 1, 'system', NULL, '', '', 1, 0, 'M', '0', '', 'system', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '系统管理目录');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (100, '用户管理', 1, 1, 'user', 'system/user/index', '', '', 1, 0, 'C', '0', 'system:user:list', 'user', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '用户管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (101, '角色管理', 1, 2, 'role', 'system/role/index', '', '', 1, 0, 'C', '0', 'system:role:list', 'peoples', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '角色管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (102, '菜单管理', 1, 3, 'menu', 'system/menu/index', '', '', 1, 0, 'C', '0', 'system:menu:list', 'tree-table', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '菜单管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (103, '部门管理', 1, 4, 'dept', 'system/dept/index', '', '', 1, 0, 'C', '0', 'system:dept:list', 'tree', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '部门管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (104, '岗位管理', 1, 5, 'post', 'system/post/index', '', '', 1, 0, 'C', '0', 'system:post:list', 'post', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '岗位管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (105, '字典管理', 1, 6, 'dict', 'system/dict/index', '', '', 1, 0, 'C', '0', 'system:dict:list', 'dict', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '字典管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (106, '参数设置', 1, 7, 'config', 'system/config/index', '', '', 1, 0, 'C', '0', 'system:config:list', 'edit', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '参数设置菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (108, '日志管理', 1, 9, 'log', '', '', '', 1, 0, 'M', '0', '', 'log', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '日志管理菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (500, '操作日志', 108, 1, 'operLog', 'system/operLog/index', '', '', 1, 0, 'C', '0', 'system:operLog:list', 'form', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '操作日志菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (501, '登录日志', 108, 2, 'loginLog', 'system/loginLog/index', '', '', 1, 0, 'C', '0', 'system:loginLog:list', 'IconDesktop', '0', 'admin', '2025-10-06 02:44:02', 'admin', '2025-10-18 02:28:34', NULL, '登录日志菜单');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1000, '用户查询', 100, 1, '', '', '', '', 1, 0, 'F', '0', 'system:user:query', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1001, '用户新增', 100, 2, '', '', '', '', 1, 0, 'F', '0', 'system:user:add', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1002, '用户修改', 100, 3, '', '', '', '', 1, 0, 'F', '0', 'system:user:edit', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1003, '用户删除', 100, 4, '', '', '', '', 1, 0, 'F', '0', 'system:user:remove', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1004, '用户导出', 100, 5, '', '', '', '', 1, 0, 'F', '0', 'system:user:export', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1005, '用户导入', 100, 6, '', '', '', '', 1, 0, 'F', '0', 'system:user:import', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1006, '重置密码', 100, 7, '', '', '', '', 1, 0, 'F', '0', 'system:user:resetPwd', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1007, '角色查询', 101, 1, '', '', '', '', 1, 0, 'F', '0', 'system:role:query', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1008, '角色新增', 101, 2, '', '', '', '', 1, 0, 'F', '0', 'system:role:add', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1009, '角色修改', 101, 3, '', '', '', '', 1, 0, 'F', '0', 'system:role:edit', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1010, '角色删除', 101, 4, '', '', '', '', 1, 0, 'F', '0', 'system:role:remove', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1011, '角色导出', 101, 5, '', '', '', '', 1, 0, 'F', '0', 'system:role:export', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1012, '菜单查询', 102, 1, '', '', '', '', 1, 0, 'F', '0', 'system:menu:query', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1013, '菜单新增', 102, 2, '', '', '', '', 1, 0, 'F', '0', 'system:menu:add', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1014, '菜单修改', 102, 3, '', '', '', '', 1, 0, 'F', '0', 'system:menu:edit', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1015, '菜单删除', 102, 4, '', '', '', '', 1, 0, 'F', '0', 'system:menu:remove', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1016, '部门查询', 103, 1, '', '', '', '', 1, 0, 'F', '0', 'system:dept:query', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1017, '部门新增', 103, 2, '', '', '', '', 1, 0, 'F', '0', 'system:dept:add', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1018, '部门修改', 103, 3, '', '', '', '', 1, 0, 'F', '0', 'system:dept:edit', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1019, '部门删除', 103, 4, '', '', '', '', 1, 0, 'F', '0', 'system:dept:remove', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1020, '岗位查询', 104, 1, '', '', '', '', 1, 0, 'F', '0', 'system:post:query', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1021, '岗位新增', 104, 2, '', '', '', '', 1, 0, 'F', '0', 'system:post:add', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1022, '岗位修改', 104, 3, '', '', '', '', 1, 0, 'F', '0', 'system:post:edit', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1023, '岗位删除', 104, 4, '', '', '', '', 1, 0, 'F', '0', 'system:post:remove', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1024, '岗位导出', 104, 5, '', '', '', '', 1, 0, 'F', '0', 'system:post:export', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1025, '字典查询', 105, 1, '#', '', '', '', 1, 0, 'F', '0', 'system:dict:query', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1026, '字典新增', 105, 2, '#', '', '', '', 1, 0, 'F', '0', 'system:dict:add', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1027, '字典修改', 105, 3, '#', '', '', '', 1, 0, 'F', '0', 'system:dict:edit', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1028, '字典删除', 105, 4, '#', '', '', '', 1, 0, 'F', '0', 'system:dict:remove', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1029, '字典导出', 105, 5, '#', '', '', '', 1, 0, 'F', '0', 'system:dict:export', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1030, '参数查询', 106, 1, '#', '', '', '', 1, 0, 'F', '0', 'system:config:query', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1031, '参数新增', 106, 2, '#', '', '', '', 1, 0, 'F', '0', 'system:config:add', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1032, '参数修改', 106, 3, '#', '', '', '', 1, 0, 'F', '0', 'system:config:edit', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1033, '参数删除', 106, 4, '#', '', '', '', 1, 0, 'F', '0', 'system:config:remove', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1034, '参数导出', 106, 5, '#', '', '', '', 1, 0, 'F', '0', 'system:config:export', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1039, '操作查询', 500, 1, '#', '', '', '', 1, 0, 'F', '0', 'system:operLog:query', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1040, '操作删除', 500, 2, '#', '', '', '', 1, 0, 'F', '0', 'system:operLog:remove', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1041, '日志导出', 500, 3, '#', '', '', '', 1, 0, 'F', '0', 'system:operLog:export', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1042, '登录查询', 501, 1, '#', '', '', '', 1, 0, 'F', '0', 'system:loginLog:query', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1043, '登录删除', 501, 2, '#', '', '', '', 1, 0, 'F', '0', 'system: loginLog:remove', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1044, '日志导出', 501, 3, '#', '', '', '', 1, 0, 'F', '0', 'system: loginLog:export', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query`, `route_name`, `is_frame`, `is_cache`, `menu_type`, `visible`, `perms`, `icon`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1045, '账户解锁', 501, 4, '#', '', '', '', 1, 0, 'F', '0', 'system: loginLog:unlock', '#', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log` (
  `oper_id` bigint NOT NULL AUTO_INCREMENT COMMENT '日志id',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '模块标题',
  `business_type` tinyint NOT NULL DEFAULT '0' COMMENT '业务类型：0-其它；1-新增；2-修改；3-删除',
  `method` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '请求方式',
  `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '请求url',
  `oper_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '请求参数',
  `json_result` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '返回参数',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '操作状态：0-正常；1-异常',
  `error_msg` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime NOT NULL COMMENT '操作时间',
  `cost_time` bigint NOT NULL DEFAULT '0' COMMENT '消耗时间',
  PRIMARY KEY (`oper_id`) USING BTREE,
  KEY `idx_sys_oper_log_bt` (`business_type`) USING BTREE,
  KEY `idx_sys_oper_log_s` (`status`) USING BTREE,
  KEY `idx_sys_oper_log_ot` (`oper_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='操作日志记录';

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
BEGIN;
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (1, '新增字典类型', 1, 'github.com/hugo8680/goat/api/controller/admin.(*DictTypeController).Create-fm', 'POST', '若依', '研发部门', '/api/system/dict/type', '127.0.0.1', '内网地址', '{\"dictName\":\"test\",\"dictType\":\"test\",\"status\":\"0\"}', '{\"code\":200,\"msg\":\"成功\"}', '0', '', '2025-10-17 03:37:58', 3);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (2, '修改菜单', 2, 'github.com/hugo8680/goat/api/controller/admin.(*RoleController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/menu', '127.0.0.1', '内网地址', '{\"component\":\"system/loginLog/index\",\"icon\":\"IconBug\",\"isCache\":0,\"isFrame\":1,\"menuId\":501,\"menuName\":\"登录日志\",\"menuType\":\"C\",\"orderNum\":2,\"parentId\":108,\"path\":\"loginLog\",\"perms\":\"system:loginLog:list\",\"query\":\"\",\"routeName\":\"\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":500,\"msg\":\"参数错误\"}', '1', '参数错误', '2025-10-18 02:22:46', 0);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (3, '修改菜单', 2, 'github.com/hugo8680/goat/api/controller/admin.(*RoleController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/menu', '127.0.0.1', '内网地址', '{\"component\":\"system/loginLog/index\",\"icon\":\"IconBug\",\"isCache\":0,\"isFrame\":1,\"menuId\":501,\"menuName\":\"登录日志\",\"menuType\":\"C\",\"orderNum\":2,\"parentId\":108,\"path\":\"loginLog\",\"perms\":\"system:loginLog:list\",\"query\":\"\",\"routeName\":\"\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":500,\"msg\":\"参数错误\"}', '1', '参数错误', '2025-10-18 02:22:55', 0);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (4, '修改菜单', 2, 'github.com/hugo8680/goat/api/controller/admin.(*RoleController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/menu', '127.0.0.1', '内网地址', '{\"component\":\"system/loginLog/index\",\"icon\":\"IconBug\",\"isCache\":0,\"isFrame\":1,\"menuId\":501,\"menuName\":\"登录日志\",\"menuType\":\"C\",\"orderNum\":2,\"parentId\":108,\"path\":\"loginLog\",\"perms\":\"system:loginLog:list\",\"query\":\"\",\"routeName\":\"\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":500,\"msg\":\"参数错误\"}', '1', '参数错误', '2025-10-18 02:23:51', 0);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (5, '修改菜单', 2, 'github.com/hugo8680/goat/api/controller/admin.(*RoleController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/menu', '127.0.0.1', '内网地址', '{\"component\":\"system/loginLog/index\",\"icon\":\"IconBug\",\"isCache\":0,\"isFrame\":1,\"menuId\":501,\"menuName\":\"登录日志\",\"menuType\":\"C\",\"orderNum\":2,\"parentId\":108,\"path\":\"loginLog\",\"perms\":\"system:loginLog:list\",\"query\":\"\",\"routeName\":\"\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":500,\"msg\":\"参数错误\"}', '1', '参数错误', '2025-10-18 02:24:09', 0);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (6, '修改菜单', 2, 'github.com/hugo8680/goat/api/controller/admin.(*MenuController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/menu', '127.0.0.1', '内网地址', '{\"component\":\"system/loginLog/index\",\"icon\":\"IconBug\",\"isCache\":0,\"isFrame\":1,\"menuId\":501,\"menuName\":\"登录日志\",\"menuType\":\"C\",\"orderNum\":2,\"parentId\":108,\"path\":\"loginLog\",\"perms\":\"system:loginLog:list\",\"query\":\"\",\"routeName\":\"\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":200,\"msg\":\"成功\"}', '0', '', '2025-10-18 02:27:32', 5);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (7, '修改菜单', 2, 'github.com/hugo8680/goat/api/controller/admin.(*MenuController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/menu', '127.0.0.1', '内网地址', '{\"component\":\"system/loginLog/index\",\"icon\":\"IconDesktop\",\"isCache\":0,\"isFrame\":1,\"menuId\":501,\"menuName\":\"登录日志\",\"menuType\":\"C\",\"orderNum\":2,\"parentId\":108,\"path\":\"loginLog\",\"perms\":\"system:loginLog:list\",\"query\":\"\",\"routeName\":\"\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":200,\"msg\":\"成功\"}', '0', '', '2025-10-18 02:28:34', 2);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (8, '导出用户', 5, 'github.com/hugo8680/goat/api/controller/admin.(*UserController).Export-fm', 'POST', '超级管理员', '研发部门', '/api/system/user/export', '127.0.0.1', '内网地址', '{\"{\\\"pageNum\\\":1,\\\"pageSize\\\":10}\":\"\"}', '', '1', '', '2025-10-18 23:31:50', 14640);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (9, '导出用户', 5, 'github.com/hugo8680/goat/api/controller/admin.(*UserController).Export-fm', 'POST', '超级管理员', '研发部门', '/api/system/user/export', '127.0.0.1', '内网地址', '{\"{\\\"pageNum\\\":1,\\\"pageSize\\\":10}\":\"\"}', '', '1', '', '2025-10-18 23:50:03', 1075216);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (10, '导出用户', 5, 'github.com/hugo8680/goat/api/controller/admin.(*UserController).Export-fm', 'POST', '超级管理员', '研发部门', '/api/system/user/export', '127.0.0.1', '内网地址', '{\"{\\\"pageNum\\\":1,\\\"pageSize\\\":10}\":\"\"}', '', '1', '', '2025-10-18 23:52:32', 136205);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (11, '导出用户', 5, 'github.com/hugo8680/goat/api/controller/admin.(*UserController).Export-fm', 'POST', '超级管理员', '研发部门', '/api/system/user/export', '127.0.0.1', '内网地址', '{}', '', '1', '', '2025-10-18 23:58:26', 8219);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (12, '导出用户', 5, 'github.com/hugo8680/goat/api/controller/admin.(*UserController).Export-fm', 'POST', '超级管理员', '研发部门', '/api/system/user/export', '127.0.0.1', '内网地址', '{\"pageNum\":1,\"pageSize\":10}', '', '0', 'OK', '2025-10-19 00:48:55', 13);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (13, '导出用户', 5, 'github.com/hugo8680/goat/api/controller/admin.(*UserController).Export-fm', 'POST', '超级管理员', '研发部门', '/api/system/user/export', '127.0.0.1', '内网地址', '{\"pageNum\":1,\"pageSize\":10}', '', '0', 'OK', '2025-10-19 00:49:38', 68320);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (14, '修改部门', 2, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/dept', '127.0.0.1', '内网地址', '{\"ancestors\":\"0\",\"createTime\":\"2025-10-06 02:44:02\",\"deptId\":100,\"deptName\":\"总公司\",\"email\":\"zhang8680@outlook.com\",\"leader\":\"hugo\",\"orderNum\":0,\"parentId\":0,\"phone\":\"18243088680\",\"status\":\"0\"}', '{\"code\":200,\"msg\":\"成功\"}', '0', '', '2025-10-19 00:53:06', 3);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (15, '修改部门', 2, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/dept', '127.0.0.1', '内网地址', '{\"ancestors\":\"0,100\",\"createTime\":\"2025-10-06 02:44:02\",\"deptId\":101,\"deptName\":\"深圳总公司\",\"email\":\"zhang8680@outlook.com\",\"leader\":\"hugo\",\"orderNum\":1,\"parentId\":100,\"phone\":\"18243088680\",\"status\":\"0\"}', '{\"code\":200,\"msg\":\"成功\"}', '0', '', '2025-10-19 00:53:34', 3);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (16, '修改部门', 2, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/dept', '127.0.0.1', '内网地址', '{\"ancestors\":\"0,100\",\"createTime\":\"2025-10-06 02:44:02\",\"deptId\":102,\"deptName\":\"长沙分公司\",\"email\":\"zhang8680@outlook.com\",\"leader\":\"hugo\",\"orderNum\":2,\"parentId\":100,\"phone\":\"18243088680\",\"status\":\"0\"}', '{\"code\":200,\"msg\":\"成功\"}', '0', '', '2025-10-19 00:53:55', 3);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (17, '删除部门', 3, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Delete-fm', 'DELETE', '超级管理员', '研发部门', '/api/system/dept/102', '127.0.0.1', '内网地址', '{}', '', '0', 'OK', '2025-10-19 00:54:05', 1);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (18, '删除部门', 3, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Delete-fm', 'DELETE', '超级管理员', '研发部门', '/api/system/dept/108', '127.0.0.1', '内网地址', '{}', '', '0', 'OK', '2025-10-19 00:54:09', 2);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (19, '删除部门', 3, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Delete-fm', 'DELETE', '超级管理员', '研发部门', '/api/system/dept/109', '127.0.0.1', '内网地址', '{}', '', '0', 'OK', '2025-10-19 00:54:15', 3);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (20, '删除部门', 3, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Delete-fm', 'DELETE', '超级管理员', '研发部门', '/api/system/dept/102', '127.0.0.1', '内网地址', '{}', '', '0', 'OK', '2025-10-19 00:54:17', 3);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (21, '删除部门', 3, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Delete-fm', 'DELETE', '超级管理员', '研发部门', '/api/system/dept/104', '127.0.0.1', '内网地址', '{}', '', '0', 'OK', '2025-10-19 00:54:23', 2);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (22, '删除部门', 3, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Delete-fm', 'DELETE', '超级管理员', '研发部门', '/api/system/dept/105', '127.0.0.1', '内网地址', '{}', '', '0', 'OK', '2025-10-19 00:54:26', 1);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (23, '删除部门', 3, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Delete-fm', 'DELETE', '超级管理员', '研发部门', '/api/system/dept/107', '127.0.0.1', '内网地址', '{}', '', '0', 'OK', '2025-10-19 00:54:28', 2);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (24, '删除部门', 3, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Delete-fm', 'DELETE', '超级管理员', '研发部门', '/api/system/dept/106', '127.0.0.1', '内网地址', '{}', '', '0', 'OK', '2025-10-19 00:54:34', 3);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (25, '删除部门', 3, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Delete-fm', 'DELETE', '超级管理员', '研发部门', '/api/system/dept/105', '127.0.0.1', '内网地址', '{}', '', '0', 'OK', '2025-10-19 00:54:36', 1);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (26, '修改部门', 2, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/dept', '127.0.0.1', '内网地址', '{\"ancestors\":\"0,100,101\",\"createTime\":\"2025-10-06 02:44:02\",\"deptId\":103,\"deptName\":\"研发部门\",\"email\":\"zhang8680@outlook.com\",\"leader\":\"hugo\",\"orderNum\":1,\"parentId\":101,\"phone\":\"18243088680\",\"status\":\"0\"}', '{\"code\":200,\"msg\":\"成功\"}', '0', '', '2025-10-19 00:55:01', 3);
INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`, `cost_time`) VALUES (27, '修改部门', 2, 'github.com/hugo8680/goat/api/controller/admin.(*DeptController).Update-fm', 'PUT', '超级管理员', '研发部门', '/api/system/dept', '127.0.0.1', '内网地址', '{\"ancestors\":\"0,100,101\",\"createTime\":\"2025-10-06 02:44:02\",\"deptId\":105,\"deptName\":\"测试部门\",\"email\":\"zhang8680@outlook.com\",\"leader\":\"hugo\",\"orderNum\":3,\"parentId\":101,\"phone\":\"18243088680\",\"status\":\"0\"}', '{\"code\":200,\"msg\":\"成功\"}', '0', '', '2025-10-19 00:55:15', 2);
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` bigint NOT NULL AUTO_INCREMENT COMMENT '岗位id',
  `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int NOT NULL DEFAULT '0' COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '状态：0-正常；1-停用',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='岗位信息表';

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_post` (`post_id`, `post_code`, `post_name`, `post_sort`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1, 'ceo', '董事长', 1, '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_post` (`post_id`, `post_code`, `post_name`, `post_sort`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (2, 'se', '项目经理', 2, '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_post` (`post_id`, `post_code`, `post_name`, `post_sort`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (3, 'hr', '人力资源', 3, '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
INSERT INTO `sys_post` (`post_id`, `post_code`, `post_name`, `post_sort`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (4, 'user', '普通员工', 4, '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `role_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色权限字符串',
  `role_sort` int NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '1' COMMENT '数据范围：1-全部数据权限；2-自定数据权限；3-本部门数据权限；4-本部门及以下数据权限',
  `menu_check_strictly` tinyint(1) NOT NULL DEFAULT '1' COMMENT '菜单树选择项是否关联显示',
  `dept_check_strictly` tinyint(1) NOT NULL DEFAULT '1' COMMENT '部门树选择项是否关联显示',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '状态：0-正常；1-停用',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色信息表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`role_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1, '超级管理员', 'admin', 1, '1', 1, 1, '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '超级管理员');
INSERT INTO `sys_role` (`role_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `menu_check_strictly`, `dept_check_strictly`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (2, '普通角色', 'common', 2, '2', 1, 1, '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '普通角色');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` bigint NOT NULL COMMENT '角色id',
  `dept_id` bigint NOT NULL COMMENT '部门id',
  PRIMARY KEY (`role_id`,`dept_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色和部门关联表';

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`) VALUES (2, 100);
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`) VALUES (2, 101);
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`) VALUES (2, 105);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` bigint NOT NULL COMMENT '角色id',
  `menu_id` bigint NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`role_id`,`menu_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色和菜单关联表';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 2);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 3);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 4);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 100);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 101);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 102);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 103);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 104);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 105);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 106);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 108);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 500);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 501);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1000);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1001);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1002);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1003);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1004);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1005);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1006);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1007);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1008);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1009);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1010);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1011);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1012);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1013);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1014);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1015);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1016);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1017);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1018);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1019);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1020);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1021);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1022);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1023);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1024);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1025);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1026);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1027);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1028);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1029);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1030);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1031);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1032);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1033);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1034);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1035);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1036);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1037);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1038);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1039);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1040);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1041);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1042);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1043);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1044);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 1045);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `dept_id` bigint NOT NULL DEFAULT '0' COMMENT '部门id',
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户账号',
  `nick_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户昵称',
  `user_type` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '00' COMMENT '用户类型：00-系统用户',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `phone_number` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `sex` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '用户性别：0-男；1-女；2-未知',
  `avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码',
  `login_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `login_date` datetime DEFAULT NULL COMMENT '最后登录时间',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '状态：0-正常；1-停用',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`user_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phone_number`, `sex`, `avatar`, `password`, `login_ip`, `login_date`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (1, 103, 'admin', '超级管理员', '00', 'zhang8680@outlook.com', '18243088680', '1', '', '$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2', '127.0.0.1', '2025-10-18 23:01:59', '0', 'admin', '2025-10-06 02:44:02', '', '2025-10-18 23:01:59', NULL, '管理员');
INSERT INTO `sys_user` (`user_id`, `dept_id`, `user_name`, `nick_name`, `user_type`, `email`, `phone_number`, `sex`, `avatar`, `password`, `login_ip`, `login_date`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`, `remark`) VALUES (2, 105, 'hugo', '雨果', '00', 'zhang8680@outlook.com', '18243088680', '1', '', '$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2', '127.0.0.1', '2025-10-06 02:44:02', '0', 'admin', '2025-10-06 02:44:02', '', NULL, NULL, '测试员');
COMMIT;

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post` (
  `user_id` bigint NOT NULL COMMENT '用户id',
  `post_id` bigint NOT NULL COMMENT '岗位id',
  PRIMARY KEY (`user_id`,`post_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户与岗位关联表';

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_post` (`user_id`, `post_id`) VALUES (1, 1);
INSERT INTO `sys_user_post` (`user_id`, `post_id`) VALUES (2, 2);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `user_id` bigint NOT NULL COMMENT '用户id',
  `role_id` bigint NOT NULL COMMENT '角色id',
  PRIMARY KEY (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户和角色关联表';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` (`user_id`, `role_id`) VALUES (1, 1);
INSERT INTO `sys_user_role` (`user_id`, `role_id`) VALUES (2, 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
