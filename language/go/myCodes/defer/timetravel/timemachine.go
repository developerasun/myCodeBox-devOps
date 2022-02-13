package timetravel

import "fmt"

// Create an abstract class : RemoteControl
type RemoteControl interface {
	Press() Inspector
}

// Create an abstract class : Inspector
type Inspector interface {
	CheckState()
	Break()
	InitTimeMachine()
}

// Create a embodying class : Person
type Person struct{}

func (p Person) Press() Inspector {
	fmt.Println("You have pressed a button")
	return &Cup{}
}

// Create a embodying class : Cup
type Cup struct {
	state string
}

// Create a function that restores the broken cup
func (c *Cup) InitTimeMachine() {
	defer func() {
		c.state = "=== Your cup is: Restored"
		fmt.Println(c.state)
	}()

	fmt.Println(" ... Initiate time machine ...")

}

// Create a function that checks the cup state if it is broken or not
func (c *Cup) CheckState() {
	c.state = "=== Your cup is: Not Broken"
	fmt.Println(c.state)
}

// Create a function that breaks the cup
func (c *Cup) Break() {
	c.state = "=== Your cup is: Broken"
	fmt.Println(c.state)
}
