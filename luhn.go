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
//

package goluhn

func LuhnValidate(str string) bool {
	l := len(str) - 1
	return l > 0 && LuhnChecksum(str[:l]) == str[l:]
}

func LuhnChecksum(str string) string {
	var sum uint
	p := ^len(str) & 0x1
	for i, r := range str {
		// Convert to number and check if the supplied values are not digits
		if r -= '0'; uint(r) > 9 {
			return ""
		}
		if i&p != 0 { // Double the number in case of even iteration
			if r <<= 1; r > 9 {
				r -= 9 // Sum the digits
			}
		}
		sum += uint(r)
	}
	return string((sum*9)%10 + '0')
}
