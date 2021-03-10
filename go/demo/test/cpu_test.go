package test

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

func testCPUPercent(percpu bool) {
	numcpu := runtime.NumCPU()
	testCount := 3

	if runtime.GOOS != "windows" {
		testCount = 100
		v, err := cpu.Percent(time.Millisecond, percpu)
		if err != nil {
			log.Fatalf("error %v", err)
		}
		// Skip CircleCI which CPU num is different
		if os.Getenv("CIRCLECI") != "true" {
			if (percpu && len(v) != numcpu) || (!percpu && len(v) != 1) {
				log.Fatalf("wrong number of entries from CPUPercent: %v", v)
			}
		}
	}
	for i := 0; i < testCount; i++ {
		duration := time.Duration(10) * time.Microsecond
		v, err := cpu.Percent(duration, percpu)
		if err != nil {
			log.Fatalf("error %v", err)
		}
		for _, percent := range v {
			// Check for slightly greater then 100% to account for any rounding issues.
			if percent < 0.0 || percent > 100.0001*float64(numcpu) {
				log.Fatalf("CPUPercent value is invalid: %f", percent)
			}
			log.Printf("CPUPercent value is invalid: %f", percent)
		}
	}
}

func TestCPUPercent(t *testing.T) {
	testCPUPercent(false)
}

func TestCPUPercentPerCpu(t *testing.T) {
	testCPUPercent(true)
}

func TestCpu_times(t *testing.T) {
	v, err := cpu.Times(false)
	if err != nil {
		t.Errorf("error %v", err)
	}
	if len(v) == 0 {
		t.Error("could not get CPUs ", err)
	}
	empty := cpu.TimesStat{}
	for _, vv := range v {
		if vv == empty {
			t.Errorf("could not get CPU User: %v", vv)
		}
		log.Printf("TimesStat value is: %v", vv)
	}

	// test sum of per cpu stats is within margin of error for cpu total stats
	cpuTotal, err := cpu.Times(false)
	if err != nil {
		t.Errorf("error %v", err)
	}
	if len(cpuTotal) == 0 {
		t.Error("could not get CPUs ", err)
	}
	perCPU, err := cpu.Times(true)
	if err != nil {
		t.Errorf("error %v", err)
	}
	if len(perCPU) == 0 {
		t.Error("could not get CPUs ", err)
	}
	var perCPUUserTimeSum float64
	var perCPUSystemTimeSum float64
	var perCPUIdleTimeSum float64
	for _, pc := range perCPU {
		perCPUUserTimeSum += pc.User
		perCPUSystemTimeSum += pc.System
		perCPUIdleTimeSum += pc.Idle
	}
	margin := 2.0
	assert.InEpsilon(t, cpuTotal[0].User, perCPUUserTimeSum, margin)
	assert.InEpsilon(t, cpuTotal[0].System, perCPUSystemTimeSum, margin)
	assert.InEpsilon(t, cpuTotal[0].Idle, perCPUIdleTimeSum, margin)
	log.Println(cpuTotal[0].User, perCPUUserTimeSum, margin)
	log.Println(cpuTotal[0].System, perCPUSystemTimeSum, margin)
	log.Println(cpuTotal[0].Idle, perCPUIdleTimeSum, margin)
	result := float32(cpuTotal[0].User+cpuTotal[0].System) / float32(cpuTotal[0].User+cpuTotal[0].System+cpuTotal[0].Idle)
	log.Println(result)
}

func TestCpu_counts(t *testing.T) {
	numcpu := runtime.NumCPU()
	v, err := cpu.Counts(true)
	if err != nil {
		t.Errorf("error %v", err)
	}
	if v == 0 {
		t.Errorf("could not get CPU counts: %v", v)
	}
	log.Printf("cpuCount value is : %v-%v", v, numcpu)
}

func TestCPUTimeStat_String(t *testing.T) {
	v := cpu.TimesStat{
		CPU:    "cpu0",
		User:   100.1,
		System: 200.1,
		Idle:   300.1,
	}
	e := `{"cpu":"cpu0","user":100.1,"system":200.1,"idle":300.1,"nice":0.0,"iowait":0.0,"irq":0.0,"softirq":0.0,"steal":0.0,"guest":0.0,"guestNice":0.0}`
	if e != fmt.Sprintf("%v", v) {
		t.Errorf("CPUTimesStat string is invalid: %v", v)
	}
	log.Printf("cpu TimesStat value is : %v", v)
}

func TestCpuInfo(t *testing.T) {
	v, err := cpu.Info()
	if err != nil {
		t.Errorf("error %v", err)
	}
	if len(v) == 0 {
		t.Errorf("could not get CPU Info")
	}
	for _, vv := range v {
		if vv.ModelName == "" {
			t.Errorf("could not get CPU Info: %v", vv)
		}
		log.Printf("cpu Info value is : %+v", vv)
	}
}
