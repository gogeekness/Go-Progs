package main

import (
  "os/exec"
  "fmt"
  "os"
  "math/rand"
  "time"
  //"math"
  //"strings"		
)

// Position of Agent in the Room
type Pos struct {
	row, col int 
}


type Block struct {
	glyph rune
	color int
	blocking bool
}

// Weapon label and damage
type Weapon struct {
	Name string
	Dam int
}

// Agent decsripton for Player and Monster
type Agent struct {
	Health int
	Name string
	Location Pos
	Buffblock Block
	Charblock Block
	Ouch Weapon
	Gold int
}

// Door
// type Door Struct {
//	Front, Back Pos
// 	Frontid, Backid int
// }

// Room definement 
type Room struct {
	Size Pos
	Loc Pos
	roomid int 
	Doors int 	//array of doors
	Creatures int 	//array of creatures
	Treasures int 	//array of creatures
}

// ============ Functionals ==================

func Screen() (int, int){
	// Set up display and keyboard controlls
	var scrh, scrw int 

        // ## Get size of terminal ##
        cmd:= exec.Command("stty", "size")
        cmd.Stdin = os.Stdin
        d, _ := cmd.Output()
        fmt.Sscan(string(d), &scrh, &scrw)

        // ## Clear the srceen ##
        cmd2:= exec.Command("clear")
        cmd2.Stdout = os.Stdout
        cmd2.Run()

	//Keyboard Buffering
        // disable input buffering
        exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
        // do not display entered characters on the screen
        exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

        return scrh, scrw
}

func IsRoomGood(Buffer [][]Block, Chamber Room) bool{
        //check for Room in incoming room site
	
	//chcek for unicode #9608 (block) and #9617 (light shade tile) 
        for l:= Chamber.Loc.row-1; l <= (Chamber.Size.row+Chamber.Loc.row+1); l++ {
                for k:= Chamber.Loc.col-1; k <= (Chamber.Size.col+Chamber.Loc.col+1); k++ {
                        if Buffer[l][k].glyph == rune('█') || Buffer[l][k].glyph == rune('░') {
                                return false
                        }
                }
        }
	return true 
}



func AddRoom(Buffer [][]Block, Chamber Room){
	// This func adds a room defined by Room Struct
	// The walls are outside the def of the room
	//need to add the block of text to the buffer
	//chcek for unicode #9608 (block) and #9617 (light shade tile)

	good := IsRoomGood(Buffer, Chamber)
	 
	if good { 		
		// Make Walls 	
		for l:=0; l < (Chamber.Size.row+1); l++ {
			for k:=0; k < (Chamber.Size.col+1); k++ {
				Buffer[Chamber.Loc.row+l][Chamber.Loc.col+k].glyph = rune('█')
				Buffer[Chamber.Loc.row+l][Chamber.Loc.col+k].color = 1
				Buffer[Chamber.Loc.row+l][Chamber.Loc.col+k].blocking = true
               		}
            	}
		// Carve out floor tiles
                for l:=1; l < (Chamber.Size.row); l++ {
                        for k:=1; k < (Chamber.Size.col); k++ {
                                Buffer[Chamber.Loc.row+l][Chamber.Loc.col+k].glyph = rune('░')
                                Buffer[Chamber.Loc.row+l][Chamber.Loc.col+k].color = 3
				Buffer[Chamber.Loc.row+l][Chamber.Loc.col+k].blocking = false
                        }
                }

	}
	fmt.Printf("End %#v    %b\n", Chamber, good)
	
}



func printagent(Buffer [][]Block, Thing Agent){
	
	trow := Thing.Location.row  
	tcol := Thing.Location.col 
	
	//Save the data from the Playfield
	Thing.Buffblock.glyph = Buffer[trow][tcol].glyph
	Thing.Buffblock.color = Buffer[trow][tcol].color
	Thing.Buffblock.blocking = Buffer[trow][tcol].blocking

	//Insert Agent into the Playfield copying over the old tile
	Buffer[trow][tcol].glyph =  Thing.Charblock.glyph 
	Buffer[trow][tcol].color =  Thing.Charblock.color
	Buffer[trow][tcol].blocking =  Thing.Charblock.blocking
}




func bufferagent(Buffer [][]Block, Thing Agent){

        trow := Thing.Location.row
        tcol := Thing.Location.col

        //Remove Agent into the Playfield and put back the old tile.
        Buffer[trow][tcol].glyph =  Thing.Buffblock.glyph
        Buffer[trow][tcol].color =  Thing.Buffblock.color
        Buffer[trow][tcol].blocking =  Thing.Buffblock.blocking
}
	




/// =============   Display   ===========================
func display(Buffer [][]Block, usrrow, usrcol *int, scrh, scrw int){
        // ## Clear the srceen ##
        cmd2:= exec.Command("clear")
        cmd2.Stdout = os.Stdout
        cmd2.Run()

	if *usrrow <= 0+scrh/2 {
		*usrrow = 0+scrh/2  }
	if *usrrow > cap(Buffer)-(scrh/2) {
		*usrrow = cap(Buffer)-(scrh/2) }
	if *usrcol <= 0+scrw/2 {
                *usrcol = 0+scrw/2 }
        if *usrcol > cap(Buffer[0])-(scrh/2) {
                *usrcol = cap(Buffer[0])-(scrh/2) }

	var row, col int
	refrow := *usrrow-scrh/2
	refcol := *usrcol-scrw/2
	const displaylines = 8 //allow to dispaly info text for the game

	for row=refrow; row<=(refrow+scrh-displaylines); row++ {
		for col=refcol; col<=(refcol+scrw-1); col++ {
			prtchar:= fmt.Sprintf("\x1b[%dm%s\x1b[1m", 30+Buffer[row][col].color, string(Buffer[row][col].glyph))
			fmt.Printf("%s", prtchar)
		}
		//fmt.Println()
	
	}	
} 
		

//////////////////////////////////////////////////////////////////////////////////

func main() {

//poll for random seed to start.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var scrh, scrw int
	Fieldsize := 512
	Rooms := 200

	//Terminal window size
        scrh, scrw = Screen()

	//center of the dispaly or the Playfield
	diw := int(scrw/2) 
	dih := int(scrh/2)
	
	//Slice for key input
	var keys []byte = make([]byte, 1)


        // Initize the Playfield to a set char (spc)

        Playfield := make([][]Block, Fieldsize+scrh)
        for i := range Playfield {
                Playfield[i] = make([]Block, Fieldsize+scrw)
        }
        for i:= range Playfield {
                for j:= range Playfield[0] {
                        Playfield[i][j].glyph = rune('x') 	//Blank
			Playfield[i][j].color = 1  		//Black
			Playfield[i][j].blocking = true
                }
        }

	Player := Agent{
		Health: 10 + r.Intn(10),
		Name: "Richard",
		Ouch: Weapon{Name: "dagger", Dam: 10},
		Location: Pos{row: dih,col: diw},
		Gold: 10,
		Charblock: Block{glyph: rune('☻'), color: 7, blocking: true},
		}

	Mon := Agent{
		Health: 5 + r.Intn(6),
		Name: "Kolbal",
		Ouch: Weapon{Name: "fangs", Dam: 5},
		Location: Pos{1, 2},
                Charblock: Block{glyph: rune('☸'), color: 4, blocking: true},
		}

	Chambers := make([]Room, Rooms)
	

	Chambers[0] = Room {
		Loc: Pos{dih-5, diw-5},
		Size: Pos{15, 15},
		Doors: 3,
		Creatures: r.Intn(3),
 		Treasures: r.Intn(100),
		roomid: 0,
		}
	
        Chambers[1] = Room {
                Loc: Pos{dih+30, diw-5},
                Size: Pos{5, 5},
                Doors: 3,
                Creatures: r.Intn(3),
                Treasures: r.Intn(100),
                roomid: 0,
                }

        Chambers[2] = Room {
                Loc: Pos{dih-5, diw+20},
                Size: Pos{10, 10},
                Doors: 3,
                Creatures: r.Intn(3),
                Treasures: r.Intn(100),
                roomid: 0,
                }



	//for i:= range Chambers {
                AddRoom(Playfield, Chambers[0])
               	AddRoom(Playfield, Chambers[1])
                AddRoom(Playfield, Chambers[2])

	//}

	
	// ************* Main Loop ***************
	for keys[0] != byte('q') {
	        os.Stdin.Read(keys)
	
		// ## Get size of terminal ##
        	cmd:= exec.Command("stty", "size")
        	cmd.Stdin = os.Stdin
        	d, _ := cmd.Output()
        	fmt.Sscan(string(d), &scrh, &scrw)

		switch keys[0] {
                case byte('a'):  
                        diw--
                        //west player
	
                case byte('w'):  
                      	dih--
                       	//up player
               		
                case byte('s'): 
                        dih++
                        //up player
                        
                case byte('d'):  
                        diw++
                        //up player
                        
                case byte('e'):  
                        //use
                        
                case byte('r'):  
                        //Attack
                        
                default:
		}

		printagent(Playfield, Player)		

		display(Playfield, &dih, &diw, scrh, scrw)
		fmt.Printf("\x1b[32m")

		bufferagent(Playfield, Player)

		Player.Location.row = dih;
		Player.Location.col = diw;


		fmt.Println("Hieght ",dih, "Width ",diw)
		fmt.Printf("End %#v\n", Chambers[1])
        	fmt.Printf("%#v\n", Player)
        	fmt.Printf("%#v\n", Mon)
			
	}  //Main Loop
	

	exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	
}
