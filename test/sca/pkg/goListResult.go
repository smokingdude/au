package pkg

import "time"

type GoListResult struct {
	Path      string    `json:"Path"`
	Main      bool      `json:"Main"`
	Dir       string    `json:"Dir"`
	GoMod     string    `json:"GoMod"`
	GoVersion string    `json:"GoVersion"`
	Version   string    `json:"Version"`
	Time      time.Time `json:"Time"`
	Indirect  bool      `json:"Indirect"`
	Error     struct {
		Err string `json:"Err"`
	} `json:"Error"`
}
