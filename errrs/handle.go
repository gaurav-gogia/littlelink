package errrs

import "log"

func Handle(err error) {
	if err != nil {
		log.Println(err)
	}
}
