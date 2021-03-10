package osG

/*
#include <windows.h>

int CurrentCPUState(double *dTotalCPUTime, double *dIdleCPUTime)
{
	FILETIME idleTime;
	FILETIME kernelTime;
	FILETIME userTime;
	BOOL res = GetSystemTimes(&idleTime, &kernelTime, &userTime);

	// 得到idleTime,//can use ULARGE_INTEGER get lang time
	INT64 i64Tmp = idleTime.dwHighDateTime;
	i64Tmp = i64Tmp << 32;
	i64Tmp = i64Tmp | idleTime.dwLowDateTime;

	*dIdleCPUTime = (double)i64Tmp;

	// 得到kernelTime
	i64Tmp = kernelTime.dwHighDateTime;
	i64Tmp = i64Tmp << 32;
	i64Tmp = i64Tmp | kernelTime.dwLowDateTime;

	*dTotalCPUTime = (double)i64Tmp;

	// 得到userTime
	i64Tmp = userTime.dwHighDateTime;
	i64Tmp = i64Tmp << 32;
	i64Tmp = i64Tmp | userTime.dwLowDateTime;

	*dTotalCPUTime += (double)i64Tmp;

	return 0;
}
*/
import "C"

import (
	"sync"
	"time"
	"context"
)

var once sync.Once
var cancel context.CancelFunc
var percent int

func InitCPU(){
	once.Do(createCPUInfo)
}

func UninitCPU(){
	cancel()
}

func CPUPrecent() int {
	return percent
}

func createCPUInfo(){
	var ctx context.Context
	ctx,cancel = context.WithCancel(context.Background())
	go run(ctx)
}

func run(ctx context.Context) {
	tick := time.NewTicker(1*time.Second)
	var totallast C.double
	var idlelast C.double
	var total C.double
	var idle C.double
	start := true
	for{
		select{
		case <- ctx.Done():
			return
		case <-tick.C:
			if start{
				start = false
				C.CurrentCPUState(&totallast,&idlelast)
			}else{
				C.CurrentCPUState(&total,&idle)
				percent = int(1+100*((total-totallast) - (idle - idlelast))/(total-totallast))
				totallast = total
				idlelast = idle
			}
		}
	}
}
