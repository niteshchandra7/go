package main

import "strings"

type Food struct {
	Name    string
	Cuisine string
	Rating  int
}

type FoodRatings struct {
	Foods   []Food
	FoodMap map[string]int
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		Foods:   []Food{},
		FoodMap: make(map[string]int),
	}

	for i := 0; i < len(foods); i++ {
		f := Food{
			Name:    foods[i],
			Cuisine: cuisines[i],
			Rating:  ratings[i],
		}
		fr.Foods = append(fr.Foods, f)
		fr.FoodMap[foods[i]] = i
	}
	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	idx, _ := this.FoodMap[food]
	this.Foods[idx].Rating = newRating
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	maxRating := 0
	bestFood := ""
	for i := 0; i < len(this.Foods); i++ {
		if this.Foods[i].Cuisine != cuisine {
			continue
		}
		if this.Foods[i].Rating > maxRating {
			maxRating = this.Foods[i].Rating
			bestFood = this.Foods[i].Name
		} else if this.Foods[i].Rating == maxRating && strings.Compare(this.Foods[i].Name, bestFood) == -1 {
			bestFood = this.Foods[i].Name

		}
	}
	return bestFood
}

func main() {

}
