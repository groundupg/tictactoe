package tictactoe

import "core:fmt"
import "core:os"
import "vendor:sdl2"

W_HEIGHT :: 480
W_WIDTH :: 640

INIT_FLAGS :: sdl2.INIT_EVERYTHING
W_FLAGS :: sdl2.WINDOW_SHOWN
R_FLAGS :: sdl2.RENDERER_ACCELERATED

Game :: struct {
	window:   ^sdl2.Window,
	renderer: ^sdl2.Renderer,
}

init :: proc() -> bool {
	if sdl2.Init(INIT_FLAGS) != 0 {
		fmt.eprintf("Error initialising SDL2: %s", sdl2.GetError())
		return false
	}
	return true
}

game :: proc() -> Game {
	g: Game
	g.window = sdl2.CreateWindow(
		"TIC TAC TOE WINDOW",
		sdl2.WINDOWPOS_CENTERED,
		sdl2.WINDOWPOS_CENTERED,
		W_WIDTH,
		W_HEIGHT,
		W_FLAGS,
	)
	g.renderer = sdl2.CreateRenderer(g.window, -1, R_FLAGS)
	return g
}

cleanup :: proc(g: ^Game, e: int) {
	sdl2.DestroyRenderer(g.renderer)
	sdl2.DestroyWindow(g.window)
	sdl2.Quit()
	os.exit(e)
}

draw :: proc(g: ^Game) {
	sdl2.RenderClear(g.renderer)
	sdl2.RenderPresent(g.renderer)
}

main :: proc() {
	exit_status := 0
	if !init() {
		exit_status = 1
		return
	}
	game := game()
	defer cleanup(&game, exit_status)
	draw(&game)
	sdl2.Delay(10000)
}

