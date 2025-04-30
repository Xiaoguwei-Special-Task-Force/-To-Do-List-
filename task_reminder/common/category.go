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


// 带颜色的分类显示
func ColorCategory(c TaskCategory) string {
    colors := map[TaskCategory]string{
        Work:     "\033[34m工作\033[0m",  // 蓝色
        Personal: "\033[32m个人\033[0m",  // 绿色
        System:   "\033[31m系统\033[0m",  // 红色
        Other:    "\033[33m其他\033[0m",  // 黄色
    }
    return colors[c]
}