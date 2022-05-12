package main

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

type Queen struct 
{
	posx, posy int
}
func (n1 *Queen) anotherQueen(n2 Queen) bool 
{
	return n1.posx == n2.posx 
	|| n1.posy == n2.posy 
	|| math.Abs(float64(n1.posx-n2.posx)) == math.Abs(float64(n1.posy-n2.posy))
}

type Board struct 
{
	Sz    int
	Queen []Queen
}

func (t *Board) fullBoard() bool 
{
	return t.Sz == len(t.Queen)
}

func (t *Board) addQueen(princess Queen) {
	if !t.fullBoard() {
		t.Queen = append(t.Queen, princess)
	}
}

func (t *Board) convertQueen(princess Queen) 
{
	for _, Queentemp := range t.Queen 
	{
		if princess.anotherQueen(Queentemp) {
			return false
		}
	}
	return true
}


func (t *Board) clonQueen(clon Board) bool 
{
	myQueen := t.Queen
	Queens := clon.Queen

	if len(myQueen) != len(Queens) 
	{
		return false
	}
	for i, Queen := range Queens {
		if myQueen[i].posx != Queen.posx 
		|| myQueen[i].posy != Queen.posy {
			return false
		}
	}
	return true
}

type mainBoard struct 
{
	Solutio []Board
}

func (p *mainBoard) addSolution(solution Board) 
{
	p.Solutio = append(p.Solutio, solution)
}

var wg sync.WaitGroup

func (p *mainBoard) paintQueens() 
{
	for _, boardi := range p.Solutio 
	{
		fmt.Println(boardi.Queen)
	}
	fmt.Println(p.paintQueens)
}

func (p *mainBoard) everyPossibility() int 
{
	return len(p.Solutio)
}

func Game(n int, main_board *mainBoard) 
{
	for i := 0; i < n; i++
	{
		firstPrincess := Queen{posx: i, posy: 0}
		boardTemp := Board{Sz: n}
		anticheat(firstPrincess, boardTemp, main_board)
	}
}
func GameConcurrency (n int, main_board *mainBoard){
	wg.Add(n)
	for i:=0;i<n;i++{
		go func (i int)  
		{
			firstQueen := Queen{posx: i, posy: 0}
			boardTemp := Board{Sz: n}
			anticheat(firstQueen, boardTemp, main_board)
		}(i)
	}
	wg.Wait()
}
func anticheat(theQueen Queen, theBoard Board, main_board *mainBoard) {
	if theBoard.convertQueen(theQueen) {
		theBoard.addQueen(theQueen)
		if theBoard.fullBoard() {
			todasQueens := make([]Queen, theBoard.Sz)
			for i, Queen := range theBoard.Queen {
				todasQueens[i] = Queen
			}
			posSoluTablero := Board{Sz: theBoard.Sz, Queen: todasQueens}
			main_board.addSolution(posSoluTablero)
		} else {
			for i := 0; i < theBoard.Sz; i++ {
				for j := theQueen.posy; j < theBoard.Sz; j++ {
					nextQueen := Queen{posx: i, posy: j}
					anticheat(nextQueen, theBoard, main_board)
				}
			}
		}
	}
}

func main() {
	start := time.Now()
	themainBoard := mainBoard{}
	n := 4
	GameConcurrency(n, &themainBoard)
	themainBoard.paintQueens()
	timeExec := time.Since(start)
	log.Printf("Tomo %s en compilar", timeExec)
}
