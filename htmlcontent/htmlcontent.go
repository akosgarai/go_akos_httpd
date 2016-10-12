// Package htmlcontent provides a simple html page.
package htmlcontent

import (
	"log"
	"os"
)

type Htmlpage struct {
	content string
	logger  *log.Logger
}

func New() *Htmlpage {
	return &Htmlpage{
		content: "<html>" +
			"<head>" +
			"<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css\">" +
			"</head>" +
			"<body>" +
			"<h1>It is working man!</h1><input type=\"button\" name=\"btn\" value=\"Click\" />" +
			"</body>" +
			"</html>",
		logger: log.New(os.Stderr, "[htmlpage NEW] ", log.LstdFlags),
	}
}

func (p *Htmlpage) Getpage() string {
	p.logger.Println("Page opened")
	return p.content
}
