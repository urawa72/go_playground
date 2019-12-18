package main

import (
    "fmt"
    "flag"
    "image"
    "image/jpeg"
    "os"
    // "reflect"

    "golang.org/x/image/draw"
)

func main() {
  flag.Parse()
  if flag.NArg() != 2 {
     fmt.Fprintln(os.Stderr, os.ErrInvalid)
     return
  }

  // 引数のイメージファイルを開く
  file, err := os.Open(flag.Arg(0))
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
  }
  defer file.Close()

  // イメージをデコード
  img, t, err := image.Decode(file)
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
  }
  fmt.Println("Type of image:", t)

  // イメージのレクタングル情報
  rct := img.Bounds()
  fmt.Println("Width:", rct.Dx())
  fmt.Println("Height:", rct.Dy())

  imgDst := image.NewRGBA(image.Rect(0, 0, rct.Dx()/2, rct.Dy()/2))
  draw.CatmullRom.Scale(imgDst, imgDst.Bounds(), img, rct, draw.Over, nil)

  dst, err := os.Create(flag.Arg(1))
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
  }
  defer dst.Close()

  switch t {
  case "jpeg":
    if err := jpeg.Encode(dst, imgDst, &jpeg.Options{Quality: 100}); err != nil {
      fmt.Fprintln(os.Stderr, err)
      return
    }
  default:
    fmt.Fprintln(os.Stderr, "format error")
  }
}

