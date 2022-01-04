package lib

import (
	"github.com/binganao/dailyCVE/model"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func GetCVE() {
	cveUrl := "https://cassandra.cerias.purdue.edu/CVE_changes/today.html"
	resp, err := http.Get(cveUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(respBody))
	re, _ := regexp.Compile(`<A HREF = '(.*?)'>(.*?)</A><br />\n`)
	cves := re.FindAllStringSubmatch(string(respBody), -1)
	re, _ = regexp.Compile(`<HTML><BODY><BR>date: (.*?)<BR>New entries:<br />\n`)
	date := re.FindAllStringSubmatch(string(respBody), -1)[0][1]

	log.Println("今天是 " + date + " 正在搜集新的 CVE...")
	for _, cve := range cves {
		cveName := "CVE-" + cve[2]
		if checkCVEExist(cveName) {
			log.Println("发现CVE项目 " + cveName + " 已重复，跳过..")
			continue
		}
		cveUrl := cve[1]
		cveResp, _ := http.Get(cveUrl)
		defer cveResp.Body.Close()
		cveBody, _ := ioutil.ReadAll(cveResp.Body)
		//fmt.Println(string(cveBody))
		re, _ = regexp.Compile(`<tr>
		<td colspan="2">(.*?)

</td>`)
		desc := re.FindAllStringSubmatch(string(cveBody), -1)[0][1]
		//break
		saveCVE(cveName, cveUrl, date, desc)
	}
}

func checkCVEExist(name string) bool {
	var cve []model.CVE
	db := GetDB()
	db.Where("name = ?", name).First(&cve)
	if len(cve) > 0 {
		return true
	}
	return false
}

func QueryCVE(date string) []model.CVE {
	var cves []model.CVE
	db := GetDB()
	db.Where("date = ?", date).Find(&cves)
	return cves
}

func saveCVE(name, url, date, desc string) {
	log.Println("发现新的CVE项目 " + name + " ,正在写入数据库...")
	db := GetDB()
	cve := model.CVE{Name: name, Url: url, Date: date, Description: desc}
	db.Create(&cve)
}
