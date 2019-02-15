package masker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Given_An_Empty_String_Mask_Does_Not_Error(t *testing.T) {
	str := ""
	masked := Mask(str, 0, '*')
	assert.Equal(t, str, masked)
}

func Test_Given_A_Skip_Greater_Than_Length_Mask_Does_Not_Error(t *testing.T) {
	str := "hello"
	masked := Mask(str, len(str)+100, '*')
	assert.Equal(t, str, masked)
}

func Test_Given_Inavlid_Skip_Mask_Does_Not_Error(t *testing.T) {
	str := "hello"
	masked := Mask(str, -100, '*')
	assert.Equal(t, str, masked)
}

func Test_Masks_OK(t *testing.T) {
	str := "1234567890"
	expectedMask := "123*******"
	masked := Mask(str, 3, '*')
	assert.Equal(t, expectedMask, masked)
}

func Test_Reverse_Masks_OK(t *testing.T) {
	str := "1234567890"
	expectedMask := "*******890"
	masked := MaskReverse(str, 3, '*')
	assert.Equal(t, expectedMask, masked)
}

func Test_Given_An_Empty_String_Reverese_Mask_Does_Not_Error(t *testing.T) {
	str := ""
	masked := MaskReverse(str, 0, '*')
	assert.Equal(t, str, masked)
}

func Test_Given_A_Skip_Greater_Than_Length_Reverese_Mask_Does_Not_Error(t *testing.T) {
	str := "hello"
	masked := MaskReverse(str, len(str)+100, '*')
	assert.Equal(t, str, masked)
}

func Test_Given_Inavlid_Skip_Reverese_Mask_Does_Not_Error(t *testing.T) {
	str := "hello"
	masked := MaskReverse(str, -100, '*')
	assert.Equal(t, str, masked)
}

func Test_Non_Ascii_Mask_OK(t *testing.T) {
	str := "こんにちは世界"
	expectedMask := "こんに****"
	masked := Mask(str, 3, '*')
	assert.Equal(t, expectedMask, masked)
}

func Test_Non_Ascii_Reverse_Mask_OK(t *testing.T) {
	str := "こんにちは世界"
	expectedMask := "****は世界"
	masked := MaskReverse(str, 3, '*')
	assert.Equal(t, expectedMask, masked)
}

