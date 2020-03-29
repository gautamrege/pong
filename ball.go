package main

import (
	"time"

	tl "github.com/JoelOtter/termloop"
)

// Inital size of ball, which reduces over time.
const (
	BALL_FONT_SIZE = 4
	RIGHT          = 1
	DOWN           = 1
	LEFT           = -1
	UP             = -1
)

type Ball struct {
	*tl.Entity
	DirHorizontal int // 1: right, -1: left
	DirVertical   int // 1: down, -1: up
	speed         int
}

func (ball *Ball) new_position(x, y int) (nx, ny int) {
	if x <= 6 { // left edge
		ball.DirHorizontal = RIGHT
	} else if x > PLAY_WIDTH+3 { // right edge
		ball.DirHorizontal = LEFT
	}

	if y <= 6 { // top edge
		ball.DirVertical = DOWN
	} else if y > PLAY_HEIGHT+4 {
		ball.DirVertical = UP
	}

	nx = x + (ball.speed * ball.DirHorizontal)
	ny = y + (ball.speed * ball.DirVertical)

	return
}

func (ball *Ball) Draw(screen *tl.Screen) {
	x, y := ball.Position()

	nx, ny := ball.new_position(x, y)
	ball.SetPosition(nx, ny)

	ball.Entity.Draw(screen)
	time.Sleep(50 * time.Millisecond)
}

func (ball *Ball) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*Brick); ok {
		// check if it's on a brick.
	}
}

func NewBall() Ball {
	ball := Ball{tl.NewEntity(90, 24, PLAY_WIDTH, PLAY_HEIGHT), 1, 1, 1}
	ball.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '🔴'})
	return ball
}
