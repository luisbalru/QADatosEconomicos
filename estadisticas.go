package estadisticas

import (
          "encoding/json"
          "io/ioutil"
          "log"
)

type EstadisticaParo struct {
    Year Integer `json:"year"`
    NumParados Integer `json:"parados"`
}

type Datos struct{
  Estadisticas []EstadisticaParo `json:"paro"`
}

var(
  datos_estadisticos Datos
)

const default_data_file_name = "./data/ejemplito.json"

func incorporarEstadistica(){
  // Incorpora los datos del paro. El fichero debe estar en la carpeta data
  LeeDatos(default_data_file_name)
}

func LeeDatos(nombre_archivo string){
  archivo, e = ioutil.ReadFile(nombre_archivo)
  if e != nil{
    log.Fatal("No se puede leer el archivo de datos")
  }

  if err := json.Unmarshal(archivo, &datos_estadisticos); err != nil{
    log.Fatal("Error en el JSON ", err)
  }


}

func Estadisticas() Datos{
  return datos_estadisticos
}

func HowManyYears() uint{
  return uint(len(datos_estadisticos.Estadisticas))
}
