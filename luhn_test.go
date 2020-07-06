//
// Copyright (C) 2020 Diego Augusto Molina
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 of the License.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package goluhn

import (
	"fmt"
	"testing"
)

var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func testByLength(data *string, dataLength int, check *string, checkLength int,
) string {
	for i := 0; i < dataLength-checkLength; i++ {
		str := (*data)[i : i+checkLength]
		digit := (*check)[i : i+1]
		calculated := LuhnChecksum(str)
		if calculated != digit {
			return fmt.Sprintf("Bad check digit: %v; Expected: %v\n; string: %v",
				calculated, digit, str)
		}
		for _, d := range digits {
			r := LuhnValidate(str + d)
			if (d != digit || !r) && (d == digit || r) {
				return fmt.Sprintf("Bad double check: %v; Current Digit: %v; Correct "+
					"Digit: %v; Validation: %v\n", str, d, digit, r)
			}
		}
	}
	return ""
}

func TestLuhnGoldenLen8(t *testing.T) {
	t.Parallel()
	load.Do(LoadTestData)

	s := testByLength(TestData, TestDataLen, TestDataGoldenLen8, 8)
	if len(s) > 0 {
		t.Fatal(s)
	}
}

func TestLuhnGoldenLen15(t *testing.T) {
	t.Parallel()
	load.Do(LoadTestData)

	s := testByLength(TestData, TestDataLen, TestDataGoldenLen15, 15)
	if len(s) > 0 {
		t.Fatal(s)
	}
}

func TestLuhnGolden(t *testing.T) {
	t.Parallel()
	load.Do(LoadTestData)

	s := testByLength(TestData, TestDataLen, TestDataGolden, TestDataLen)
	if len(s) > 0 {
		t.Fatal(s)
	}
}

func benchmarkValidateByLength(data *string, dataLength int, check *string,
	checkLength int) {
	for i := 0; i <= dataLength-checkLength; i++ {
		_ = LuhnValidate((*data)[i : i+checkLength])
	}
}

func BenchmarkLuhnLen8(b *testing.B) {
	b.StopTimer()
	load.Do(LoadTestData)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		benchmarkValidateByLength(TestData, TestDataLen,
			TestDataGoldenLen8, 8)
	}
	b.ReportMetric(float64(TestDataLen-7), "Validations/op")
	b.SetBytes(int64((8 * (TestDataLen - 7))))
}

func BenchmarkLuhnLen15(b *testing.B) {
	b.StopTimer()
	load.Do(LoadTestData)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		benchmarkValidateByLength(TestData, TestDataLen,
			TestDataGoldenLen15, 15)
	}
	b.ReportMetric(float64(TestDataLen-14), "Validations/op")
	b.SetBytes(int64((15 * (TestDataLen - 14))))
}

func BenchmarkLuhnFull(b *testing.B) {
	b.StopTimer()
	load.Do(LoadTestData)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		benchmarkValidateByLength(TestData, TestDataLen, TestDataGolden,
			TestDataLen)
	}
	b.ReportMetric(1, "Validations/op")
	b.SetBytes(int64((TestDataLen)))
}
