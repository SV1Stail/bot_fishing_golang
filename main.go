package main

import (
	"fmt"
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
	// path := "2.jpg"
	// path2 := "3.jpg"

	time.Sleep(3 * time.Second)
	Fishing_down_left()
	// if ok := Fishing_down_left(); ok != nil {
	// 	fmt.Println("fick")
	// 	return
	// }

}

func Fishing_down_left() error {
	prev_mat := gocv.NewMat()
	defer prev_mat.Close()
	for {
		// x_rand := funcs.Rand_number(60, 533)
		// y_rand := funcs.Rand_number(500, 800)
		// funcs.Move_mouse_slow(x_rand, y_rand)
		// funcs.Mouse_left_long_rand_click(0.3, 1.0)
		// time.Sleep(time.Duration(1.4 * float64(time.Second)))

		fmt.Println(1)
		img := funcs.Capture_screen()

		fmt.Println(2)
		x, y, err := funcs.Find_poplavok(img, 3)

		fmt.Println(3)
		if err != nil {
			img.Close()
			fmt.Println(4)
			continue
		}
		poplavok_croped := funcs.CropImage(img, x-5, y-5, 40, 50)

		fmt.Println(5)
		if !prev_mat.Empty() {
			// fmt.Println(prev_mat.Rows(), "!=", poplavok_croped.Rows(), prev_mat.Cols(), "!=", poplavok_croped.Cols(), prev_mat.Channels(), "!=", poplavok_croped.Channels())

			// sum := суммарную разницу (int),  meanDiff := среднее отклонение пикселей (float32)
			fmt.Println(8)
			var sum int = 0
			var meanDiff float32 = 0
			if prev_mat.Rows() == poplavok_croped.Rows() && prev_mat.Cols() == poplavok_croped.Cols() && prev_mat.Channels() == poplavok_croped.Channels() {
				sum, meanDiff = CompareImages(prev_mat, poplavok_croped)
				fmt.Println(9)
			}
			if sum > 35000 && meanDiff >= 20 {
				funcs.Mouse_left_long_rand_click(0.05, 0.06)
				fmt.Println(10)
			}
			fmt.Println(sum, meanDiff)

			prev_mat.Close() // Закрытие предыдущей матрицы

			fmt.Println(11)

		}

		prev_mat = poplavok_croped.Clone() // Безопасное клонирование новой матрицы

		fmt.Println(12)

		// funcs.Move_mouse_slow(x, y)
		img.Close()

		fmt.Println(13)
		poplavok_croped.Close()

		fmt.Println(14)
		// funcs.Mouse_left_long_rand_click(0.05, 0.06)
		time.Sleep(time.Duration(0.3 * float64(time.Second)))
	}
	return nil
}

func Fishing_top() {
	time.Sleep(3 * time.Second)
	for {
		x_rand := funcs.Rand_number(533, 1066)
		y_rand := funcs.Rand_number(100, 280)
		funcs.Move_mouse_slow(x_rand, y_rand)
		funcs.Mouse_left_long_rand_click(0.3, 1.0)

		time.Sleep(time.Second)

	}
}

func Fishing_left() {
	funcs.Move_mouse_slow(funcs.Rand_number(1, 700), funcs.Rand_number(280, 500))
	funcs.Mouse_left_long_rand_click(0.3, 1.0)

}
func Fishing_right() {
	funcs.Move_mouse_slow(funcs.Rand_number(900, 1599), funcs.Rand_number(280, 500))
	funcs.Mouse_left_long_rand_click(0.3, 1.0)

}

func Fishing_down_right() {
	funcs.Move_mouse_slow(funcs.Rand_number(1066, 1450), funcs.Rand_number(500, 800))
	funcs.Mouse_left_long_rand_click(0.3, 1.0)

}
func Fishing_down() {
	funcs.Move_mouse_slow(funcs.Rand_number(533, 1066), funcs.Rand_number(500, 800))
	funcs.Mouse_left_long_rand_click(0.3, 1.0)

}
func Fishing_top_left() {
	funcs.Move_mouse_slow(funcs.Rand_number(60, 533), funcs.Rand_number(70, 280))
	funcs.Mouse_left_long_rand_click(0.3, 1.0)

}
func Fishing_top_right() {
	funcs.Move_mouse_slow(funcs.Rand_number(1066, 1450), funcs.Rand_number(70, 280))
	funcs.Mouse_left_long_rand_click(0.3, 1.0)

}

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
func CompareImages(img1, img2 gocv.Mat) (int, float32) {
	if img1.Empty() || img2.Empty() {
		return 0, 0.0 // Возвращает 0, если одно из изображений пустое
	}

	// Создаем новую матрицу для хранения различий
	diff := gocv.NewMat()
	defer diff.Close()

	// Вычисляем разницу между двумя изображениями
	gocv.AbsDiff(img1, img2, &diff)

	// Вычисляем суммарное значение пикселей различий
	sum := diff.Sum()
	// Рассчитываем среднее значение разницы на пиксель
	totalPixels := diff.Rows() * diff.Cols()
	meanDiff := float32(sum.Val1) / float32(totalPixels)

	return int(sum.Val1), meanDiff
}
