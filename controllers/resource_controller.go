package controllers

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
)

type ResourceController struct {
	BaseController
}

func (this *ResourceController) GetThumbnailImage() {
	gid := this.GetString(":gid")
	size := this.GetString(":size")
	sizeParams := strings.Split(strings.ToLower(size), "x")
	var width, height int
	if len(sizeParams) == 2 {
		width, _ = strconv.Atoi(sizeParams[0])
		height, _ = strconv.Atoi(sizeParams[1])
	}

	if width <= 0 {
		width = 120
	}

	if height <= 0 {
		height = 80
	}

	thumbImageUrl := ".\\static\\img\\product\\thumbnails\\" + gid + "-" + size + ".jpg"

	//首先检查本地有没有该文件
	if _, err := os.Stat(thumbImageUrl); err == nil {
		this.ResponseFile("image/jpeg", thumbImageUrl)
		this.StopRun()
	} else {
		imageUrl := ".\\static\\img\\product\\" + gid + ".jpg"
		//首先检查本地有没有原图
		if _, err := os.Stat(imageUrl); err == nil {
			//生成缩略图
			if err := createThumbPic(imageUrl, thumbImageUrl, uint(width), uint(height)); err == nil {
				this.ResponseFile("image/jpeg", thumbImageUrl)
				this.StopRun()
			}
		}
	}

	this.Abort("404")
}

func (this *ResourceController) ResponseFile(fileType, filePath string) {
	var buf []byte = make([]byte, 255)
	var fileBody []byte = make([]byte, 0)
	if file, err := os.Open(filePath); err == nil {
		for {
			if length, _ := file.Read(buf); length == 0 {
				break
			}
			for _, v := range buf {
				fileBody = append(fileBody, v)
			}
		}
		file.Close()
	} else {
		log.Println(err.Error())
	}

	this.Ctx.Output.Header("Accept-Ranges", "bytes")
	this.Ctx.Output.Header("Content-Type", fileType)
	this.Ctx.ResponseWriter.Write(fileBody)
}

func createThumbPic(source, to string, width, height uint) error {
	source_file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer source_file.Close()

	thumb_pic, err := os.Create(to)
	if err != nil {
		return err
	}
	defer thumb_pic.Close()

	// decode jpeg into image.Image
	img, err := jpeg.Decode(source_file)
	if err != nil {
		return err
	}

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Thumbnail(width, height, img, resize.Lanczos3)

	// write new image to file
	if err := jpeg.Encode(thumb_pic, m, nil); err != nil {
		return err
	}
	return nil
}
