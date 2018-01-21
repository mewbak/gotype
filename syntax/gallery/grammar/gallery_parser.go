// Generated from Gallery.g4 by ANTLR 4.7.

package grammar // Gallery
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 326,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 62, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4,
	68, 10, 4, 12, 4, 14, 4, 71, 11, 4, 3, 4, 3, 4, 5, 4, 75, 10, 4, 3, 4,
	3, 4, 3, 4, 7, 4, 80, 10, 4, 12, 4, 14, 4, 83, 11, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 91, 10, 4, 5, 4, 93, 10, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 7, 5, 99, 10, 5, 12, 5, 14, 5, 102, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7,
	5, 108, 10, 5, 12, 5, 14, 5, 111, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 5, 5, 120, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 127,
	10, 7, 12, 7, 14, 7, 130, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 144, 10, 11, 3, 12, 3, 12,
	3, 12, 6, 12, 149, 10, 12, 13, 12, 14, 12, 150, 3, 13, 3, 13, 3, 13, 6,
	13, 156, 10, 13, 13, 13, 14, 13, 157, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 7, 15, 168, 10, 15, 12, 15, 14, 15, 171, 11, 15, 3,
	16, 3, 16, 3, 16, 5, 16, 176, 10, 16, 3, 16, 3, 16, 3, 16, 7, 16, 181,
	10, 16, 12, 16, 14, 16, 184, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 6, 17,
	190, 10, 17, 13, 17, 14, 17, 191, 3, 17, 5, 17, 195, 10, 17, 3, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	6, 19, 209, 10, 19, 13, 19, 14, 19, 210, 7, 19, 213, 10, 19, 12, 19, 14,
	19, 216, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 6, 20, 253, 10, 20, 13, 20, 14,
	20, 254, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 263, 10, 20,
	3, 21, 5, 21, 266, 10, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 5,
	22, 274, 10, 22, 3, 23, 3, 23, 5, 23, 278, 10, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 5, 23, 296, 10, 23, 3, 24, 3, 24, 3, 24, 7, 24, 301,
	10, 24, 12, 24, 14, 24, 304, 11, 24, 3, 24, 3, 24, 3, 24, 7, 24, 309, 10,
	24, 12, 24, 14, 24, 312, 11, 24, 3, 24, 5, 24, 315, 10, 24, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 5, 25, 322, 10, 25, 3, 26, 3, 26, 3, 26, 2, 5, 28,
	30, 36, 27, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 2, 8, 3, 2, 19, 21, 11, 2, 3, 3, 6,
	6, 15, 21, 24, 25, 28, 31, 33, 34, 39, 39, 41, 44, 46, 46, 3, 2, 15, 16,
	3, 2, 17, 18, 3, 2, 30, 31, 3, 2, 46, 47, 2, 354, 2, 52, 3, 2, 2, 2, 4,
	61, 3, 2, 2, 2, 6, 92, 3, 2, 2, 2, 8, 119, 3, 2, 2, 2, 10, 121, 3, 2, 2,
	2, 12, 128, 3, 2, 2, 2, 14, 131, 3, 2, 2, 2, 16, 135, 3, 2, 2, 2, 18, 137,
	3, 2, 2, 2, 20, 143, 3, 2, 2, 2, 22, 145, 3, 2, 2, 2, 24, 152, 3, 2, 2,
	2, 26, 159, 3, 2, 2, 2, 28, 161, 3, 2, 2, 2, 30, 175, 3, 2, 2, 2, 32, 185,
	3, 2, 2, 2, 34, 196, 3, 2, 2, 2, 36, 199, 3, 2, 2, 2, 38, 262, 3, 2, 2,
	2, 40, 265, 3, 2, 2, 2, 42, 273, 3, 2, 2, 2, 44, 295, 3, 2, 2, 2, 46, 314,
	3, 2, 2, 2, 48, 321, 3, 2, 2, 2, 50, 323, 3, 2, 2, 2, 52, 53, 5, 12, 7,
	2, 53, 54, 7, 2, 2, 3, 54, 3, 3, 2, 2, 2, 55, 62, 5, 14, 8, 2, 56, 62,
	5, 6, 4, 2, 57, 62, 5, 18, 10, 2, 58, 62, 5, 20, 11, 2, 59, 62, 5, 8, 5,
	2, 60, 62, 5, 16, 9, 2, 61, 55, 3, 2, 2, 2, 61, 56, 3, 2, 2, 2, 61, 57,
	3, 2, 2, 2, 61, 58, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 60, 3, 2, 2, 2,
	62, 5, 3, 2, 2, 2, 63, 64, 7, 5, 2, 2, 64, 69, 7, 46, 2, 2, 65, 66, 7,
	10, 2, 2, 66, 68, 7, 46, 2, 2, 67, 65, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2,
	69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 93, 3, 2, 2, 2, 71, 69, 3,
	2, 2, 2, 72, 74, 7, 26, 2, 2, 73, 75, 7, 5, 2, 2, 74, 73, 3, 2, 2, 2, 74,
	75, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 81, 7, 46, 2, 2, 77, 78, 7, 10,
	2, 2, 78, 80, 7, 46, 2, 2, 79, 77, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81,
	79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 93, 3, 2, 2, 2, 83, 81, 3, 2, 2,
	2, 84, 85, 7, 4, 2, 2, 85, 90, 5, 46, 24, 2, 86, 87, 5, 30, 16, 2, 87,
	88, 7, 38, 2, 2, 88, 89, 5, 30, 16, 2, 89, 91, 3, 2, 2, 2, 90, 86, 3, 2,
	2, 2, 90, 91, 3, 2, 2, 2, 91, 93, 3, 2, 2, 2, 92, 63, 3, 2, 2, 2, 92, 72,
	3, 2, 2, 2, 92, 84, 3, 2, 2, 2, 93, 7, 3, 2, 2, 2, 94, 95, 7, 43, 2, 2,
	95, 100, 7, 46, 2, 2, 96, 97, 7, 10, 2, 2, 97, 99, 7, 46, 2, 2, 98, 96,
	3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 120, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 104, 7, 44, 2, 2, 104,
	109, 7, 46, 2, 2, 105, 106, 7, 10, 2, 2, 106, 108, 7, 46, 2, 2, 107, 105,
	3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2,
	2, 2, 110, 120, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112, 113, 7, 42, 2, 2,
	113, 120, 7, 50, 2, 2, 114, 115, 7, 45, 2, 2, 115, 116, 5, 26, 14, 2, 116,
	117, 7, 7, 2, 2, 117, 118, 7, 32, 2, 2, 118, 120, 3, 2, 2, 2, 119, 94,
	3, 2, 2, 2, 119, 103, 3, 2, 2, 2, 119, 112, 3, 2, 2, 2, 119, 114, 3, 2,
	2, 2, 120, 9, 3, 2, 2, 2, 121, 122, 7, 3, 2, 2, 122, 11, 3, 2, 2, 2, 123,
	124, 5, 4, 3, 2, 124, 125, 7, 9, 2, 2, 125, 127, 3, 2, 2, 2, 126, 123,
	3, 2, 2, 2, 127, 130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2,
	2, 2, 129, 13, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132, 7, 24, 2, 2,
	132, 133, 5, 12, 7, 2, 133, 134, 7, 25, 2, 2, 134, 15, 3, 2, 2, 2, 135,
	136, 3, 2, 2, 2, 136, 17, 3, 2, 2, 2, 137, 138, 5, 46, 24, 2, 138, 139,
	7, 6, 2, 2, 139, 140, 5, 28, 15, 2, 140, 19, 3, 2, 2, 2, 141, 144, 5, 22,
	12, 2, 142, 144, 5, 24, 13, 2, 143, 141, 3, 2, 2, 2, 143, 142, 3, 2, 2,
	2, 144, 21, 3, 2, 2, 2, 145, 148, 5, 28, 15, 2, 146, 147, 7, 7, 2, 2, 147,
	149, 5, 28, 15, 2, 148, 146, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 148,
	3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 23, 3, 2, 2, 2, 152, 155, 5, 30,
	16, 2, 153, 154, 9, 2, 2, 2, 154, 156, 5, 30, 16, 2, 155, 153, 3, 2, 2,
	2, 156, 157, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158,
	25, 3, 2, 2, 2, 159, 160, 9, 3, 2, 2, 160, 27, 3, 2, 2, 2, 161, 162, 8,
	15, 1, 2, 162, 163, 5, 30, 16, 2, 163, 169, 3, 2, 2, 2, 164, 165, 12, 3,
	2, 2, 165, 166, 7, 41, 2, 2, 166, 168, 5, 30, 16, 2, 167, 164, 3, 2, 2,
	2, 168, 171, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170,
	29, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 172, 173, 8, 16, 1, 2, 173, 176,
	5, 36, 19, 2, 174, 176, 5, 32, 17, 2, 175, 172, 3, 2, 2, 2, 175, 174, 3,
	2, 2, 2, 176, 182, 3, 2, 2, 2, 177, 178, 12, 4, 2, 2, 178, 179, 9, 4, 2,
	2, 179, 181, 5, 36, 19, 2, 180, 177, 3, 2, 2, 2, 181, 184, 3, 2, 2, 2,
	182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 31, 3, 2, 2, 2, 184, 182,
	3, 2, 2, 2, 185, 189, 5, 36, 19, 2, 186, 187, 5, 10, 6, 2, 187, 188, 5,
	36, 19, 2, 188, 190, 3, 2, 2, 2, 189, 186, 3, 2, 2, 2, 190, 191, 3, 2,
	2, 2, 191, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 194, 3, 2, 2, 2,
	193, 195, 5, 34, 18, 2, 194, 193, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195,
	33, 3, 2, 2, 2, 196, 197, 5, 10, 6, 2, 197, 198, 7, 40, 2, 2, 198, 35,
	3, 2, 2, 2, 199, 200, 8, 19, 1, 2, 200, 201, 5, 38, 20, 2, 201, 214, 3,
	2, 2, 2, 202, 203, 12, 4, 2, 2, 203, 204, 9, 5, 2, 2, 204, 213, 5, 38,
	20, 2, 205, 208, 12, 3, 2, 2, 206, 207, 7, 39, 2, 2, 207, 209, 5, 38, 20,
	2, 208, 206, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210,
	211, 3, 2, 2, 2, 211, 213, 3, 2, 2, 2, 212, 202, 3, 2, 2, 2, 212, 205,
	3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2,
	2, 2, 215, 37, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217, 218, 7, 32, 2, 2,
	218, 263, 5, 44, 23, 2, 219, 220, 5, 40, 21, 2, 220, 221, 5, 44, 23, 2,
	221, 263, 3, 2, 2, 2, 222, 223, 5, 42, 22, 2, 223, 224, 7, 13, 2, 2, 224,
	225, 5, 30, 16, 2, 225, 226, 7, 10, 2, 2, 226, 227, 5, 30, 16, 2, 227,
	228, 7, 14, 2, 2, 228, 263, 3, 2, 2, 2, 229, 230, 5, 44, 23, 2, 230, 231,
	7, 13, 2, 2, 231, 232, 5, 30, 16, 2, 232, 233, 7, 10, 2, 2, 233, 234, 5,
	30, 16, 2, 234, 235, 7, 14, 2, 2, 235, 263, 3, 2, 2, 2, 236, 263, 5, 44,
	23, 2, 237, 238, 7, 27, 2, 2, 238, 263, 5, 38, 20, 2, 239, 240, 7, 36,
	2, 2, 240, 241, 5, 30, 16, 2, 241, 242, 7, 37, 2, 2, 242, 243, 5, 38, 20,
	2, 243, 263, 3, 2, 2, 2, 244, 245, 7, 34, 2, 2, 245, 263, 5, 38, 20, 2,
	246, 247, 7, 33, 2, 2, 247, 248, 5, 30, 16, 2, 248, 249, 7, 37, 2, 2, 249,
	250, 5, 38, 20, 2, 250, 263, 3, 2, 2, 2, 251, 253, 7, 28, 2, 2, 252, 251,
	3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 254, 255, 3, 2,
	2, 2, 255, 256, 3, 2, 2, 2, 256, 263, 5, 38, 20, 2, 257, 258, 9, 6, 2,
	2, 258, 263, 5, 46, 24, 2, 259, 260, 7, 28, 2, 2, 260, 261, 7, 29, 2, 2,
	261, 263, 5, 36, 19, 2, 262, 217, 3, 2, 2, 2, 262, 219, 3, 2, 2, 2, 262,
	222, 3, 2, 2, 2, 262, 229, 3, 2, 2, 2, 262, 236, 3, 2, 2, 2, 262, 237,
	3, 2, 2, 2, 262, 239, 3, 2, 2, 2, 262, 244, 3, 2, 2, 2, 262, 246, 3, 2,
	2, 2, 262, 252, 3, 2, 2, 2, 262, 257, 3, 2, 2, 2, 262, 259, 3, 2, 2, 2,
	263, 39, 3, 2, 2, 2, 264, 266, 9, 4, 2, 2, 265, 264, 3, 2, 2, 2, 265, 266,
	3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 268, 5, 42, 22, 2, 268, 41, 3, 2,
	2, 2, 269, 270, 7, 48, 2, 2, 270, 271, 7, 18, 2, 2, 271, 274, 7, 48, 2,
	2, 272, 274, 7, 48, 2, 2, 273, 269, 3, 2, 2, 2, 273, 272, 3, 2, 2, 2, 274,
	43, 3, 2, 2, 2, 275, 277, 7, 48, 2, 2, 276, 278, 7, 22, 2, 2, 277, 276,
	3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 296, 3, 2, 2, 2, 279, 296, 5, 46,
	24, 2, 280, 281, 7, 11, 2, 2, 281, 282, 5, 30, 16, 2, 282, 283, 7, 10,
	2, 2, 283, 284, 5, 30, 16, 2, 284, 285, 7, 12, 2, 2, 285, 296, 3, 2, 2,
	2, 286, 287, 7, 11, 2, 2, 287, 288, 5, 30, 16, 2, 288, 289, 7, 12, 2, 2,
	289, 296, 3, 2, 2, 2, 290, 291, 7, 24, 2, 2, 291, 292, 5, 12, 7, 2, 292,
	293, 5, 30, 16, 2, 293, 294, 7, 25, 2, 2, 294, 296, 3, 2, 2, 2, 295, 275,
	3, 2, 2, 2, 295, 279, 3, 2, 2, 2, 295, 280, 3, 2, 2, 2, 295, 286, 3, 2,
	2, 2, 295, 290, 3, 2, 2, 2, 296, 45, 3, 2, 2, 2, 297, 302, 7, 47, 2, 2,
	298, 301, 5, 48, 25, 2, 299, 301, 5, 50, 26, 2, 300, 298, 3, 2, 2, 2, 300,
	299, 3, 2, 2, 2, 301, 304, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302, 303,
	3, 2, 2, 2, 303, 315, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 305, 310, 7, 46,
	2, 2, 306, 309, 5, 48, 25, 2, 307, 309, 5, 50, 26, 2, 308, 306, 3, 2, 2,
	2, 308, 307, 3, 2, 2, 2, 309, 312, 3, 2, 2, 2, 310, 308, 3, 2, 2, 2, 310,
	311, 3, 2, 2, 2, 311, 315, 3, 2, 2, 2, 312, 310, 3, 2, 2, 2, 313, 315,
	7, 23, 2, 2, 314, 297, 3, 2, 2, 2, 314, 305, 3, 2, 2, 2, 314, 313, 3, 2,
	2, 2, 315, 47, 3, 2, 2, 2, 316, 322, 7, 48, 2, 2, 317, 318, 7, 13, 2, 2,
	318, 319, 5, 30, 16, 2, 319, 320, 7, 14, 2, 2, 320, 322, 3, 2, 2, 2, 321,
	316, 3, 2, 2, 2, 321, 317, 3, 2, 2, 2, 322, 49, 3, 2, 2, 2, 323, 324, 9,
	7, 2, 2, 324, 51, 3, 2, 2, 2, 35, 61, 69, 74, 81, 90, 92, 100, 109, 119,
	128, 143, 150, 157, 169, 175, 182, 191, 194, 210, 212, 214, 254, 262, 265,
	273, 277, 295, 300, 302, 308, 310, 314, 321,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'parameter'", "", "':='", "'='", "':'", "';'", "','", "'('", "')'",
	"'['", "']'", "'+'", "'-'", "'*'", "'/'", "'||'", "'|-'", "'~'", "", "'@'",
	"'begingroup'", "'endgroup'", "'local'", "", "", "'edge'", "'frame'", "'box'",
	"", "'subpath'", "'reverse'", "'with'", "'point'", "'of'", "'to'", "",
	"'cycle'", "", "'proof'", "'save'", "'show'", "'let'", "", "", "", "'.'",
}
var symbolicNames = []string{
	"", "PATHJOIN", "PARAMETER", "TYPE", "ASSIGN", "EQUALS", "COLON", "SEMIC",
	"COMMA", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "PLUS", "MINUS", "TIMES",
	"OVER", "PARALLEL", "PERPENDIC", "CONGRUENT", "UNIT", "LAMBDAARG", "BEGINGROUP",
	"ENDGROUP", "LOCAL", "PAIRPART", "EDGECONSTR", "EDGE", "FRAME", "BOX",
	"MATHFUNC", "SUBPATH", "REVERSE", "WITH", "POINT", "OF", "TO", "TRANSFORM",
	"CYCLE", "PATHCLIPOP", "PROOF", "SAVE", "SHOW", "LET", "TAG", "MIXEDTAG",
	"DECIMALTOKEN", "DOT", "LABEL", "WS",
}

var ruleNames = []string{
	"program", "statement", "declaration", "command", "pathjoin", "statementlist",
	"compound", "empty", "assignment", "constraint", "equation", "orientation",
	"token", "expression", "tertiary", "path", "cycle", "secondary", "primary",
	"scalarmulop", "numtokenatom", "atom", "variable", "subscript", "anytag",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GalleryParser struct {
	*antlr.BaseParser
}

func NewGalleryParser(input antlr.TokenStream) *GalleryParser {
	this := new(GalleryParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gallery.g4"

	return this
}

// GalleryParser tokens.
const (
	GalleryParserEOF          = antlr.TokenEOF
	GalleryParserPATHJOIN     = 1
	GalleryParserPARAMETER    = 2
	GalleryParserTYPE         = 3
	GalleryParserASSIGN       = 4
	GalleryParserEQUALS       = 5
	GalleryParserCOLON        = 6
	GalleryParserSEMIC        = 7
	GalleryParserCOMMA        = 8
	GalleryParserLPAREN       = 9
	GalleryParserRPAREN       = 10
	GalleryParserLBRACKET     = 11
	GalleryParserRBRACKET     = 12
	GalleryParserPLUS         = 13
	GalleryParserMINUS        = 14
	GalleryParserTIMES        = 15
	GalleryParserOVER         = 16
	GalleryParserPARALLEL     = 17
	GalleryParserPERPENDIC    = 18
	GalleryParserCONGRUENT    = 19
	GalleryParserUNIT         = 20
	GalleryParserLAMBDAARG    = 21
	GalleryParserBEGINGROUP   = 22
	GalleryParserENDGROUP     = 23
	GalleryParserLOCAL        = 24
	GalleryParserPAIRPART     = 25
	GalleryParserEDGECONSTR   = 26
	GalleryParserEDGE         = 27
	GalleryParserFRAME        = 28
	GalleryParserBOX          = 29
	GalleryParserMATHFUNC     = 30
	GalleryParserSUBPATH      = 31
	GalleryParserREVERSE      = 32
	GalleryParserWITH         = 33
	GalleryParserPOINT        = 34
	GalleryParserOF           = 35
	GalleryParserTO           = 36
	GalleryParserTRANSFORM    = 37
	GalleryParserCYCLE        = 38
	GalleryParserPATHCLIPOP   = 39
	GalleryParserPROOF        = 40
	GalleryParserSAVE         = 41
	GalleryParserSHOW         = 42
	GalleryParserLET          = 43
	GalleryParserTAG          = 44
	GalleryParserMIXEDTAG     = 45
	GalleryParserDECIMALTOKEN = 46
	GalleryParserDOT          = 47
	GalleryParserLABEL        = 48
	GalleryParserWS           = 49
)

// GalleryParser rules.
const (
	GalleryParserRULE_program       = 0
	GalleryParserRULE_statement     = 1
	GalleryParserRULE_declaration   = 2
	GalleryParserRULE_command       = 3
	GalleryParserRULE_pathjoin      = 4
	GalleryParserRULE_statementlist = 5
	GalleryParserRULE_compound      = 6
	GalleryParserRULE_empty         = 7
	GalleryParserRULE_assignment    = 8
	GalleryParserRULE_constraint    = 9
	GalleryParserRULE_equation      = 10
	GalleryParserRULE_orientation   = 11
	GalleryParserRULE_token         = 12
	GalleryParserRULE_expression    = 13
	GalleryParserRULE_tertiary      = 14
	GalleryParserRULE_path          = 15
	GalleryParserRULE_cycle         = 16
	GalleryParserRULE_secondary     = 17
	GalleryParserRULE_primary       = 18
	GalleryParserRULE_scalarmulop   = 19
	GalleryParserRULE_numtokenatom  = 20
	GalleryParserRULE_atom          = 21
	GalleryParserRULE_variable      = 22
	GalleryParserRULE_subscript     = 23
	GalleryParserRULE_anytag        = 24
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Statementlist() IStatementlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementlistContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GalleryParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *GalleryParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GalleryParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Statementlist()
	}
	{
		p.SetState(51)
		p.Match(GalleryParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Compound() ICompoundContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundContext)
}

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Constraint() IConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstraintContext)
}

func (s *StatementContext) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *StatementContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *GalleryParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GalleryParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Compound()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Assignment()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Constraint()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.Command()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(58)
			p.Empty()
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) CopyFrom(ctx *DeclarationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypedeclContext struct {
	*DeclarationContext
}

func NewTypedeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedeclContext {
	var p = new(TypedeclContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *TypedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GalleryParserTYPE, 0)
}

func (s *TypedeclContext) AllTAG() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserTAG)
}

func (s *TypedeclContext) TAG(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserTAG, i)
}

func (s *TypedeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserCOMMA)
}

func (s *TypedeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserCOMMA, i)
}

func (s *TypedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterTypedecl(s)
	}
}

func (s *TypedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitTypedecl(s)
	}
}

type LocaldeclContext struct {
	*DeclarationContext
}

func NewLocaldeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocaldeclContext {
	var p = new(LocaldeclContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *LocaldeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocaldeclContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(GalleryParserLOCAL, 0)
}

func (s *LocaldeclContext) AllTAG() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserTAG)
}

func (s *LocaldeclContext) TAG(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserTAG, i)
}

func (s *LocaldeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GalleryParserTYPE, 0)
}

func (s *LocaldeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserCOMMA)
}

func (s *LocaldeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserCOMMA, i)
}

func (s *LocaldeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterLocaldecl(s)
	}
}

func (s *LocaldeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitLocaldecl(s)
	}
}

type ParameterdeclContext struct {
	*DeclarationContext
}

func NewParameterdeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParameterdeclContext {
	var p = new(ParameterdeclContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *ParameterdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterdeclContext) PARAMETER() antlr.TerminalNode {
	return s.GetToken(GalleryParserPARAMETER, 0)
}

func (s *ParameterdeclContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ParameterdeclContext) AllTertiary() []ITertiaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITertiaryContext)(nil)).Elem())
	var tst = make([]ITertiaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITertiaryContext)
		}
	}

	return tst
}

func (s *ParameterdeclContext) Tertiary(i int) ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *ParameterdeclContext) TO() antlr.TerminalNode {
	return s.GetToken(GalleryParserTO, 0)
}

func (s *ParameterdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterParameterdecl(s)
	}
}

func (s *ParameterdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitParameterdecl(s)
	}
}

func (p *GalleryParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GalleryParserRULE_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GalleryParserTYPE:
		localctx = NewTypedeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Match(GalleryParserTYPE)
		}
		{
			p.SetState(62)
			p.Match(GalleryParserTAG)
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GalleryParserCOMMA {
			{
				p.SetState(63)
				p.Match(GalleryParserCOMMA)
			}
			{
				p.SetState(64)
				p.Match(GalleryParserTAG)
			}

			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case GalleryParserLOCAL:
		localctx = NewLocaldeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Match(GalleryParserLOCAL)
		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GalleryParserTYPE {
			{
				p.SetState(71)
				p.Match(GalleryParserTYPE)
			}

		}
		{
			p.SetState(74)
			p.Match(GalleryParserTAG)
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GalleryParserCOMMA {
			{
				p.SetState(75)
				p.Match(GalleryParserCOMMA)
			}
			{
				p.SetState(76)
				p.Match(GalleryParserTAG)
			}

			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case GalleryParserPARAMETER:
		localctx = NewParameterdeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.Match(GalleryParserPARAMETER)
		}
		{
			p.SetState(83)
			p.Variable()
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GalleryParserLPAREN)|(1<<GalleryParserPLUS)|(1<<GalleryParserMINUS)|(1<<GalleryParserLAMBDAARG)|(1<<GalleryParserBEGINGROUP)|(1<<GalleryParserPAIRPART)|(1<<GalleryParserEDGECONSTR)|(1<<GalleryParserFRAME)|(1<<GalleryParserBOX)|(1<<GalleryParserMATHFUNC)|(1<<GalleryParserSUBPATH))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GalleryParserREVERSE-32))|(1<<(GalleryParserPOINT-32))|(1<<(GalleryParserTAG-32))|(1<<(GalleryParserMIXEDTAG-32))|(1<<(GalleryParserDECIMALTOKEN-32)))) != 0) {
			{
				p.SetState(84)
				p.tertiary(0)
			}
			{
				p.SetState(85)
				p.Match(GalleryParserTO)
			}
			{
				p.SetState(86)
				p.tertiary(0)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) CopyFrom(ctx *CommandContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ShowcmdContext struct {
	*CommandContext
}

func NewShowcmdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShowcmdContext {
	var p = new(ShowcmdContext)

	p.CommandContext = NewEmptyCommandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CommandContext))

	return p
}

func (s *ShowcmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowcmdContext) SHOW() antlr.TerminalNode {
	return s.GetToken(GalleryParserSHOW, 0)
}

func (s *ShowcmdContext) AllTAG() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserTAG)
}

func (s *ShowcmdContext) TAG(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserTAG, i)
}

func (s *ShowcmdContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserCOMMA)
}

func (s *ShowcmdContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserCOMMA, i)
}

func (s *ShowcmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterShowcmd(s)
	}
}

func (s *ShowcmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitShowcmd(s)
	}
}

type ProofcmdContext struct {
	*CommandContext
}

func NewProofcmdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProofcmdContext {
	var p = new(ProofcmdContext)

	p.CommandContext = NewEmptyCommandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CommandContext))

	return p
}

func (s *ProofcmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProofcmdContext) PROOF() antlr.TerminalNode {
	return s.GetToken(GalleryParserPROOF, 0)
}

func (s *ProofcmdContext) LABEL() antlr.TerminalNode {
	return s.GetToken(GalleryParserLABEL, 0)
}

func (s *ProofcmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterProofcmd(s)
	}
}

func (s *ProofcmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitProofcmd(s)
	}
}

type SavecmdContext struct {
	*CommandContext
}

func NewSavecmdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SavecmdContext {
	var p = new(SavecmdContext)

	p.CommandContext = NewEmptyCommandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CommandContext))

	return p
}

func (s *SavecmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SavecmdContext) SAVE() antlr.TerminalNode {
	return s.GetToken(GalleryParserSAVE, 0)
}

func (s *SavecmdContext) AllTAG() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserTAG)
}

func (s *SavecmdContext) TAG(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserTAG, i)
}

func (s *SavecmdContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserCOMMA)
}

func (s *SavecmdContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserCOMMA, i)
}

func (s *SavecmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterSavecmd(s)
	}
}

func (s *SavecmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitSavecmd(s)
	}
}

type LetcmdContext struct {
	*CommandContext
}

func NewLetcmdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetcmdContext {
	var p = new(LetcmdContext)

	p.CommandContext = NewEmptyCommandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CommandContext))

	return p
}

func (s *LetcmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetcmdContext) LET() antlr.TerminalNode {
	return s.GetToken(GalleryParserLET, 0)
}

func (s *LetcmdContext) Token() ITokenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITokenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITokenContext)
}

func (s *LetcmdContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GalleryParserEQUALS, 0)
}

func (s *LetcmdContext) MATHFUNC() antlr.TerminalNode {
	return s.GetToken(GalleryParserMATHFUNC, 0)
}

func (s *LetcmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterLetcmd(s)
	}
}

func (s *LetcmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitLetcmd(s)
	}
}

func (p *GalleryParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GalleryParserRULE_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GalleryParserSAVE:
		localctx = NewSavecmdContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(GalleryParserSAVE)
		}
		{
			p.SetState(93)
			p.Match(GalleryParserTAG)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GalleryParserCOMMA {
			{
				p.SetState(94)
				p.Match(GalleryParserCOMMA)
			}
			{
				p.SetState(95)
				p.Match(GalleryParserTAG)
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case GalleryParserSHOW:
		localctx = NewShowcmdContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Match(GalleryParserSHOW)
		}
		{
			p.SetState(102)
			p.Match(GalleryParserTAG)
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GalleryParserCOMMA {
			{
				p.SetState(103)
				p.Match(GalleryParserCOMMA)
			}
			{
				p.SetState(104)
				p.Match(GalleryParserTAG)
			}

			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case GalleryParserPROOF:
		localctx = NewProofcmdContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(GalleryParserPROOF)
		}
		{
			p.SetState(111)
			p.Match(GalleryParserLABEL)
		}

	case GalleryParserLET:
		localctx = NewLetcmdContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(112)
			p.Match(GalleryParserLET)
		}
		{
			p.SetState(113)
			p.Token()
		}
		{
			p.SetState(114)
			p.Match(GalleryParserEQUALS)
		}
		{
			p.SetState(115)
			p.Match(GalleryParserMATHFUNC)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPathjoinContext is an interface to support dynamic dispatch.
type IPathjoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathjoinContext differentiates from other interfaces.
	IsPathjoinContext()
}

type PathjoinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathjoinContext() *PathjoinContext {
	var p = new(PathjoinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_pathjoin
	return p
}

func (*PathjoinContext) IsPathjoinContext() {}

func NewPathjoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathjoinContext {
	var p = new(PathjoinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_pathjoin

	return p
}

func (s *PathjoinContext) GetParser() antlr.Parser { return s.parser }

func (s *PathjoinContext) PATHJOIN() antlr.TerminalNode {
	return s.GetToken(GalleryParserPATHJOIN, 0)
}

func (s *PathjoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathjoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathjoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterPathjoin(s)
	}
}

func (s *PathjoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitPathjoin(s)
	}
}

func (p *GalleryParser) Pathjoin() (localctx IPathjoinContext) {
	localctx = NewPathjoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GalleryParserRULE_pathjoin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(GalleryParserPATHJOIN)
	}

	return localctx
}

// IStatementlistContext is an interface to support dynamic dispatch.
type IStatementlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementlistContext differentiates from other interfaces.
	IsStatementlistContext()
}

type StatementlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementlistContext() *StatementlistContext {
	var p = new(StatementlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_statementlist
	return p
}

func (*StatementlistContext) IsStatementlistContext() {}

func NewStatementlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementlistContext {
	var p = new(StatementlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_statementlist

	return p
}

func (s *StatementlistContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementlistContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementlistContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementlistContext) AllSEMIC() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserSEMIC)
}

func (s *StatementlistContext) SEMIC(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserSEMIC, i)
}

func (s *StatementlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterStatementlist(s)
	}
}

func (s *StatementlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitStatementlist(s)
	}
}

func (p *GalleryParser) Statementlist() (localctx IStatementlistContext) {
	localctx = NewStatementlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GalleryParserRULE_statementlist)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(121)
				p.Statement()
			}
			{
				p.SetState(122)
				p.Match(GalleryParserSEMIC)
			}

		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// ICompoundContext is an interface to support dynamic dispatch.
type ICompoundContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompoundContext differentiates from other interfaces.
	IsCompoundContext()
}

type CompoundContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundContext() *CompoundContext {
	var p = new(CompoundContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_compound
	return p
}

func (*CompoundContext) IsCompoundContext() {}

func NewCompoundContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundContext {
	var p = new(CompoundContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_compound

	return p
}

func (s *CompoundContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundContext) BEGINGROUP() antlr.TerminalNode {
	return s.GetToken(GalleryParserBEGINGROUP, 0)
}

func (s *CompoundContext) Statementlist() IStatementlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementlistContext)
}

func (s *CompoundContext) ENDGROUP() antlr.TerminalNode {
	return s.GetToken(GalleryParserENDGROUP, 0)
}

func (s *CompoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterCompound(s)
	}
}

func (s *CompoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitCompound(s)
	}
}

func (p *GalleryParser) Compound() (localctx ICompoundContext) {
	localctx = NewCompoundContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GalleryParserRULE_compound)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(GalleryParserBEGINGROUP)
	}
	{
		p.SetState(130)
		p.Statementlist()
	}
	{
		p.SetState(131)
		p.Match(GalleryParserENDGROUP)
	}

	return localctx
}

// IEmptyContext is an interface to support dynamic dispatch.
type IEmptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyContext differentiates from other interfaces.
	IsEmptyContext()
}

type EmptyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyContext() *EmptyContext {
	var p = new(EmptyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_empty
	return p
}

func (*EmptyContext) IsEmptyContext() {}

func NewEmptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyContext {
	var p = new(EmptyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_empty

	return p
}

func (s *EmptyContext) GetParser() antlr.Parser { return s.parser }
func (s *EmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterEmpty(s)
	}
}

func (s *EmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitEmpty(s)
	}
}

func (p *GalleryParser) Empty() (localctx IEmptyContext) {
	localctx = NewEmptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GalleryParserRULE_empty)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GalleryParserASSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *GalleryParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GalleryParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Variable()
	}
	{
		p.SetState(136)
		p.Match(GalleryParserASSIGN)
	}
	{
		p.SetState(137)
		p.expression(0)
	}

	return localctx
}

// IConstraintContext is an interface to support dynamic dispatch.
type IConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstraintContext differentiates from other interfaces.
	IsConstraintContext()
}

type ConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstraintContext() *ConstraintContext {
	var p = new(ConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_constraint
	return p
}

func (*ConstraintContext) IsConstraintContext() {}

func NewConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstraintContext {
	var p = new(ConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_constraint

	return p
}

func (s *ConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstraintContext) Equation() IEquationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEquationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEquationContext)
}

func (s *ConstraintContext) Orientation() IOrientationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrientationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrientationContext)
}

func (s *ConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterConstraint(s)
	}
}

func (s *ConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitConstraint(s)
	}
}

func (p *GalleryParser) Constraint() (localctx IConstraintContext) {
	localctx = NewConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GalleryParserRULE_constraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.Equation()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Orientation()
		}

	}

	return localctx
}

// IEquationContext is an interface to support dynamic dispatch.
type IEquationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEquationContext differentiates from other interfaces.
	IsEquationContext()
}

type EquationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquationContext() *EquationContext {
	var p = new(EquationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_equation
	return p
}

func (*EquationContext) IsEquationContext() {}

func NewEquationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquationContext {
	var p = new(EquationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_equation

	return p
}

func (s *EquationContext) GetParser() antlr.Parser { return s.parser }

func (s *EquationContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EquationContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EquationContext) AllEQUALS() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserEQUALS)
}

func (s *EquationContext) EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserEQUALS, i)
}

func (s *EquationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EquationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterEquation(s)
	}
}

func (s *EquationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitEquation(s)
	}
}

func (p *GalleryParser) Equation() (localctx IEquationContext) {
	localctx = NewEquationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GalleryParserRULE_equation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.expression(0)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GalleryParserEQUALS {
		{
			p.SetState(144)
			p.Match(GalleryParserEQUALS)
		}
		{
			p.SetState(145)
			p.expression(0)
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrientationContext is an interface to support dynamic dispatch.
type IOrientationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrientationContext differentiates from other interfaces.
	IsOrientationContext()
}

type OrientationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrientationContext() *OrientationContext {
	var p = new(OrientationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_orientation
	return p
}

func (*OrientationContext) IsOrientationContext() {}

func NewOrientationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrientationContext {
	var p = new(OrientationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_orientation

	return p
}

func (s *OrientationContext) GetParser() antlr.Parser { return s.parser }

func (s *OrientationContext) AllTertiary() []ITertiaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITertiaryContext)(nil)).Elem())
	var tst = make([]ITertiaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITertiaryContext)
		}
	}

	return tst
}

func (s *OrientationContext) Tertiary(i int) ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *OrientationContext) AllPARALLEL() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserPARALLEL)
}

func (s *OrientationContext) PARALLEL(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserPARALLEL, i)
}

func (s *OrientationContext) AllPERPENDIC() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserPERPENDIC)
}

func (s *OrientationContext) PERPENDIC(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserPERPENDIC, i)
}

func (s *OrientationContext) AllCONGRUENT() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserCONGRUENT)
}

func (s *OrientationContext) CONGRUENT(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserCONGRUENT, i)
}

func (s *OrientationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrientationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrientationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterOrientation(s)
	}
}

func (s *OrientationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitOrientation(s)
	}
}

func (p *GalleryParser) Orientation() (localctx IOrientationContext) {
	localctx = NewOrientationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GalleryParserRULE_orientation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.tertiary(0)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GalleryParserPARALLEL)|(1<<GalleryParserPERPENDIC)|(1<<GalleryParserCONGRUENT))) != 0) {
		p.SetState(151)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GalleryParserPARALLEL)|(1<<GalleryParserPERPENDIC)|(1<<GalleryParserCONGRUENT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(152)
			p.tertiary(0)
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITokenContext is an interface to support dynamic dispatch.
type ITokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTokenContext differentiates from other interfaces.
	IsTokenContext()
}

type TokenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTokenContext() *TokenContext {
	var p = new(TokenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_token
	return p
}

func (*TokenContext) IsTokenContext() {}

func NewTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TokenContext {
	var p = new(TokenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_token

	return p
}

func (s *TokenContext) GetParser() antlr.Parser { return s.parser }

func (s *TokenContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GalleryParserPLUS, 0)
}

func (s *TokenContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GalleryParserMINUS, 0)
}

func (s *TokenContext) TIMES() antlr.TerminalNode {
	return s.GetToken(GalleryParserTIMES, 0)
}

func (s *TokenContext) OVER() antlr.TerminalNode {
	return s.GetToken(GalleryParserOVER, 0)
}

func (s *TokenContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GalleryParserASSIGN, 0)
}

func (s *TokenContext) PARALLEL() antlr.TerminalNode {
	return s.GetToken(GalleryParserPARALLEL, 0)
}

func (s *TokenContext) PERPENDIC() antlr.TerminalNode {
	return s.GetToken(GalleryParserPERPENDIC, 0)
}

func (s *TokenContext) CONGRUENT() antlr.TerminalNode {
	return s.GetToken(GalleryParserCONGRUENT, 0)
}

func (s *TokenContext) BEGINGROUP() antlr.TerminalNode {
	return s.GetToken(GalleryParserBEGINGROUP, 0)
}

func (s *TokenContext) ENDGROUP() antlr.TerminalNode {
	return s.GetToken(GalleryParserENDGROUP, 0)
}

func (s *TokenContext) EDGECONSTR() antlr.TerminalNode {
	return s.GetToken(GalleryParserEDGECONSTR, 0)
}

func (s *TokenContext) PATHCLIPOP() antlr.TerminalNode {
	return s.GetToken(GalleryParserPATHCLIPOP, 0)
}

func (s *TokenContext) PATHJOIN() antlr.TerminalNode {
	return s.GetToken(GalleryParserPATHJOIN, 0)
}

func (s *TokenContext) EDGE() antlr.TerminalNode {
	return s.GetToken(GalleryParserEDGE, 0)
}

func (s *TokenContext) FRAME() antlr.TerminalNode {
	return s.GetToken(GalleryParserFRAME, 0)
}

func (s *TokenContext) BOX() antlr.TerminalNode {
	return s.GetToken(GalleryParserBOX, 0)
}

func (s *TokenContext) REVERSE() antlr.TerminalNode {
	return s.GetToken(GalleryParserREVERSE, 0)
}

func (s *TokenContext) SUBPATH() antlr.TerminalNode {
	return s.GetToken(GalleryParserSUBPATH, 0)
}

func (s *TokenContext) PROOF() antlr.TerminalNode {
	return s.GetToken(GalleryParserPROOF, 0)
}

func (s *TokenContext) SAVE() antlr.TerminalNode {
	return s.GetToken(GalleryParserSAVE, 0)
}

func (s *TokenContext) SHOW() antlr.TerminalNode {
	return s.GetToken(GalleryParserSHOW, 0)
}

func (s *TokenContext) TRANSFORM() antlr.TerminalNode {
	return s.GetToken(GalleryParserTRANSFORM, 0)
}

func (s *TokenContext) TAG() antlr.TerminalNode {
	return s.GetToken(GalleryParserTAG, 0)
}

func (s *TokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterToken(s)
	}
}

func (s *TokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitToken(s)
	}
}

func (p *GalleryParser) Token() (localctx ITokenContext) {
	localctx = NewTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GalleryParserRULE_token)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(157)
	_la = p.GetTokenStream().LA(1)

	if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GalleryParserPATHJOIN)|(1<<GalleryParserASSIGN)|(1<<GalleryParserPLUS)|(1<<GalleryParserMINUS)|(1<<GalleryParserTIMES)|(1<<GalleryParserOVER)|(1<<GalleryParserPARALLEL)|(1<<GalleryParserPERPENDIC)|(1<<GalleryParserCONGRUENT)|(1<<GalleryParserBEGINGROUP)|(1<<GalleryParserENDGROUP)|(1<<GalleryParserEDGECONSTR)|(1<<GalleryParserEDGE)|(1<<GalleryParserFRAME)|(1<<GalleryParserBOX)|(1<<GalleryParserSUBPATH))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GalleryParserREVERSE-32))|(1<<(GalleryParserTRANSFORM-32))|(1<<(GalleryParserPATHCLIPOP-32))|(1<<(GalleryParserPROOF-32))|(1<<(GalleryParserSAVE-32))|(1<<(GalleryParserSHOW-32))|(1<<(GalleryParserTAG-32)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Tertiary() ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *ExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) PATHCLIPOP() antlr.TerminalNode {
	return s.GetToken(GalleryParserPATHCLIPOP, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GalleryParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *GalleryParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, GalleryParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.tertiary(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GalleryParserRULE_expression)
			p.SetState(162)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(163)
				p.Match(GalleryParserPATHCLIPOP)
			}
			{
				p.SetState(164)
				p.tertiary(0)
			}

		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// ITertiaryContext is an interface to support dynamic dispatch.
type ITertiaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTertiaryContext differentiates from other interfaces.
	IsTertiaryContext()
}

type TertiaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTertiaryContext() *TertiaryContext {
	var p = new(TertiaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_tertiary
	return p
}

func (*TertiaryContext) IsTertiaryContext() {}

func NewTertiaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TertiaryContext {
	var p = new(TertiaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_tertiary

	return p
}

func (s *TertiaryContext) GetParser() antlr.Parser { return s.parser }

func (s *TertiaryContext) CopyFrom(ctx *TertiaryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TertiaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TertiaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PathtertiaryContext struct {
	*TertiaryContext
}

func NewPathtertiaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PathtertiaryContext {
	var p = new(PathtertiaryContext)

	p.TertiaryContext = NewEmptyTertiaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TertiaryContext))

	return p
}

func (s *PathtertiaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathtertiaryContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *PathtertiaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterPathtertiary(s)
	}
}

func (s *PathtertiaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitPathtertiary(s)
	}
}

type TermContext struct {
	*TertiaryContext
}

func NewTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermContext {
	var p = new(TermContext)

	p.TertiaryContext = NewEmptyTertiaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TertiaryContext))

	return p
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) Secondary() ISecondaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISecondaryContext)
}

func (s *TermContext) Tertiary() ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *TermContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GalleryParserPLUS, 0)
}

func (s *TermContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GalleryParserMINUS, 0)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *GalleryParser) Tertiary() (localctx ITertiaryContext) {
	return p.tertiary(0)
}

func (p *GalleryParser) tertiary(_p int) (localctx ITertiaryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTertiaryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITertiaryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, GalleryParserRULE_tertiary, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTermContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(171)
			p.secondary(0)
		}

	case 2:
		localctx = NewPathtertiaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(172)
			p.Path()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTermContext(p, NewTertiaryContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, GalleryParserRULE_tertiary)
			p.SetState(175)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(176)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GalleryParserPLUS || _la == GalleryParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(177)
				p.secondary(0)
			}

		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllSecondary() []ISecondaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISecondaryContext)(nil)).Elem())
	var tst = make([]ISecondaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISecondaryContext)
		}
	}

	return tst
}

func (s *PathContext) Secondary(i int) ISecondaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISecondaryContext)
}

func (s *PathContext) AllPathjoin() []IPathjoinContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPathjoinContext)(nil)).Elem())
	var tst = make([]IPathjoinContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPathjoinContext)
		}
	}

	return tst
}

func (s *PathContext) Pathjoin(i int) IPathjoinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathjoinContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPathjoinContext)
}

func (s *PathContext) Cycle() ICycleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICycleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICycleContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitPath(s)
	}
}

func (p *GalleryParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GalleryParserRULE_path)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.secondary(0)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(184)
				p.Pathjoin()
			}
			{
				p.SetState(185)
				p.secondary(0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(191)
			p.Cycle()
		}

	}

	return localctx
}

// ICycleContext is an interface to support dynamic dispatch.
type ICycleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCycleContext differentiates from other interfaces.
	IsCycleContext()
}

type CycleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCycleContext() *CycleContext {
	var p = new(CycleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_cycle
	return p
}

func (*CycleContext) IsCycleContext() {}

func NewCycleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CycleContext {
	var p = new(CycleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_cycle

	return p
}

func (s *CycleContext) GetParser() antlr.Parser { return s.parser }

func (s *CycleContext) Pathjoin() IPathjoinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathjoinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathjoinContext)
}

func (s *CycleContext) CYCLE() antlr.TerminalNode {
	return s.GetToken(GalleryParserCYCLE, 0)
}

func (s *CycleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CycleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CycleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterCycle(s)
	}
}

func (s *CycleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitCycle(s)
	}
}

func (p *GalleryParser) Cycle() (localctx ICycleContext) {
	localctx = NewCycleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GalleryParserRULE_cycle)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Pathjoin()
	}
	{
		p.SetState(195)
		p.Match(GalleryParserCYCLE)
	}

	return localctx
}

// ISecondaryContext is an interface to support dynamic dispatch.
type ISecondaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSecondaryContext differentiates from other interfaces.
	IsSecondaryContext()
}

type SecondaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySecondaryContext() *SecondaryContext {
	var p = new(SecondaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_secondary
	return p
}

func (*SecondaryContext) IsSecondaryContext() {}

func NewSecondaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SecondaryContext {
	var p = new(SecondaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_secondary

	return p
}

func (s *SecondaryContext) GetParser() antlr.Parser { return s.parser }

func (s *SecondaryContext) CopyFrom(ctx *SecondaryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SecondaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecondaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TransformContext struct {
	*SecondaryContext
}

func NewTransformContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TransformContext {
	var p = new(TransformContext)

	p.SecondaryContext = NewEmptySecondaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SecondaryContext))

	return p
}

func (s *TransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransformContext) Secondary() ISecondaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISecondaryContext)
}

func (s *TransformContext) AllTRANSFORM() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserTRANSFORM)
}

func (s *TransformContext) TRANSFORM(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserTRANSFORM, i)
}

func (s *TransformContext) AllPrimary() []IPrimaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryContext)(nil)).Elem())
	var tst = make([]IPrimaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryContext)
		}
	}

	return tst
}

func (s *TransformContext) Primary(i int) IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *TransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterTransform(s)
	}
}

func (s *TransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitTransform(s)
	}
}

type FactorContext struct {
	*SecondaryContext
}

func NewFactorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FactorContext {
	var p = new(FactorContext)

	p.SecondaryContext = NewEmptySecondaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SecondaryContext))

	return p
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *FactorContext) Secondary() ISecondaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISecondaryContext)
}

func (s *FactorContext) TIMES() antlr.TerminalNode {
	return s.GetToken(GalleryParserTIMES, 0)
}

func (s *FactorContext) OVER() antlr.TerminalNode {
	return s.GetToken(GalleryParserOVER, 0)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *GalleryParser) Secondary() (localctx ISecondaryContext) {
	return p.secondary(0)
}

func (p *GalleryParser) secondary(_p int) (localctx ISecondaryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSecondaryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISecondaryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, GalleryParserRULE_secondary, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewFactorContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(198)
		p.Primary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFactorContext(p, NewSecondaryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GalleryParserRULE_secondary)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(201)
				_la = p.GetTokenStream().LA(1)

				if !(_la == GalleryParserTIMES || _la == GalleryParserOVER) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(202)
					p.Primary()
				}

			case 2:
				localctx = NewTransformContext(p, NewSecondaryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GalleryParserRULE_secondary)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				p.SetState(206)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(204)
							p.Match(GalleryParserTRANSFORM)
						}
						{
							p.SetState(205)
							p.Primary()
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(208)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
				}

			}

		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) CopyFrom(ctx *PrimaryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EdgeconstraintContext struct {
	*PrimaryContext
}

func NewEdgeconstraintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EdgeconstraintContext {
	var p = new(EdgeconstraintContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *EdgeconstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeconstraintContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *EdgeconstraintContext) AllEDGECONSTR() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserEDGECONSTR)
}

func (s *EdgeconstraintContext) EDGECONSTR(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserEDGECONSTR, i)
}

func (s *EdgeconstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterEdgeconstraint(s)
	}
}

func (s *EdgeconstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitEdgeconstraint(s)
	}
}

type InterpolationContext struct {
	*PrimaryContext
}

func NewInterpolationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InterpolationContext {
	var p = new(InterpolationContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *InterpolationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterpolationContext) Numtokenatom() INumtokenatomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumtokenatomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumtokenatomContext)
}

func (s *InterpolationContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GalleryParserLBRACKET, 0)
}

func (s *InterpolationContext) AllTertiary() []ITertiaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITertiaryContext)(nil)).Elem())
	var tst = make([]ITertiaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITertiaryContext)
		}
	}

	return tst
}

func (s *InterpolationContext) Tertiary(i int) ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *InterpolationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GalleryParserCOMMA, 0)
}

func (s *InterpolationContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GalleryParserRBRACKET, 0)
}

func (s *InterpolationContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *InterpolationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterInterpolation(s)
	}
}

func (s *InterpolationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitInterpolation(s)
	}
}

type SimpleatomContext struct {
	*PrimaryContext
}

func NewSimpleatomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleatomContext {
	var p = new(SimpleatomContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *SimpleatomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleatomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *SimpleatomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterSimpleatom(s)
	}
}

func (s *SimpleatomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitSimpleatom(s)
	}
}

type FuncatomContext struct {
	*PrimaryContext
}

func NewFuncatomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncatomContext {
	var p = new(FuncatomContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *FuncatomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncatomContext) MATHFUNC() antlr.TerminalNode {
	return s.GetToken(GalleryParserMATHFUNC, 0)
}

func (s *FuncatomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *FuncatomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterFuncatom(s)
	}
}

func (s *FuncatomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitFuncatom(s)
	}
}

type PointofContext struct {
	*PrimaryContext
}

func NewPointofContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PointofContext {
	var p = new(PointofContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *PointofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointofContext) POINT() antlr.TerminalNode {
	return s.GetToken(GalleryParserPOINT, 0)
}

func (s *PointofContext) Tertiary() ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *PointofContext) OF() antlr.TerminalNode {
	return s.GetToken(GalleryParserOF, 0)
}

func (s *PointofContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PointofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterPointof(s)
	}
}

func (s *PointofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitPointof(s)
	}
}

type SubpathContext struct {
	*PrimaryContext
}

func NewSubpathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubpathContext {
	var p = new(SubpathContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *SubpathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubpathContext) SUBPATH() antlr.TerminalNode {
	return s.GetToken(GalleryParserSUBPATH, 0)
}

func (s *SubpathContext) Tertiary() ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *SubpathContext) OF() antlr.TerminalNode {
	return s.GetToken(GalleryParserOF, 0)
}

func (s *SubpathContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *SubpathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterSubpath(s)
	}
}

func (s *SubpathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitSubpath(s)
	}
}

type PairpartContext struct {
	*PrimaryContext
}

func NewPairpartContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PairpartContext {
	var p = new(PairpartContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *PairpartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairpartContext) PAIRPART() antlr.TerminalNode {
	return s.GetToken(GalleryParserPAIRPART, 0)
}

func (s *PairpartContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PairpartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterPairpart(s)
	}
}

func (s *PairpartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitPairpart(s)
	}
}

type BoxContext struct {
	*PrimaryContext
}

func NewBoxContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoxContext {
	var p = new(BoxContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *BoxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoxContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *BoxContext) FRAME() antlr.TerminalNode {
	return s.GetToken(GalleryParserFRAME, 0)
}

func (s *BoxContext) BOX() antlr.TerminalNode {
	return s.GetToken(GalleryParserBOX, 0)
}

func (s *BoxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterBox(s)
	}
}

func (s *BoxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitBox(s)
	}
}

type ReversepathContext struct {
	*PrimaryContext
}

func NewReversepathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReversepathContext {
	var p = new(ReversepathContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *ReversepathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReversepathContext) REVERSE() antlr.TerminalNode {
	return s.GetToken(GalleryParserREVERSE, 0)
}

func (s *ReversepathContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ReversepathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterReversepath(s)
	}
}

func (s *ReversepathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitReversepath(s)
	}
}

type EdgepathContext struct {
	*PrimaryContext
}

func NewEdgepathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EdgepathContext {
	var p = new(EdgepathContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *EdgepathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgepathContext) EDGECONSTR() antlr.TerminalNode {
	return s.GetToken(GalleryParserEDGECONSTR, 0)
}

func (s *EdgepathContext) EDGE() antlr.TerminalNode {
	return s.GetToken(GalleryParserEDGE, 0)
}

func (s *EdgepathContext) Secondary() ISecondaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISecondaryContext)
}

func (s *EdgepathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterEdgepath(s)
	}
}

func (s *EdgepathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitEdgepath(s)
	}
}

type ScalaratomContext struct {
	*PrimaryContext
}

func NewScalaratomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScalaratomContext {
	var p = new(ScalaratomContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *ScalaratomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalaratomContext) Scalarmulop() IScalarmulopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarmulopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarmulopContext)
}

func (s *ScalaratomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ScalaratomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterScalaratom(s)
	}
}

func (s *ScalaratomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitScalaratom(s)
	}
}

func (p *GalleryParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GalleryParserRULE_primary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFuncatomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.Match(GalleryParserMATHFUNC)
		}
		{
			p.SetState(216)
			p.Atom()
		}

	case 2:
		localctx = NewScalaratomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Scalarmulop()
		}
		{
			p.SetState(218)
			p.Atom()
		}

	case 3:
		localctx = NewInterpolationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.Numtokenatom()
		}
		{
			p.SetState(221)
			p.Match(GalleryParserLBRACKET)
		}
		{
			p.SetState(222)
			p.tertiary(0)
		}
		{
			p.SetState(223)
			p.Match(GalleryParserCOMMA)
		}
		{
			p.SetState(224)
			p.tertiary(0)
		}
		{
			p.SetState(225)
			p.Match(GalleryParserRBRACKET)
		}

	case 4:
		localctx = NewInterpolationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(227)
			p.Atom()
		}
		{
			p.SetState(228)
			p.Match(GalleryParserLBRACKET)
		}
		{
			p.SetState(229)
			p.tertiary(0)
		}
		{
			p.SetState(230)
			p.Match(GalleryParserCOMMA)
		}
		{
			p.SetState(231)
			p.tertiary(0)
		}
		{
			p.SetState(232)
			p.Match(GalleryParserRBRACKET)
		}

	case 5:
		localctx = NewSimpleatomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(234)
			p.Atom()
		}

	case 6:
		localctx = NewPairpartContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(235)
			p.Match(GalleryParserPAIRPART)
		}
		{
			p.SetState(236)
			p.Primary()
		}

	case 7:
		localctx = NewPointofContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(237)
			p.Match(GalleryParserPOINT)
		}
		{
			p.SetState(238)
			p.tertiary(0)
		}
		{
			p.SetState(239)
			p.Match(GalleryParserOF)
		}
		{
			p.SetState(240)
			p.Primary()
		}

	case 8:
		localctx = NewReversepathContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(242)
			p.Match(GalleryParserREVERSE)
		}
		{
			p.SetState(243)
			p.Primary()
		}

	case 9:
		localctx = NewSubpathContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(244)
			p.Match(GalleryParserSUBPATH)
		}
		{
			p.SetState(245)
			p.tertiary(0)
		}
		{
			p.SetState(246)
			p.Match(GalleryParserOF)
		}
		{
			p.SetState(247)
			p.Primary()
		}

	case 10:
		localctx = NewEdgeconstraintContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(249)
					p.Match(GalleryParserEDGECONSTR)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
		}
		{
			p.SetState(254)
			p.Primary()
		}

	case 11:
		localctx = NewBoxContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		p.SetState(255)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GalleryParserFRAME || _la == GalleryParserBOX) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(256)
			p.Variable()
		}

	case 12:
		localctx = NewEdgepathContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(257)
			p.Match(GalleryParserEDGECONSTR)
		}
		{
			p.SetState(258)
			p.Match(GalleryParserEDGE)
		}
		{
			p.SetState(259)
			p.secondary(0)
		}

	}

	return localctx
}

// IScalarmulopContext is an interface to support dynamic dispatch.
type IScalarmulopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarmulopContext differentiates from other interfaces.
	IsScalarmulopContext()
}

type ScalarmulopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarmulopContext() *ScalarmulopContext {
	var p = new(ScalarmulopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_scalarmulop
	return p
}

func (*ScalarmulopContext) IsScalarmulopContext() {}

func NewScalarmulopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarmulopContext {
	var p = new(ScalarmulopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_scalarmulop

	return p
}

func (s *ScalarmulopContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarmulopContext) Numtokenatom() INumtokenatomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumtokenatomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumtokenatomContext)
}

func (s *ScalarmulopContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GalleryParserPLUS, 0)
}

func (s *ScalarmulopContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GalleryParserMINUS, 0)
}

func (s *ScalarmulopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarmulopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarmulopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterScalarmulop(s)
	}
}

func (s *ScalarmulopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitScalarmulop(s)
	}
}

func (p *GalleryParser) Scalarmulop() (localctx IScalarmulopContext) {
	localctx = NewScalarmulopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GalleryParserRULE_scalarmulop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GalleryParserPLUS || _la == GalleryParserMINUS {
		p.SetState(262)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GalleryParserPLUS || _la == GalleryParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	{
		p.SetState(265)
		p.Numtokenatom()
	}

	return localctx
}

// INumtokenatomContext is an interface to support dynamic dispatch.
type INumtokenatomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumtokenatomContext differentiates from other interfaces.
	IsNumtokenatomContext()
}

type NumtokenatomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumtokenatomContext() *NumtokenatomContext {
	var p = new(NumtokenatomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_numtokenatom
	return p
}

func (*NumtokenatomContext) IsNumtokenatomContext() {}

func NewNumtokenatomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumtokenatomContext {
	var p = new(NumtokenatomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_numtokenatom

	return p
}

func (s *NumtokenatomContext) GetParser() antlr.Parser { return s.parser }

func (s *NumtokenatomContext) AllDECIMALTOKEN() []antlr.TerminalNode {
	return s.GetTokens(GalleryParserDECIMALTOKEN)
}

func (s *NumtokenatomContext) DECIMALTOKEN(i int) antlr.TerminalNode {
	return s.GetToken(GalleryParserDECIMALTOKEN, i)
}

func (s *NumtokenatomContext) OVER() antlr.TerminalNode {
	return s.GetToken(GalleryParserOVER, 0)
}

func (s *NumtokenatomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumtokenatomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumtokenatomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterNumtokenatom(s)
	}
}

func (s *NumtokenatomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitNumtokenatom(s)
	}
}

func (p *GalleryParser) Numtokenatom() (localctx INumtokenatomContext) {
	localctx = NewNumtokenatomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GalleryParserRULE_numtokenatom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)
			p.Match(GalleryParserDECIMALTOKEN)
		}
		{
			p.SetState(268)
			p.Match(GalleryParserOVER)
		}
		{
			p.SetState(269)
			p.Match(GalleryParserDECIMALTOKEN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			p.Match(GalleryParserDECIMALTOKEN)
		}

	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VaratomContext struct {
	*AtomContext
}

func NewVaratomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VaratomContext {
	var p = new(VaratomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *VaratomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VaratomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VaratomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterVaratom(s)
	}
}

func (s *VaratomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitVaratom(s)
	}
}

type ExprgroupContext struct {
	*AtomContext
}

func NewExprgroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprgroupContext {
	var p = new(ExprgroupContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ExprgroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprgroupContext) BEGINGROUP() antlr.TerminalNode {
	return s.GetToken(GalleryParserBEGINGROUP, 0)
}

func (s *ExprgroupContext) Statementlist() IStatementlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementlistContext)
}

func (s *ExprgroupContext) Tertiary() ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *ExprgroupContext) ENDGROUP() antlr.TerminalNode {
	return s.GetToken(GalleryParserENDGROUP, 0)
}

func (s *ExprgroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterExprgroup(s)
	}
}

func (s *ExprgroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitExprgroup(s)
	}
}

type DecimalContext struct {
	*AtomContext
}

func NewDecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalContext {
	var p = new(DecimalContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *DecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalContext) DECIMALTOKEN() antlr.TerminalNode {
	return s.GetToken(GalleryParserDECIMALTOKEN, 0)
}

func (s *DecimalContext) UNIT() antlr.TerminalNode {
	return s.GetToken(GalleryParserUNIT, 0)
}

func (s *DecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterDecimal(s)
	}
}

func (s *DecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitDecimal(s)
	}
}

type SubexpressionContext struct {
	*AtomContext
}

func NewSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubexpressionContext {
	var p = new(SubexpressionContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *SubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubexpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GalleryParserLPAREN, 0)
}

func (s *SubexpressionContext) Tertiary() ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *SubexpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GalleryParserRPAREN, 0)
}

func (s *SubexpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterSubexpression(s)
	}
}

func (s *SubexpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitSubexpression(s)
	}
}

type LiteralpairContext struct {
	*AtomContext
}

func NewLiteralpairContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralpairContext {
	var p = new(LiteralpairContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *LiteralpairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralpairContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GalleryParserLPAREN, 0)
}

func (s *LiteralpairContext) AllTertiary() []ITertiaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITertiaryContext)(nil)).Elem())
	var tst = make([]ITertiaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITertiaryContext)
		}
	}

	return tst
}

func (s *LiteralpairContext) Tertiary(i int) ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *LiteralpairContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GalleryParserCOMMA, 0)
}

func (s *LiteralpairContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GalleryParserRPAREN, 0)
}

func (s *LiteralpairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterLiteralpair(s)
	}
}

func (s *LiteralpairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitLiteralpair(s)
	}
}

func (p *GalleryParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GalleryParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDecimalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.Match(GalleryParserDECIMALTOKEN)
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(274)
				p.Match(GalleryParserUNIT)
			}

		}

	case 2:
		localctx = NewVaratomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Variable()
		}

	case 3:
		localctx = NewLiteralpairContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(278)
			p.Match(GalleryParserLPAREN)
		}
		{
			p.SetState(279)
			p.tertiary(0)
		}
		{
			p.SetState(280)
			p.Match(GalleryParserCOMMA)
		}
		{
			p.SetState(281)
			p.tertiary(0)
		}
		{
			p.SetState(282)
			p.Match(GalleryParserRPAREN)
		}

	case 4:
		localctx = NewSubexpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(284)
			p.Match(GalleryParserLPAREN)
		}
		{
			p.SetState(285)
			p.tertiary(0)
		}
		{
			p.SetState(286)
			p.Match(GalleryParserRPAREN)
		}

	case 5:
		localctx = NewExprgroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(288)
			p.Match(GalleryParserBEGINGROUP)
		}
		{
			p.SetState(289)
			p.Statementlist()
		}
		{
			p.SetState(290)
			p.tertiary(0)
		}
		{
			p.SetState(291)
			p.Match(GalleryParserENDGROUP)
		}

	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) MIXEDTAG() antlr.TerminalNode {
	return s.GetToken(GalleryParserMIXEDTAG, 0)
}

func (s *VariableContext) AllSubscript() []ISubscriptContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubscriptContext)(nil)).Elem())
	var tst = make([]ISubscriptContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubscriptContext)
		}
	}

	return tst
}

func (s *VariableContext) Subscript(i int) ISubscriptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubscriptContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubscriptContext)
}

func (s *VariableContext) AllAnytag() []IAnytagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnytagContext)(nil)).Elem())
	var tst = make([]IAnytagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnytagContext)
		}
	}

	return tst
}

func (s *VariableContext) Anytag(i int) IAnytagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnytagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnytagContext)
}

func (s *VariableContext) TAG() antlr.TerminalNode {
	return s.GetToken(GalleryParserTAG, 0)
}

func (s *VariableContext) LAMBDAARG() antlr.TerminalNode {
	return s.GetToken(GalleryParserLAMBDAARG, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *GalleryParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GalleryParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(312)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GalleryParserMIXEDTAG:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)
			p.Match(GalleryParserMIXEDTAG)
		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(298)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case GalleryParserLBRACKET, GalleryParserDECIMALTOKEN:
					{
						p.SetState(296)
						p.Subscript()
					}

				case GalleryParserTAG, GalleryParserMIXEDTAG:
					{
						p.SetState(297)
						p.Anytag()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			}
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
		}

	case GalleryParserTAG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(303)
			p.Match(GalleryParserTAG)
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(306)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case GalleryParserLBRACKET, GalleryParserDECIMALTOKEN:
					{
						p.SetState(304)
						p.Subscript()
					}

				case GalleryParserTAG, GalleryParserMIXEDTAG:
					{
						p.SetState(305)
						p.Anytag()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			}
			p.SetState(310)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
		}

	case GalleryParserLAMBDAARG:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(311)
			p.Match(GalleryParserLAMBDAARG)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISubscriptContext is an interface to support dynamic dispatch.
type ISubscriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubscriptContext differentiates from other interfaces.
	IsSubscriptContext()
}

type SubscriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubscriptContext() *SubscriptContext {
	var p = new(SubscriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_subscript
	return p
}

func (*SubscriptContext) IsSubscriptContext() {}

func NewSubscriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubscriptContext {
	var p = new(SubscriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_subscript

	return p
}

func (s *SubscriptContext) GetParser() antlr.Parser { return s.parser }

func (s *SubscriptContext) DECIMALTOKEN() antlr.TerminalNode {
	return s.GetToken(GalleryParserDECIMALTOKEN, 0)
}

func (s *SubscriptContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GalleryParserLBRACKET, 0)
}

func (s *SubscriptContext) Tertiary() ITertiaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITertiaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITertiaryContext)
}

func (s *SubscriptContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GalleryParserRBRACKET, 0)
}

func (s *SubscriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubscriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubscriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterSubscript(s)
	}
}

func (s *SubscriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitSubscript(s)
	}
}

func (p *GalleryParser) Subscript() (localctx ISubscriptContext) {
	localctx = NewSubscriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GalleryParserRULE_subscript)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(319)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GalleryParserDECIMALTOKEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(314)
			p.Match(GalleryParserDECIMALTOKEN)
		}

	case GalleryParserLBRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.Match(GalleryParserLBRACKET)
		}
		{
			p.SetState(316)
			p.tertiary(0)
		}
		{
			p.SetState(317)
			p.Match(GalleryParserRBRACKET)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAnytagContext is an interface to support dynamic dispatch.
type IAnytagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnytagContext differentiates from other interfaces.
	IsAnytagContext()
}

type AnytagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnytagContext() *AnytagContext {
	var p = new(AnytagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GalleryParserRULE_anytag
	return p
}

func (*AnytagContext) IsAnytagContext() {}

func NewAnytagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnytagContext {
	var p = new(AnytagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GalleryParserRULE_anytag

	return p
}

func (s *AnytagContext) GetParser() antlr.Parser { return s.parser }

func (s *AnytagContext) TAG() antlr.TerminalNode {
	return s.GetToken(GalleryParserTAG, 0)
}

func (s *AnytagContext) MIXEDTAG() antlr.TerminalNode {
	return s.GetToken(GalleryParserMIXEDTAG, 0)
}

func (s *AnytagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnytagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnytagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.EnterAnytag(s)
	}
}

func (s *AnytagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GalleryListener); ok {
		listenerT.ExitAnytag(s)
	}
}

func (p *GalleryParser) Anytag() (localctx IAnytagContext) {
	localctx = NewAnytagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GalleryParserRULE_anytag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(321)
	_la = p.GetTokenStream().LA(1)

	if !(_la == GalleryParserTAG || _la == GalleryParserMIXEDTAG) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

func (p *GalleryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 14:
		var t *TertiaryContext = nil
		if localctx != nil {
			t = localctx.(*TertiaryContext)
		}
		return p.Tertiary_Sempred(t, predIndex)

	case 17:
		var t *SecondaryContext = nil
		if localctx != nil {
			t = localctx.(*SecondaryContext)
		}
		return p.Secondary_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GalleryParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GalleryParser) Tertiary_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GalleryParser) Secondary_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
