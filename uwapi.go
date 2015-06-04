package uwapi

import (
	"fmt"
	"github.com/jeffail/gabs"
	"io/ioutil"
	"net/http"
	"strings"
)

const URLPrefix = "https://api.uwaterloo.ca/v2/"

/*
Given the API endpoint (url), a call is made to the UW API with the url
and the key. If any errors occur, an empty gabs.container and the error
is returned. Otherwise, the API response is returned as a
*gabs.container .
*/
func callAPI(url string) (*gabs.Container, error) {
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

/*
Given the API key and the full path of the URL (with parameters) as separate
strings, a properly formatted URL is returned.
*/
func formatURL(key string, args ...string) string {
	argsJoined := strings.Join(args, "/")
	return fmt.Sprintf("%s%s.json?key=%s", URLPrefix, argsJoined, key)
}

// Container for FoodServices methods.
type FoodServices struct {
	Key string
}

/*
Get food menu in the current week.
*/
func (f FoodServices) Menu() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", "menu"))
	return response, err
}

/*
Get additional notes regarding food served in the current week.
*/
func (f FoodServices) Notes() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", "notes"))
	return response, err
}

/*
Get a list of all diets.
*/
func (f FoodServices) Diets() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", "diets"))
	return response, err
}

/*
Get a list of all outlets and their unique IDs, names and breakfast/lunch/dinner
meal service indicators.
*/
func (f FoodServices) Outlets() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", "outlets"))
	return response, err
}

/*
Get a list of all outlets and their operating hour data.
*/
func (f FoodServices) Locations() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", "locations"))
	return response, err
}

/*
Get a list of all WatCard locations according to Food Services.
*/
func (f FoodServices) Watcard() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", "watcard"))
	return response, err
}

/*
Get additional announcements regarding food served in the current week.
*/
func (f FoodServices) Announcements() (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", "announcements"))
	return response, err
}

/*
Get a product's nutritional information from a product ID.
*/
func (f FoodServices) Products(product_ID string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", "products", product_ID))
	return response, err
}

/*
Get the food menu given a year and week.
*/
func (f FoodServices) Menu_dated(year, week string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", year, week, "menu"))
	return response, err
}

/*
Get additional notes regarding food given a year and week.
*/
func (f FoodServices) Notes_dated(year, week string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", year, week, "notes"))
	return response, err
}

/*
Get additional announcements regarding food served given a year and week.
*/
func (f FoodServices) Announcements_dated(year, week string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(f.Key, "foodservices", year, week, "announcements"))
	return response, err
}

// Container for Courses methods.
type Courses struct {
	Key string
}

/*
Get all the courses offered under a given subject.
*/
func (c Courses) CoursesBySubject(subject string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "courses", subject))
	return response, err
}

/*
Get all available information for a course, given its ID (e.g. 7407).
*/
func (c Courses) InfoByID(id string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "courses", id))
	return response, err
}

/*
Get the class schedule for a course, given its ID (e.g. 5377).
*/
func (c Courses) ScheduleByID(id string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "courses", id, "schedule"))
	return response, err
}

/*
Get all available information for a given course, given its subject short code
(e.g. PHYS) and catalog number (e.g. 234).
*/
func (c Courses) InfoByCatalogNumber(subject, catnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "courses", subject, catnum))
	return response, err
}

/*
Get the schedule for a given course, given its subject short code (e.g. PHYS)
and catalog number (e.g. 234).
*/
func (c Courses) ScheduleByCatalogNumber(subject, catnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "courses", subject, catnum, "schedule"))
	return response, err
}

/*
Get the prerequisites for a given course, given its subject short code (e.g.
PHYS) and catalog number (e.g. 234).
*/
func (c Courses) PrereqsByCatalogNumber(subject, catnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "courses", subject, catnum, "prerequisites"))
	return response, err
}

/*
Get the exam schedule for a given course, given its subject short code (e.g.
PHYS) and catalog number (e.g. 234).
*/
func (c Courses) ExamScheduleByCatalogNumber(subject, catnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "courses", subject, catnum, "examschedule"))
	return response, err
}

// Container for Events methods.
type Events struct {
	Key string
}

/*
Get a list of the upcoming University of Waterloo events.
*/
func (e Events) All() (*gabs.Container, error) {
	response, err := callAPI(formatURL(e.Key, "events"))
	return response, err
}

/*
Get a list of the upcoming University of Waterloo events at a given site, e.g.
"engineering".
*/
func (e Events) EventsBySite(site string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(e.Key, "events", site))
	return response, err
}

/*
Get information on a specific University of Waterloo event given its site, e.g.
"engineering", and unique ID.
*/
func (e Events) EventsBySiteAndID(site, id string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(e.Key, "events", site, id))
	return response, err
}

/*
Get a list of upcoming holidays.
*/
func (e Events) Holidays() (*gabs.Container, error) {
	response, err := callAPI(formatURL(e.Key, "events", "holidays"))
	return response, err
}

// Container for News methods.
type News struct {
	Key string
}

/*
Get news from all sites, e.g. "engineering".
*/
func (n News) All() (*gabs.Container, error) {
	response, err := callAPI(formatURL(n.Key, "news"))
	return response, err
}

/*
Get news from a given site, e.g. "engineering".
*/
func (n News) NewsBySite(site string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(n.Key, "news", site))
	return response, err
}

/*
Get information on a news item given its site, e.g. "engineering", and its
unique ID.
*/
func (n News) NewsBySiteAndID(site, id string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(n.Key, "news", site, id))
	return response, err
}

// Container for Services methods.
type Services struct {
	Key string
}

/*
Get associated services for a given site, e.g. "engineering".
*/
func (s Services) ServicesBySite(site string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(s.Key, "services", site))
	return response, err
}

// Container for Weather methods.
type Weather struct {
	Key string
}

/*
Get weather details from the University of Waterloo weather station.
*/
func (w Weather) Current() (*gabs.Container, error) {
	response, err := callAPI(formatURL(w.Key, "weather", "current"))
	return response, err
}

// Container for Terms methods.
type Terms struct {
	Key string
}

/*
Get the current, previous and next term's id along with a list of terms in
the past year and the next year.
*/
func (t Terms) List() (*gabs.Container, error) {
	response, err := callAPI(formatURL(t.Key, "terms", "list"))
	return response, err
}

/*
Get the exam schedule for a given term.
*/
func (t Terms) ExamScheduleByTerm(term string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(t.Key, "terms", term, "examschedule"))
	return response, err
}

/*
Get the subject schedule for a given term and subject.
*/
func (t Terms) SubjectScheduleByTerm(term, sub string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(t.Key, "terms", term, sub, "schedule"))
	return response, err
}

/*
Get the class schedule for a given term, subject and catalog number.
*/
func (t Terms) ClassScheduleByTerm(term, sub, catnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(t.Key, "terms", term, sub, catnum, "schedule"))
	return response, err
}

/*
Get the employee information sessions for a given term.
*/
func (t Terms) InfoSessionsByTerm(term string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(t.Key, "terms", term, "infosessions"))
	return response, err
}

// Container for Resources methods.
type Resources struct {
	Key string
}

/*
Get a list of all the tutors available to help for courses.
*/
func (r Resources) Tutors() (*gabs.Container, error) {
	response, err := callAPI(formatURL(r.Key, "resources", "tutors"))
	return response, err
}

/*
Get a list of printers on campus.
*/
func (r Resources) Printers() (*gabs.Container, error) {
	response, err := callAPI(formatURL(r.Key, "resources", "printers"))
	return response, err
}

/*
Get a list of employee information sessions.
*/
func (r Resources) InfoSessions() (*gabs.Container, error) {
	response, err := callAPI(formatURL(r.Key, "resources", "infosessions"))
	return response, err
}

/*
Get a list of geese nests during their spring mating season.
*/
func (r Resources) Goosewatch() (*gabs.Container, error) {
	response, err := callAPI(formatURL(r.Key, "resources", "goosewatch"))
	return response, err
}

// Container for Codes methods.
type Codes struct {
	Key string
}

/*
Get a list of all code lookups and their respective descriptions for
organizations.
*/
func (c Codes) Units() (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "codes", "units"))
	return response, err
}

/*
Get a list of all code lookups for terms.
*/
func (c Codes) Terms() (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "codes", "terms"))
	return response, err
}

/*
Get a list of all code lookups for groups.
*/
func (c Codes) Groups() (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "codes", "groups"))
	return response, err
}

/*
Get a list of all code lookups for subjects.
*/
func (c Codes) Subjects() (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "codes", "subjects"))
	return response, err
}

/*
Get a list of Instructions.
*/
func (c Codes) Instructions() (*gabs.Container, error) {
	response, err := callAPI(formatURL(c.Key, "codes", "instructions"))
	return response, err
}

// Container for Buildings methods.
type Buildings struct {
	Key string
}

/*
Get a list of official building names, codes, numbers, and their lat/long
coordinates.
*/
func (b Buildings) List() (*gabs.Container, error) {
	response, err := callAPI(formatURL(b.Key, "buildings", "list"))
	return response, err
}

/*
Get the official building name, its unique number, and its lat/long coordinates
given a building code.
*/
func (b Buildings) DetailsByCode(code string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(b.Key, "buildings", code))
	return response, err
}

/*
Get the all the courses offered in a given classroom (building code and room
number).
*/
func (b Buildings) CoursesInRoom(code, roomnum string) (*gabs.Container, error) {
	response, err := callAPI(formatURL(b.Key, "buildings", code, roomnum, "courses"))
	return response, err
}

// Container for API methods.
type API struct {
	Key string
}

/*
Get user's API usage statistics.
*/
func (a API) Usage() (*gabs.Container, error) {
	response, err := callAPI(formatURL(a.Key, "api", "usage"))
	return response, err
}

/*
Get all API services available to use.
*/
func (a API) Services() (*gabs.Container, error) {
	response, err := callAPI(formatURL(a.Key, "api", "services"))
	return response, err
}

/*
Get all API endpoint methods available to use.
*/
func (a API) Methods() (*gabs.Container, error) {
	response, err := callAPI(formatURL(a.Key, "api", "methods"))
	return response, err
}

/*
Get information for all API subversions.
*/
func (a API) Versions() (*gabs.Container, error) {
	response, err := callAPI(formatURL(a.Key, "api", "versions"))
	return response, err
}

/*
Get a list of changes made to the API.
*/
func (a API) Changelog() (*gabs.Container, error) {
	response, err := callAPI(formatURL(a.Key, "api", "changelog"))
	return response, err
}

// Container for Server methods.
type Server struct {
	Key string
}

/*
Get time information about the server.
*/
func (s Server) Time() (*gabs.Container, error) {
	response, err := callAPI(formatURL(s.Key, "server", "time"))
	return response, err
}

/*
Get a list of all possible API error codes.
*/
func (s Server) Codes() (*gabs.Container, error) {
	response, err := callAPI(formatURL(s.Key, "server", "codes"))
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

/*
Create a UWAPI struct so that the user doesn't have to do it themself, because
it would definitely be annoying.
*/
func Create(key string) UWAPI {
	return UWAPI{
		FoodServices: FoodServices{key},
		Courses:      Courses{key},
		Events:       Events{key},
		News:         News{key},
		Services:     Services{key},
		Weather:      Weather{key},
		Terms:        Terms{key},
		Resources:    Resources{key},
		Codes:        Codes{key},
		Buildings:    Buildings{key},
		API:          API{key},
		Server:       Server{key},
	}
}
