package main

import (
	"fmt"
	"math"
)

type Coordinates struct {
	x, y, z float64
}

func find_lenght(x1 float64, y1 float64, z1 float64, x2 float64, y2 float64, z2 float64) float64 {
	var lenght float64
	lenght = math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2) + math.Pow(z1-z2, 2))
	return lenght

}

func find_cos(len1 float64, len2 float64, len3 float64) float64 {
	var cos float64
	cos = (math.Pow(len1, 2) - math.Pow(len2, 2) + math.Pow(len3, 2)) / (2 * (len1 + len2))
	return cos
}

func main() {
	var r int
	var a, b, c, d Coordinates
	var lab, lac, lad, lbc, lbd, lcd, min1, min2, min3, min4, p float64

	fmt.Println("Choose dimension plane\n1. Two-dimension\n2. Three-dimension")
	fmt.Scanf("%d\n", &r)
	for r < 1 || r > 2 {
		fmt.Println("Error. Incorrect mode.")
		fmt.Println("Choose dimension plane\n1. Two-dimension\n2. Three-dimension")
		fmt.Scanf("%d\n", &r)
	}
	if r == 1 {
		fmt.Println("Insert coordinates (in format \"x y\"):")
		fmt.Println("Sector AB:")
		fmt.Println("Dot A:")
		fmt.Scanf("%f%f", &a.x, &a.y)
		fmt.Println("Dot B:")
		fmt.Scanf("%f%f", &b.x, &b.y)

		fmt.Println("Sector CD:")
		fmt.Println("Dot C:")
		fmt.Scanf("%f%f", &c.x, &c.y)
		fmt.Println("Dot D:")
		fmt.Scanf("%f%f", &d.x, &d.y)
		a.z = 0.0
		b.z = 0.0
		c.z = 0.0
		d.z = 0.0
	}
	if r == 2 {
		fmt.Println("Insert coordinates (in format \"x y z\"):")
		fmt.Println("Sector AB:")
		fmt.Println("Dot A:")
		fmt.Scanf("%f%f%f", &a.x, &a.y, &a.z)
		fmt.Println("Dot B:")
		fmt.Scanf("%f%f%f", &b.x, &b.y, &b.z)

		fmt.Println("Sector CD:")
		fmt.Println("Dot C:")
		fmt.Scanf("%f%f%f", &c.x, &c.y, &c.z)
		fmt.Println("Dot D:")
		fmt.Scanf("%f%f%f", &d.x, &d.y, &d.z)
	}

	lab = find_lenght(a.x, a.y, a.z, b.x, b.y, b.z) //Ищем длины расстояний от каждой точки до каждой
	lac = find_lenght(a.x, a.y, a.z, c.x, c.y, c.z)
	lad = find_lenght(a.x, a.y, a.z, d.x, d.y, d.z)
	lbc = find_lenght(b.x, b.y, b.z, c.x, c.y, c.z)
	lbd = find_lenght(b.x, b.y, b.z, d.x, d.y, d.z)
	lcd = find_lenght(c.x, c.y, c.z, d.x, d.y, d.z)

	if find_cos(lab, lac, lbc) > 0 && find_cos(lab, lbc, lac) > 0 { //Условие cos>0 означает, что угол треугольника острый.
		p = (lab + lac + lbc) / 2                               //Если два угла основания острые, тогда можем посчитать
		min1 = (2 * math.Sqrt(p*(p-lab)*(p-lac)*(p-lbc))) / lab //минимальное рассотяние (высоту) по трём сторонам треугольника
	}
	if (find_cos(lab, lac, lbc) > 0 && find_cos(lab, lbc, lac) <= 0) || (find_cos(lab, lac, lbc) <= 0 && find_cos(lab, lbc, lac) > 0) { //Если один из углов 90 градусов или больше,
		if lbc < lac {
			min1 = lbc
		} else { //то просто сравниваем между собой прилежащие к этому углу катеты.
			min1 = lac
		} //Более короткий катет записываем в ответ

	}

	if find_cos(lab, lad, lbd) > 0 && find_cos(lab, lbd, lad) > 0 { //Алалогично
		p = (lab + lad + lbd) / 2
		min2 = (2 * math.Sqrt(p*(p-lab)*(p-lad)*(p-lbd))) / lab
		if min1 > min2 {
			min1 = min2
		} //Сравниваем полученный ранее результат с новыми вычислениями, наименьший результат идёт в ответ
	}
	if (find_cos(lab, lad, lbd) > 0 && find_cos(lab, lbd, lad) <= 0) || (find_cos(lab, lad, lbd) <= 0 && find_cos(lab, lbd, lad) > 0) {
		if lbd < lad {
			min2 = lbd
		} else {
			min2 = lad
		}
		if min1 > min2 {
			min1 = min2
		}
	}

	if find_cos(lcd, lbc, lbd) > 0 && find_cos(lcd, lbd, lbc) > 0 {
		p = (lcd + lbc + lbd) / 2
		min3 = (2 * math.Sqrt(p*(p-lcd)*(p-lbc)*(p-lbd))) / lcd
		if min1 > min3 {
			min1 = min3
		}
	}
	if (find_cos(lcd, lbc, lbd) > 0 && find_cos(lcd, lbd, lbc) <= 0) || (find_cos(lcd, lbc, lbd) <= 0 && find_cos(lcd, lbd, lbc) > 0) {
		if lbd < lbc {
			min3 = lbd
		} else {
			min3 = lbc
		}
		if min1 > min3 {
			min1 = min3
		}
	}

	if find_cos(lcd, lac, lad) > 0 && find_cos(lcd, lad, lac) > 0 {
		p = (lcd + lac + lad) / 2
		min4 = (2 * math.Sqrt(p*(p-lcd)*(p-lac)*(p-lad))) / lcd
		if min1 > min4 {
			min1 = min4
		}
	}
	if (find_cos(lcd, lac, lad) > 0 && find_cos(lcd, lad, lac) <= 0) || (find_cos(lcd, lac, lad) <= 0 && find_cos(lcd, lad, lac) > 0) {
		if lad < lac {
			min4 = lad
		} else {
			min4 = lac
		}
		if min1 > min4 {
			min1 = min4
		}

	}
	fmt.Println("Minimal distance:") //Вывод ответа на экран
	fmt.Printf("%f", min1)
}
