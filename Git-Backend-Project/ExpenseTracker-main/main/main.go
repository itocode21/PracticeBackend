package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

// --------------------------------------------------
// Expense struct
type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"text"`
	Amount      int       `json:"amount"`
	ExpenseTime time.Time `json:"expense_time"`
}

var expenses []Expense

//--------------------------------------------------
// read .json file

func loadExpense() {
	file, err := os.ReadFile("expense.json")
	if err != nil {
		return
	}
	json.Unmarshal(file, &expenses)
}

// --------------------------------------------------
//savew new expense

func saveExpense() {
	file, _ := json.MarshalIndent(expenses, "", "	")
	_ = os.WriteFile("expense.json", file, 0644)
}

//--------------------------------------------------
// add expense

func addExpense(text string, amount int) {
	expense := Expense{
		ID:          len(expenses) + 1,
		Description: text,
		Amount:      amount,
		ExpenseTime: time.Now()}

	expenses = append(expenses, expense)
	saveExpense()
	fmt.Printf("Expense added successfully [ID: %d]", expense.ID)
}

//--------------------------------------------------
// delete expense

func deleteExpense(id int) {
	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			saveExpense()
			fmt.Printf("Delete expense successfully [ID: %d]", expense.ID)
			return
		}
	}
}

//--------------------------------------------------
// show summary

func showSummary() {
	var totalAmount int
	for _, expense := range expenses {
		totalAmount += expense.Amount
	}
	fmt.Printf("Total expenses: $%d\n", totalAmount)
}

//--------------------------------------------------
// show expense for n-month

func showExpenceMoth(month time.Month) {
	var totalAmount int
	for _, expense := range expenses {
		if expense.ExpenseTime.Month() == month {
			totalAmount += expense.Amount
		}
	}
	fmt.Printf("Total expenses for %s: $%d\n", month.String(), totalAmount)
}

// --------------------------------------------------
// list all expense

func listExpense() {
	fmt.Println("# ID  Date       Description          Amount")
	for _, expense := range expenses {
		fmt.Printf("# %d   %s  %-20s  $%d\n",
			expense.ID, expense.ExpenseTime.Format("2006-01-02"),
			expense.Description,
			expense.Amount)
	}
}

// --------------------------------------------------
// download json

func handleFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=expense.json")
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "expense.json")
}

func main() {
	loadExpense()
	if len(os.Args) < 2 {
		fmt.Println("No Command provided")
		return
	}

	// --------------------------------------------------
	// add expense

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: add [description] [amount]")
			return
		}
		text := os.Args[2]
		amount, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Invalid amount. It must be a number.")
			return
		}
		addExpense(text, amount)

	// --------------------------------------------------
	// delete expense by ID

	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		deleteExpense(id)

	// --------------------------------------------------
	// list summary all expenses

	case "summary":
		showSummary()

	// --------------------------------------------------
	// list expense for n-month

	case "month":
		if len(os.Args) < 3 {
			fmt.Println("Usage: month [month number]")
			return
		}
		monthNum, err := strconv.Atoi(os.Args[2])
		if err != nil || monthNum < 1 || monthNum > 12 {
			fmt.Println("Invalid month number. It must be a number between 1 and 12.")
			return
		}
		showExpenceMoth(time.Month(monthNum))

	// --------------------------------------------------
	// list all expense

	case "list":
		listExpense()

	// --------------------------------------------------
	// download json

	case "export":
		http.HandleFunc("/download", handleFile)

		fmt.Println("Server running on http://localhost:8080/download")
		http.ListenAndServe(":8080", nil)

	default:
		fmt.Println("Invalid command")
	}
}
