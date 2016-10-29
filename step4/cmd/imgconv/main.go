package main

import (
	"fmt"
	"os"
	"image"
	"image/jpeg"
	"image/png"	
	"path/filepath"
	"strings"
	"io"
)

func convertToJpeg(dest io.Writer, img image.Image) error {
	return jpeg.Encode(dest, img, &jpeg.Options{jpeg.DefaultQuality})	
}

func convertToPng(dest io.Writer, img image.Image) error {
	return png.Encode(dest, img)
}


func run() error {

	if len(os.Args) < 3 {
		return fmt.Errorf("引数が足りません。")
	}

	src, dest := os.Args[1], os.Args[2]

	// 拡張子取得
	ext := filepath.Ext(dest)

	file, err := os.Open(src)

	if err != nil {
		return fmt.Errorf("ファイルが開けませんでした。%s", src)
	}

	// 関数終了時にファイルを閉じる
	defer file.Close()

	// io.Readerをimage.Image形式にデコード
	img, _, err := image.Decode(file)

	if err != nil {
		return fmt.Errorf("画像ファイルとして開けませんでした。%s", src)
	}

	outFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("ファイルを書き出せませんでした %s", dest)
	}
	
	// 関数終了時にファイルを閉じる
	defer file.Close()

	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		err = convertToJpeg(outFile, img)
	case "png":
		err = convertToPng(outFile, img)
	default: fmt.Println("対応していない画像形式です。")
	}

	if err != nil {
		return fmt.Errorf("画像ファイルを書き出せませんでした。 %s", dest)
	}

	return nil
}



func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
