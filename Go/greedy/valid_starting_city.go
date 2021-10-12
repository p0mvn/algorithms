package main

// Problem: city i is distances[i] away from city i + 1. Since the ities are connected via a
// cirular road, the last city is conneted to the first city. fuel[i] is how much fuel you can get in gallons
// in city i, mpg = miles per gallon. The total amount of fuel is exactly enough to travel to all cities.
// Write a function that returns the indix of the valid starting city from which you can travel to every city. 
//
// Solution: using the following facts:
// 1) The total amount of fuel is exactly enough to travel to all cities.
// 2) There is always exactly one valid starting city.
//
// Keep track of how much gas you have as you enter city i before fueling up. The city with the
// lowest amount of gas is the solution because there is no other city that you would enter with less gass,
// no matter what city you start from.
func ValidStartingCity(distances []int, fuel []int, mpg int) int {
    fuelAtEntry :=  0 + fuel[len(fuel) - 1] * mpg - distances[len(fuel) - 1]
    
    minFuelAtEntry := fuelAtEntry
    minFuelAtEntryIdx := 0
    
    for i := 0; i < len(fuel); i++ {
        if fuelAtEntry < minFuelAtEntry {
            minFuelAtEntry = fuelAtEntry
            minFuelAtEntryIdx = i
        }
        fuelAtEntry += fuel[i] * mpg - distances[i]
    }
    
    return minFuelAtEntryIdx
}
