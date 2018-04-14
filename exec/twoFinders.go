package main

// import local package of finder functions

func main() {
 theMine := [5]string{“rock”, “ore”, “ore”, “rock”, “ore”}
 go finder1(theMine)
 go finder2(theMine)
 <-time.After(time.Second * 5) //you can ignore this for now
}
