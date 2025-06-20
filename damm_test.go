package damm_test

import (
	"fmt"
	"testing"

	"github.com/go-oss/damm"
)

// Ordnung 64: x*y:=2x+y mit Rechnung in GF(2^6)
// http://www.md-software.de/math/DAMM_Quasigruppen.txt
var matrix64 = [64][64]int{
	{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 3, 1, 7, 5, 11, 9, 15, 13, 19, 17, 23, 21, 27, 25, 31, 29, 35, 33, 39, 37, 43, 41, 47, 45, 51, 49, 55, 53, 59, 57, 63, 61},
	{2, 0, 6, 4, 10, 8, 14, 12, 18, 16, 22, 20, 26, 24, 30, 28, 34, 32, 38, 36, 42, 40, 46, 44, 50, 48, 54, 52, 58, 56, 62, 60, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57, 59, 61, 63},
	{4, 6, 0, 2, 12, 14, 8, 10, 20, 22, 16, 18, 28, 30, 24, 26, 36, 38, 32, 34, 44, 46, 40, 42, 52, 54, 48, 50, 60, 62, 56, 58, 7, 5, 3, 1, 15, 13, 11, 9, 23, 21, 19, 17, 31, 29, 27, 25, 39, 37, 35, 33, 47, 45, 43, 41, 55, 53, 51, 49, 63, 61, 59, 57},
	{6, 4, 2, 0, 14, 12, 10, 8, 22, 20, 18, 16, 30, 28, 26, 24, 38, 36, 34, 32, 46, 44, 42, 40, 54, 52, 50, 48, 62, 60, 58, 56, 5, 7, 1, 3, 13, 15, 9, 11, 21, 23, 17, 19, 29, 31, 25, 27, 37, 39, 33, 35, 45, 47, 41, 43, 53, 55, 49, 51, 61, 63, 57, 59},
	{8, 10, 12, 14, 0, 2, 4, 6, 24, 26, 28, 30, 16, 18, 20, 22, 40, 42, 44, 46, 32, 34, 36, 38, 56, 58, 60, 62, 48, 50, 52, 54, 11, 9, 15, 13, 3, 1, 7, 5, 27, 25, 31, 29, 19, 17, 23, 21, 43, 41, 47, 45, 35, 33, 39, 37, 59, 57, 63, 61, 51, 49, 55, 53},
	{10, 8, 14, 12, 2, 0, 6, 4, 26, 24, 30, 28, 18, 16, 22, 20, 42, 40, 46, 44, 34, 32, 38, 36, 58, 56, 62, 60, 50, 48, 54, 52, 9, 11, 13, 15, 1, 3, 5, 7, 25, 27, 29, 31, 17, 19, 21, 23, 41, 43, 45, 47, 33, 35, 37, 39, 57, 59, 61, 63, 49, 51, 53, 55},
	{12, 14, 8, 10, 4, 6, 0, 2, 28, 30, 24, 26, 20, 22, 16, 18, 44, 46, 40, 42, 36, 38, 32, 34, 60, 62, 56, 58, 52, 54, 48, 50, 15, 13, 11, 9, 7, 5, 3, 1, 31, 29, 27, 25, 23, 21, 19, 17, 47, 45, 43, 41, 39, 37, 35, 33, 63, 61, 59, 57, 55, 53, 51, 49},
	{14, 12, 10, 8, 6, 4, 2, 0, 30, 28, 26, 24, 22, 20, 18, 16, 46, 44, 42, 40, 38, 36, 34, 32, 62, 60, 58, 56, 54, 52, 50, 48, 13, 15, 9, 11, 5, 7, 1, 3, 29, 31, 25, 27, 21, 23, 17, 19, 45, 47, 41, 43, 37, 39, 33, 35, 61, 63, 57, 59, 53, 55, 49, 51},
	{16, 18, 20, 22, 24, 26, 28, 30, 0, 2, 4, 6, 8, 10, 12, 14, 48, 50, 52, 54, 56, 58, 60, 62, 32, 34, 36, 38, 40, 42, 44, 46, 19, 17, 23, 21, 27, 25, 31, 29, 3, 1, 7, 5, 11, 9, 15, 13, 51, 49, 55, 53, 59, 57, 63, 61, 35, 33, 39, 37, 43, 41, 47, 45},
	{18, 16, 22, 20, 26, 24, 30, 28, 2, 0, 6, 4, 10, 8, 14, 12, 50, 48, 54, 52, 58, 56, 62, 60, 34, 32, 38, 36, 42, 40, 46, 44, 17, 19, 21, 23, 25, 27, 29, 31, 1, 3, 5, 7, 9, 11, 13, 15, 49, 51, 53, 55, 57, 59, 61, 63, 33, 35, 37, 39, 41, 43, 45, 47},
	{20, 22, 16, 18, 28, 30, 24, 26, 4, 6, 0, 2, 12, 14, 8, 10, 52, 54, 48, 50, 60, 62, 56, 58, 36, 38, 32, 34, 44, 46, 40, 42, 23, 21, 19, 17, 31, 29, 27, 25, 7, 5, 3, 1, 15, 13, 11, 9, 55, 53, 51, 49, 63, 61, 59, 57, 39, 37, 35, 33, 47, 45, 43, 41},
	{22, 20, 18, 16, 30, 28, 26, 24, 6, 4, 2, 0, 14, 12, 10, 8, 54, 52, 50, 48, 62, 60, 58, 56, 38, 36, 34, 32, 46, 44, 42, 40, 21, 23, 17, 19, 29, 31, 25, 27, 5, 7, 1, 3, 13, 15, 9, 11, 53, 55, 49, 51, 61, 63, 57, 59, 37, 39, 33, 35, 45, 47, 41, 43},
	{24, 26, 28, 30, 16, 18, 20, 22, 8, 10, 12, 14, 0, 2, 4, 6, 56, 58, 60, 62, 48, 50, 52, 54, 40, 42, 44, 46, 32, 34, 36, 38, 27, 25, 31, 29, 19, 17, 23, 21, 11, 9, 15, 13, 3, 1, 7, 5, 59, 57, 63, 61, 51, 49, 55, 53, 43, 41, 47, 45, 35, 33, 39, 37},
	{26, 24, 30, 28, 18, 16, 22, 20, 10, 8, 14, 12, 2, 0, 6, 4, 58, 56, 62, 60, 50, 48, 54, 52, 42, 40, 46, 44, 34, 32, 38, 36, 25, 27, 29, 31, 17, 19, 21, 23, 9, 11, 13, 15, 1, 3, 5, 7, 57, 59, 61, 63, 49, 51, 53, 55, 41, 43, 45, 47, 33, 35, 37, 39},
	{28, 30, 24, 26, 20, 22, 16, 18, 12, 14, 8, 10, 4, 6, 0, 2, 60, 62, 56, 58, 52, 54, 48, 50, 44, 46, 40, 42, 36, 38, 32, 34, 31, 29, 27, 25, 23, 21, 19, 17, 15, 13, 11, 9, 7, 5, 3, 1, 63, 61, 59, 57, 55, 53, 51, 49, 47, 45, 43, 41, 39, 37, 35, 33},
	{30, 28, 26, 24, 22, 20, 18, 16, 14, 12, 10, 8, 6, 4, 2, 0, 62, 60, 58, 56, 54, 52, 50, 48, 46, 44, 42, 40, 38, 36, 34, 32, 29, 31, 25, 27, 21, 23, 17, 19, 13, 15, 9, 11, 5, 7, 1, 3, 61, 63, 57, 59, 53, 55, 49, 51, 45, 47, 41, 43, 37, 39, 33, 35},
	{32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 35, 33, 39, 37, 43, 41, 47, 45, 51, 49, 55, 53, 59, 57, 63, 61, 3, 1, 7, 5, 11, 9, 15, 13, 19, 17, 23, 21, 27, 25, 31, 29},
	{34, 32, 38, 36, 42, 40, 46, 44, 50, 48, 54, 52, 58, 56, 62, 60, 2, 0, 6, 4, 10, 8, 14, 12, 18, 16, 22, 20, 26, 24, 30, 28, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57, 59, 61, 63, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31},
	{36, 38, 32, 34, 44, 46, 40, 42, 52, 54, 48, 50, 60, 62, 56, 58, 4, 6, 0, 2, 12, 14, 8, 10, 20, 22, 16, 18, 28, 30, 24, 26, 39, 37, 35, 33, 47, 45, 43, 41, 55, 53, 51, 49, 63, 61, 59, 57, 7, 5, 3, 1, 15, 13, 11, 9, 23, 21, 19, 17, 31, 29, 27, 25},
	{38, 36, 34, 32, 46, 44, 42, 40, 54, 52, 50, 48, 62, 60, 58, 56, 6, 4, 2, 0, 14, 12, 10, 8, 22, 20, 18, 16, 30, 28, 26, 24, 37, 39, 33, 35, 45, 47, 41, 43, 53, 55, 49, 51, 61, 63, 57, 59, 5, 7, 1, 3, 13, 15, 9, 11, 21, 23, 17, 19, 29, 31, 25, 27},
	{40, 42, 44, 46, 32, 34, 36, 38, 56, 58, 60, 62, 48, 50, 52, 54, 8, 10, 12, 14, 0, 2, 4, 6, 24, 26, 28, 30, 16, 18, 20, 22, 43, 41, 47, 45, 35, 33, 39, 37, 59, 57, 63, 61, 51, 49, 55, 53, 11, 9, 15, 13, 3, 1, 7, 5, 27, 25, 31, 29, 19, 17, 23, 21},
	{42, 40, 46, 44, 34, 32, 38, 36, 58, 56, 62, 60, 50, 48, 54, 52, 10, 8, 14, 12, 2, 0, 6, 4, 26, 24, 30, 28, 18, 16, 22, 20, 41, 43, 45, 47, 33, 35, 37, 39, 57, 59, 61, 63, 49, 51, 53, 55, 9, 11, 13, 15, 1, 3, 5, 7, 25, 27, 29, 31, 17, 19, 21, 23},
	{44, 46, 40, 42, 36, 38, 32, 34, 60, 62, 56, 58, 52, 54, 48, 50, 12, 14, 8, 10, 4, 6, 0, 2, 28, 30, 24, 26, 20, 22, 16, 18, 47, 45, 43, 41, 39, 37, 35, 33, 63, 61, 59, 57, 55, 53, 51, 49, 15, 13, 11, 9, 7, 5, 3, 1, 31, 29, 27, 25, 23, 21, 19, 17},
	{46, 44, 42, 40, 38, 36, 34, 32, 62, 60, 58, 56, 54, 52, 50, 48, 14, 12, 10, 8, 6, 4, 2, 0, 30, 28, 26, 24, 22, 20, 18, 16, 45, 47, 41, 43, 37, 39, 33, 35, 61, 63, 57, 59, 53, 55, 49, 51, 13, 15, 9, 11, 5, 7, 1, 3, 29, 31, 25, 27, 21, 23, 17, 19},
	{48, 50, 52, 54, 56, 58, 60, 62, 32, 34, 36, 38, 40, 42, 44, 46, 16, 18, 20, 22, 24, 26, 28, 30, 0, 2, 4, 6, 8, 10, 12, 14, 51, 49, 55, 53, 59, 57, 63, 61, 35, 33, 39, 37, 43, 41, 47, 45, 19, 17, 23, 21, 27, 25, 31, 29, 3, 1, 7, 5, 11, 9, 15, 13},
	{50, 48, 54, 52, 58, 56, 62, 60, 34, 32, 38, 36, 42, 40, 46, 44, 18, 16, 22, 20, 26, 24, 30, 28, 2, 0, 6, 4, 10, 8, 14, 12, 49, 51, 53, 55, 57, 59, 61, 63, 33, 35, 37, 39, 41, 43, 45, 47, 17, 19, 21, 23, 25, 27, 29, 31, 1, 3, 5, 7, 9, 11, 13, 15},
	{52, 54, 48, 50, 60, 62, 56, 58, 36, 38, 32, 34, 44, 46, 40, 42, 20, 22, 16, 18, 28, 30, 24, 26, 4, 6, 0, 2, 12, 14, 8, 10, 55, 53, 51, 49, 63, 61, 59, 57, 39, 37, 35, 33, 47, 45, 43, 41, 23, 21, 19, 17, 31, 29, 27, 25, 7, 5, 3, 1, 15, 13, 11, 9},
	{54, 52, 50, 48, 62, 60, 58, 56, 38, 36, 34, 32, 46, 44, 42, 40, 22, 20, 18, 16, 30, 28, 26, 24, 6, 4, 2, 0, 14, 12, 10, 8, 53, 55, 49, 51, 61, 63, 57, 59, 37, 39, 33, 35, 45, 47, 41, 43, 21, 23, 17, 19, 29, 31, 25, 27, 5, 7, 1, 3, 13, 15, 9, 11},
	{56, 58, 60, 62, 48, 50, 52, 54, 40, 42, 44, 46, 32, 34, 36, 38, 24, 26, 28, 30, 16, 18, 20, 22, 8, 10, 12, 14, 0, 2, 4, 6, 59, 57, 63, 61, 51, 49, 55, 53, 43, 41, 47, 45, 35, 33, 39, 37, 27, 25, 31, 29, 19, 17, 23, 21, 11, 9, 15, 13, 3, 1, 7, 5},
	{58, 56, 62, 60, 50, 48, 54, 52, 42, 40, 46, 44, 34, 32, 38, 36, 26, 24, 30, 28, 18, 16, 22, 20, 10, 8, 14, 12, 2, 0, 6, 4, 57, 59, 61, 63, 49, 51, 53, 55, 41, 43, 45, 47, 33, 35, 37, 39, 25, 27, 29, 31, 17, 19, 21, 23, 9, 11, 13, 15, 1, 3, 5, 7},
	{60, 62, 56, 58, 52, 54, 48, 50, 44, 46, 40, 42, 36, 38, 32, 34, 28, 30, 24, 26, 20, 22, 16, 18, 12, 14, 8, 10, 4, 6, 0, 2, 63, 61, 59, 57, 55, 53, 51, 49, 47, 45, 43, 41, 39, 37, 35, 33, 31, 29, 27, 25, 23, 21, 19, 17, 15, 13, 11, 9, 7, 5, 3, 1},
	{62, 60, 58, 56, 54, 52, 50, 48, 46, 44, 42, 40, 38, 36, 34, 32, 30, 28, 26, 24, 22, 20, 18, 16, 14, 12, 10, 8, 6, 4, 2, 0, 61, 63, 57, 59, 53, 55, 49, 51, 45, 47, 41, 43, 37, 39, 33, 35, 29, 31, 25, 27, 21, 23, 17, 19, 13, 15, 9, 11, 5, 7, 1, 3},
	{3, 1, 7, 5, 11, 9, 15, 13, 19, 17, 23, 21, 27, 25, 31, 29, 35, 33, 39, 37, 43, 41, 47, 45, 51, 49, 55, 53, 59, 57, 63, 61, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62},
	{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57, 59, 61, 63, 2, 0, 6, 4, 10, 8, 14, 12, 18, 16, 22, 20, 26, 24, 30, 28, 34, 32, 38, 36, 42, 40, 46, 44, 50, 48, 54, 52, 58, 56, 62, 60},
	{7, 5, 3, 1, 15, 13, 11, 9, 23, 21, 19, 17, 31, 29, 27, 25, 39, 37, 35, 33, 47, 45, 43, 41, 55, 53, 51, 49, 63, 61, 59, 57, 4, 6, 0, 2, 12, 14, 8, 10, 20, 22, 16, 18, 28, 30, 24, 26, 36, 38, 32, 34, 44, 46, 40, 42, 52, 54, 48, 50, 60, 62, 56, 58},
	{5, 7, 1, 3, 13, 15, 9, 11, 21, 23, 17, 19, 29, 31, 25, 27, 37, 39, 33, 35, 45, 47, 41, 43, 53, 55, 49, 51, 61, 63, 57, 59, 6, 4, 2, 0, 14, 12, 10, 8, 22, 20, 18, 16, 30, 28, 26, 24, 38, 36, 34, 32, 46, 44, 42, 40, 54, 52, 50, 48, 62, 60, 58, 56},
	{11, 9, 15, 13, 3, 1, 7, 5, 27, 25, 31, 29, 19, 17, 23, 21, 43, 41, 47, 45, 35, 33, 39, 37, 59, 57, 63, 61, 51, 49, 55, 53, 8, 10, 12, 14, 0, 2, 4, 6, 24, 26, 28, 30, 16, 18, 20, 22, 40, 42, 44, 46, 32, 34, 36, 38, 56, 58, 60, 62, 48, 50, 52, 54},
	{9, 11, 13, 15, 1, 3, 5, 7, 25, 27, 29, 31, 17, 19, 21, 23, 41, 43, 45, 47, 33, 35, 37, 39, 57, 59, 61, 63, 49, 51, 53, 55, 10, 8, 14, 12, 2, 0, 6, 4, 26, 24, 30, 28, 18, 16, 22, 20, 42, 40, 46, 44, 34, 32, 38, 36, 58, 56, 62, 60, 50, 48, 54, 52},
	{15, 13, 11, 9, 7, 5, 3, 1, 31, 29, 27, 25, 23, 21, 19, 17, 47, 45, 43, 41, 39, 37, 35, 33, 63, 61, 59, 57, 55, 53, 51, 49, 12, 14, 8, 10, 4, 6, 0, 2, 28, 30, 24, 26, 20, 22, 16, 18, 44, 46, 40, 42, 36, 38, 32, 34, 60, 62, 56, 58, 52, 54, 48, 50},
	{13, 15, 9, 11, 5, 7, 1, 3, 29, 31, 25, 27, 21, 23, 17, 19, 45, 47, 41, 43, 37, 39, 33, 35, 61, 63, 57, 59, 53, 55, 49, 51, 14, 12, 10, 8, 6, 4, 2, 0, 30, 28, 26, 24, 22, 20, 18, 16, 46, 44, 42, 40, 38, 36, 34, 32, 62, 60, 58, 56, 54, 52, 50, 48},
	{19, 17, 23, 21, 27, 25, 31, 29, 3, 1, 7, 5, 11, 9, 15, 13, 51, 49, 55, 53, 59, 57, 63, 61, 35, 33, 39, 37, 43, 41, 47, 45, 16, 18, 20, 22, 24, 26, 28, 30, 0, 2, 4, 6, 8, 10, 12, 14, 48, 50, 52, 54, 56, 58, 60, 62, 32, 34, 36, 38, 40, 42, 44, 46},
	{17, 19, 21, 23, 25, 27, 29, 31, 1, 3, 5, 7, 9, 11, 13, 15, 49, 51, 53, 55, 57, 59, 61, 63, 33, 35, 37, 39, 41, 43, 45, 47, 18, 16, 22, 20, 26, 24, 30, 28, 2, 0, 6, 4, 10, 8, 14, 12, 50, 48, 54, 52, 58, 56, 62, 60, 34, 32, 38, 36, 42, 40, 46, 44},
	{23, 21, 19, 17, 31, 29, 27, 25, 7, 5, 3, 1, 15, 13, 11, 9, 55, 53, 51, 49, 63, 61, 59, 57, 39, 37, 35, 33, 47, 45, 43, 41, 20, 22, 16, 18, 28, 30, 24, 26, 4, 6, 0, 2, 12, 14, 8, 10, 52, 54, 48, 50, 60, 62, 56, 58, 36, 38, 32, 34, 44, 46, 40, 42},
	{21, 23, 17, 19, 29, 31, 25, 27, 5, 7, 1, 3, 13, 15, 9, 11, 53, 55, 49, 51, 61, 63, 57, 59, 37, 39, 33, 35, 45, 47, 41, 43, 22, 20, 18, 16, 30, 28, 26, 24, 6, 4, 2, 0, 14, 12, 10, 8, 54, 52, 50, 48, 62, 60, 58, 56, 38, 36, 34, 32, 46, 44, 42, 40},
	{27, 25, 31, 29, 19, 17, 23, 21, 11, 9, 15, 13, 3, 1, 7, 5, 59, 57, 63, 61, 51, 49, 55, 53, 43, 41, 47, 45, 35, 33, 39, 37, 24, 26, 28, 30, 16, 18, 20, 22, 8, 10, 12, 14, 0, 2, 4, 6, 56, 58, 60, 62, 48, 50, 52, 54, 40, 42, 44, 46, 32, 34, 36, 38},
	{25, 27, 29, 31, 17, 19, 21, 23, 9, 11, 13, 15, 1, 3, 5, 7, 57, 59, 61, 63, 49, 51, 53, 55, 41, 43, 45, 47, 33, 35, 37, 39, 26, 24, 30, 28, 18, 16, 22, 20, 10, 8, 14, 12, 2, 0, 6, 4, 58, 56, 62, 60, 50, 48, 54, 52, 42, 40, 46, 44, 34, 32, 38, 36},
	{31, 29, 27, 25, 23, 21, 19, 17, 15, 13, 11, 9, 7, 5, 3, 1, 63, 61, 59, 57, 55, 53, 51, 49, 47, 45, 43, 41, 39, 37, 35, 33, 28, 30, 24, 26, 20, 22, 16, 18, 12, 14, 8, 10, 4, 6, 0, 2, 60, 62, 56, 58, 52, 54, 48, 50, 44, 46, 40, 42, 36, 38, 32, 34},
	{29, 31, 25, 27, 21, 23, 17, 19, 13, 15, 9, 11, 5, 7, 1, 3, 61, 63, 57, 59, 53, 55, 49, 51, 45, 47, 41, 43, 37, 39, 33, 35, 30, 28, 26, 24, 22, 20, 18, 16, 14, 12, 10, 8, 6, 4, 2, 0, 62, 60, 58, 56, 54, 52, 50, 48, 46, 44, 42, 40, 38, 36, 34, 32},
	{35, 33, 39, 37, 43, 41, 47, 45, 51, 49, 55, 53, 59, 57, 63, 61, 3, 1, 7, 5, 11, 9, 15, 13, 19, 17, 23, 21, 27, 25, 31, 29, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30},
	{33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57, 59, 61, 63, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 34, 32, 38, 36, 42, 40, 46, 44, 50, 48, 54, 52, 58, 56, 62, 60, 2, 0, 6, 4, 10, 8, 14, 12, 18, 16, 22, 20, 26, 24, 30, 28},
	{39, 37, 35, 33, 47, 45, 43, 41, 55, 53, 51, 49, 63, 61, 59, 57, 7, 5, 3, 1, 15, 13, 11, 9, 23, 21, 19, 17, 31, 29, 27, 25, 36, 38, 32, 34, 44, 46, 40, 42, 52, 54, 48, 50, 60, 62, 56, 58, 4, 6, 0, 2, 12, 14, 8, 10, 20, 22, 16, 18, 28, 30, 24, 26},
	{37, 39, 33, 35, 45, 47, 41, 43, 53, 55, 49, 51, 61, 63, 57, 59, 5, 7, 1, 3, 13, 15, 9, 11, 21, 23, 17, 19, 29, 31, 25, 27, 38, 36, 34, 32, 46, 44, 42, 40, 54, 52, 50, 48, 62, 60, 58, 56, 6, 4, 2, 0, 14, 12, 10, 8, 22, 20, 18, 16, 30, 28, 26, 24},
	{43, 41, 47, 45, 35, 33, 39, 37, 59, 57, 63, 61, 51, 49, 55, 53, 11, 9, 15, 13, 3, 1, 7, 5, 27, 25, 31, 29, 19, 17, 23, 21, 40, 42, 44, 46, 32, 34, 36, 38, 56, 58, 60, 62, 48, 50, 52, 54, 8, 10, 12, 14, 0, 2, 4, 6, 24, 26, 28, 30, 16, 18, 20, 22},
	{41, 43, 45, 47, 33, 35, 37, 39, 57, 59, 61, 63, 49, 51, 53, 55, 9, 11, 13, 15, 1, 3, 5, 7, 25, 27, 29, 31, 17, 19, 21, 23, 42, 40, 46, 44, 34, 32, 38, 36, 58, 56, 62, 60, 50, 48, 54, 52, 10, 8, 14, 12, 2, 0, 6, 4, 26, 24, 30, 28, 18, 16, 22, 20},
	{47, 45, 43, 41, 39, 37, 35, 33, 63, 61, 59, 57, 55, 53, 51, 49, 15, 13, 11, 9, 7, 5, 3, 1, 31, 29, 27, 25, 23, 21, 19, 17, 44, 46, 40, 42, 36, 38, 32, 34, 60, 62, 56, 58, 52, 54, 48, 50, 12, 14, 8, 10, 4, 6, 0, 2, 28, 30, 24, 26, 20, 22, 16, 18},
	{45, 47, 41, 43, 37, 39, 33, 35, 61, 63, 57, 59, 53, 55, 49, 51, 13, 15, 9, 11, 5, 7, 1, 3, 29, 31, 25, 27, 21, 23, 17, 19, 46, 44, 42, 40, 38, 36, 34, 32, 62, 60, 58, 56, 54, 52, 50, 48, 14, 12, 10, 8, 6, 4, 2, 0, 30, 28, 26, 24, 22, 20, 18, 16},
	{51, 49, 55, 53, 59, 57, 63, 61, 35, 33, 39, 37, 43, 41, 47, 45, 19, 17, 23, 21, 27, 25, 31, 29, 3, 1, 7, 5, 11, 9, 15, 13, 48, 50, 52, 54, 56, 58, 60, 62, 32, 34, 36, 38, 40, 42, 44, 46, 16, 18, 20, 22, 24, 26, 28, 30, 0, 2, 4, 6, 8, 10, 12, 14},
	{49, 51, 53, 55, 57, 59, 61, 63, 33, 35, 37, 39, 41, 43, 45, 47, 17, 19, 21, 23, 25, 27, 29, 31, 1, 3, 5, 7, 9, 11, 13, 15, 50, 48, 54, 52, 58, 56, 62, 60, 34, 32, 38, 36, 42, 40, 46, 44, 18, 16, 22, 20, 26, 24, 30, 28, 2, 0, 6, 4, 10, 8, 14, 12},
	{55, 53, 51, 49, 63, 61, 59, 57, 39, 37, 35, 33, 47, 45, 43, 41, 23, 21, 19, 17, 31, 29, 27, 25, 7, 5, 3, 1, 15, 13, 11, 9, 52, 54, 48, 50, 60, 62, 56, 58, 36, 38, 32, 34, 44, 46, 40, 42, 20, 22, 16, 18, 28, 30, 24, 26, 4, 6, 0, 2, 12, 14, 8, 10},
	{53, 55, 49, 51, 61, 63, 57, 59, 37, 39, 33, 35, 45, 47, 41, 43, 21, 23, 17, 19, 29, 31, 25, 27, 5, 7, 1, 3, 13, 15, 9, 11, 54, 52, 50, 48, 62, 60, 58, 56, 38, 36, 34, 32, 46, 44, 42, 40, 22, 20, 18, 16, 30, 28, 26, 24, 6, 4, 2, 0, 14, 12, 10, 8},
	{59, 57, 63, 61, 51, 49, 55, 53, 43, 41, 47, 45, 35, 33, 39, 37, 27, 25, 31, 29, 19, 17, 23, 21, 11, 9, 15, 13, 3, 1, 7, 5, 56, 58, 60, 62, 48, 50, 52, 54, 40, 42, 44, 46, 32, 34, 36, 38, 24, 26, 28, 30, 16, 18, 20, 22, 8, 10, 12, 14, 0, 2, 4, 6},
	{57, 59, 61, 63, 49, 51, 53, 55, 41, 43, 45, 47, 33, 35, 37, 39, 25, 27, 29, 31, 17, 19, 21, 23, 9, 11, 13, 15, 1, 3, 5, 7, 58, 56, 62, 60, 50, 48, 54, 52, 42, 40, 46, 44, 34, 32, 38, 36, 26, 24, 30, 28, 18, 16, 22, 20, 10, 8, 14, 12, 2, 0, 6, 4},
	{63, 61, 59, 57, 55, 53, 51, 49, 47, 45, 43, 41, 39, 37, 35, 33, 31, 29, 27, 25, 23, 21, 19, 17, 15, 13, 11, 9, 7, 5, 3, 1, 60, 62, 56, 58, 52, 54, 48, 50, 44, 46, 40, 42, 36, 38, 32, 34, 28, 30, 24, 26, 20, 22, 16, 18, 12, 14, 8, 10, 4, 6, 0, 2},
	{61, 63, 57, 59, 53, 55, 49, 51, 45, 47, 41, 43, 37, 39, 33, 35, 29, 31, 25, 27, 21, 23, 17, 19, 13, 15, 9, 11, 5, 7, 1, 3, 62, 60, 58, 56, 54, 52, 50, 48, 46, 44, 42, 40, 38, 36, 34, 32, 30, 28, 26, 24, 22, 20, 18, 16, 14, 12, 10, 8, 6, 4, 2, 0},
}

func damm64(digits []int) (checkDigit int) {
	for _, digit := range digits {
		checkDigit = matrix64[checkDigit][digit]
	}
	return checkDigit
}

func TestDamm64_matrix(t *testing.T) {
	t.Parallel()
	d64 := damm.New64()
	m := genMatrix(d64)
	assertWeaklyTotallyAntiSymmetric(t, m, d64.Modulus())
	for i := range matrix64 {
		for j := range matrix64[i] {
			if got := m[i][j]; got != matrix64[i][j] {
				t.Fatalf("d64.Generate([]int{r[%d], %d}) = %d; want %d", i, j, got, matrix64[i][j])
			}
		}
	}
}

func TestDamm32_matrix(t *testing.T) {
	t.Parallel()
	d32 := damm.New32()
	m := genMatrix(d32)
	assertWeaklyTotallyAntiSymmetric(t, m, d32.Modulus())
}

func TestDamm64_Generate(t *testing.T) {
	t.Parallel()
	d64 := damm.New64()
	tests := []struct {
		digits []int
	}{
		{digits: []int{0}},
		{digits: []int{0, 1}},
		{digits: []int{0, 1, 2}},
		{digits: []int{0, 1, 2, 3}},
		{digits: []int{0, 1, 2, 3, 4}},
		{digits: []int{0, 1, 2, 3, 4, 5}},
		{digits: []int{0, 1, 2, 3, 4, 5, 6}},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{digits: testDigits64},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.digits), func(t *testing.T) {
			got := d64.Generate(tt.digits)
			want := damm64(tt.digits)
			if got != want {
				t.Errorf("d64.Generate(%v) = %d; want %d", tt.digits, got, want)
			}
		})
	}
}

func TestDamm64_Verify(t *testing.T) {
	t.Parallel()
	d64 := damm.New64()
	tests := []struct {
		digits []int
		want   bool
	}{
		{digits: []int{0}, want: true},
		{digits: []int{0, 0}, want: true},
		{digits: []int{0, 1, 2}, want: true},
		{digits: []int{0, 1, 2, 0}, want: true},
		{digits: []int{0, 1, 2, 3, 6}, want: true},
		{digits: []int{0, 1, 2, 3, 4, 4}, want: true},
		{digits: []int{0, 1, 2, 3, 4, 5, 2}, want: true},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 8}, want: true},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7, 30}, want: true},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 44}, want: true},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9}, want: true},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 6}, want: true},
		{digits: append(testDigits64, 50), want: true},
		{digits: []int{0, 1}, want: false},
		{digits: []int{0, 1, 3}, want: false},
		{digits: []int{0, 1, 2, 3}, want: false},
		{digits: []int{0, 1, 2, 3, 4}, want: false},
		{digits: []int{0, 1, 2, 3, 4, 5}, want: false},
		{digits: []int{0, 1, 2, 3, 4, 5, 6}, want: false},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7}, want: false},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, want: false},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, want: false},
		{digits: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.digits), func(t *testing.T) {
			got := d64.Verify(tt.digits)
			if got != tt.want {
				t.Errorf("d64.Generate(%v) = %v; want %v", tt.digits, got, tt.want)
			}
		})
	}
}

func assertWeaklyTotallyAntiSymmetric(t *testing.T, m [][]int, modulus int) {
	t.Helper()
	if len(m) != modulus {
		t.Fatalf("len(m) = %d; want %d", len(m), modulus)
	}
	for y := range modulus {
		if len(m[y]) != modulus {
			t.Fatalf("len(m[%d]) = %d; want %d", y, len(m[y]), modulus)
		}
		for x := range modulus {
			if m[y][x] < 0 || m[y][x] >= modulus {
				t.Fatalf("should be 0 < x * y <= %d: x=%d, y=%d, x * y=%d", modulus, x, y, m[y][x])
			}
			if y == x && m[y][x] != 0 {
				t.Fatalf("should be x * x = 0, x=%d, x * x = %d", x, m[y][x])
			}
		}
	}
	for c := range modulus {
		for y := range modulus {
			for x := range modulus {
				if y != x && m[m[c][y]][x] == m[m[c][x]][y] {
					t.Fatalf("should be (c * x) * y = (c * y) * x => x = y: c=%d, x=%d, y=%d, (c * x) * y=%d, (c * y) * x=%d", c, x, y, m[m[c][y]][x], m[m[c][x]][y])
				}
			}
		}
	}
}

func genMatrix(d damm.Damm) [][]int {
	m := make([][]int, d.Modulus())
	r := make([]int, d.Modulus())
	for i := range d.Modulus() {
		r[d.Generate([]int{i})] = i
	}
	for i := range d.Modulus() {
		m[i] = make([]int, d.Modulus())
		for j := range d.Modulus() {
			m[i][j] = d.Generate([]int{r[i], j})
		}
	}
	return m
}

var testDigits64 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63}

func BenchmarkDamm64Matrix(b *testing.B) {
	for range b.N {
		damm64(testDigits64)
	}
}

func BenchmarkDamm64(b *testing.B) {
	d64 := damm.New64()
	for range b.N {
		d64.Generate(testDigits64)
	}
}
