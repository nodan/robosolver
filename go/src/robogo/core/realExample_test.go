package core

import "fmt"

var b = NewWalledBoard(16)

//   09 01 01 03|09 01 01 01 01 03|09 01 01 01 01 03
//
//   08 00 00 00 00 02|*c 00 00 00 00 00 00 02|08 02
//                     --
//   08 04 00 00 00 00 01 00 00 00 02|0c 00 00 00 02
//      --                            --
//   08 03|08 00 00 04 00 00 00 00 00 01 00 00 00 06
//                  --                            --
//   08 00 00 00 02|09 00 00 00 00 00 00 00 00 00 03
//
//   08 00 06|08 00 00 00 06|08 00 00 00 00 00 00 02
//         --             --
//   0c 00 01 00 00 00 00 05 04 00 04 00 00 06|08 02
//   --                   -- --    --       --
//   09 00 00 00 00 00 02|09 03|08 03|08 00 01 00 02
//
//   08 00 00 00 00 00 02|0c 06|08 00 00 04 00 00 02
//                        -- --          --
//   08 04 00 02|0c 00 00 01 01 00 00 00 03|08 00 06
//      --       --                               --
//   08 03|08 00 01 00 00 00 02|0c 00 00 00 00 00 03
//                              --
//   0c 00 00 00 00 00 00 00 00 01 00 00 00 00 00 02
//   --
//   09 00 00 00 00 00 04 00 00 00 40 30 10 20 00 02
//                     --
//   08 00 00 00 00 02|09 00 00 04 00 00 00 00 06|0a
//                              --             --
//   08 00 06|08 00 00 00 00 02|09 00 00 00 00 01 02
//         --
//   0c 04 05 06|0c 04 04 04 04 04 06|0c 04 04 04 06
func ExampleSolve() {
	addRightOf(3, 0)
	addRightOf(9, 0)
	addRightOf(5, 1)
	addBelow(6, 1)
	addRightOf(13, 1)
	addRightOf(10, 2)
	addBelow(11, 2)
	addRightOf(1, 3)
	addBelow(5, 3)
	addBelow(15, 3)
	addRightOf(4, 4)
	addRightOf(2, 5)
	addBelow(2, 5)
	addRightOf(7, 5)
	addBelow(7, 5)
	addBelow(0, 6)
	addBelow(7, 6)
	addBelow(8, 6)
	addBelow(10, 6)
	addBelow(13, 6)
	addRightOf(13, 6)
	addRightOf(6, 7)
	addRightOf(8, 7)
	addRightOf(10, 7)
	addRightOf(6, 8)
	addBelow(7, 8)
	addRightOf(8, 8)
	addBelow(8, 8)
	addBelow(12, 8)
	addBelow(1, 9)
	addRightOf(3, 9)
	addBelow(4, 9)
	addRightOf(12, 9)
	addBelow(15, 9)
	addRightOf(1, 10)
	addRightOf(8, 10)
	addBelow(9, 10)
	addBelow(0, 11)
	addBelow(6, 12)
	addRightOf(5, 13)
	addBelow(9, 13)
	addBelow(14, 13)
	addRightOf(14, 13)
	addRightOf(2, 14)
	addRightOf(8, 14)
	addRightOf(3, 15)
	addRightOf(10, 15)

	*b.fieldAt(10, 12) |= EncodeColor(4)
	*b.fieldAt(11, 12) |= EncodeColor(3)
	*b.fieldAt(12, 12) |= EncodeColor(1)
	*b.fieldAt(13, 12) |= EncodeColor(2)

	p := NewPosition(b, 22)
	ok := p.Solve(10)
	fmt.Println(ok)
	fmt.Println(p.move[1:])
	// Output:
	// true
	// [{1 1} {1 8} {1 1} {1 2} {1 4} {4 1} {4 8} {4 1} {1 1} {1 8}]

}

func addRightOf(x, y uint) {
	*b.fieldAt(x, y) |= Wall(EAST)
	*b.fieldAt(x + 1, y) |= Wall(WEST)
}

func addBelow(x, y uint) {
	*b.fieldAt(x, y) |= Wall(SOUTH)
	*b.fieldAt(x, y + 1) |= Wall(NORTH)
}