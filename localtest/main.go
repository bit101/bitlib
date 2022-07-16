package main

import (
	"fmt"

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

	p := geom.NewPolygonFromCoords(100, 100, 200, 120, 190, 200, 80, 300)
	p.Draw(context)
	p.DrawVertices(context, 2)
	p.Centroid().Fill(context, 2)

	q := geom.NewPolygonFromCoords(100, 100, 200, 120, 190, 200, 80, 300)
	fmt.Printf("p.Equals(q) = %+v\n", p.Equals(q))
}
