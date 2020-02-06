package function

const (
	// 每分钟60秒
	SecondsPerMinute = 60
	// 小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// 天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

// 传入秒
func ResolveTime(seconds int) (day, hour, minute int)  {  // day int, hour int, minute int
	//fmt.Println("解析为时间单位...")
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}

