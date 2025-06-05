package main

import (
	"Rp/backend/config"
	"Rp/backend/controllers"
	"Rp/backend/routes"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	// "Rp/backend/services"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize the database connection first
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer config.DB.Close()

	// Initialize admin configuration and load current values
	if err := config.InitAdminDB(); err != nil {
		log.Printf("Warning: Admin configuration not initialized: %v", err)
	}

	// Load the current config
	if err := config.LoadConfigFromDB(); err != nil {
		log.Printf("Warning: Failed to load initial config: %v", err)
	}

	// notificationService := services.NewNotificationService(config.DB)
	// 	notificationService.Start()

	if len(os.Args) > 1 && os.Args[1] == "server" {
		startServer()
	} else {
		consoleMode()
	}
}

func startServer() {
	port := "8080"

	// Setup routes
	r := routes.SetupRoutes()

	// Start server
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func consoleMode() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Reward Points Allocation System")
	fmt.Println("Enter reward points for students (comma separated or 'exit' to quit):")

	for {
		fmt.Print("\nEnter points: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			break
		}

		pointsList := strings.Split(input, ",")
		for _, pointStr := range pointsList {
			pointStr = strings.TrimSpace(pointStr)
			points, err := strconv.Atoi(pointStr)
			if err != nil {
				fmt.Printf("Invalid input '%s'. Please enter numbers only.\n", pointStr)
				continue
			}

			if points < 0 {
				fmt.Println("Points cannot be negative. Please enter a positive number.")
				continue
			}

			allocation := controllers.AllocatePoints(points, "CSE", 1)
			controllers.PrintAllocation(allocation)
		}
	}

	fmt.Println("Program exited.")
}
