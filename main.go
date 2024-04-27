package main

import (
	"fmt"
	"image"
	"time"

	"fishing/funcs"

	"gocv.io/x/gocv"
)

// x, y := robotgo.Location() расположение мышки
// robotgo.Move(x,y) переместить мышку
// // Двойной клик правой кнопкой мыши
// robotgo.Click("right", true)

// // Нажатие клавиши 'enter' с модификатором 'alt'
// robotgo.KeyTap("enter", "alt")
// fmt.Println(time.Now().Format("15:04:05"))
func main() {
	// Загрузка исходного и целевого изображений
	// path := "images/Untitled.png"
	// path2 := "images/forest_rivers/Screenshot from 2024-04-23 18-41-44.png"
	time.Sleep(2 * time.Second)
	Fishing_down_left()

}
func findTemplate(screen, template gocv.Mat) (image.Point, bool) {
	result := gocv.NewMat()
	defer result.Close()

	// Параметры MatchTemplate
	method := gocv.TmCcoeffNormed

	gocv.MatchTemplate(screen, template, &result, method, gocv.NewMat())
	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)

	if maxVal >= 0.5 { // пороговое значение для определения, найден ли шаблон
		return maxLoc, true
	}

	return image.Point{}, false
}
func Fishing_down_left() error {
	// path := "images/Untitled.png"

	// template_m_g := gocv.IMRead(path, gocv.IMReadColor)

	prev_mat_popl := gocv.NewMat()
	i := 0
	for {
		i++
		fmt.Println("-------", i)
		// x_rand := funcs.Rand_number(60, 533)
		// y_rand := funcs.Rand_number(500, 800)
		// funcs.Move_mouse_slow(x_rand, y_rand)
		// funcs.Mouse_left_long_rand_click(0.3, 1.0)
		// time.Sleep(time.Duration(1.4 * float64(time.Second)))

		img := funcs.Capture_screen()

		x, y, err := funcs.Find_poplavok(img, 3)
		if err != nil {
			img.Close()
			fmt.Println(4)
			continue
		}
		popl_croped := funcs.CropImage(img, x-5, y-5, 40, 50)
		if !prev_mat_popl.Empty() {

			// fmt.Println(prev_mat_popl.Rows(), "!=", popl_croped.Rows(), prev_mat_popl.Cols(), "!=", popl_croped.Cols(), prev_mat_popl.Channels(), "!=", popl_croped.Channels())
			var sum_popl int = 0
			var meanDiff_popl float32 = 0
			if prev_mat_popl.Rows() == popl_croped.Rows() && prev_mat_popl.Cols() == popl_croped.Cols() && prev_mat_popl.Channels() == popl_croped.Channels() {
				sum_popl, meanDiff_popl = funcs.CompareImages(prev_mat_popl, popl_croped)
			}
			// начать тянуть удочку
			prev_mat_m_g := funcs.CropImage(img, 714, 454, 170, 30)
			if sum_popl > 35000 && meanDiff_popl >= 18 {
				funcs.Mouse_left_long_rand_click(0.05, 0.06)
				time.Sleep(time.Duration(1 * float64(time.Second)))
				img2 := funcs.Capture_screen()
				m_g_cropped := funcs.CropImage(img2, 714, 454, 170, 30)
				img2.Close()
				var sum_m_g int = 0
				var meanDiff_m_g float32 = 0
				fmt.Println("3")
				sum_m_g, meanDiff_m_g = funcs.CompareImages(prev_mat_m_g, m_g_cropped)
				fmt.Println("2")
				if !prev_mat_m_g.Empty() {
					gocv.IMWrite("1.jpg", prev_mat_m_g)
				}
				gocv.IMWrite("2.jpg", m_g_cropped)
				fmt.Println("m_g", sum_m_g, meanDiff_m_g)

				if sum_m_g > 190000 && meanDiff_m_g >= 39 {
					fmt.Println("kk")
					// for i := 0; i < 20; i++ {
					// 	topLeft, found := findTemplate(m_g_cropped, template_m_g)
					// 	if found {
					// 		fmt.Printf("Object found at: %v\n", topLeft)
					// 	} else {
					// 		fmt.Println("Object not found")
					// 	}
					// 	time.Sleep(time.Duration(0.5 * float64(time.Second)))

					// }
				}

				m_g_cropped.Close()
			}
			prev_mat_m_g.Close()
			fmt.Println("popl", sum_popl, meanDiff_popl)

			prev_mat_popl.Close()
		}
		img.Close()

		prev_mat_popl = popl_croped.Clone()
		// funcs.Move_mouse_slow(x, y)
		popl_croped.Close()
		// funcs.Mouse_left_long_rand_click(0.05, 0.06)
		time.Sleep(time.Duration(0.3 * float64(time.Second)))
	}
	return nil
}

// func Fishing_top() {
// 	time.Sleep(3 * time.Second)
// 	for {
// 		x_rand := funcs.Rand_number(533, 1066)
// 		y_rand := funcs.Rand_number(100, 280)
// 		funcs.Move_mouse_slow(x_rand, y_rand)
// 		funcs.Mouse_left_long_rand_click(0.3, 1.0)

// 		time.Sleep(time.Second)

// 	}
// }

// func Fishing_left() {
// 	funcs.Move_mouse_slow(funcs.Rand_number(1, 700), funcs.Rand_number(280, 500))
// 	funcs.Mouse_left_long_rand_click(0.3, 1.0)

// }
// func Fishing_right() {
// 	funcs.Move_mouse_slow(funcs.Rand_number(900, 1599), funcs.Rand_number(280, 500))
// 	funcs.Mouse_left_long_rand_click(0.3, 1.0)

// }

// func Fishing_down_right() {
// 	funcs.Move_mouse_slow(funcs.Rand_number(1066, 1450), funcs.Rand_number(500, 800))
// 	funcs.Mouse_left_long_rand_click(0.3, 1.0)

// }
// func Fishing_down() {
// 	funcs.Move_mouse_slow(funcs.Rand_number(533, 1066), funcs.Rand_number(500, 800))
// 	funcs.Mouse_left_long_rand_click(0.3, 1.0)

// }
// func Fishing_top_left() {
// 	funcs.Move_mouse_slow(funcs.Rand_number(60, 533), funcs.Rand_number(70, 280))
// 	funcs.Mouse_left_long_rand_click(0.3, 1.0)

// }
// func Fishing_top_right() {
// 	funcs.Move_mouse_slow(funcs.Rand_number(1066, 1450), funcs.Rand_number(70, 280))
// 	funcs.Mouse_left_long_rand_click(0.3, 1.0)

// }

// func setPixelToBlack(img *gocv.Mat, x int, y int) {
// 	// Проверка на выход за пределы изображения
// 	if x >= 0 && y >= 0 && x < img.Cols() && y < img.Rows() {
// 		// Установка каждого канала (B, G, R) в 0
// 		img.SetUCharAt(y, x*img.Channels(), 0)   // Blue channel
// 		img.SetUCharAt(y, x*img.Channels()+1, 0) // Green channel
// 		img.SetUCharAt(y, x*img.Channels()+2, 0) // Red channel
// 	}
// }
// func isBlueAround(img gocv.Mat, x, y int, targetBlueLevel int, tolerance int) bool {
// 	// Проверка границ изображения
// 	if x < 0 || y < 0 || x >= img.Cols() || y >= img.Rows() {
// 		return false
// 	}

// 	// Получение значения BGR для пикселя
// 	vec3b := img.GetVecbAt(y, x)
// 	blueValue := int(vec3b[0])

// 	// Сравнение синего канала с заданным уровнем и допуском
// 	return blueValue >= targetBlueLevel-tolerance
// }

// CompareImages вычисляет и возвращает суммарную разницу и среднее отклонение пикселей между двумя изображениями
