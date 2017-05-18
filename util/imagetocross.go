package util

func ImageJpegToCross(filename string, scale int, pathOut string) error {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  img, err := jpeg.Decode(file)
  if err != nil {
    panic(err)
  }

  outputImage := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
  for x := 0; x < img.Bounds().Dx(); x++ {
    for y := 0; y < img.Bounds().Dy(); y++ {
      outputImage.Set(x, y, img.At((x/scale)*scale, (y/scale)*scale))
      if x % scale == 0 {
        outputImage.Set(x, y, color.RGBA{0, 0, 0, 1})
      }
      if y % scale == 0 {
        outputImage.Set(x, y, color.RGBA{0, 0, 0, 1})
      }
    }
  }

  fileOutput = "image" + strconv.Itoa(scale) + ".jpeg"

  newFile, err := os.Create(pathOut + "image" + strconv.Itoa(scale) + ".jpeg")
  if err != nil {
    panic(err)
  }
  defer newFile.Close()

  err = jpeg.Encode(newFile, outputImage, nil)
  if err != nil {
    panic(err)
  }
  return err
}


func ImagePngToCross(filename string, scale int, pathOut string) error {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  img, err := png.Decode(file)
  if err != nil {
    panic(err)
  }

  outputImage := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
  for x := 0; x < img.Bounds().Dx(); x++ {
    for y := 0; y < img.Bounds().Dy(); y++ {
      outputImage.Set(x, y, img.At((x/scale)*scale, (y/scale)*scale))
      if x % scale == 0 {
        outputImage.Set(x, y, color.RGBA{0, 0, 0, 1})
      }
      if y % scale == 0 {
        outputImage.Set(x, y, color.RGBA{0, 0, 0, 1})
      }
    }
  }

  fileOutput = "image" + strconv.Itoa(scale) + ".png"

  newFile, err := os.Create(pathOut + "image" + strconv.Itoa(scale) + ".png")
  if err != nil {
    panic(err)
  }
  defer newFile.Close()

  err = png.Encode(newFile, outputImage)
  if err != nil {
    panic(err)
  }
  return err
}
