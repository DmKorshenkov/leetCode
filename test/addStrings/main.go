package main

import (
	"bytes"
)

func main() {
	print(
		func(num1 string, num2 string) string {
			amount := 0
			x := 1

			for in := len(num1) - 1; in > -1; in-- {
				amount += int(num1[in]-48) * x
				x *= 10
			}
			x = 1
			for in := len(num2) - 1; in > -1; in-- {
				amount += int(num2[in]-48) * x
				x *= 10
			}
			bb := bytes.Buffer{}

			x = 10
			for {

				bb.WriteByte(byte(amount%x + 48))
				amount /= 10
				if amount == 0 {
					b2 := bytes.Buffer{}
					for in := len(bb.Bytes()) - 1; in > -1; in-- {
						b2.WriteByte(bb.Bytes()[in])
					}
					return b2.String()
				}
			}

		}("49", "51"))
}
