package main

import
(
    "flag"; // command line option parser
    "fmt";  // package implementing formatted I/O.
    "os";   // allow access to Stdout
)

type SMazeCell struct
{
    north_open bool;
    west_open  bool;
    south_open bool;
    east_open  bool;
};

func (inpMazeCell *SMazeCell) display()
{
    if inpMazeCell.north_open == true
    {
        os.Stdout.WriteString("Nop\n")
    }
    else
    {
        os.Stdout.WriteString("Ncl\n")
    }

    if inpMazeCell.west_open == true
    {
        os.Stdout.WriteString("Wop\n")
    }
    else
    {
        os.Stdout.WriteString("Wcl\n")
    }

    if inpMazeCell.south_open == true
    {
        os.Stdout.WriteString("Sop\n")
    }
    else
    {
        os.Stdout.WriteString("Scl\n")
    }

    if inpMazeCell.east_open == true
    {
        os.Stdout.WriteString("Eop\n")
    }
    else
    {
        os.Stdout.WriteString("Ecl\n")
    }
}

func main()
{
    flag.Parse(); // scans the arg list and sets up flags

    if flag.NArg() > 0
    {
        fmt.Printf("Args given :\n");
    }
    else
    {
        fmt.Printf("No args\n")
    }

    for i := 0; i < flag.NArg(); i++
    {
        os.Stdout.WriteString(flag.Arg(i) + "\n")
    }

    var caze *SMazeCell = &SMazeCell{false, false, false, false};

    caze.display();
}

