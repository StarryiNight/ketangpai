package controller

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/go-emix/utils"
	"go.uber.org/zap"
	"io"
	"ketangpai/dao/redis"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func UploadHandler(c *gin.Context) {
	total := 0
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		ResponseErrorWithMsg(c,CodeRequestFileError,CodeRequestFileError.Msg())
	}
	path:=c.PostForm("path")+"/"
	zap.L().Info("path",zap.String("path",path+header.Filename))
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path,os.ModePerm)
	}
	saveFile, err := os.OpenFile(path+header.Filename, os.O_RDWR, os.ModePerm)
	if err != nil {
		saveFile, _ = os.Create(path+header.Filename)
	}

	buf := make([]byte, 2048)
	var str string
	if result, _ := redis.GetClient().Get(header.Filename).Int64(); result != 0 {
		total = int(result)
		file.Seek(int64(total), io.SeekStart)

		str= "该文件之前已上传"+ strconv.Itoa(total)
	}
	for true {
		read, err := file.Read(buf)
		if err == io.EOF {
			break
		}

		saveFile.WriteAt(buf, int64(total))

		total += read
		if err = redis.GetClient().Set(header.Filename, total, 0).Err(); err != nil {
			ResponseErrorWithMsg(c,CodeRedisSaveFiled,CodeRedisSaveFiled.Msg())
		}
	}
	saveFile.Close()
	str=str+"     ...上传完毕,总共"+ strconv.Itoa(total)
	ResponseSuccess(c,str)
}

func DownloadHandler(c *gin.Context) {

	myUrl := c.PostForm("Url")
	dir := c.PostForm("dir")
	fileName := c.PostForm("fileName")
	//文件夹是否存在
	if !utils.FileIsExist(dir) {
		//不存在则创建文件夹
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		return
	}
	//下载文件路径
	dfn := dir + "/" + fileName

	var file *os.File
	var size int64

	if utils.FileIsExist(dfn) {
		//如果文件存在未下载完  则使file指向该文件
		fi, err := os.OpenFile(dfn, os.O_RDWR, os.ModePerm)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		//指针指向末尾
		stat,_:=fi.Stat()
		size=stat.Size()
		sk,err:=fi.Seek(size, 0)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			_=fi.Close()
			return
		}

		if sk!= size {
			ResponseErrorWithMsg(c,CodeFileSeekFailed,CodeFileSeekFailed.Msg())
		}

		file=fi
	} else{
		//没有则以路径dfn创建文件 并使file指向新创建的文件
		create,err:=os.Create(dfn)

		if err != nil {
			ResponseErrorWithMsg(c,CodeServerBusy,err.Error())
			return
		}

		file=create
	}
	client:=&http.Client{}
	client.Timeout=time.Hour
	request:=http.Request{}
	request.Method=http.MethodGet

	if size!= 0 {
		//指向上次位置
		header:=http.Header{}
		header.Set("Range","bytes="+strconv.FormatInt(size,10)+"-")
		request.Header=header
	}
	parse,err:= url.Parse(myUrl)
	request.URL=parse
	get,err:=client.Do(&request)
	if err != nil {
		ResponseErrorWithMsg(c,CodeServerBusy,err.Error())
		return
	}

	defer func() {
		//关闭打开的流
		err:=get.Body.Close()
		if err != nil {
			ResponseErrorWithMsg(c,CodeServerBusy,err.Error())
			return
		}
		err=file.Close()
		if err != nil {
			ResponseErrorWithMsg(c,CodeServerBusy,err.Error())
			return
		}
	}()

	if get.ContentLength== 0 {
		//下载完成
		ResponseSuccess(c,"下载完成")
		return
	}
	body:=get.Body
	writer:=bufio.NewWriter(file)
	buf:=make([]byte,10*1024*1024)
	for true {
		var read int
		read,err=body.Read(buf)
		if err != nil {
			if err != io.EOF {
				ResponseErrorWithMsg(c,CodeServerBusy,err.Error())
				return
			} else {
				err=nil
			}
			break
		}
		_,err=writer.Write(buf[:read])
		if err != nil {
			ResponseErrorWithMsg(c,CodeServerBusy,err.Error())
			break
		}

	}
	if err != nil {
		return
	}
	err=writer.Flush()
	if err != nil {
		ResponseErrorWithMsg(c,CodeServerBusy,err.Error())
		return
	}
	ResponseSuccess(c,"下载完成")
}

