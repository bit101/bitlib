package main

import (
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/blgg/blgg"
	"github.com/bit101/blgg/render"
	"github.com/bit101/blgg/util"
)

func main() {
	target := render.Image

	switch target {
	case render.Image:
		render.RenderImage(800, 800, "out.png", renderFrame, 0.5)
		util.ViewImage("out.png")
		break

	case render.Gif:
		render.RenderFrames(400, 400, 120, "frames", renderFrame)
		util.MakeGIF("ffmpeg", "frames", "out.gif", 30)
		util.ViewImage("out.gif")
		break

	}
}

func renderFrame(context *blgg.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.Translate(400, 400)
	r := geom.NewRectXY(0, 0, 100, 50)
	r.Stroke(context)

	a := geom.NewAffineTransform(1, 1, 0, 0, 1, 0)
	r1 := a.ApplyToRect(r)
	r1.Stroke(context)

	b := a.Inverse()
	r2 := b.ApplyToRect(r)
	r2.Stroke(context)

}
