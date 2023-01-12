package backtrack

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/12/6 11:51
 */

func Test_GenerateParenthesis(t *testing.T) {
	Convey("GenerateParenthesis", t, func() {
		Convey("GenerateParenthesis test1", func() {
			fmt.Println(generateParenthesis(1))
			fmt.Println(generateParenthesis(3))
			fmt.Println(generateParenthesis(4))
		})
	})
}

func Test_SolveNQueen(t *testing.T) {
	Convey("SolveNQueen", t, func() {
		Convey("SolveNQueen test1", func() {
			fmt.Println(solveNQueens(4))
		})
	})
}
