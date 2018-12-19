package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYear(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(1, Roman("I"), "罗马数字测试I")
	ast.Equal(5, Roman("V"), "罗马数字测试V")
	ast.Equal(10, Roman("X"), "罗马数字测试X")
	ast.Equal(50, Roman("L"), "罗马数字测试L")
	ast.Equal(100, Roman("C"), "罗马数字测试C")
	ast.Equal(500, Roman("D"), "罗马数字测试D")
	ast.Equal(1000, Roman("M"), "罗马数字测试M")

	ast.Equal(2, Roman("II"), "罗马数字测试II")
	ast.Equal(3, Roman("III"), "罗马数字测试III")
	ast.Equal(4, Roman("IV"), "罗马数字测试IV")
	ast.Equal(6, Roman("VI"), "罗马数字测试VI")
	ast.Equal(7, Roman("VII"), "罗马数字测试VII")
	ast.Equal(9, Roman("IX"), "罗马数字测试IX")
	ast.Equal(45, Roman("XLV"), "罗马数字测试XLV")
	ast.Equal(1900, Roman("MCM"), "罗马数字测试MCM")
	ast.Equal(99, Roman("XCIX"), "罗马数字测试XCIX")
	ast.Equal(14, Roman("XIV"), "罗马数字测试XIV")
	ast.Equal(3000, Roman("MMM"), "罗马数字测试MMM")
	ast.Equal(19, Roman("XIX"), "罗马数字测试XIX")

	ast.Equal(0, Roman("IIII"), "罗马数字测试IIII")
	ast.Equal(0, Roman("VIIII"), "罗马数字测试VIIII")
	ast.Equal(0, Roman("VL"), "罗马数字测试VL")
	ast.Equal(0, Roman("IC"), "罗马数字测试IC")
	ast.Equal(0, Roman("IIX"), "罗马数字测试IIX")
	ast.Equal(0, Roman("XIIII"), "罗马数字测试XIIII")
}
