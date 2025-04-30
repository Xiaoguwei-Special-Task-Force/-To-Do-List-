package common

import "encoding/json"

// 新增任务分类枚举
type TaskCategory int

const (
    Work TaskCategory = iota + 1
    Personal
    System
    Other
)

func (c TaskCategory) String() string {
    switch c {
    case Work:
        return "工作"
    case Personal:
        return "个人"
    case System:
        return "系统"
    default:
        return "其他"
    }
}

// JSON序列化方法
func (c TaskCategory) MarshalJSON() ([]byte, error) {
    return json.Marshal(c.String())
}