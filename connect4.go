package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"regexp"
	"github.com/inancgumus/screen"
	"math/rand"
	"time"
)

var s = rand.NewSource(time.Now().UnixNano());
var r = rand.New(s);

const width = 7;
const height = 6;
var reg = regexp.MustCompile("[^0-9]");

type board struct{
	state [height][width]int;
	moves int;
	winner int;
}

func (b *board) print(){

	screen.Clear();
	screen.MoveTopLeft();

	for i:=0; i<width; i++ {
		fmt.Print(i," ");
	}

	fmt.Println("");

	for i:=0; i<height; i++ {
		for j:=0; j<width; j++ {
			switch b.state[i][j]{
			case 0: fmt.Print("  ");
			case 1: fmt.Print("X ");
			case 2: fmt.Print("O ");
			}
		}
		fmt.Println("");
	}
}

func (b *board) place(col int, player int) int {

	if (b.state[height-1][col] == 0){
		b.state[height-1][col] = player;
		return height-1;
	}

	for i:=0; i<height-1; i++{
		if (b.state[i+1][col] != 0 && b.state[i][col] == 0){
			b.state[i][col] = player;
			return i;
		}
	}
	return -1;
}

func (b *board) randomInput(player int) int {

	var validPlay = -1;
	var col = 0;

	for (validPlay < 0){
		col = r.Intn(width)
		if (col < 0 || col >= width){
			continue;
		}
		validPlay = b.place(col, player);
	}

	if (b.checkState(player)){
		b.winner = player;
	}

	if (player == 1){
		player = 2; 
	} else {
		player = 1; 
	}

	return player;

}

func (b *board) parseInput(player int) int {

	reader := bufio.NewReader(os.Stdin);
	var validPlay = -1;
	var col = 0;
	var currentPlayer = "";

	if (player == 1){ 
		currentPlayer  ="X";
	} else { 
		currentPlayer = "O";
	}

	for (validPlay < 0){
		fmt.Print("Please select a column ",currentPlayer," > ");
		text, _ := reader.ReadString('\n');
		text = reg.ReplaceAllString(text, "");
		col, _ = strconv.Atoi(text);
		if (col < 0 || col >= width){
			continue;
		}
		validPlay = b.place(col, player);
	}

	if (b.checkState(player)){
		b.winner = player;
	}

	if (player == 1){
		player = 2; 
	} else {
		player = 1; 
	}

	return player;

}

func (b *board) checkState(player int) bool {
	for i:=0; i<height; i++{
		for j:=0; j<width; j++{
			if(b.HR(i, j, player) ||
			   b.DR(i, j, player) ||
			   b.VD(i, j, player)){
				   return true;
			   }
		}
	}
	return false;
}

func (b *board) HR(row int, col int, player int) bool {
	if (col+3 >= width || b.state[row][col] != player){
		return false;
	}

	if (b.state[row][col+1] == player &&
		b.state[row][col+2] == player &&
		b.state[row][col+3] == player){
		return true;
	}

	return false;
}

func (b *board) VD(row int, col int, player int) bool {
	if (row+3 >= height || b.state[row][col] != player){
		return false;
	}

	if (b.state[row+1][col] == player &&
		b.state[row+2][col] == player &&
		b.state[row+3][col] == player){
		return true;
	}

	return false;
}

func (b *board) DR(row int, col int, player int) bool {
	if (row+3 >= height || col+3 >= width || b.state[row][col] != player){
		return false;
	}

	if (b.state[row+1][col+1] == player &&
		b.state[row+2][col+2] == player &&
		b.state[row+3][col+3] == player){
		return true;
	}

	return false;
}

func main() {

	var b board;
	var playing = true;
	var currentPlayer = 1;
	var players [2]int;

	players[0] = 1;
	players[1] = 1;

	for(playing){
		switch players[currentPlayer-1]{
			case 0:
				b.print();
				currentPlayer = b.parseInput(currentPlayer);
			case 1:
				currentPlayer = b.randomInput(currentPlayer);
		}

		if (b.winner == 1){
			b.print();
			fmt.Println("X Wins!");
			playing = false;
		} else if(b.winner == 2){
			b.print();
			fmt.Println("O Wins!");
			playing = false;
		}
		if (b.moves >= 42){
			b.print();
			fmt.Println("Tie Game!");
			playing = false;
		}
		
	}
}