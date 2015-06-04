package main

import (
	"fmt"
	"github.com/jeffail/gabs"
	"io/ioutil"
	"net/http"
	"strings"
)

const URLPrefix = "https://api.uwaterloo.ca/v2/"

func callAPI(url string) (*gabs.Container, error) {
	/*
		Given the API key and the API endpoint (url), a call is made to the UW
		API with the url and the key. If any errors occur, an empty
		gabs.container and the error is returned. Otherwise, the API response
		is returned as a gabs.container
	*/
	var empty *gabs.Container

	// send the get request to the UW API...
	res, err := http.Get(url)
	if err != nil {
		return empty, err
	}
	defer res.Body.Close()

	// read the response (it's a byte array)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return empty, err
	}

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return empty, err
	}

	return jsonParsed, nil
}

func formatURL(key string, args ...string) string {
	argsJoined := strings.Join(args, "/")
	return fmt.Sprintf("%s%s.json?key=%s", URLPrefix, argsJoined, key)
}

// FOOD SERVICES ===========================
type FoodServices struct {
	key string
}

func (f FoodServices) Menu() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", "menu"))
	return response, err
}

func (f FoodServices) Notes() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", "notes"))
	return response, err
}

func (f FoodServices) Diets() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", "diets"))
	return response, err
}

func (f FoodServices) Outlets() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", "outlets"))
	return response, err
}

func (f FoodServices) Locations() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", "locations"))
	return response, err
}

func (f FoodServices) watcard() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", "watcard"))
	return response, err
}

func (f FoodServices) announcements() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", "announcements"))
	return response, err
}

func (f FoodServices) products(product_ID string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", "products", product_ID))
	return response, err
}

func (f FoodServices) menu_dated(year, week string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", year, week, "menu"))
	return response, err
}

func (f FoodServices) notes_dated(year, week string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", year, week, "notes"))
	return response, err
}

func (f FoodServices) announcements_dated(year, week string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.key, "foodservices", year, week, "announcements"))
	return response, err
}

// COURSES =================================
type Courses struct {
	key string
}

func (c Courses) courses_by_subject(subject string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "courses", subject))
	return response, err
}

func (c Courses) info_by_id(course_id string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "courses", course_id))
	return response, err
}

func (c Courses) schedule_by_classnum(classnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "courses", classnum, "schedule"))
	return response, err
}

func (c Courses) info_by_catnum(subject, catnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "courses", subject, catnum))
	return response, err
}

func (c Courses) schedule_by_catnum(subject, catnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "courses", subject, catnum, "schedule"))
	return response, err
}

func (c Courses) prereqs_by_catnum(subject, catnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "courses", subject, catnum, "prerequisites"))
	return response, err
}

func (c Courses) exam_schedule_by_catnum(subject, catnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "courses", subject, catnum, "examschedule"))
	return response, err
}

// EVENTS ==================================
type Events struct {
	key string
}

func (e Events) all() (*gabs.Container, error) {
	response, err := callAPI(formatURL(e.key, "events"))
	return response, err
}

func (e Events) events_by_site(site string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(e.key, "events", site))
	return response, err
}

func (e Events) events_by_site_and_id(site, id string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(e.key, "events", site, id))
	return response, err
}

func (e Events) holidays() (*gabs.Container, error) {
	response, err := callAPI(formatURL(e.key, "events", "holidays"))
	return response, err
}

// NEWS ====================================
type News struct {
	key string
}

func (n News) all() (*gabs.Container, error) {
	response, err := callAPI(formatURL(n.key, "news"))
	return response, err
}

func (n News) news_by_site(site string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(n.key, "news", site))
	return response, err
}

func (n News) news_by_site_and_id(site, id string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(n.key, "news", site, id))
	return response, err
}

// SERVICES ================================
type Services struct {
	key string
}

func (s Services) services_by_site(site string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(s.key, "services", site))
	return response, err
}

// WEATHER =================================
type Weather struct {
	key string
}

func (w Weather) current() (*gabs.Container, error) {
	response, err := callAPI(formatURL(w.key, "weather", "current"))
	return response, err
}

// TERMS ===================================
type Terms struct {
	key string
}

func (t Terms) list() (*gabs.Container, error) {
	response, err := callAPI(formatURL(t.key, "terms", "list"))
	return response, err
}

func (t Terms) exam_schedule_by_term(term string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(t.key, "terms", term, "examschedule"))
	return response, err
}

func (t Terms) subject_schedule_by_term(term, sub string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(t.key, "terms", term, sub, "schedule"))
	return response, err
}

func (t Terms) class_schedule_by_term(term, sub, catnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(t.key, "terms", term, sub, catnum, "schedule"))
	return response, err
}

func (t Terms) info_sessions_by_term(term string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(t.key, "terms", term, "infosessions"))
	return response, err
}

// RESOURCES ===============================
type Resources struct {
	key string
}

func (r Resources) tutors() (*gabs.Container, error) {
	response, err := callAPI(formatURL(r.key, "resources", "tutors"))
	return response, err
}

func (r Resources) printers() (*gabs.Container, error) {
	response, err := callAPI(formatURL(r.key, "resources", "printers"))
	return response, err
}

func (r Resources) infosessions() (*gabs.Container, error) {
	response, err := callAPI(formatURL(r.key, "resources", "infosessions"))
	return response, err
}

func (r Resources) goosewatch() (*gabs.Container, error) {
	response, err := callAPI(formatURL(r.key, "resources", "goosewatch"))
	return response, err
}

// CODES ===================================
type Codes struct {
	key string
}

func (c Codes) units() (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "codes", "units"))
	return response, err
}

func (c Codes) terms() (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "codes", "terms"))
	return response, err
}

func (c Codes) groups() (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "codes", "groups"))
	return response, err
}

func (c Codes) subjects() (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "codes", "subjects"))
	return response, err
}

func (c Codes) instructions() (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.key, "codes", "instructions"))
	return response, err
}

// BUILDINGS ===============================
type Buildings struct {
	key string
}

func (b Buildings) list() (*gabs.Container, error) {
	response, err := callAPI(formatURL(b.key, "buildings", "list"))
	return response, err
}

func (b Buildings) details_by_code(code string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(b.key, "buildings", code))
	return response, err
}

func (b Buildings) courses_in_room(code, roomnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(b.key, "buildings", code, roomnum, "courses"))
	return response, err
}

// API =====================================
type API struct {
	key string
}

func (a API) usage() (*gabs.Container, error) {
	response, err := callAPI(formatURL(a.key, "api", "usage"))
	return response, err
}

func (a API) services() (*gabs.Container, error) {
	response, err := callAPI(formatURL(a.key, "api", "services"))
	return response, err
}

func (a API) methods() (*gabs.Container, error) {
	response, err := callAPI(formatURL(a.key, "api", "methods"))
	return response, err
}

func (a API) versions() (*gabs.Container, error) {
	response, err := callAPI(formatURL(a.key, "api", "versions"))
	return response, err
}

func (a API) changelog() (*gabs.Container, error) {
	response, err := callAPI(formatURL(a.key, "api", "changelog"))
	return response, err
}

// SERVER ==================================
type Server struct {
	key string
}

func (s Server) time() (*gabs.Container, error) {
	response, err := callAPI(formatURL(s.key, "server", "time"))
	return response, err
}

func (s Server) codes() (*gabs.Container, error) {
	response, err := callAPI(formatURL(s.key, "server", "codes"))
	return response, err
}

// "wrapper" object
type UWAPI struct {
	FoodServices
	Courses
	Events
	News
	Services
	Weather
	Terms
	Resources
	Codes
	Buildings
	API
	Server
}

func main() {
	API_KEY := "YOUR_API_KEY_HERE"
	uw := UWAPI{
		FoodServices: FoodServices{key: API_KEY},
		Courses:      Courses{key: API_KEY},
		Events:       Events{key: API_KEY},
		News:         News{key: API_KEY},
		Services:     Services{key: API_KEY},
		Weather:      Weather{key: API_KEY},
		Terms:        Terms{key: API_KEY},
		Resources:    Resources{key: API_KEY},
		Codes:        Codes{key: API_KEY},
		Buildings:    Buildings{key: API_KEY},
		API:          API{key: API_KEY},
		Server:       Server{key: API_KEY},
	}
}
