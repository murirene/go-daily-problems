package bar 

import (
    "math"
    "sort"
)


func getDrinks(preferences map[int][]int) []int {
    d := make(map[int]bool, 0)
    for _, drinks := range preferences{
        for _, drink := range drinks {
            if _, ok := d[drink]; !ok {
                d[drink] = true
            }
        }
    }

    drinkList := make([]int, len(d))
    for idx, _ := range d {
        drinkList[idx] = idx 
    }

    drinkList = sort.IntSlice(drinkList)
    return drinkList
}

func getClientDrinkSet(preferences map[int][]int) ([]int, map[int] map[int]bool) {
    clients := make([]int, 0)
    clientDrinkSet := make(map[int]map[int]bool, 0)

    for client, drinks := range preferences {
        clients = append(clients, client)
        for _, drink := range drinks {
            drinkSet, ok := clientDrinkSet[client]
            if ok == false {
                clientDrinkSet[client] = make(map[int]bool, 0)
                drinkSet = clientDrinkSet[client]
            }
            drinkSet[drink] = true            
        }
    } 
    
    return clients, clientDrinkSet
}

func getMinRecipeHelper(clients, drinks []int , preferences map[int] map[int]bool, cIdx, dIdx int) int {
    if cIdx == len(clients) {
        return 0
    }

    if dIdx == len(drinks) {
        return 0
    }

    drinkSet := preferences[clients[cIdx]]
    include := 0
    if drinkSet[drinks[dIdx]] {
        include = 1
    }
    countInclude := getMinRecipeHelper(clients, drinks, preferences, cIdx + 1, dIdx + 1) + include 
    countExclude := getMinRecipeHelper(clients, drinks, preferences, cIdx, dIdx + 1) 

    return int(math.Min(float64(countInclude), float64(countExclude)))
}

func GetMinRecipes(preferences map[int][]int) int {
    drinks := getDrinks(preferences)
    clients, clientDrinkSet := getClientDrinkSet(preferences)    

    return getMinRecipeHelper(clients, drinks, clientDrinkSet, 0, 0)
}
