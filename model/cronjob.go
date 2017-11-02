package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type cron struct {
	base
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Command     command  `json:"command"`
	Schedule    schedule `json:"schedule"`
}

func GetAllCrons() ([]cron, error) {

	return []cron{
		cron{
			base: base{
				Id: 1234,
			},
			Name:        "test cron 1",
			Description: "description of test cron 1",
			Command:     NewShellCommand("echo hello world"),
			Schedule:    NewSchedule("* * * * *"),
		},
		cron{
			base: base{
				Id: 1235,
			},
			Name:        "test cron 2",
			Description: "description of test cron 2",
			Command:     NewShellCommand("echo this is cron 2"),
			Schedule:    NewSchedule("* * * * *"),
		},
	}, nil

}

func GetCronjob(id int64) (cron, error) {
	return cron{
		base: base{
			Id: id,
		},
		Name:        "test cron 2",
		Description: "description of test cron but only for one ... todo :)",
		Command:     NewShellCommand("echo test cron get"),
		Schedule:    NewSchedule("* * * * *"),
	}, nil
}
