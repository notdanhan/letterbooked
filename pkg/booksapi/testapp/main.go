package main

// BFT for one of the API endpoints
import (
	"bufio"
	"fmt"
	"os"

	"github.com/dedobbin/letterbooked/pkg/booksapi"
)

func main() {
	fmt.Print("Enter a book name: ")
	var buff string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		buff = scanner.Text()
	}

	query, err := booksapi.QueryTitle(buff, -1)
	if err != nil {
		fmt.Println("Query failed!", err)
		os.Exit(1)
	}

	fmt.Println(query)

}
