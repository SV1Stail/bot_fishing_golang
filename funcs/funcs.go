package funcs

import (
	"fmt"
	"image"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	"gocv.io/x/gocv"
)

func Write_mouse_pos_in_console() {
	for {
		x, y := robotgo.Location()
		fmt.Println(x, y)
		time.Sleep(100 * time.Millisecond)
		if x == 0 {
			Move_mouse_on_pers_center()
		}
	}
}

func Mouse_left_long_rand_click(min, max float64) {
	robotgo.MouseDown("left")
	time.Sleep(time.Duration(RandFloat64(min, max) * float64(time.Second)))
	robotgo.MouseUp("left")
}

func Move_mouse_on_pers_center() {
	Move_mouse_slow(800, 366)
}

// возвращает случайное число от минимального до максимального
func Rand_number(min_numb, max_numb int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max_numb-min_numb+1) + min_numb
}
func RandFloat64(minNumb, maxNumb float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(maxNumb-minNumb) + minNumb
}

// передвигает мышку медленно в заданную координату
func Move_mouse_slow(x, y int) {
	x1, y1 := robotgo.Location()
	for x != x1 || y != y1 {
		if x > x1 {
			x1++
		} else if x < x1 {
			x1--
		}
		if y > y1 {
			y1++
		} else if y < y1 {
			y1--
		}
		robotgo.Move(x1, y1)
		time.Sleep(time.Duration(float64(time.Microsecond) * 2))
	}

}

// захватить экран и получить матрицу для gocv
func Capture_screen() gocv.Mat {
	bitmap := robotgo.CaptureScreen()
	if bitmap == nil {
		return gocv.NewMat()
	}
	defer robotgo.FreeBitmap(bitmap)
	img := robotgo.ToImage(bitmap)
	mat, err := gocv.ImageToMatRGB(img)
	if err != nil {
		fmt.Println("Error converting image to Mat:", err)
		return gocv.NewMat()
	}
	return mat
}

// вырезать кусок по заданным координатам заданного размера
func CropImage(img gocv.Mat, x, y, width, height int) gocv.Mat {
	if img.Empty() {
		panic("Пустая матрица изображения")
	}
	var rect image.Rectangle
	// Определяем область интереса
	if x+width > 1600 && y+height > 800 {
		rect = image.Rect(x, y, 1600, 800)
	} else if x+width > 1600 {
		rect = image.Rect(x, y, 1600, y+height)
	} else if y+height > 800 {
		rect = image.Rect(x, y, x+width, 800)
	} else {
		rect = image.Rect(x, y, x+width, y+height)
	}

	// Обрезаем изображение
	croppedImg := img.Region(rect)

	return croppedImg
}

// 2 | 1
// --|--
// 3 | 4
// возвращает положение красной точки поплавка (в идеале поплавка)
// trturn x,y,error
func Find_poplavok(img gocv.Mat, quarter int) (int, int, error) {

	min_x := 0
	min_y := 0
	max_x := 0
	max_y := 0

	switch quarter {
	case 1:
		min_x = img.Cols() / 2
		min_y = 0
		max_x = img.Cols()
		max_y = img.Rows() / 2
	case 2:
		min_x = 0
		min_y = 0
		max_x = img.Cols() / 2
		max_y = img.Rows() / 2
	case 3:
		min_x = 0
		min_y = img.Rows() / 2
		max_x = img.Cols()/2 - 100
		max_y = img.Rows()
	default:
		min_x = img.Cols()/2 + 100
		min_y = img.Rows() / 2
		max_x = img.Cols()
		max_y = img.Rows()
	}
	for y := min_y; y < max_y; y++ {
		for x := min_x; x < max_x; x++ {
			// Получаем значения каналов BGR
			vecRGB := img.GetVecbAt(y, x)
			// Проверяем значение красного канала
			if int(vecRGB[2]) > 240 && int(vecRGB[1]) < 100 && int(vecRGB[0]) < 180 {
				if x >= 706 && y >= 530 && y <= 560 && x <= 890 {
					continue
				}
				return x, y, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("cant find")
}
func Kanny(input interface{}) (gocv.Mat, error) {
	var template gocv.Mat
	switch path := input.(type) {
	case string:
		template = gocv.IMRead(path, gocv.IMReadColor)
		defer template.Close()
		if template.Empty() {
			return gocv.NewMat(), fmt.Errorf("no image")
		}
	case gocv.Mat:
		template = path.Clone()
		defer template.Close()
		if template.Empty() {
			return gocv.NewMat(), fmt.Errorf("no image")
		}
	default:
		return gocv.NewMat(), fmt.Errorf("unsupported input type")
	}
	defer template.Close()
	grey := gocv.NewMat()
	defer grey.Close()
	gocv.CvtColor(template, &grey, gocv.ColorBGRToGray)

	edges := gocv.NewMat()
	defer edges.Close()
	gocv.Canny(grey, &edges, 35.0, 75.0)
	defer edges.Close()
	Bigger_edges(edges)

	if ok := gocv.IMWrite("edges.jpg", edges); !ok {
		return gocv.NewMat(), fmt.Errorf("cant save")
	}
	return edges.Clone(), nil
}

func Bigger_edges(img gocv.Mat) {
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(2, 1))
	defer kernel.Close()
	gocv.Dilate(img, &img, kernel)
}
