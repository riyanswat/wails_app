package main

import (
	"embed"
	"math/rand"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// GeneratePassword generates a random password based on user-specified criteria
func (a *App) GeneratePassword(length int, includeLowercase, includeUppercase, includeNumbers, includeSpecialChars bool) string {
	// Define character sets for password generation
	lowercaseChars := "abcdefghijklmnopqrstuvwxyz"
	uppercaseChars := strings.ToUpper(lowercaseChars)
	numberChars := "0123456789"
	specialChars := "!@#$%&*()"

	// Combine character sets based on user preferences
	charset := ""
	if includeLowercase {
		charset += lowercaseChars
	}
	if includeUppercase {
		charset += uppercaseChars
	}
	if includeNumbers {
		charset += numberChars
	}
	if includeSpecialChars {
		charset += specialChars
	}

	// Generate the password
	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Password Generator",
		Width:  1024,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// ===========================================================

// package main

// import (
// 	"embed"

// 	"github.com/wailsapp/wails/v2"
// 	"github.com/wailsapp/wails/v2/pkg/options"
// 	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
// )

// //go:embed all:frontend/dist
// var assets embed.FS

// func main() {
// 	// Create an instance of the app structure
// 	app := NewApp()

// 	// Create application with options
// 	err := wails.Run(&options.App{
// 		Title:  "myproject",
// 		Width:  1024,
// 		Height: 600,
// 		AssetServer: &assetserver.Options{
// 			Assets: assets,
// 		},
// 		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
// 		OnStartup:        app.startup,
// 		Bind: []interface{}{
// 			app,
// 		},
// 	})

// 	if err != nil {
// 		println("Error:", err.Error())
// 	}
// }
