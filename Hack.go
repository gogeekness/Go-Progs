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
	X, Y int 
}


type Block struct {
	glyph int
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
	Ouch Weapon
	Gold int
	buffer byte //to buffer Playfield on movement	
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
	Treasure int 	//array of creatures
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

func AddRoom(Buffer [][]byte, Chamber Room){
	// This func adds a room defined by Room Struct
	// The walls are outside the def of the room
	//need to add the block of text to the buffer
	
	var Seg string //Middle of the room	
	Lid := make([]Block, len(Playfield) //Top and bottom of the room
 
	//check for Room in incoming room site
	for l:= Chamber.Loc.Y-1; l <= (Chamber.Size.Y+Chamber.Loc.Y+1); l++ {
		for k:= Chamber.Loc.X-1; k <= (Chamber.Size.X+Chamber.Loc.X+1); k++ {
			if Buffer[l][k].glyph == byte('#') || Buffer[l][k].glyph == byte('.') { 
				Chamber.roomid = -1
			}
		}
	}


	//Need to rework Adding rooms//


	if Chamber.roomid != -1 { 
		Seg = "#"
		// Upper Wall coppied to PlayField 	
		for l:=0; l < (Chamber.Size.X+1); l++{
               		Seg = Seg + "#"
            	}
		Lid = []Block.glyph(Seg) var Seg string //Middle of the room
        var Lid []Block //Top and bottom of the room

		copy(Buffer[Chamber.Loc.Y-1][Chamber.Loc.X:], Lid)
	
		// Midroom segments -> #.....# 
		for k:=0; k < Chamber.Size.Y; k++{
			Seg = "#"	
			for l:=0; l < Chamber.Size.X; l++{
				Seg = Seg + "." 
			}
			Seg = Seg + "#"
			Wall:= []byte(Seg)
			copy(Buffer[Chamber.Loc.Y+k][Chamber.Loc.X:], Wall)
		}
		//Add bottom of the room
        	copy(Buffer[Chamber.Loc.Y+Chamber.Size.Y][Chamber.Loc.X:], Lid)
	}
	fmt.Printf("%#v\n", Chamber)
}

/// =============   Display   ===========================
func display(Buffer [][]byte, usrrow, usrcol *int, h, w int){
        // ## Clear the srceen ##
        cmd2:= exec.Command("clear")
        cmd2.Stdout = os.Stdout
        cmd2.Run()

	if *usrrow <= 0+h {
		*usrrow = 0+h  }
	if *usrrow > cap(Buffer)-(h*2) {
		*usrrow = cap(Buffer)-(h*2) }
	if *usrcol <= 0+w {
                *usrcol = 0+w }
        if *usrcol > cap(Buffer[0])-(w*2) {
                *usrcol = cap(Buffer[0])-(w*2) }

	var row int
	refrow := *usrrow
	refcol := *usrcol
	const displaylines = 8 //allow to dispaly info text for the game

	for row=refrow; row<=(refrow+((h*2)-displaylines)); row++ {
		fmt.Println(string(Buffer[row][refcol:int(refcol)+(w*2)]))
	}	
} 
		

//////////////////////////////////////////////////////////////////////////////////

func main() {

//poll for random seed to start.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var scrh, scrw int
	Fieldsize := 512

	//Terminal window size
        scrh, scrw = Screen()

	//center of the dispaly or the Playfield
	diw := int(scrw/2) 
	dih := int(scrh/2)
	
	//Slice for key input
	var keys []byte = make([]byte, 1)


        // Initize the Playfield to a set char "="

        Playfield := make([][]Block, Fieldsize+scrh)
        for i := range Playfield {
                Playfield[i] = make([]Block, Fieldsize+scrw)
        }
        for i:= range Playfield {
                for j:= range Playfield[0] {
                        Playfield[i][j].glyph = int(" ") /Blank
			Playfield[i][j].color = 0  /Black
			Playfield[i][j].blocking = True
                }
        }

	Player := Agent{
		Health: 10 + r.Intn(10),
		Name: "Richard",
		Ouch: Weapon{Name: "dagger", Dam: 10},
		Location: Pos{X: 5,Y: 5},
		Gold: 10,
		}

	Mon := Agent{
		Health: 5 + r.Intn(6),
		Name: "Kolbal",
		Ouch: Weapon{Name: "fangs", Dam: 5},
		Location: Pos{X: 1, Y: 2},
		}

	Chambers := make([]Room, 400)

	for i:= range Chambers {
		Chambers[i] = Room {
			Loc: Pos{r.Intn(Fieldsize)+1, r.Intn(Fieldsize)+1},
			Size: Pos{r.Intn(20)+3, r.Intn(20)+3},
			Doors: 1,
			Creatures: r.Intn(3),
 			Treasure: r.Intn(100),
			roomid: i,
		}
	}

	for i:= range Chambers {
		AddRoom(Playfield, Chambers[i])
	}

	
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

		display(Playfield, &dih, &diw, scrh/2, scrw/2)
        	fmt.Printf("%#v\n", Player)
        	fmt.Printf("%#v\n", Mon)
			
	}  //Main Loop
	

	exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	
}
