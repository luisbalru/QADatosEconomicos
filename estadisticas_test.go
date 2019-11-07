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
  } else{
    t.Error("El número de años es incorrecto")
  }
}

func TestTodosDatos(t *testing.T){
  t.Log("Test todos")
  statistics := Estadisticas()
  if reflect.TypeOf(statistics).String() == "estadisticas.Datos"{
    t.Log("El tipo de dato es correcto")
  } else{
    t.Error("El tipo de dato no es correcto")
  }
}

func TestMedia(t *testing.T){
  t.Log("Test Media")
  m := Media()
  if m != 2750000.0 || reflect.TypeOf(m).String() != "float64"{
    t.Error("La función media no va o tipo de dato incorrecto")
  } else{
    t.Log("Media correcta")
  }

}

func TestFatal(t *testing.T) {
    origLogFatalf := logFatalf

    // After this test, replace the original fatal function
    defer func() { logFatalf = origLogFatalf } ()

    errors := []string{}
    logFatalf = func(format string, args ...interface{}) {
        if len(args) > 0 {
            errors = append(errors, fmt.Sprintf(format, args))
        } else {
            errors = append(errors, format)
        }
    }
    if len(errors) != 1 {
        t.Errorf("excepted one error, actual %v", len(errors))
    }
}
