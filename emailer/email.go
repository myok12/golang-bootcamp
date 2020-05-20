package emailer

import "fmt"

func SendEmail(to, url string, failure bool) {
	title := fmt.Sprintf("Site %v is alive", url)
	if failure {
		title = fmt.Sprintf("Site %v is dead", url)
	}
	fmt.Printf(fmt.Sprintf("Would be sending email to %v, with: %v\n", to, title))
}

