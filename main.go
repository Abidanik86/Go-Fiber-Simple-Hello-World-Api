package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// নতুন Fiber অ্যাপ তৈরি করা
	app := fiber.New()

	//  রুট রাউট সেট করা
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// সার্ভার চালানো (পোর্ট 3000 এ)
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("সার্ভার চালু করতে সমস্যা হয়েছে:", err)
	}

}
