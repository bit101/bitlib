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

	context.SetLineWidth(0.5)

	a := geom.NewPoint(200, 200)
	b := geom.NewPoint(250, 220)
	c := geom.NewPoint(100, 230)

	circle := geom.NewCircleFromPoints(a, b, c)
	circle.Stroke(context)
	circle.ToPolygon(10).DrawVertices(context, 2)
}
