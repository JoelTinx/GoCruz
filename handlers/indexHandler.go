package handlers

import (
  "html/template"
  "image"
  "image/color"
  "image/jpeg"
  "image/png"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "path/filepath"
  "strconv"
  "strings"
)

var templates map[string]*template.Template
var fileOutput string
var tmpFileName string

func init()  {
  templates = parseTemplates()
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  templates["index.html"].Execute(w, struct{ Title string }{Title: "GoCruz - Imagen a Punto Cruz"})
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
  file, header, err := r.FormFile("imagen")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  r.ParseForm()
  tmpFileName = header.Filename
  scale, _ := strconv.Atoi(r.Form.Get("scale"))

  out, err := os.Create("public/images/original/" + tmpFileName)
 	if err != nil {
    panic(err)
 	}
  defer out.Close()

 	_, err = io.Copy(out, file)
 	if err != nil {
 		panic(err)
 	}

  if strings.ToUpper(filepath.Ext(tmpFileName)) == ".JPG" || strings.ToUpper(filepath.Ext(tmpFileName)) == ".JPEG" {
    ImageJpegToCross("public/images/original/" + tmpFileName, scale, "public/images/converted/")
    http.Redirect(w, r, "/display", http.StatusMovedPermanently)
  } else if strings.ToUpper(filepath.Ext(tmpFileName)) == ".PNG" {
    ImagePngToCross("public/images/original/" + tmpFileName, scale, "public/images/converted/")
    http.Redirect(w, r, "/display", http.StatusMovedPermanently)
  } else {
    http.Redirect(w, r, "/", http.StatusMovedPermanently)
  }
}

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
  data := struct {
    Title string
    Imagen string
  }{
    Title: "GoCruz - Imagen a Punto Cruz",
    Imagen: fileOutput,
  }

  templates["cross.html"].Execute(w, data)
}

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
  templates["gallery.html"].Execute(w, struct{ Title string }{Title: "Galeria de Imagenes - GoCruz"})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
  templates["about.html"].Execute(w, struct{ Title string }{Title: "Acerca de - GoCruz"})
}

func parseTemplates() map[string]*template.Template {
  result := make(map[string]*template.Template)
  layout, _ := template.ParseFiles("templates/layout.html")
  dir, _ := os.Open("templates/index")
  defer dir.Close()

  fis, _ := dir.Readdir(-1)

  for _, fi := range fis {
    if fi.IsDir() {
      continue
    }
    f, _ := os.Open("templates/index/" + fi.Name())
    content, _ := ioutil.ReadAll(f)
    f.Close()
    tmpl, _ := layout.Clone()
    tmpl.Parse(string(content))
    result[fi.Name()] = tmpl
  }
  return result
}

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
