package golang

import (
	"fmt"
	"strconv"
	"time"
)

// a *int is similar to int &a in C++
func printResults(
	a *int,
	e *int,
	h *int,
	l *int,
	p *int,
	t *int,
	eat *int,
	that *int,
	apple *int) {
	fmt.Print("A: " + strconv.Itoa(*a) + " ")
	fmt.Print("E: " + strconv.Itoa(*e) + " ")
	fmt.Print("H: " + strconv.Itoa(*h) + " ")
	fmt.Print("L: " + strconv.Itoa(*l) + " ")
	fmt.Print("P: " + strconv.Itoa(*p) + " ")
	fmt.Print("T: " + strconv.Itoa(*t) + "\n")

	fmt.Print("EAT: " + strconv.Itoa(*eat) + " ")
	fmt.Print("THAT: " + strconv.Itoa(*that) + " ")
	fmt.Print("APPLE: " + strconv.Itoa(*apple) + "\n")
}

func digiEat(eat *int, e *int, a *int, t *int) {
	var rem int = *eat
	*e = rem / 100
	rem %= 100
	*a = rem / 10
	*t = rem % 10
}

func digiThat(that *int, t1 *int, h *int, a *int, t2 *int) {
	var rem int = *that
	*t1 = rem / 1000
	rem %= 1000
	*h = rem / 100
	rem %= 100
	*a = rem / 10
	*t2 = rem % 10
}

func digiApple(apple *int, a *int, p1 *int, p2 *int, l *int, e *int) {
	var rem int = *apple
	*a = rem / 10000
	rem %= 10000
	*p1 = rem / 1000
	rem %= 1000
	*p2 = rem / 100
	rem %= 100
	*l = rem / 10
	*e = rem % 10
}

func findApple() {
	var foundApple bool = false
	var eat, e1, a1, t1 int
	var that, t21, h2, a2, t22 int
	var apple, a3, p31, p32, l3, e3 int

	for eat = 100; eat < 1000; eat++ {
		digiEat(&eat, &e1, &a1, &t1)
		if e1 != a1 && e1 != t1 && a1 != t1 {
			for that = 1000; that < 10000; that++ {
				digiThat(&that, &t21, &h2, &a2, &t22)
				if t22 == t21 && t1 == t21 && a1 == a2 && t21 != h2 && h2 != a2 {
					if eat+that >= 10000 {
						apple = eat + that
						digiApple(&apple, &a3, &p31, &p32, &l3, &e3)
						if a1 == a3 && p31 == p32 && e1 == e3 && t21 != l3 && a3 != p31 && a3 != l3 && a3 != e3 && p31 != l3 && p31 != e3 && l3 != e3 {
							printResults(&a1, &e1, &h2, &l3, &p31, &t1, &eat, &that, &apple)
							foundApple = true
						} else {
							continue
						}
					} else {
						continue
					}
				} else {
					continue
				}

				if foundApple {
					break
				}
			}
		} else {
			continue
		}

		if foundApple {
			break
		}
	}

}

func Eat() {
	start := time.Now()
	findApple()
	execution := time.Since(start)
	printExecutionTime(execution)
}
