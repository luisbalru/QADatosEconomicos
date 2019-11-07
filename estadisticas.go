package estadisticas



import (
          "encoding/json"
          "io/ioutil"
          "github.com/montanaflynn/stats"
          "fmt"
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
    fmt.Println("No se puede leer el archivo de datos")
  }

  if err := json.Unmarshal(archivo, &datos_estadisticos); err != nil{
    fmt.Println("Error en el JSON ", err)
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

func Mediana() float64{
  d := stats.LoadRawData([]int{2000000,10,3500000})
  mediana, _ := stats.Median(d)
  return mediana
}


func HowManyYears() int{
  return int(len(datos_estadisticos.Estadisticas))
}
