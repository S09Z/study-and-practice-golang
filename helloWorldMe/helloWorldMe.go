package helloworldme

import (
	"fmt"

	internal "github.com/S09Z/study-and-practice-golang/helloWorldMe/internal"
	"github.com/google/uuid"
)

func SayHelloWorldMe() {
	uuid := uuid.New()
	generateTest() // Can Call function in same level of package folder
	fmt.Printf("Hello Me [%s] ==> %d \n", uuid, 5555)
	internal.SayInternal()
}
