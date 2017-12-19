package main

import(
    "math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var x = []float64{0, 3.5, 25, 25, -5,-5, 15, -0.5, 19.5, 7, 1.5}
var y = []float64{0, 36, 25, 1.5, 3, 33, 11, 35, 15.5, 0, 10.5}
var waga = []float64{1,6,4,2,3,2,2,1,5,4,1}

func Newton(n int, k int) float64{
   var rez float64 = 1
   	for i := 1; i < k; i++ {
		rez= rez* float64(n-i+1.0) / float64(i)
	}
   return rez
}

func Bernstein(n int, i int, t float64) float64{
    var rez = math.Pow(float64(t),float64(i)) * math.Pow(float64(1-t),float64(n-i))
    return Newton(n,i) * rez
}

func Points() plotter.XYs {
   var n int = 10
   pts := make(plotter.XYs, 101)

   for z:=0;z<=100;z++ {
    var current_x,current_y,t float64  = 0.0, 0.0, float64(z)/100.0
    var temp float64
     
      for i:=0;i<=n;i++ {
        temp = waga[i] * x[i] * Bernstein(n,i,t)
        current_x += temp        
        }

      for i:=0;i<=n;i++ {
        temp = waga[i] * y[i] * Bernstein(n,i,t)
        current_y += temp        
        }
     
    var denominator float64 = 0.0
    for i:=0;i<=n;i++ {
      temp = waga[i] * Bernstein(n,i,t)
     denominator += temp
     }
     pts[z].X = current_x / denominator
	 pts[z].Y = current_y / denominator
   }
    
   return pts
}
func main(){

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"Bézier Curve", Points())
	if err != nil {
		panic(err)
	}

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}








