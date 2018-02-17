package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Restaurant struct {
	Id          string
	Name        string
	Longitude   float64
	Latitude    float64
	Rating      float32
	CuisineType string
}

var restaurants []Restaurant

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi There</h1>")
}

func GetNearbyRestaurants(w http.ResponseWriter, req *http.Request) {

	var radius float64 = 5000 //radius 5km
	var result []Restaurant

	params := mux.Vars(req)
	longitude, err := strconv.ParseFloat(params["longitude"], 64)
	latitude, err := strconv.ParseFloat(params["latitude"], 64)

	if err == nil {
		fmt.Println("Find nearby restaurant in radius", radius/1000, "km")

		for idx, restaurant := range restaurants {
			var dist = Distance(latitude, longitude, restaurant.Latitude, restaurant.Longitude)

			if dist < radius {
				result = append(result, restaurant)
				fmt.Println((idx + 1), restaurant.Name, dist, "meters")
			} else {
				fmt.Println((idx + 1), restaurant.Name, dist, "meters <outside radius>")
			}

		}

	} else {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(result)
}

func PopulateData() {
	restaurants = append(restaurants, Restaurant{Id: "1", Name: "McD", Longitude: -6.237714, Latitude: 106.765751, Rating: 5, CuisineType: "Fast Food"})
	restaurants = append(restaurants, Restaurant{Id: "2", Name: "KFC", Longitude: -6.232738, Latitude: 106.780292, Rating: 4.5, CuisineType: "Fast Food"})
	restaurants = append(restaurants, Restaurant{Id: "3", Name: "Sushi Lover", Longitude: -6.239784, Latitude: 106.780709, Rating: 2, CuisineType: "Japanese"})
	restaurants = append(restaurants, Restaurant{Id: "4", Name: "Super Bento", Longitude: -6.244634, Latitude: 106.783794, Rating: 3, CuisineType: "Japanese"})
	restaurants = append(restaurants, Restaurant{Id: "5", Name: "Ramen Tsubasa", Longitude: -6.251845, Latitude: 106.791383, Rating: 5, CuisineType: "Japanese"})
	restaurants = append(restaurants, Restaurant{Id: "6", Name: "Warteg Matahari", Longitude: -6.256734, Latitude: 106.789918, Rating: 5, CuisineType: "Indonesian"})
	restaurants = append(restaurants, Restaurant{Id: "7", Name: "Warung war", Longitude: -6.259425, Latitude: 106.789269, Rating: 5, CuisineType: "Indonesian"})
	restaurants = append(restaurants, Restaurant{Id: "8", Name: "Indonesian Taste Cafe", Longitude: -6.259317, Latitude: 106.781799, Rating: 5, CuisineType: "Indonesian"})
	restaurants = append(restaurants, Restaurant{Id: "9", Name: "Spoon n Spoon", Longitude: -6.266061, Latitude: 106.783639, Rating: 3.5, CuisineType: "Indonesian"})
	restaurants = append(restaurants, Restaurant{Id: "10", Name: "Your Thai", Longitude: -6.269002, Latitude: 106.765812, Rating: 4, CuisineType: "Thai"})
	restaurants = append(restaurants, Restaurant{Id: "11", Name: "Nong Kitchen", Longitude: -6.271334, Latitude: 106.768049, Rating: 2, CuisineType: "Thai"})
	restaurants = append(restaurants, Restaurant{Id: "12", Name: "Baba's Kitchen", Longitude: -6.240161, Latitude: 106.796161, Rating: 1.5, CuisineType: "Western"})
	restaurants = append(restaurants, Restaurant{Id: "13", Name: "Steak together", Longitude: -6.238798, Latitude: 106.813375, Rating: 1, CuisineType: "Western"})
	restaurants = append(restaurants, Restaurant{Id: "14", Name: "Steak on you", Longitude: -6.240269, Latitude: 106.822397, Rating: 5, CuisineType: "Western"})
	restaurants = append(restaurants, Restaurant{Id: "15", Name: "Fish n friend", Longitude: -6.255120, Latitude: 106.814638, Rating: 5, CuisineType: "Western"})
	restaurants = append(restaurants, Restaurant{Id: "16", Name: "Uncle Bro", Longitude: -6.262402, Latitude: 106.796775, Rating: 5, CuisineType: "Western"})
	restaurants = append(restaurants, Restaurant{Id: "17", Name: "My Aunt", Longitude: -6.268950, Latitude: 106.797264, Rating: 5, CuisineType: "Western"})
	restaurants = append(restaurants, Restaurant{Id: "18", Name: "Just Juice", Longitude: -6.263575, Latitude: 106.808294, Rating: 4, CuisineType: "Drink"})
	restaurants = append(restaurants, Restaurant{Id: "19", Name: "Tea Tea Tea", Longitude: -6.261485, Latitude: 106.815718, Rating: 3, CuisineType: "Drink"})
	restaurants = append(restaurants, Restaurant{Id: "20", Name: "Ice Ice Baby", Longitude: -6.268609, Latitude: 106.815675, Rating: 4, CuisineType: "Dessert"})
}

func main() {
	fmt.Println("Program Start...")
	//kudo office long= -6.253520, lat = 106.790803

	router := mux.NewRouter()

	PopulateData()

	router.HandleFunc("/getnearbyrestaurant/{longitude}/{latitude}", GetNearbyRestaurants).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", router))
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}
