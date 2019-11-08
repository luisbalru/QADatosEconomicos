package estadisticas

import(
  "reflect"
  "testing"
  "os"
  "fmt"
)

func TestMain(m *testing.M){
  LeeDatos("./data/estadisticas_test.json")
  os.Exit(m.Run())
}

func TestImpresion(t *testing.T){
  t.Log("Test de fmt println")
  n, _ := fmt.Println("a")
  if n < 0 || reflect.TypeOf(n).String() != "int"{
    t.Error("Error en la impresión")
  } else{
    t.Log("Todo bien")
  }
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

func TestMediana(t *testing.T){
  t.Log("Test Mediana")
  m := Mediana()
  if m != 2000000.0 || reflect.TypeOf(m).String() != "float64"{
    t.Error("La función mediana no va o no es el tipo de dato correcto")
  } else{
    t.Log("Mediana correcta")
  }
}

func TestMin(t *testing.T){
  t.Log("Test Mínimo")
  m := Minimo()
  if m != 10.0 || reflect.TypeOf(m).String() != "float64"{
    t.Error("La función mínimo no va o no es el tipo de dato correcto")
  } else{
    t.Log("mínimo correcto")
  }
}

func TestMax(t *testing.T){
  t.Log("Test Máximo")
  m := Maximo()
  if m != 3500000.0 || reflect.TypeOf(m).String() != "float64"{
    t.Error("La función máximo no va o no es el tipo de dato correcto")
  } else{
    t.Log("Máximo correcto")
  }
}
