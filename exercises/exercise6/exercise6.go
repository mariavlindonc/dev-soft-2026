package main

import (
	"fmt"
	"os"
	"strings"
)

const add = 1
const display = 2
const exit = 0
const path = "contacts.json"

func addContact(name string, email string, phone int) {
	contact := fmt.Sprintf(`{"name":"%s","email":"%s","phone":%d}`,
		name, email, phone)

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			newData := fmt.Sprintf("[%s]", contact)
			if err := os.WriteFile(path, []byte(newData), 0644); err != nil {
				fmt.Println("Error creating contacts file:", err)
			}
			return
		}
		fmt.Println("Error reading contacts:", err)
		return
	}

	content := strings.TrimSpace(string(data))
	if content == "" || content == "[]" {
		content = fmt.Sprintf("[%s]", contact)
	} else {
		if !strings.HasSuffix(content, "]") {
			fmt.Println("Invalid contacts file.")
			return
		}
		trimmed := strings.TrimSpace(content[:len(content)-1])
		if strings.HasSuffix(trimmed, "[") {
			content = fmt.Sprintf("%s%s]", trimmed, contact)
		} else {
			content = fmt.Sprintf("%s,%s]", trimmed, contact)
		}
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		fmt.Println("Error writing contacts:", err)
	}
}

func displayAllContacts() {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No contacts found.")
			return
		}
		fmt.Println("Error reading contacts:", err)
		return
	}

	content := strings.TrimSpace(string(data))
	if content == "" || content == "[]" {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("Contacts:")
	fmt.Println(content)
}

func main() {
	option := 0
	var name, email string
	var phone int

	for {
		fmt.Println("|| Contact Management ||")
		fmt.Println("[1] Add a contact.")
		fmt.Println("[2] Display all contacts")
		fmt.Println("[0] Exit")
		fmt.Scanln(&option)

		switch option {
		case add:
			fmt.Println("Enter the contact name (no spaces):")
			fmt.Scanln(&name)
			for {
				fmt.Println("Enter the email:")
				fmt.Scanln(&email)
				if strings.Contains(email, "@") {
					break
				}
				fmt.Scanln("Enter a valid email.")
			}
			for {
				fmt.Println("Enter the phone:")
				fmt.Scanln(&phone)
				if phone < 10000000000 && phone > 999999999 {
					break
				}
				fmt.Scanln("Enter a valid number (10 digits).")
			}
			addContact(name, email, phone)
		case display:
			displayAllContacts()

		case exit:
			fmt.Println("Exiting the app.")
			return
		default:
			fmt.Println("Invalid input, try again.")
		}
	}
}
