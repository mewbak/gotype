package testimages

import (
	"testing"

	"github.com/npillmayer/gotype/gtbackend/gfx"
	_ "github.com/npillmayer/gotype/gtbackend/gfx/png"
	"github.com/npillmayer/gotype/gtcore/config"
	"github.com/npillmayer/gotype/gtcore/config/tracing"
	"github.com/npillmayer/gotype/gtcore/path"
	"github.com/sirupsen/logrus"
)

var T tracing.Trace = tracing.GraphicsTracer

func savePNG(pic *gfx.Picture) {
	pic.Shipout()
	/*
		f, err := os.Create(pic.Name + ".png")
		if err != nil {
			log.Fatal(err)
		}
		img := pic.AsImage()
		if err := ospng.Encode(f, img); err != nil {
			f.Close()
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	*/
}

func TestEnvironment(t *testing.T) {
	config.Initialize()
	T.SetLevel(logrus.DebugLevel)
}

func TestSimplePath1(t *testing.T) {
	pic := gfx.NewPicture("simple1", 100, 100, "PNG")
	savePNG(pic)
}

func TestPath1(t *testing.T) {
	pic := gfx.NewPicture("path1", 100, 100, "PNG")
	p, controls := path.Nullpath().Knot(path.P(0, 0)).Curve().Knot(path.P(50, 50)).Curve().
		Knot(path.P(100, 65)).End()
	controls = path.FindHobbyControls(p, controls)
	pic.Draw(gfx.NewDrawablePath(p, controls))
	savePNG(pic)
}

func TestPath2(t *testing.T) {
	pic := gfx.NewPicture("path2", 100, 100, "PNG")
	p, controls := path.Nullpath().Knot(path.P(10, 50)).Curve().Knot(path.P(50, 90)).Curve().
		Knot(path.P(90, 50)).End()
	controls = path.FindHobbyControls(p, controls)
	pic.Draw(gfx.NewDrawablePath(p, controls))
	savePNG(pic)
}

func TestPath3(t *testing.T) {
	pic := gfx.NewPicture("path3", 100, 100, "PNG")
	p, controls := path.Nullpath().Knot(path.P(10, 50)).Curve().Knot(path.P(50, 90)).Curve().
		Knot(path.P(90, 50)).Curve().Knot(path.P(50, 10)).Curve().Cycle()
	controls = path.FindHobbyControls(p, controls)
	pic.Draw(gfx.NewDrawablePath(p, controls))
	savePNG(pic)
}
