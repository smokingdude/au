package main

import (
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func main() {
	result := RawBom{}
	path := "./bom.json"
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(file, &result)
	if err != nil {
		fmt.Println(err)
	}
	for _, comp := range result.Components {
		fmt.Println(apiSearchTrivy(comp.Name, comp.Version, "Node.js"))
	}
}

//
//func getFileForLog() *os.File {
//	if _, err := os.Stat(config.Logdir); errors.Is(err, os.ErrNotExist) {
//		err := os.MkdirAll(config.Logdir, os.ModePerm)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//
//	filename := filepath.Join(config.Logdir, "recdep.log")
//	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0644) // os.Create(filename)
//	if err != nil {
//		log.Fatalf("Failed to open log file %s: %s", filename, err.Error())
//	}
//	return file
//}
//
//func logInit(d bool, f *os.File) *zap.SugaredLogger {
//	pe := zap.NewDevelopmentEncoderConfig()
//	fileEncoder := zapcore.NewConsoleEncoder(pe)
//	consoleEncoder := zapcore.NewConsoleEncoder(pe)
//
//	level := zap.InfoLevel
//	if d {
//		level = zap.DebugLevel
//	}
//
//	core := zapcore.NewTee(
//		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
//		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
//	)
//
//	return zap.New(core, zap.AddCaller()).Sugar()
//}

/*
func searchVulnerabilities(dependencies *dag.DAG, rootProjectLanguage string) map[string]string {
	vulns := make(map[string]string)
	curDepNum := 1
	allDeps := len(dependencies.GetVertices())
	var lock sync.Mutex
	wg := sizedwaitgroup.New(19)
	for _, dep := range dependencies.GetVertices() {
		sugar.Infof("added %d of %d", curDepNum, allDeps)
		curDepNum++
		wg.Add()
		go func(depName string, depVersion string, rootProjectLanguage string) {
			lock.Lock()
			defer lock.Unlock()
			if depName != "" && depVersion != "" && depName != "PROJECT" && depVersion != "1" {
				cves, err := apiSearchTrivy(depName, depVersion, rootProjectLanguage)
				if err != nil {
					sugar.Errorf("Failed to fetch vulnerabilities from Trivy for (%s, %s, %s): %s", depName, depVersion, rootProjectLanguage, err.Error())
					vulns[depName] = "[]"
				} else if len(cves) != 0 {
					sugar.Infof("%s %s %s", depName, depVersion, cves)
					vulns[depName] = cves
				} else {
					vulns[depName] = "[]"
				}
			}
			wg.Done()
		}(dep.(Dependency).Name, dep.(Dependency).Version, rootProjectLanguage)
	}
	wg.Wait()
	lock.Lock()

	sugar.Infof("END OF VULN SEARCH")
	return vulns
}
*/
