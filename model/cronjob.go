package model

type cronjob struct {
	base
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Command     command  `json:"command"`
	Schedule    schedule `json:"schedule"`
}

func GetAllCronjobs() ([]cronjob, error) {

	return []cronjob{
		cronjob{
			base: base{
				Id: 1234,
			},
			Name:        "test cronjob 1",
			Description: "description of test cronjob 1",
			Command:     NewShellCommand("echo hello world"),
			Schedule:    NewSchedule("* * * * *"),
		},
		cronjob{
			base: base{
				Id: 1235,
			},
			Name:        "test cronjob 2",
			Description: "description of test cronjob 2",
			Command:     NewShellCommand("echo this is cronjob 2"),
			Schedule:    NewSchedule("* * * * *"),
		},
	}, nil

}

func GetCronjob(id int64) (cronjob, error) {
	return cronjob{
		base: base{
			Id: id,
		},
		Name:        "test cronjob 2",
		Description: "description of test cronjob but only for one ... todo :)",
		Command:     NewShellCommand("echo test cronjob get"),
		Schedule:    NewSchedule("* * * * *"),
	}, nil
}
