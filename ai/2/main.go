package main

import (
  "fmt"
  //Used this to set some delay's, to make the game more aesthetically pleasing (did I spell that right? :D)
  "time"
)


//To save me from debugging errors where it was 'x' instead of 'X', I set globals vars as the standard

var xPlayer = "X"
var oPlayer = "O"


//Controls whose turn it is

var aiTurn = true


//Used to show how many levels the program searches

var deep int = 0


func main() {


  //The playing board (a map, used because it gives me some more control over the values and makes using it easier)

  playingBoard := map[uint]string{1: "-", 2: "-", 3: "-",
                                  4: "-", 5: "-", 6: "-", 
                                  7: "-", 8: "-", 9: "-"}


  //Pretty self explanatory, when the game is over this is set to false... and the game ends

  var gameNotEnded bool = true


  //The main game loop, as long as GameNotEnded is true, it keeps going

  for gameNotEnded {


    // AiTurn = false means it's the players turn

    if (aiTurn == false) {


      //For the spot the player choses
      var spot uint


      //Gets the spot, uses a label to come back here if the input is not valid

      GETSPOT: fmt.Println("Please enter a spot. Numbering starts from 1 in the top-left")


      fmt.Scanln(&spot)

      //If spot is valid, asign it

      if (playingBoard[spot] == "-") {

        playingBoard[spot] = oPlayer

      } else { //Otherwise ask again

        fmt.Println("Please enter a valid value")

        goto GETSPOT
      }
    } else { //Now if it's the Ai's turn...


      //Set initial depth to zero (increment everytime minimax is called)

      deep = 0


      //Let player now (2 seconds delay to make it seem like the bot it thinkin' hard :D)

      fmt.Println("Bot's turn.....")
      time.Sleep(2 * time.Second)


      //computerMove returns an integer, used as the index of it's move

      computer_move := computerMove(playingBoard)


      //Outputs some data about the choice

      fmt.Println("Chose", computer_move)
      fmt.Println("Searched", deep, "levels deep")


      //Sets position to AI's mark

      playingBoard[computer_move] = xPlayer
    }


    //Draws the board

    drawBoard(playingBoard)


    //Checks for wins for xPlayer (bot) and oPlayer (you)

    if (checkWin(playingBoard, oPlayer)) {

      fmt.Println("Man triumphed over Machine!!")

      gameNotEnded = false

    } else if (checkWin(playingBoard, xPlayer)) {

      fmt.Println("Machine triumphed over Man!!")

      gameNotEnded = false

    } else if (checkTie(playingBoard)) { //Then checks for ties

      fmt.Println("It's a tie :(")
      gameNotEnded = false
    }


    //Switches turns (let aiPlayer equal the opposite of aiPlyer)
      aiTurn = !aiTurn
  }
}


//The meat of the program
//Takes the board as input, the depth as input (though it really is not that useful) and isMaxing (is it the maximizing player or not)

func miniMax( board map[uint]string, depth int, isMaxing bool ) int {
  
  //Increments the depth variable
  deep++

    //If any of these states is reached, it returns a heuristic value, which is then passed back up the tree

  if (checkWin(board, xPlayer)) {
    return 1
  } else if (checkWin(board, oPlayer)) {
    return -1
  } else if (checkTie(board)) {
    return 0
  }


  //For the "X" player

  if isMaxing {

    //Sets bestScore, just has to be lower than any of the other values
    bestScore := -1000

     //Another plus to using maps: It's very easy to loop over all the elements, and you get both the key and the value

    for key, value := range board {

      //If the value is '-' the place is empty

      if value == "-" {

        //Set the place to the piece of player
        board[key] = xPlayer

        //Calls minimax to fiund the OTHER players move and gets it's score (it goes through all this for every time it gets called)

        score := miniMax(board, depth+1, false)

        //We have the score, so reset the position to empty
        board[key] = "-"


        //If the score we got was HIGHER that bestScore, keep that move and set bestScore to the result we got. You'll notice we don't get the key. We use a different function, called computerMove to do that

        if score > bestScore {
          bestScore = score
        }
      }
    } 
    return bestScore


    //This just repeats everything above, but for the other player

  } else {
    bestScore := 1000
    for key, value := range board {
      if value == "-" {
        board[key] = oPlayer
        score := miniMax(board, depth+1, true) 
        board[key] ="-"
        if score < bestScore {
          bestScore = score
        }
      }
    }
    return bestScore
  }
}


//Checks if there are any wins. Takes a board and the character of the player ("X" or "0")

func checkWin (board map[uint]string, char string) bool {
  if ( 
      (board[1] == board[2] && board[1] == board[3] && board[1] == char) ||
      (board[4] == board[5] && board[4] == board[6] && board[4] == char) ||
      (board[7] == board[8] && board[7] == board[9] && board[7] == char) ||
      (board[1] == board[4] && board[1] == board[7] && board[1] == char) ||
      (board[2] == board[5] && board[2] == board[8] && board[2] == char) ||
      (board[3] == board[6] && board[3] == board[9] && board[3] == char) ||
      (board[1] == board[5] && board[1] == board[9] && board[1] == char) ||
      (board[3] == board[5] && board[3] == board[7] && board[3] == char)) {
        return true
      } else {
        return false
      }
}


//Checks ties
//If there are any '-' spots left, the board is not full, so there is no ties

func checkTie(board map[uint]string) bool {
  for _, v := range board {
    if (v == "-") {
      return false
    }
    
  }
  return true
}


//This function just gets the computers move for us, and makes life easier it collecting values. It only takes a board as input

func computerMove(board map[uint]string) uint {

  //Notice we have a new variable: bestMove. This returns the key of the best move.

  var bestMove uint


  //Because we know the the BOT is always the maximizing player, we can just set this to -1000 without worrying about the minimizing player

  var bestScore = -1000
  bestMove = 0 
  
  for key, value := range board {

    //Again, if position is empty...
    if value == "-" {
      
       //Set position to ours...
      board[key] = xPlayer
      
      //get minimax...
      score := miniMax(board, 0, false)

      //Reset...
      board[key] = "-"

      //If the score is better, keep it and the move

      if score > bestScore {
        bestScore = score

        //Set the bestMove to the key of the (heh) best move

        bestMove = key
      }
    }
  }
  return bestMove
}


//Draws the board

func drawBoard (board map[uint]string) {
  fmt.Println(board[1] + "  |  " + board[2] + "  |  " + board[3])
  fmt.Println("-------------")
  fmt.Println(board[4] + "  |  " + board[5] + "  |  " + board[6])
  fmt.Println("-------------")
  fmt.Println(board[7] + "  |  " + board[8] + "  |  " + board[9])
}
