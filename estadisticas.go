package estadisticas

import (
          "encoding/json"
          "io/ioutil"
)

type EstadisticaParo struct {
    Year Integer
    NumParados Integer
}

type Datos struct{
  Estadisticas []EstadisticaParo
}

const default_data_file_name = "./data/ejemplito.json"

func incorporarEstadistica(){
  // Incorpora los datos del paro. El fichero debe estar en la carpeta data
  LeeDatos(default_data_file_name)
}
