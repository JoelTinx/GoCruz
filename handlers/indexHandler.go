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
  templates["index.tmpl"].Execute(w, struct{ Title string }{Title: "GoCruz - Imagen a Punto Cruz"})
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

  templates["cross.tmpl"].Execute(w, data)
}

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
  templates["gallery.tmpl"].Execute(w, struct{ Title string }{Title: "Galeria de Imagenes - GoCruz"})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
  templates["about.tmpl"].Execute(w, struct{ Title string }{Title: "Acerca de - GoCruz"})
}

func parseTemplates() map[string]*template.Template {
  result := make(map[string]*template.Template)
  layout, _ := template.ParseFiles("templates/layout.tmpl")
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
