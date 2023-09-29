# Hi, dear auditor 👋

# As you know, the project is "ascii-art"
It is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.

## To run a program type to the terminal:

$ go run . "any text"  banner
or 
$ go run main.go "any text"" banner

## The program works with ASCII symbols using a banners "standard", "shadow" and "thinkertoy"

$ go run . "hello" standard | cat -e
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
$ go run . "Hello There!" shadow | cat -e
                                                                                         
_|    _|          _| _|                _|_|_|_|_| _|                                  _| 
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| 
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| 
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| 
                                                                                         
                                                                                         
$ go run . "Hello There!" thinkertoy | cat -e
                                                
o  o     o o           o-O-o o                o 
|  |     | |             |   |                | 
O--O o-o | | o-o         |   O--o o-o o-o o-o o 
|  | |-' | | | |         |   |  | |-' |   |-'   
o  o o-o o o o-o         o   o  o o-o o   o-o O 
                                                
                                                

## The characteristics of this program:
- Each character has a height of 8 lines.
- Characters are separated by a new line \n.
