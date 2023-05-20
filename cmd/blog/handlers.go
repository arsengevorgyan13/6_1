package main

import (
	_ "database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Header             []headerData
	TopBlock           []topBlockData
	PageButtons        []pageButtonsData
	FeaturedPostsTitle string
	FeaturedPosts      []featuredPostsData
	RecentPagesTitle   string
	RecentPages        []recentPagesData
	Footer             []footerData
}

type postPage struct {
	PostHeader   []postHeaderData
	PostTopBlock []postTopBlockData
	PostImage    string
	Pharagraphs  []pharagraphsData
	PostFooter   []postFooterData
}

type mainImageData struct {
	Image string
}

type postHeaderData struct {
	Escape               string
	PostHeaderNavButtons []postHeaderNavButtonsData
}

type postHeaderNavButtonsData struct {
	Home       string
	Categories string
	About      string
	Contact    string
}

type postTopBlockData struct {
	Title  string
	Phrase string
}

type pharagraphsData struct {
	FirstPart  string
	SecondPart string
	ThirdPart  string
	FourthPart string
}

type postFooterData struct {
	Background         string
	PostFooterContacts []postFooterContactsData
	PostFooterBottom   []postFooterBottomData
}

type postFooterContactsData struct {
	Title  string
	Button string
}

type postFooterBottomData struct {
	Escape            string
	PostFooterButtons []postFooterButtonsData
}

type postFooterButtonsData struct {
	Home       string
	Categories string
	About      string
	Contact    string
}

type headerData struct {
	Image            string
	Escape           string
	HeaderNavButtons []headerNavButtonsData
}

type headerNavButtonsData struct {
	Home       string
	Categories string
	About      string
	Contact    string
}

type topBlockData struct {
	Title  string
	Text   string
	Button string
}

type pageButtonsData struct {
	PageButtonsList []pageButtonsListData
}

type pageButtonsListData struct {
	Nature      string
	Photography string
	Relaxation  string
	Vacation    string
	Travel      string
	Adventure   string
}

type featuredPostsData struct {
	Image       string `db:"image_url"`
	Title       string `db:"title"`
	Description string `db:"subtitle"`
	AuthorImage string `db:"author_img"`
	AuthorName  string `db:"author_name"`
	Date        string `db:"publish_date"`
}

type recentPagesData struct {
	Image       string `db:"image_url"`
	Title       string `db:"title"`
	Description string `db:"subtitle"`
	AuthorImage string `db:"author_img"`
	AuthorName  string `db:"author_name"`
	Date        string `db:"publish_date"`
}

type footerData struct {
	Background     string
	FooterContacts []footerContactsData
	FooterBottom   []footerBottomData
}

type footerContactsData struct {
	Title  string
	Button string
}

type footerBottomData struct {
	Escape        string
	FooterButtons []footerButtonsData
}

type footerButtonsData struct {
	Home       string
	Categories string
	About      string
	Contact    string
}

func Footer() []footerData {
	return []footerData{
		{
			Background:     "../static/image/footer.svg",
			FooterContacts: FooterContacts(),
			FooterBottom:   FooterBottom(),
		},
	}
}

func FooterContacts() []footerContactsData {
	return []footerContactsData{
		{
			Title:  "Stay in Touch",
			Button: "Submit",
		},
	}
}

func FooterBottom() []footerBottomData {
	return []footerBottomData{
		{
			Escape:        "../static/image/Escape.svg",
			FooterButtons: FooterButtons(),
		},
	}
}

func FooterButtons() []footerButtonsData {
	return []footerButtonsData{
		{
			Home:       "HOME",
			Categories: "CATEGORIES",
			About:      "ABOUT",
			Contact:    "CONTACT",
		},
	}
}

func PostHeader() []postHeaderData {
	return []postHeaderData{
		{
			Escape:               "../static/image/blackescape.svg",
			PostHeaderNavButtons: PostHeaderNavButtons(),
		},
	}
}

func PostHeaderNavButtons() []postHeaderNavButtonsData {
	return []postHeaderNavButtonsData{
		{
			Home:       "HOME",
			Categories: "CATEGORIES",
			About:      "ABOUT",
			Contact:    "CONTACT",
		},
	}
}

func PostTopBlock() []postTopBlockData {
	return []postTopBlockData{
		{
			Title:  "The Road Ahead",
			Phrase: "The road ahead might be paved - it might not be.",
		},
	}
}

func Pharagraphs() []pharagraphsData {
	return []pharagraphsData{
		{
			FirstPart:  "Dark spruce forest frowned on either side the frozen waterway. The trees had been stripped by a recent wind of their white covering of frost, and they seemed to lean towards each other, black and ominous, in the fading light. A vast silence reigned over the land. The land itself was a desolation, lifeless, without movement, so lone and cold that the spirit of it was not even that of sadness. There was a hint in it of laughter, but of a laughter more terrible than any sadness—a laughter that was mirthless as the smile of the sphinx, a laughter cold as the frost and partaking of the grimness of infallibility. It was the masterful and incommunicable wisdom of eternity laughing at the futility of life and the effort of life. It was the Wild, the savage, frozen-hearted Northland Wild.",
			SecondPart: "But there was life, abroad in the land and defiant. Down the frozen waterway toiled a string of wolfish dogs. Their bristly fur was rimed with frost. Their breath froze in the air as it left their mouths, spouting forth in spumes of vapour that settled upon the hair of their bodies and formed into crystals of frost. Leather harness was on the dogs, and leather traces attached them to a sled which dragged along behind. The sled was without runners. It was made of stout birch-bark, and its full surface rested on the snow. The front end of the sled was turned up, like a scroll, in order to force down and under the bore of soft snow that surged like a wave before it. On the sled, securely lashed, was a long and narrow oblong box. There were other things on the sled—blankets, an axe, and a coffee-pot and frying-pan; but prominent, occupying most of the space, was the long and narrow oblong box.",
			ThirdPart:  "In advance of the dogs, on wide snowshoes, toiled a man. At the rear of the sled toiled a second man. On the sled, in the box, lay a third man whose toil was over,—a man whom the Wild had conquered and beaten down until he would never move nor struggle again. It is not the way of the Wild to like movement. Life is an offence to it, for life is movement; and the Wild aims always to destroy movement. It freezes the water to prevent it running to the sea; it drives the sap out of the trees till they are frozen to their mighty hearts; and most ferociously and terribly of all does the Wild harry and crush into submission man—man who is the most restless of life, ever in revolt against the dictum that all movement must in the end come to the cessation of movement.",
			FourthPart: "But at front and rear, unawed and indomitable, toiled the two men who were not yet dead. Their bodies were covered with fur and soft-tanned leather. Eyelashes and cheeks and lips were so coated with the crystals from their frozen breath that their faces were not discernible. This gave them the seeming of ghostly masques, undertakers in a spectral world at the funeral of some ghost. But under it all they were men, penetrating the land of desolation and mockery and silence, puny adventurers bent on colossal adventure, pitting themselves against the might of a world as remote and alien and pulseless as the abysses of space.",
		},
	}
}

func PostFooter() []postFooterData {
	return []postFooterData{
		{
			Background:         "../static/image/footer.svg",
			PostFooterContacts: PostFooterContacts(),
			PostFooterBottom:   PostFooterBottom(),
		},
	}
}

func PostFooterContacts() []postFooterContactsData {
	return []postFooterContactsData{
		{
			Title:  "Stay in Touch",
			Button: "Submit",
		},
	}
}

func PostFooterBottom() []postFooterBottomData {
	return []postFooterBottomData{
		{
			Escape:            "../static/image/Escape.svg",
			PostFooterButtons: PostFooterButtons(),
		},
	}
}

func PostFooterButtons() []postFooterButtonsData {
	return []postFooterButtonsData{
		{
			Home:       "HOME",
			Categories: "CATEGORIES",
			About:      "ABOUT",
			Contact:    "CONTACT",
		},
	}
}

func Header() []headerData {
	return []headerData{
		{
			Image:            "../static/image/dust.png",
			Escape:           "../static/image/Escape.svg",
			HeaderNavButtons: HeaderNavButtons(),
		},
	}
}

func HeaderNavButtons() []headerNavButtonsData {
	return []headerNavButtonsData{
		{
			Home:       "HOME",
			Categories: "CATEGORIES",
			About:      "ABOUT",
			Contact:    "CONTACT",
		},
	}
}

func TopBlock() []topBlockData {
	return []topBlockData{
		{
			Title:  "Lets do it together.",
			Text:   "We travel the world in search of stories. Come along for the ride.",
			Button: "View Latest Posts",
		},
	}
}

func PageButtons() []pageButtonsData {
	return []pageButtonsData{
		{
			PageButtonsList: PageButtonsList(),
		},
	}
}

func PageButtonsList() []pageButtonsListData {
	return []pageButtonsListData{
		{
			Nature:      "Nature",
			Photography: "Photography",
			Relaxation:  "Relaxation",
			Vacation:    "Vacation",
			Travel:      "Travel",
			Adventure:   "Adventure",
		},
	}
}

/*func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Header:        Header(),
		TopBlock:      TopBlock(),
		PageButtons:   PageButtons(),
		FeaturedPosts: FeaturedPosts(),
		RecentPages:   RecentPages(),
		Footer:        Footer(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}*/

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := postPage{
		PostHeader:   PostHeader(),
		PostTopBlock: PostTopBlock(),
		PostImage:    "../static/image/POLAR.png",
		Pharagraphs:  Pharagraphs(),
		PostFooter:   PostFooter(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredposts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Error1", 500)
			log.Println(err)
			return
		}

		recentpages, err := recentPages(db)
		if err != nil {
			http.Error(w, "Error2", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошибки в консоль
			return                                      // Выполнение ф-ии
		}

		data := indexPage{
			Header:             Header(),
			TopBlock:           TopBlock(),
			PageButtons:        PageButtons(),
			FeaturedPostsTitle: "Featured Posts",
			FeaturedPosts:      featuredposts,
			RecentPagesTitle:   "Recent Pages",
			RecentPages:        recentpages,
			Footer:             Footer(),
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func featuredPosts(db *sqlx.DB) ([]featuredPostsData, error) {
	const query = `
		SELECT
			image_url,
			title,
			subtitle,
			author_img,
			author_name,
			publish_date
		FROM
			post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var featuredposts []featuredPostsData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&featuredposts, query) // Делаем запрос в базу данных
	if err != nil {                         // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return featuredposts, nil
}

func recentPages(db *sqlx.DB) ([]recentPagesData, error) {
	const query = `
		SELECT
			image_url,
			title,
			subtitle,
			author_img,
			author_name,
			publish_date
		FROM
			post
		WHERE featured = 0
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var recentpages []recentPagesData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&recentpages, query) // Делаем запрос в базу данных
	if err != nil {                       // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return recentpages, nil
}
