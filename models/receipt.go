package models

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Receipt represents the structure of the receipt JSON.
type Receipt struct {
	Retailer      string  `json:"retailer"`
	PurchaseDate  string  `json:"purchaseDate"`
	PurchaseTime  string  `json:"purchaseTime"`
	Items         []Item  `json:"items"`
	Total         string  `json:"total"`
}


func (receipt *Receipt) CalculatePoints() int {
	points := 0

	// One point for every alphanumeric character in the retailer name
	retailer := strings.ReplaceAll(receipt.Retailer, " ", "")
	retailerPoints := 0
	for _, char := range retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			retailerPoints++
		}
	}
	points += retailerPoints


	// 50 points if the total is a round dollar amount with no cents
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil {
		if amount := math.Floor(total); amount == total {
			points += 50
		}
	}

	// 25 points if the total is a multiple of 0.25
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
		if math.Mod(total, 0.25) == 0 {
			points += 25
		}
	}

	// 5 points for every two items on the receipt
	itemCount := len(receipt.Items)
	itemPoints := (itemCount / 2) * 5
	points += itemPoints

	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				itemDescPoints := int(math.Ceil(price * 0.2))
				points += itemDescPoints
			}
		}
	}

	// 6 points if the day in the purchase date is odd
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
		
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.After(time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)) && purchaseTime.Before(time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)) {
		points += 10
		
	}


	return points
}

