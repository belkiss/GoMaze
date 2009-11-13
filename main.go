package main

import
(
    "flag"; // command line option parser
    "fmt";  // package implementing formatted I/O.
    "os";   // allow access to Stdout
)

func main()
{
    flag.Parse(); // scans the arg list and sets up flags ??
    //var pec string = "";
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
}

