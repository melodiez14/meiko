//Package cron contains all cron schedule used by reputationapp
package cron

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	rcron "github.com/robfig/cron"
)

var (
	enabled     = false
	errStopCron = errors.New("Stop cron job")
)

func init() {
	enabled, _ = strconv.ParseBool(os.Getenv("CRON_ENABLED"))
}

type cron struct {
	rCron       *rcron.Cron
	listenErrCh chan error
	jobs        []job
}

type job struct {
	name    string
	rule    string
	handler func() error
}

func Init() {
	cronJob := new()
	cronJob.run()
	go func() {
		for {
			err := <-cronJob.ListenError()
			if err == errStopCron {
				//do sometihing if cron stopped
			}
		}
	}()
	//so that when we stop the service, it will signal cronjob to stop
	defer cronJob.stop()
}

func new() *cron {
	c := &cron{
		rCron:       rcron.New(),
		listenErrCh: make(chan error),
	}
	/*
		Example to register a job
		c.Register(Job{
			Name:    "Reputation Auto Score Listing",
			Rule:    "0 0 * * * *",
			Handler: jobs.AutoScoreList,
		})
	*/
	return c
}

func (c *cron) register(j job) {
	c.jobs = append(c.jobs, j)
}

func (c *cron) run() {
	if !enabled {
		log.Println("Crontab is not enabled by setting. please make sure you already run this command: export CRON_ENABLED=true")
	}
	for _, j := range c.jobs {
		c.rCron.AddFunc(j.rule, func(j job) func() {
			return func() {
				log.Printf("Cron [%s] invoked", j.name)
				err := j.handler()
				if err != nil {
					c.listenErrCh <- fmt.Errorf("Cron [%s] Error: %s", j.name, err.Error())
				}
				log.Printf("Cron [%s] executed", j.name)
			}
		}(j))
	}

	c.rCron.Start()
}

func (c *cron) ListenError() <-chan error {
	return c.listenErrCh
}

func (c *cron) stop() {
	c.rCron.Stop()
	c.listenErrCh <- errStopCron
}
