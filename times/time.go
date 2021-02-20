package times

// 精度: 秒级
// 例如: 一天 24小时, 一周 168小时, 7天
const (
	// Second 秒, 作为最基础的单位
	Second int = 1

	// Minute 分钟, 表示1分钟持续的秒时长
	Minute = 60 * Second

	// Hour 小时, 表示1小时持续的秒时长
	Hour = 60 * Minute

	// Day 天, 表示1天持续的秒时长
	// 这里不考虑夏令时问题, 泛指国际基础单位制(民用日)所理解的时间
	Day = 24 * Hour

	// Week 周, 表示1周持续的秒时长
	// 这里不考虑夏令时问题, 泛指国际基础单位制(民用日)所理解的时间
	Week = 7 * Day
)
