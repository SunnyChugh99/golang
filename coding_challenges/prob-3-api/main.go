package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const apiEndpoint = "https://jsonmock.hackerrank.com/api/food_outlets?city=%s"

type FoodOutlet struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	City          string `json:"city"`
	EstimatedCost int    `json:"estimated_cost"`
	UserRating    struct {
		AverageRating float64 `json:"average_rating"`
		Votes         int     `json:"votes"`
	} `json:"user_rating"`
}

type Response struct {
	Page       int          `json:"page"`
	PerPage    int          `json:"per_page"`
	Total      int          `json:"total"`
	TotalPages int          `json:"total_pages"`
	Data       []FoodOutlet `json:"data"`
}

func getFoodOutlets(city string, votes int) ([]FoodOutlet, error) {
	resp, err := http.Get(fmt.Sprintf(apiEndpoint, city))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	filteredOutlets := make([]FoodOutlet, 0)
	for _, outlet := range response.Data {
		if outlet.UserRating.Votes >= votes {
			filteredOutlets = append(filteredOutlets, outlet)
		}
	}
	return filteredOutlets, nil
}

func finestFoodOutlet(city string, votes int) (string, error) {
	outlets, err := getFoodOutlets(city, votes)
	fmt.Println("outlets")
	fmt.Println(outlets)
	if err != nil {
		return "", err
	}

	var maxRatingOutlet FoodOutlet
	for _, outlet := range outlets {
		if outlet.UserRating.AverageRating > maxRatingOutlet.UserRating.AverageRating ||
			(outlet.UserRating.AverageRating == maxRatingOutlet.UserRating.AverageRating &&
				outlet.UserRating.Votes > maxRatingOutlet.UserRating.Votes) {
			maxRatingOutlet = outlet
		}
	}

	return maxRatingOutlet.Name, nil
}

func main() {
	// Example usage:
	city := "Seattle"
	votes := 500
	winningRestaurant, err := finestFoodOutlet(city, votes)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("Winning Restaurant:", winningRestaurant)
}
