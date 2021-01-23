package book

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "Mindf*ck: Cambridge Analytica and the Plot to Break America",
		Author:        "Christopher Wylie",
		YearPublished: 2019,
	},
	Book{
		ID:            2,
		Title:         "The Myth of Sisyphus",
		Author:        "Albert Camus",
		YearPublished: 1942,
	},
	Book{
		ID:            3,
		Title:         "The Sixth Extinction: An Unnatural History",
		Author:        "Elizabeth Kolbert",
		YearPublished: 2014,
	},
	Book{
		ID:            4,
		Title:         "The Stranger",
		Author:        "Albert Camus",
		YearPublished: 1942,
	},
	Book{
		ID:            5,
		Title:         "Sophie's World",
		Author:        "Jostein Gaarder",
		YearPublished: 1991,
	},
	Book{
		ID:            6,
		Title:         "The Solitaire Mystery",
		Author:        "Jostein Gaarder",
		YearPublished: 1990,
	},
	Book{
		ID:            7,
		Title:         "Maya",
		Author:        "Jostein Gaarder",
		YearPublished: 1999,
	},
	Book{
		ID:            8,
		Title:         "The Puppeteer",
		Author:        "Jostein Gaarder",
		YearPublished: 2016,
	},
	Book{
		ID:            9,
		Title:         "Permanent Record",
		Author:        "Edward Snowden",
		YearPublished: 2019,
	},
	Book{
		ID:            10,
		Title:         "A History of God",
		Author:        "Karen Armstrong",
		YearPublished: 1993,
	},
}

func queryCache(id int, cache map[int]Book) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int, cache map[int]Book) (Book, bool) {
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
