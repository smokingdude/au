package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RawBom struct {
	Components []struct {
		Type     string `json:"type"`
		BomRef   string `json:"bomRef"`
		Name     string `json:"name"`
		Version  string `json:"version"`
		Scope    string `json:"scope"`
		Purl     string `json:"purl"`
		Hashes   string `json:"hashes"`
		Licences []struct {
			Id string `json:"id"`
		} `json:"licences"`
	} `json:"components"`
}

type TrivyPacketInfoV2 struct {
	Vulnerabilities []struct {
		CveId            string   `json:"cveId"`
		CweIds           []string `json:"cweIds"`
		Name             string   `json:"name"`
		CvssScore        float32  `json:"cvssScore"`
		CvssAttackVector string   `json:"cvssAttackVector"`
	} `json:"vulnerabilities"`
}

type VulndbStruct struct {
	Id    string
	CweId string
	Table string
	Cvss  float32
}

//		A function to get ids from VulnDB API.
//
//		\param[in] pkgname - package name
//		\param[in] pkgversion - package version
//		\return An string array

func apiSearchTrivy(pkgName string, pkgVersion string, language string) (string, error) {
	pkgName = url.QueryEscape(pkgName)
	pkgVersion = url.QueryEscape(pkgVersion)
	packetManager := ""
	switch language {
	case "Node.js":
		packetManager = "npm"
	case ".NET":
		packetManager = "nuget"
	case "Java":
		packetManager = "maven"
		if strings.Contains(pkgName, ":") {
			pkgName = strings.Replace(pkgName, ":", "/", 1)
		}
		if strings.Contains(pkgName, "%3A") {
			pkgName = strings.Replace(pkgName, "%3A", "/", 1)
		}
		if strings.Contains(pkgName, "%3a") {
			pkgName = strings.Replace(pkgName, "%3a", "/", 1)
		}
	case "Go":
		packetManager = "gomod"
	case "Python":
		packetManager = "pip"
	default:
		return "", errors.New(fmt.Sprintf("unable to find packet manager for %s", language))
	}

	reqUrl, _ := url.Parse(config.VulnDB.Baseurl)
	reqUrl.Path = fmt.Sprintf("v1/%s/%s/%s", packetManager, pkgName, pkgVersion)
	sugar.Debugf("HTTP GET %s", reqUrl)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = getTLSClientConfig()
	resp, err := http.Get(reqUrl.String())
	if err != nil {
		return "", errors.New(fmt.Sprintf("http get error: %s", err.Error()))
	}
	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("status code not 200, but %d", resp.StatusCode))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to read response body: %s", err.Error()))
	}
	defer resp.Body.Close()
	sugar.Debugf("Response body: %s", string(body))

	var result TrivyPacketInfoV2
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to unmarshal response body: %s", err.Error()))
	}

	oldStyleVulns := make([]VulndbStruct, 0)

	for _, vuln := range result.Vulnerabilities {
		oldStyleVulns = append(oldStyleVulns, VulndbStruct{vuln.CveId, "", "trivy:" + packetManager, vuln.CvssScore})
	}
	res, err := json.Marshal(oldStyleVulns)
	if err != nil {
		return "", errors.New(fmt.Sprintf("unable to marshal vulns: %s", err.Error()))
	}
	return string(res), nil
}
