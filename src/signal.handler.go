// from
// https://stackoverflow.com/questions/11268943/is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in-a-defe

c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt)
go func(){
    for sig := range c {
        // sig is a ^C, handle it
    }
}()
