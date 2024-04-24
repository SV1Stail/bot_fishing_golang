package main

import (
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
func main() {
	// Загрузка исходного и целевого изображений
	// path := "images/poplavok/Screenshot from 2024-04-24 11-54-58.png"
	// path2 := "images/image.png"
	time.Sleep(3 * time.Second)
	for {
		img := funcs.Capture_screen()
		if x, y, err := funcs.Find_poplavok(img, 3); err == nil {
			funcs.Move_mouse_slow(x, y)
			return
		}
		img.Close()
	}
}

// вырезать кусок по заданным координатам заданного размера
func CropImage(img gocv.Mat, x, y, width, height int) gocv.Mat {
	if img.Empty() {
		panic("Пустая матрица изображения")
	}

	// Определяем область интереса
	rect := image.Rect(x, y, x+width, y+height)

	// Обрезаем изображение
	croppedImg := img.Region(rect)

	return croppedImg
}
func Fishing_top() {
	time.Sleep(3 * time.Second)
	for {
		x_rand := funcs.Rand_number(533, 1066)
		y_rand := funcs.Rand_number(100, 280)
		funcs.Move_mouse_slow(x_rand, y_rand)
		funcs.Mouse_left_long_rand_click()
		time.Sleep(time.Second)

	}
}

func Fishing_left() {
	funcs.Move_mouse_slow(funcs.Rand_number(1, 700), funcs.Rand_number(280, 500))
	funcs.Mouse_left_long_rand_click()

}
func Fishing_right() {
	funcs.Move_mouse_slow(funcs.Rand_number(900, 1599), funcs.Rand_number(280, 500))
	funcs.Mouse_left_long_rand_click()

}
func Fishing_down_left() {
	funcs.Move_mouse_slow(funcs.Rand_number(60, 533), funcs.Rand_number(500, 800))
	funcs.Mouse_left_long_rand_click()

}
func Fishing_down_right() {
	funcs.Move_mouse_slow(funcs.Rand_number(1066, 1450), funcs.Rand_number(500, 800))
	funcs.Mouse_left_long_rand_click()

}
func Fishing_down() {
	funcs.Move_mouse_slow(funcs.Rand_number(533, 1066), funcs.Rand_number(500, 800))
	funcs.Mouse_left_long_rand_click()

}
func Fishing_top_left() {
	funcs.Move_mouse_slow(funcs.Rand_number(60, 533), funcs.Rand_number(70, 280))
	funcs.Mouse_left_long_rand_click()

}
func Fishing_top_right() {
	funcs.Move_mouse_slow(funcs.Rand_number(1066, 1450), funcs.Rand_number(70, 280))
	funcs.Mouse_left_long_rand_click()

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
