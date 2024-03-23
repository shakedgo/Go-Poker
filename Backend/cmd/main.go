package main

import (
	"Go-Poker/pkg/api"
	"Go-Poker/pkg/utils"
	"log"
)

func main() {
	// var playerName1 string
	// var playerName2 string
	// fmt.Println("Type a name")
	// fmt.Scanln(&playerName1)
	// fmt.Println("Type anoter name")
	// fmt.Scanln(&playerName2)

	// player1 := utils.NewPlayer(playerName1)
	// player2 := utils.NewPlayer(playerName2)
	// err = table.JoinTable(&player1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = table.JoinTable(&player2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// table.Deal(&table.Deck)
	// fmt.Println(table)

	router := api.SetupRouter()
	utils.InitTable()
	// non blocking operation
	// go func() {
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
	// }()

	// // Wait 3 seconds for the server to start
	// time.Sleep(3 * time.Second)

	// // Add players
	// postNewPlayer("Shaked")
	// postNewPlayer("Emanu")

	// // Join players to table
	// getJoinTable("1")
	// getJoinTable("2")

	// // Print table
	// getPrintTable()
}

// func postNewPlayer(name string) {
// 	url := "http://localhost:8080/new-player"
// 	jsonData := map[string]string{"Name": name}
// 	jsonValue, _ := json.Marshal(jsonData)

// 	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
// 	if err != nil {
// 		log.Println("Error sending POST request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("POST /new-player response:", resp.Status)
// }

// func getJoinTable(playerID string) {
// 	url := "http://localhost:8080/join-table/1/" + playerID

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Println("Error sending GET request to join-table:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	fmt.Printf("GET /join-table/1/%s response:%s", playerID, resp.Status)
// }

// func getPrintTable() {
// 	url := "http://localhost:8080/print-table/1"

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Println("Error sending GET request to print-table:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println("Error reading response body:", err)
// 		return
// 	}

// 	fmt.Println("GET /print-table/1 response:", string(body))
// }
