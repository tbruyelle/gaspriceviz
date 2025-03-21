package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/vicanso/go-charts/v2"
)

func genChart() image.Image {
	values := [][]float64{
		{
			120,
			132,
			101,
			// 134,
			charts.GetNullValue(),
			90,
			230,
			210,
		},
		{
			220,
			182,
			191,
			234,
			290,
			330,
			310,
		},
		{
			150,
			232,
			201,
			154,
			190,
			330,
			410,
		},
		{
			320,
			332,
			301,
			334,
			390,
			330,
			320,
		},
		{
			820,
			932,
			901,
			934,
			1290,
			1330,
			1320,
		},
	}
	p, err := charts.LineRender(
		values,
		charts.WidthOptionFunc(1000),
		charts.HeightOptionFunc(800),
		charts.TitleTextOptionFunc("Line"),
		charts.XAxisDataOptionFunc([]string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		}),
		charts.LegendLabelsOptionFunc([]string{
			"Email",
			"Union Ads",
			"Video Ads",
			"Direct",
			"Search Engine",
		}, "50"),
		func(opt *charts.ChartOption) {
			opt.Legend.Padding = charts.Box{
				Top:    5,
				Bottom: 10,
			}
			opt.YAxisOptions = []charts.YAxisOption{
				{
					SplitLineShow: charts.FalseFlag(),
				},
			}
			opt.SymbolShow = charts.FalseFlag()
			opt.LineStrokeWidth = 1
			opt.ValueFormatter = func(f float64) string {
				return fmt.Sprintf("%.0f", f)
			}
		},
	)
	if err != nil {
		panic(err)
	}

	buf, err := p.Bytes()
	if err != nil {
		panic(err)
	}
	writeFile(buf)
	img, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		panic(fmt.Errorf("fetchImage: image decode failed: %v", err))
	}
	return img
}

func writeFile(buf []byte) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0o700)
	if err != nil {
		return err
	}

	file := filepath.Join(tmpPath, "line-chart.png")
	err = os.WriteFile(file, buf, 0o600)
	if err != nil {
		return err
	}
	return nil
}
