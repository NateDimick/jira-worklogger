package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	err := createAppDir()
	if err != nil {
		fmt.Println("Warning! App directory not created!")
	}
}

func (a *App) emitError(e error) {
	fmt.Println("emitting error", e.Error())
	runtime.EventsEmit(a.ctx, "error", e.Error())
}
