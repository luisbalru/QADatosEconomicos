package estadisticas

import(
  "reflect"
  "testing"
  "os"
)

func TestMain(m *testing.M){
  LeeDatos("./data/estadisticas_test.json")
  os.Exit(m.Run())
}

func TestYears(t *testing.T){
  t.Log("Test Year")
  if HowManyYears() <= 0{
    t.Error("Sin estadísticas acumuladas")
  }
}

func TestNumYears(t *testing.T){
  t.Log("Test número de años")
  var x uint = uint(HowManyYears())
  if x == 2{
    t.Log("El número de años es correcto")
  }
  else{
    t.Error("El número de años es incorrecto")
  }
}

func TestTodosDatos(t *testing.T){
  t.Log("Test todos")
  statistics = Estadisticas()
  if reflect.TypeOf(statistics).String() == "Datos"{
    t.Log("El tipo de dato es correcto")
  }
  else{
    t.Error("El tipo de dato es incorrecto")
  }
}
