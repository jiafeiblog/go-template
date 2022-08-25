package db

import (
	"time"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// Engine orm全局引擎
var Engine *xorm.Engine

type EngineOption func(*xorm.Engine)

// WithShowSQL 是否打印 sql 语句
func WithShowSQL(showSQL bool) EngineOption {
	return func(e *xorm.Engine) {
		e.ShowSQL(showSQL)
	}
}

// WithConnMaxLifetime 设定连接最大存活时间
func WithConnMaxLifetime(d time.Duration) EngineOption {
	return func(e *xorm.Engine) {
		e.SetConnMaxLifetime(d)
	}
}

// WithMaxIdleConns 设定最大空闲连接数
func WithMaxIdleConns(c int) EngineOption {
	return func(e *xorm.Engine) {
		e.SetMaxIdleConns(c)
	}
}

// WithMaxOpenConns 设定最大连接数
func WithMaxOpenConns(c int) EngineOption {
	return func(e *xorm.Engine) {
		e.SetMaxOpenConns(c)
	}
}

// InitDB 初始化全局数据库实例
func InitDB(conn string, options ...func(*xorm.Engine)) {
	// 创建数据库连接
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		panic(err)
	}

	for _, option := range options {
		option(engine)
	}

	// 测试数据库连接
	if err = engine.Ping(); err != nil {
		panic(err)
	}

	mapper := names.GonicMapper{}
	engine.SetTableMapper(mapper)
	engine.SetColumnMapper(mapper)
	Engine = engine
}
