package delivery

import (
	"fmt"
	"html/template"
	"net/http"
	"reegle/internal/entity"
	"reegle/internal/router"
	"reegle/internal/usecase/search"
	"reegle/internal/usecase/search/repo"
	"reegle/pkg/dict"
	"reegle/pkg/parser"
	"regexp"
	"strconv"
	"time"
)

func NewSearchDelivery(router *router.Router, db *dict.WordDB) {
	uc := search.NewUseCase(repo.NewSearchRepo(db))

	router.GET("/search", func(w http.ResponseWriter, r *http.Request) {
		word := r.URL.Query().Get("regex")
		offset := r.URL.Query().Get("offset")
		limit := r.URL.Query().Get("limit")

		res := entity.ServerResponse{}

		limitInt := 20
		offsetInt := 0

		if limit != "" && regexp.MustCompile(`^\d+$`).MatchString(limit) {
			limitInt, _ = strconv.Atoi(limit)
		}
		if offset != "" && regexp.MustCompile(`^\d+$`).MatchString(offset) {
			offsetInt, _ = strconv.Atoi(offset)
		}

		timeStart := time.Now()

		result, err := uc.Search(r.Context(), word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		timeEnd := time.Since(timeStart)
		res.Total = len(result)
		if offsetInt >= 0 && offsetInt*20 < len(result) && limitInt > 0 && limitInt < len(result) {
			res.Pages = len(result) / limitInt
			res.Page = (offsetInt / limitInt) + 1
			result = result[offsetInt*20 : limitInt*(offsetInt+1)]
		} else {
			res.Pages = 1
			res.Page = 1
			offsetInt = 0
		}

		res.Page = offsetInt
		for i := 0; i < 3; i++ {
			next := fmt.Sprintf("regex=%s&offset=%d&limit=%d", word, res.Page+i, 20)
			next = fmt.Sprintf(" <li class=\"page-item\"><a class=\"page-link\" href=\"/search?%s\">%d</a></li>", next, res.Page+i+1)
			res.PagesContent = append(res.PagesContent, template.HTML(next))
		}

		res.HasPrev = res.Page > 1
		res.HasNext = res.Page < res.Pages
		if res.HasNext {
			next := fmt.Sprintf("regex=%s&offset=%d&limit=%d", word, res.Page+1, 20)
			next = fmt.Sprintf(" <li class=\"page-item\"><a class=\"page-link\" href=\"/search?%s\">Next</a></li>", next)
			res.PagesContent = append(res.PagesContent, template.HTML(next))
		}
		res.NextPage = res.Page + 1
		res.PreviousPage = res.Page - 1

		res.Regex = word
		res.Time = timeEnd.Seconds()
		res.Results = result
		res.StatusCode = http.StatusOK

		w.Write(parser.ParseTemplate("templates/results.template.html", res))
	})
}
