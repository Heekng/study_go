package scrapper

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type extractedJob struct {
	id       string
	title    string
	location string
	summary  string
}

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?search_area=main_b&search_done=y&search_optional_item=n&searchType=search&searchword=" + term + "&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting="
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	for i := 1; i <= totalPages; i++ {
		go getPage(i, baseURL, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "ID", "Title", "Location", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		linkStr := "https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&recommend_ids=eJxNjskRA0EIA6PxX1wS83Ygm38WnvWWGf%2FoaoHI6CK7rvb10jtDaQo%2FSHXzarMvliI8BrszF3b4QcmJvWt8bDR8TnmWaD5haqlsh%2FHsOhwY3HfLMUWCLP56AS4dS9%2Fj6S3zsLFUbT%2BWQrROb%2Fj%2B64fO7gCPzQzcYX4AWaRAHQ%3D%3D&view_type=search&searchword=%EC%9B%B9%EA%B0%9C%EB%B0%9C%EC%9E%90&searchType=search&gz=1&t_ref_content=generic&t_ref=search&paid_fl=n&search_uuid=deca9650-b8c6-47f4-8eb8-910ef7b91df5&rec_idx="
		jobSlice := []string{linkStr + job.id, job.id, job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}

}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := CleanString(card.Find(".job_tit>a").Text())
	location := CleanString(card.Find(".job_condition > span").First().Text())
	sectorText := card.Find(".job_sector").Text()
	index := strings.Index(sectorText, "수정일")
	if index == -1 {
		index = strings.Index(sectorText, "등록일")
	}
	summary := CleanString(sectorText[:index])
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		summary:  summary,
	}
}

// CleanString cleans a string
func CleanString(str string) string {
	// strings.TrimSpace 양쪽의 공백 제거
	// strings.Fields 공백을 모두 나누어 텍스트 배열로 생성
	// strings.Join 배열을 하나의 문자로 합치는데 separator을 포함하여 합친다.
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	/*
		req, rErr := http.NewRequest("GET", baseURL, nil)
		checkErr(rErr)
		req.Header.Add("User-Agent", "Crawler")
		client := &http.Client{}
		res, err := client.Do(req)
	*/
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	//fmt.Println(doc)
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
