package main

import (
	"github.com/gofiber/fiber/v2"
	"fuckme2000/internal/lexer"
	"fuckme2000/internal/parser"
	"fuckme2000/internal/vm"
	"github.com/gofiber/fiber/v2/middleware/cors"

)

type RunRequest struct {
	Code string `json:"code"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello from hell")
	})

	app.Post("/run", func(c *fiber.Ctx) error {
		var reqModel RunRequest
	
		if err := c.BodyParser(&reqModel); err != nil {
			return c.Status(400).SendString("Invalid JSON.")
		}
	
		sourceCode := reqModel.Code
		l := lexer.New(sourceCode)
	
		var tokens []lexer.Token
		for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
			tokens = append(tokens, tok)
		}
	
		p := parser.New(tokens)
	
		var output string
		vm.SetPrinter(func(s string) {
			output += s + "\n"
		})
	
		defer func() {
			if r := recover(); r != nil {
				output = "Parser Error: " + r.(string)
			}
		}()
	
		instructions := p.ParseProgram()
	
		machine := vm.New()
		machine.Run(instructions)
	
		return c.SendString(output)
	})	

	app.Listen(":3001")
}