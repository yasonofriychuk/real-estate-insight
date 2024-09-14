package main

import (
	"encoding/json"
	"fmt"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/scanner"
	"os"
	"runtime"
	"time"

	"github.com/paulmach/go.geojson"
)

func main() {
	start := time.Now()
	file, err := os.Open("osmfiles/RU.osm.pbf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc, err := scanner.New(file, runtime.GOMAXPROCS(-1))
	if err != nil {
		panic(err)
	}

	featureCollection := geojson.NewFeatureCollection()
	for i := 0; i < 150; i++ {
		feature, err := sc.Next()
		if err != nil {
			break
		}

		featureCollection.AddFeature(feature)
	}

	// После завершения итерации сохраняем GeoJSON в файл
	outputFile, err := os.Create("shops_output.geojson")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	geojsonBytes, err := json.MarshalIndent(featureCollection, "", "  ")
	if err != nil {
		panic(err)
	}

	// Записываем GeoJSON данные в файл
	_, err = outputFile.Write(geojsonBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println("Конвертация завершена:", time.Since(start).String())
}
