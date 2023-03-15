package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"fmt"
	"time"
)

// 消息对象：主线程跑圈，主线程每循环一次发送一次
type SysMsgRun struct {
	MsgB
}

var SysMsg_Run = &SysMsgRun{}

type timeInfo struct {
	Sec    int
	Minute int
	Hour   int
	Week   time.Weekday
	Day    int
	Month  time.Month
	Year   int
	// Unix Sec
	Time int64
	// Unix Day
	Date int64
}

func (this *timeInfo) UnixMilli() int64 {
	return CurTick
}

func (this *timeInfo) TimeToStr() string {
	return fmt.Sprintf("%02d:%02d:%02d", this.Hour, this.Minute, this.Sec)
}

func (this *timeInfo) DateToStr() string {
	return fmt.Sprintf("%02d/%02d/%02d", this.Day, this.Month, this.Year)
}

func (this *timeInfo) DateTimeToStr() string {
	return fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d", this.Day, this.Month, this.Year, this.Hour, this.Minute, this.Sec)
}

// 消息对象：秒跳变
type SysMsgSecChg struct {
	MsgB
	TimeInfo *timeInfo
}

// 消息对象：分跳变
type SysMsgMinuteChg struct {
	MsgB
	TimeInfo *timeInfo
}

// 消息对象：时跳变
type SysMsgHourChg struct {
	MsgB
	TimeInfo *timeInfo
}

// 消息对象：日跳变
type SysMsgDayChg struct {
	MsgB
	TimeInfo *timeInfo
}

// 消息对象：周跳变
type SysMsgWeekChg struct {
	MsgB
	TimeInfo *timeInfo
}

// 消息对象：月跳变
type SysMsgMonthChg struct {
	MsgB
	TimeInfo *timeInfo
}

// 消息对象：月跳变
type SysMsgYearChg struct {
	MsgB
	TimeInfo *timeInfo
}

var CurTime = time.Now()
var CurTick = CurTime.UnixMilli()
var TimeInfo = timeToInfo(&CurTime)

func timeToInfo(now *time.Time) *timeInfo {
	ret := &timeInfo{}

	ret.Sec = now.Second()
	ret.Minute = now.Minute()
	ret.Hour = now.Hour()
	ret.Week = now.Weekday()
	ret.Day, ret.Month, ret.Year = now.Date()

	ret.Time = now.Unix()
	const secondsPerDay = 60 * 60 * 24
	ret.Date = ret.Time / secondsPerDay

	return ret
}

type timeEvent struct {
	ModuleST
	saveSec    int
	saveMinute int
	saveHour   int
	saveWeek   time.Weekday
	saveDay    int
	saveMonth  time.Month
	saveYear   int
}

func (this *timeEvent) init() {
	this.saveSec = CurTime.Second()
	this.saveMinute = CurTime.Minute()
	this.saveHour = CurTime.Hour()
	this.saveWeek = CurTime.Weekday()
	this.saveDay, this.saveMonth, this.saveYear = CurTime.Date()
}

func (this *timeEvent) Run() {
	if this.saveSec == 0 {
		this.init()
		return
	}

	time.Sleep(time.Second)
	if !switchTimeMsg {
		return
	}

	now := time.Now()
	newInfo := timeToInfo(&now)
	msgs := []IMsg(nil)

	TimeInfo.Sec = newInfo.Sec
	msgs = append(msgs, &SysMsgSecChg{TimeInfo: newInfo})

	if TimeInfo.Minute != newInfo.Minute {
		TimeInfo.Minute = newInfo.Minute
		msgs = append(msgs, &SysMsgMinuteChg{TimeInfo: newInfo})
	}
	if TimeInfo.Hour != newInfo.Hour {
		TimeInfo.Hour = newInfo.Hour
		msgs = append(msgs, &SysMsgHourChg{TimeInfo: newInfo})
	}

	if TimeInfo.Day != newInfo.Day {
		TimeInfo.Day = newInfo.Day
		msgs = append(msgs, &SysMsgDayChg{TimeInfo: newInfo})
	}
	if TimeInfo.Week != newInfo.Week {
		TimeInfo.Week = newInfo.Week
		msgs = append(msgs, &SysMsgWeekChg{TimeInfo: newInfo})
	}
	if TimeInfo.Month != newInfo.Month {
		TimeInfo.Month = newInfo.Month
		msgs = append(msgs, &SysMsgMonthChg{TimeInfo: newInfo})
	}
	if TimeInfo.Year != newInfo.Year {
		TimeInfo.Year = newInfo.Year
		msgs = append(msgs, &SysMsgYearChg{TimeInfo: newInfo})
	}

	for _, msg := range msgs {
		this.SendMsg(NullCtx, msg)
	}
}
