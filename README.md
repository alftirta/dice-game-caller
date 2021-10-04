# dice-game-caller

```dice-game-caller``` is simply an app that calls out a ```dice-game``` module so that you can create and play a dice game with the following rules:

* At the start of the game, each player will get the same amount of dice as specified by the game owner.

* In each turn, all players roll all of their dice simultaneously.

* For each numbered 6 die, a player will get  1 point, and the die will be removed from the game.

* For each numbered 1 die, a player have to give the die to the next player, for example if it's player 1 then he/she has to give the die to player 2.

* If a player has lost all of his/her dice, then the player is out from the game.

* After each die has been evaluated, then comes the next turn, all players roll all of their dice.

* The game continues until at least only one player remains.

* The player with the highest point wins the game.

## Getting started

To get started, you could simply clone the repo (```dice-game-caller```), navigate into that folder, and you can run it right away by running the following command ```go run main.go```. Alternatively you can first build a binary file by running ```go build``` command, and then execute the file with ```./dice-game-caller``` (linux).

After you run the app, you will be prompted to input some game settings, including the number of players, the number of dice of each player, and animation time delay.

* ```Jumlah pemain:```

  Here you input the number of players you want in the game. It must be a non-negative integer.

* ```Jumlah dadu:```

  Here you input the number of dice for each player. It also must be a non-negative integer.

* ```Waktu delay (dalam milidetik):```

  Here you specify the time delay for the game animation. It must be a non-negative integer that represent a time duration in milliseconds.

And then, you can just press ```Enter``` one more time, and the game will start automatically until the game ends, and you will be prompted with the game result whether if there is a winner or if it is a tie game. In each turn, you will see each player with their current score and remaining dice. 
