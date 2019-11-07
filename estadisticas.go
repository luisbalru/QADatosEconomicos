package estadisticas



import (
          "encoding/json"
          "io/ioutil"
          "log"
          "github.com/montanaflynn/stats"
)

type EstadisticaParo struct {
    Year int `json:"year"`
    NumParados int `json:"parados"`
}

type Datos struct{
  Estadisticas []EstadisticaParo `json:"paro"`
}

var(
  datos_estadisticos Datos
)

const default_data_file_name = "./data/ejemplito.json"


func LeeDatos(nombre_archivo string){
  archivo, e := ioutil.ReadFile(nombre_archivo)
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

func Media() float64{
  d := stats.LoadRawData([]int{2000000,3500000})
  m,_ := stats.Mean(d)
  return m
}


func HowManyYears() int{
  return int(len(datos_estadisticos.Estadisticas))
}
