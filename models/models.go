package models

import (
	"time"

	"gorm.io/gorm"
)

type PhysicalInfo struct {
	gorm.Model
	ID           uint           `gorm:"primaryKey" json:"id"`                      // 使用uint作为主键，并设置自动生成
	Name         string         `json:"name"`                                      // 测试者姓名
	TestTime     time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"testTime"` // 测试时间，默认为当前时间
	Height       float64        `json:"height"`                                    // 身高（单位：米）
	Weight       float64        `json:"weight"`                                    // 体重（单位：千克）
	StandingJump float64        `json:"standingJump"`                              // 立定跳远（单位：米）
	LongJump     float64        `json:"longJump"`                                  // 双脚连续跳（单位：米）
	SitReach     float64        `json:"sitReach"`                                  // 坐体位前屈（单位：米）
	BalanceBeam  float64        `json:"balanceBeam"`                               // 走平衡木（单位：秒）
	GripStrength float64        `json:"gripStrength"`                              // 握力（单位：千克）
	ObstacleRun  float64        `json:"obstacleRun"`                               // 15绕障碍跑（单位：秒）
	IsDel        gorm.DeletedAt `gorm:"index" json:"isDel"`                        // 软删除字段
}

// 在初始化或更新记录时，如果没有提供TestTime，则使用当前时间
func (p *PhysicalInfo) BeforeCreate(tx *gorm.DB) (err error) {
	if p.TestTime.IsZero() {
		p.TestTime = time.Now()
	}
	return
}

func (p *PhysicalInfo) BeforeUpdate(tx *gorm.DB) (err error) {
	if p.TestTime.IsZero() {
		p.TestTime = time.Now()
	}
	return
}
