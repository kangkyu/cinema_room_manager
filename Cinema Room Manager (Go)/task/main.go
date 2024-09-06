package main

import (
	"fmt"
	"os"
)

type Ticket struct {
	seat, row int
}

var tickets []Ticket

func main() {
	maxRows, maxSeats := 9, 9

	var rows, seats int
	fmt.Println("Enter the number of rows:")
	fmt.Scan(&rows)
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scan(&seats)

	if rows > maxRows {
		rows = maxRows
	}
	if seats > maxSeats {
		seats = maxSeats
	}

	for {
		fmt.Println()
		fmt.Println("1. Show the seats")
		fmt.Println("2. Buy a ticket")
		fmt.Println("3. Statistics")
		fmt.Println("0. Exit")

		var input int
		fmt.Scan(&input)
		switch input {
		case 0:
			os.Exit(0)
		case 1:
			showSeats(rows, seats)
		case 2:
			buyTicket(rows, seats)
		case 3:
			printStatistics(rows, seats)
		default:
			continue
		}
	}
}

func showSeats(rows, seats int) {

	fmt.Println()
	fmt.Println("Cinema:")
	for i := 0; i <= rows; i += 1 {
		for j := 0; j <= seats; j += 1 {
			if i == 0 && j == 0 {
				fmt.Print(" ")
			} else if i == 0 {
				fmt.Print(j)
			} else if j == 0 {
				fmt.Print(i)
			} else if inTickets(i, j) {
				fmt.Print("B")
			} else {
				fmt.Print("S")
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func inTickets(i, j int) bool {
	for _, t := range tickets {
		if i == t.row && j == t.seat {
			return true
		}
	}
	return false
}

func buyTicket(rows, seats int) {
	var row, seat int

	for {
		fmt.Println()
		fmt.Println("Enter a row number:")
		fmt.Scan(&row)
		fmt.Println("Enter a seat number in that row:")
		fmt.Scan(&seat)

		if row > rows || seat > seats {
			fmt.Println()
			fmt.Println("Wrong input!")
			continue
		}
		if inTickets(row, seat) {
			fmt.Println()
			fmt.Println("That ticket has already been purchased!")
			continue
		}
		break
	}

	ticket := Ticket{row: row, seat: seat}
	ticketPrice := ticketPrice(rows, seats, ticket)

	fmt.Println()
	fmt.Printf("Ticket price: $%d\n", ticketPrice)

	tickets = append(tickets, ticket)
}

func printStatistics(rows, seats int) {

	ticketCount := len(tickets)
	percentage := expressPercentage(len(tickets), rows*seats)
	latestPrice := currentIncome(rows, seats, tickets)
	totalIncome := totalIncome(rows, seats)

	fmt.Println()
	fmt.Printf(`Number of purchased tickets: %d
Percentage: %s
Current income: $%d
Total income: $%d`, ticketCount, percentage, latestPrice, totalIncome)
	fmt.Println()
}

func totalIncome(rows, seats int) int {

	frontRows, backRows := rows/2, rows/2+rows%2
	if rows*seats > 60 {
		return (frontRows*10 + backRows*8) * seats
	} else {
		return rows * seats * 10
	}
}

func currentIncome(rows, seats int, tickets []Ticket) int {

	var income int
	for _, t := range tickets {
		income += ticketPrice(rows, seats, t)
	}
	return income
}

func ticketPrice(rows, seats int, ticket Ticket) int {
	frontRows := rows / 2
	if rows*seats > 60 && ticket.row > frontRows {
		return 8
	} else {
		return 10
	}
}

func expressPercentage(portion, total int) string {
	return fmt.Sprintf("%.2f", float64(portion)*100/float64(total)) + "%"
}
