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
	// path := "images/poplavok/Screenshot from 2024-04-24 11-54-58.png"
	// path2 := "images/image.png"
	time.Sleep(3 * time.Second)
	if err := Fishing_down_left(); err != nil {
		fmt.Println(err)
		return
	}

}

func Fishing_down_left() error {
	for {
		x_rand := funcs.Rand_number(60, 533)
		y_rand := funcs.Rand_number(500, 800)
		funcs.Move_mouse_slow(x_rand, y_rand)
		fmt.Println("1")
		funcs.Mouse_left_long_rand_click(0.3, 1.0)
		time.Sleep(time.Duration(1.4 * float64(time.Second)))
		img := funcs.Capture_screen()
		x, y, err := funcs.Find_poplavok(img, 3)
		fmt.Println(err)
		if err == nil {
			poplavok_croped := funcs.CropImage(img, x-5, y-5, 40, 50)
			if ok := gocv.IMWrite("image.jpg", poplavok_croped); !ok {
				img.Close()
				poplavok_croped.Close()
				return fmt.Errorf("cant save")
			}
			funcs.Move_mouse_slow(x, y)
			img.Close()
			poplavok_croped.Close()
		}
		funcs.Mouse_left_long_rand_click(0.05, 0.06)
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
