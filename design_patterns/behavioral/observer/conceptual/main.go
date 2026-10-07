package main

func main() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}

// Item Nike Shirt is now in stock
// Sending email to customer abc@gmail.com for item Nike Shirt
// Sending email to customer xyz@gmail.com for item Nike Shirt
