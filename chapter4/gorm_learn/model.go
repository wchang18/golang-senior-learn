package gorm_learn

import "time"

/*
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uid` varchar(50) NOT NULL COMMENT 'uuid',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '电子邮件',
  `telephone` varchar(30) NOT NULL DEFAULT '' COMMENT '电话',
  `del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：默认0, 1为删除',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uid` (`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户表';
*/

type User struct {
	ID          int64         `gorm:"column:id;primary_key;"`
	Uid         string        `gorm:"column:uid;unique_index"`
	Username    string        `gorm:"column:username"`
	Email       string        `gorm:"column:email"`
	Telephone   string        `gorm:"column:telephone"`
	Del         int8          `gorm:"column:del"`
	CreatedAt   time.Time     `gorm:"column:create_time"`
	UpdatedAt   time.Time     `gorm:"column:update_time"`
	UserCompany UserCompany   `gorm:"foreignKey:UserId"`
	UserAddress []UserAddress `gorm:"foreignKey:UserId"`
}

func (*User) TableName() string {
	return "user"
}

/*
CREATE TABLE `user_address` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `userId` bigint(20) NOT NULL COMMENT '用户id',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户地址';
*/

type UserCompany struct {
	ID          int64     `gorm:"column:id;primary_key;"`
	UserId      int64     `gorm:"column:userId"`
	CompanyName string    `gorm:"column:company_name"`
	CreatedAt   time.Time `gorm:"column:create_time"`
	UpdatedAt   time.Time `gorm:"column:update_time"`
}

func (*UserCompany) TableName() string {
	return "user_company"
}

/*
CREATE TABLE `user_company` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `userId` bigint(20) NOT NULL COMMENT '用户id',
  `company_name` varchar(255) NOT NULL DEFAULT '' COMMENT '公司名',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户公司';
*/

type UserAddress struct {
	ID        int64     `gorm:"column:id;primary_key;"`
	UserId    int64     `gorm:"column:userId"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
}

func (*UserAddress) TableName() string {
	return "user_address"
}
