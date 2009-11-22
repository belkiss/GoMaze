package main

import
(
    "flag"; // command line option parser
    "fmt";  // package implementing formatted I/O.
    "os";   // allow access to Stdout
    "exp/draw";     // allow access to draw primitives
    "exp/draw/x11"; // allow x11 dialog init
)


type SMazeCell struct
{
    northOpen,
    westOpen,
    southOpen,
    eastOpen  bool;
};


////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
func (inpMazeCell *SMazeCell) displayText()
{
    if inpMazeCell.northOpen
    {
        os.Stdout.WriteString("Nop\n")
    }
    else
    {
        os.Stdout.WriteString("Ncl\n")
    }

    if inpMazeCell.westOpen
    {
        os.Stdout.WriteString("Wop\n")
    }
    else
    {
        os.Stdout.WriteString("Wcl\n")
    }

    if inpMazeCell.southOpen
    {
        os.Stdout.WriteString("Sop\n")
    }
    else
    {
        os.Stdout.WriteString("Scl\n")
    }

    if inpMazeCell.eastOpen
    {
        os.Stdout.WriteString("Eop\n")
    }
    else
    {
        os.Stdout.WriteString("Ecl\n")
    }
}


////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
func (inpMazeCell *SMazeCell) draw(inContext draw.Context, inSquareSize int)
{
    vl_screen := inContext.Screen();

    upXLeft, upYLeft := 100, 100;
    squareSize := inSquareSize;

    fullCell := draw.Rect(upXLeft, upYLeft, squareSize + upXLeft, squareSize + upYLeft);
    draw.Draw(vl_screen, fullCell, draw.Black, nil, draw.ZP);

    // inset set the color of the inside
    draw.Draw(vl_screen, fullCell.Inset(1), draw.White, nil, draw.ZP);

    //hide a wall with a rectangle... POOR SOLUTION !!
    if inpMazeCell.northOpen
    {
        // up wall
        hiderup := draw.Rect(upXLeft, upYLeft,
                             upXLeft + squareSize, upYLeft + 1);
        draw.Draw(vl_screen, hiderup, draw.White, nil, draw.ZP);
    }
    else
    {
        // do not hide
    }

    if inpMazeCell.westOpen
    {
        // right wall
        hiderright := draw.Rect(upXLeft + squareSize - 1, upYLeft,
                                upXLeft + squareSize, upYLeft + squareSize);
        draw.Draw(vl_screen, hiderright, draw.White, nil, draw.ZP);
    }
    else
    {
        // do not hide
    }

    if inpMazeCell.southOpen
    {
        // down wall
        hiderdown := draw.Rect(upXLeft, upYLeft + squareSize - 1,
                              upXLeft + squareSize, upYLeft + squareSize);
        draw.Draw(vl_screen, hiderdown, draw.White, nil, draw.ZP);
    }
    else
    {
        // do not hide
    }

    if inpMazeCell.eastOpen
    {
        // left wall
        hiderleft := draw.Rect(upXLeft, upYLeft,
                               upXLeft + 1, upYLeft + squareSize);
        draw.Draw(vl_screen, hiderleft, draw.White, nil, draw.ZP);
    }
    else
    {
        // do not hide
    }
}


////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
func app(inContext draw.Context)
{
    vl_screen := inContext.Screen();

    screenr := draw.Rect(0, 0, vl_screen.Width(), vl_screen.Height());
    draw.Draw(vl_screen, screenr, draw.White, nil, draw.ZP);

    squareSize := 30;

    var cell *SMazeCell = new(SMazeCell);
    cell.draw(inContext, squareSize);

    inContext.FlushImage();

    fmt.Printf("Press the any key to exit.\n");
    for
    {
        select
        {
            case r := <- inContext.KeyboardChan():
                switch r
                {
                    case 'q', 'Q', 0x04, 0x7F, 32 :
                        fmt.Printf("Exiting because of keyboard event %d\n", r);
                        os.Exit(0);
                    default :
                        fmt.Printf("Exiting because of keyboard event %d\n", r);
                }
            case <- inContext.MouseChan():
                // No-op.
            case <- inContext.ResizeChan():
                // No-op.
            case <- inContext.QuitChan():
                fmt.Printf("Exiting because of QuitChan\n");
                return;
        }
    }
}


////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
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

//     var cell *SMazeCell = new(SMazeCell);
// //     var cell *SMazeCell = &SMazeCell{false, false, false, false};
// 
//     cell.displayText();

    // window init start
    x11Context, x11Error := x11.NewWindow();
    if x11Error != nil
    {
        fmt.Printf("Error: %v\n", x11Error);
        return;
    }
    // window init end
    app(x11Context);

}
