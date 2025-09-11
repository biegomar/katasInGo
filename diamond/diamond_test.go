package diamond

import "fmt"

func ExampleRenderDiamond() {
	diamond := RenderDiamond("E")
	fmt.Println(diamond)
	// Output:
	//     A
	//    B B
	//   C   C
	//  D     D
	// E       E
	//  D     D
	//   C   C
	//    B B
	//     A
}
