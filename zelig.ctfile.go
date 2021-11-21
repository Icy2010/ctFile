package ZeligCTFile

/*
	城通网盘的 API 实现  http://openapi.ctfile.com/
    作者 Icy
    Web zelig.cn

	目前实现的
	EMail 登录 ✔
    Token 登录 ✔
	获取用户信息 ✔
	获取网盘容量 ✔
	获取网盘直连流量 ✔
	创建文件夹 ✔
	获取文件夹信息 ✔
	获取文件夹列表 ✔
	修改文件夹信息 ✔
	获取文件列表 ✔
	获取指定文件列表 ✔
    文件上传 ✔
*/

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	//"bytes"
	"crypto/md5"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"github.com/tidwall/gjson"
	"io"
	"os"
	"path/filepath"
)

type TCTFileUserProfile struct {
	User_id     int    `json:"user_id"`     //用户ID
	User_name   string `json:"user_name"`   //用户名
	Nick_name   string `json:"nick_name"`   //昵称
	Group_type  int    `json:"group_type"`  //会员类型
	Group_name  string `json:"group_name"`  //会员名称
	Has_avatar  int    `json:"has_avatar"`  //是否有头像
	Reg_time    int64  `json:"reg_time"`    //注册时间
	Is_vip      bool   `json:"is_vip"`      //是否为VIP
	Is_realname bool   `json:"is_realname"` //是否已经通过实名验证
	Avatar_url  string `json:"avatar_url"`  //头像url
}

type TCTFileQuota struct {
	Max_storage         int64 `json:"max_storage"`         //公有云总容量 （bytes）
	Max_private_storage int64 `json:"max_private_storage"` //私有云总容量 （bytes）
	Space_used          int64 `json:"space_used"`          //公有云已用容量 （bytes）
	Private_space_used  int64 `json:"private_space_used"`  //私有云已用容量 （bytes）
	Total_files         int   `json:"total_files"`         //公有云总文件数
	Total_private_files int   `json:"total_private_files"` //私有云总文件数
}

type TCTFileBandwidth struct {
	Bandwith_total     int64 `json:"bandwith_total"`     //直连总流量（bytes）
	Bandwith_remaining int64 `json:"bandwith_remaining"` //直连剩余流量（bytes）
	Bandwith_used      int64 `json:"bandwith_used"`      //直连已用流量（bytes）
	Max_yun            int64 `json:"max_yun"`            //总共云处理数
}

type TCTFileFolder struct {
	Key  string `json:"key"`
	Icon string `json:"icon"`
	Name string `json:"name"`
	Date int64  `json:"date"`
}

type TCTFileFolders = []TCTFileFolder

type TCTFileFolderMeta struct {
	Key       string `json:"key"`       //文件ID
	Icon      string `json:"icon"`      //文件图标
	Name      string `json:"name"`      //文件名称
	Is_hidden int    `json:"is_hidden"` //是否在父目录隐藏
	Path      string `json:"path"`      //文件夹位置
}

type TCTFileFolderFiles = []TCTFileFolderFile

type TCTFileFolderFile struct {
	Key    string `json:"key"`    //文件ID
	Icon   string `json:"icon"`   //文件图标
	Imgsrc string `json:"imgsrc"` //视频/图片的缩略图URL
	Name   string `json:"name"`   //文件名称
	Size   int64  `json:"size"`   //文件大小（bytes）
	Date   int64  `json:"date"`   //文件上传时间
	Status int    `json:"status"` //1: complete, 2: incomplete。如果为2，属于incomplete的状态，那请添加个未完成的icon（未完成的文件不支持下载、打包、解压）
}

type TCTFileFolderFileRecycle struct {
	Key      string `json:"key"`      //文件ID
	Icon     string `json:"icon"`     //文件图标
	Imgsrc   string `json:"imgsrc"`   //视频/图片的缩略图URL
	Name     string `json:"name"`     //文件名称
	Size     int64  `json:"size"`     //文件大小（bytes）
	Date     int64  `json:"date"`     //文件上传时间
	Del_time int64  `json:"del_time"` //文件被回收站清空时间
}

type TCTFileFolderFileRecycles = []TCTFileFolderFileRecycle
type TCTFileFolderFileDownload struct {
	Key  string `json:"key"`  //文件ID
	Icon string `json:"icon"` //文件图标
	Name string `json:"name"` //文件名称
	Size int64  `json:"size"` //文件大小（bytes）
	Path string `json:"path"` //为对应当前文件目录的相对位置。需要自动创建文件夹，并把文件下载至对应的文件夹内
}
type TCTFileFolderFileDownloads = []TCTFileFolderFileDownload

type TCTFileFolderFileShare struct {
	Key        string `json:"key"`        //文件ID
	Icon       string `json:"icon"`       //文件图标
	Name       string `json:"name"`       //文件名称
	Size       int64  `json:"size"`       //文件大小（bytes）
	Date       int64  `json:"date"`       //文件上传时间
	Weblink    string `json:"weblink"`    //第三方网页分享地址
	Xtlink     string `json:"xtlink"`     //小通链接地址
	Directlink string `json:"directlink"` //直连分享地址
}

type TCTFileFolderFileShares = []TCTFileFolderFileShare

type TCTFileIncome struct {
	AccountMode     int     `json:"account_mode"`      //0,1,2,5
	AccountModeInfo string  `json:"account_mode_info"` //	分成模式（高收益模式已开启，临时低收益模式已开启，赚钱收益功能已关闭，低收益模式已开启）
	AccountType     string  `json:"account_type"`      //	账号类型（普通账户，问答账户）
	UserLevel       int     `json:"user_level"`        //会员等级
	GroupType       int     `json:"qroup_type"`        //会员类型
	TodayIncome     float64 `json:"today_income"`      //今日收入
	TodayClicked    int     `json:"today_clicked"`     //今日点击数
	AspireIncome    float64 `json:"aspire_income"`     //尊享卡翻倍收入
	UnpaidIncome    float64 `json:"unpaid_income"`     //未兑换佣金
	PaidIncome      float64 `json:"paid_income"`       //已兑换佣金
}

type TCTFileFileMeta struct {
	Key  string `json:"key"`  //文件ID
	Icon string `json:"icon"` //文件图标
	Name string `json:"name"` //文件名称
	Size int64  `json:"size"` //文件大小（bytes）
	Path string `json:"path"` //为对应当前文件目录的相对位置。需要自动创建文件夹，并把文件下载至对应的文件夹内
}

type TCTFile struct {
	token     string
	Profile   TCTFileUserProfile
	Quota     TCTFileQuota
	Bandwidth TCTFileBandwidth
}

func (this *TCTFile) Token() string {
	return this.token
}

func (this *TCTFile) LoginFromToken(token string) error {
	this.token = token

	err := this.getProfile()
	if err != nil {
		return err
	}

	err = this.getQuota()
	if err != nil {
		return err
	}

	err = this.getBandwidth()
	if err != nil {
		return err
	}

	return err
}

func cTFileHttp() *HttpRequest.Request {
	return HttpRequest.NewRequest().SetHeaders(map[string]string{
		"myapp-id":     "b81c58f3f33548d5f063abab00b63262",
		"HOST":         "rest.ctfile.com",
		"content-type": "application/json",
	})
}

func public_private(IsPublic bool) string {
	if IsPublic {
		return "public"
	} else {
		return "private"
	}
}

func cTFilehttpGet(url string, OnData func(JO gjson.Result) error) error {
	res, err := cTFileHttp().Get(url)
	if err == nil {
		var body []byte = nil
		body, err = res.Body()
		if err == nil {
			JO := gjson.Parse(string(body))
			if JO.Get("code").Int() == 200 {
				if OnData != nil {
					return OnData(JO)
				} else {
					return nil
				}
			} else {
				return fmt.Errorf(JO.Get("message").String())
			}

		}
	}

	return err
}

func cTFilehttpPost(url string, Body map[string]interface{}, OnData func(JO gjson.Result) error) error {
	res, err := cTFileHttp().Post(url, Body)
	if err == nil {
		var body []byte = nil
		body, err = res.Body()
		if err == nil {
			JO := gjson.Parse(string(body))
			if JO.Get("code").Int() == 200 {
				if OnData != nil {
					return OnData(JO)
				} else {
					return nil
				}
			} else {
				return fmt.Errorf(JO.Get("message").String())
			}

		}
	}

	return err
}

func (this *TCTFile) Login(EMail, Password string) error {
	url := fmt.Sprintf(`https://rest.ctfile.com/v1/user/auth/login?email=%s&password=%s`, EMail, Password)
	return cTFilehttpGet(url, func(JO gjson.Result) error {
		this.token = JO.Get("token").String()
		err := this.getProfile()
		if err != nil {
			return err
		}

		err = this.getQuota()
		if err != nil {
			return err
		}

		err = this.getBandwidth()
		if err != nil {
			return err
		}

		return nil
	})

}

func (this *TCTFile) PublicCloud() *TCTFilePublic {
	return newCTFilePublic(this)
}

func (this *TCTFile) PrivateCloud() *TCTFilePrivate {
	return newCTFilePrivate(this)
}

func (this *TCTFile) getProfile() error {
	url := fmt.Sprintf(`https://rest.ctfile.com/v1/user/info/profile?session=%s`, this.token)
	return cTFilehttpGet(url, func(JO gjson.Result) error {
		this.Profile = TCTFileUserProfile{
			User_id:     int(JO.Get("userid").Int()),
			User_name:   JO.Get("username").String(),
			Nick_name:   JO.Get("nick_name").String(),
			Group_type:  int(JO.Get("group_type").Int()),
			Group_name:  JO.Get("group_name").String(),
			Has_avatar:  int(JO.Get("has_avatar").Int()),
			Reg_time:    JO.Get("reg_time").Int(),
			Is_vip:      JO.Get("is_vip").Bool(),
			Is_realname: JO.Get("userid").Bool(),
			Avatar_url:  JO.Get("avatar_url").String(),
		}
		return nil
	})
}

func (this *TCTFile) getQuota() error {
	url := fmt.Sprintf(`https://rest.ctfile.com/v1/user/info/quota?session=%s`, this.token)
	return cTFilehttpGet(url, func(JO gjson.Result) error {
		this.Quota = TCTFileQuota{
			Max_storage:         JO.Get("max_storage").Int(),
			Max_private_storage: JO.Get("max_private_storage").Int(),
			Space_used:          JO.Get("space_used").Int(),
			Private_space_used:  JO.Get("private_space_used").Int(),
			Total_files:         int(JO.Get("total_files").Int()),
			Total_private_files: int(JO.Get("total_private_files").Int()),
		}
		return nil
	})
}

func (this *TCTFile) getBandwidth() error {
	url := fmt.Sprintf(`https://rest.ctfile.com/v1/user/info/bandwidth?session=%s`, this.token)
	return cTFilehttpGet(url, func(JO gjson.Result) error {
		this.Bandwidth = TCTFileBandwidth{
			Bandwith_total:     JO.Get("bandwith_total").Int(),
			Bandwith_remaining: JO.Get("bandwith_remaining").Int(),
			Bandwith_used:      JO.Get("bandwith_used").Int(),
			Max_yun:            JO.Get("max_yun").Int(),
		}
		return nil
	})
}

func (this *TCTFile) folderCreate(IsPublic bool, Folder_id, Name, Description string, Is_Hidden int) (map[string]string, error) {
	result := make(map[string]string)

	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/folder/create`, public_private(IsPublic)), map[string]interface{}{
		"session":     this.token,
		"name":        Name,
		"description": Description,
		"is_hidden":   Is_Hidden,
		"folder_id":   Folder_id,
	}, func(JO gjson.Result) error {
		result["folder_id"] = JO.Get("folder_id").String()
		result["folder_path"] = JO.Get("folder_path").String()
		return nil
	})

	return result, err
}

func (this *TCTFile) folderList(IsPublic bool, Folder_id string) (TCTFileFolders, error) {
	var result TCTFileFolders
	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/folder/list`, public_private(IsPublic)), map[string]interface{}{
		"folder_id": Folder_id,
		"session":   this.token,
	}, func(JO gjson.Result) error {
		JA := JO.Get("results").Array()
		if len(JA) > 0 {
			for i := 0; i < len(JA); i++ {
				result = append(result, TCTFileFolder{
					Key:  JA[i].Get("key").String(),
					Icon: JA[i].Get("icon").String(),
					Name: JA[i].Get("name").String(),
					Date: JA[i].Get("date").Int(),
				})
			}
		}
		return nil
	})
	return result, err

}

func (this *TCTFile) folderMeta(IsPublic bool, Folder_id string) (TCTFileFolderMeta, error) {
	var result TCTFileFolderMeta

	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/folder/meta`, public_private(IsPublic)), map[string]interface{}{
		"folder_id": Folder_id,
		"session":   this.token,
	},
		func(JO gjson.Result) error {
			result = TCTFileFolderMeta{
				Key:       JO.Get("key").String(),
				Icon:      JO.Get("icon").String(),
				Name:      JO.Get("name").String(),
				Is_hidden: int(JO.Get("is_hidden").Int()),
				Path:      JO.Get("path").String(),
			}
			return nil
		})

	return result, err
}

func (this *TCTFile) folderModifyMeta(IsPublic bool, Folder_id, Name, Description string, Is_Hidden int) (map[string]string, error) {
	data := map[string]interface{}{
		"folder_id": Folder_id,
		"session":   this.token,
		"is_hidden": Is_Hidden,
	}

	result := make(map[string]string)

	if Description != "" {
		data["description"] = Description
	}

	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/folder/modify_meta`, public_private(IsPublic)), data,
		func(JO gjson.Result) error {
			result["folder_id"] = JO.Get("folder_id").String()
			result["folder_path"] = JO.Get("folder_path").String()
			return nil
		})

	return result, err
}

func (this *TCTFile) fileList(IsPublic bool, Folder_id string, Start, Reload int, Orderby, Filter, Keyword string) (TCTFileFolderFiles, error) {
	var result TCTFileFolderFiles
	data := map[string]interface{}{
		"folder_id": Folder_id,
		"session":   this.token,
	}

	if Start > 0 {
		data["start"] = Start
	}

	if Reload > 0 {
		data["reload"] = Reload
	}

	if Orderby != "" {
		data["orderby"] = Orderby
	}

	if Filter != "" {
		data["filter"] = Filter
	}

	if Keyword != "" {
		data["keyword"] = Keyword
	}

	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/list`, public_private(IsPublic)), data,
		func(JO gjson.Result) error {
			JA := JO.Get("results").Array()
			if len(JA) > 0 {
				for i := 0; i < len(JA); i++ {
					result = append(result, TCTFileFolderFile{
						Key:    JA[i].Get("key").String(),
						Icon:   JA[i].Get("icon").String(),
						Name:   JA[i].Get("name").String(),
						Date:   JA[i].Get("date").Int(),
						Imgsrc: JA[i].Get("imgsrc").String(),
						Size:   JA[i].Get("size").Int(),
						Status: int(JA[i].Get("status").Int()),
					})
				}
			}

			return nil
		})

	return result, err
}

func (this *TCTFile) fileMeta(IsPublic bool, file_id string) (TCTFileFileMeta, error) {
	Data := TCTFileFileMeta{}
	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/meta`, public_private(IsPublic)),
		map[string]interface{}{"session": this.token,
			"file_id": file_id},
		func(JO gjson.Result) error {
			if JO.Get("code").Int() == 200 {
				Data = TCTFileFileMeta{
					Key:  JO.Get("key").String(),
					Icon: JO.Get("icon").String(),
					Name: JO.Get("name").String(),
					Size: JO.Get("size").Int(),
					Path: JO.Get("path").String(),
				}

				return nil
			}

			return fmt.Errorf(JO.Get("message").String())
		})
	return Data, err
}

func (this *TCTFile) fileIdsList(IsPublic bool, Ids []string) (TCTFileFolderFiles, error) {
	var result TCTFileFolderFiles

	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/ids_list`, public_private(IsPublic)), map[string]interface{}{
		"session": this.token,
		"ids":     Ids,
	}, func(JO gjson.Result) error {
		JA := JO.Get("results").Array()
		if len(JA) > 0 {
			for i := 0; i < len(JA); i++ {
				result = append(result, TCTFileFolderFile{
					Key:    JA[i].Get("key").String(),
					Icon:   JA[i].Get("icon").String(),
					Name:   JA[i].Get("name").String(),
					Date:   JA[i].Get("date").Int(),
					Imgsrc: JA[i].Get("imgsrc").String(),
					Size:   JA[i].Get("size").Int(),
					Status: int(JA[i].Get("status").Int()),
				})
			}
		}

		return nil
	})

	return result, err
}

func (this *TCTFile) fileRecycle(IsPublic bool, Start, Reload int) (TCTFileFolderFileRecycles, error) {
	var result TCTFileFolderFileRecycles
	data := map[string]interface{}{
		"session": this.token,
	}

	if Start > 0 {
		data["start"] = Start
	}

	if Reload > 0 {
		data["reload"] = Reload
	}

	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/recycle`, public_private(IsPublic)), data,
		func(JO gjson.Result) error {
			JA := JO.Get("results").Array()
			if len(JA) > 0 {
				for i := 0; i < len(JA); i++ {
					result = append(result, TCTFileFolderFileRecycle{
						Key:      JA[i].Get("key").String(),
						Icon:     JA[i].Get("icon").String(),
						Name:     JA[i].Get("name").String(),
						Date:     JA[i].Get("date").Int(),
						Imgsrc:   JA[i].Get("imgsrc").String(),
						Size:     JA[i].Get("size").Int(),
						Del_time: JA[i].Get("del_time").Int(),
					})
				}
			}

			return nil
		})

	return result, err
}

func (this *TCTFile) fileRecycle_empty(IsPublic bool, Ids []string) error {
	return cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/ids_list`, public_private(IsPublic)), map[string]interface{}{
		"session": this.token,
		"ids":     Ids,
	}, nil)
}

func (this *TCTFile) fileRecycle_empty_all(IsPublic bool) error {
	return cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/recycle_empty_all`, public_private(IsPublic)), map[string]interface{}{
		"session": this.token,
	}, nil)
}

func (this *TCTFile) fileDownload(IsPublic bool, Ids []string) (TCTFileFolderFileDownloads, error) {
	var result TCTFileFolderFileDownloads

	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/download`, public_private(IsPublic)), map[string]interface{}{
		"session": this.token,
		"ids":     Ids,
	}, func(JO gjson.Result) error {
		JA := JO.Get("results").Array()
		if len(JA) > 0 {
			for i := 0; i < len(JA); i++ {
				result = append(result, TCTFileFolderFileDownload{
					Key:  JA[i].Get("key").String(),
					Icon: JA[i].Get("icon").String(),
					Name: JA[i].Get("name").String(),
					Size: JA[i].Get("size").Int(),
					Path: JO.Get("path").String(),
				})
			}
		}

		return nil
	})

	return result, err
}

func (this *TCTFile) fileFetch_urlb(IsPublic bool, File_id string) (string, error) {
	url := ""
	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/fetch_url`, public_private(IsPublic)), map[string]interface{}{
		"session": this.token,
		"file_id": File_id,
	}, func(JO gjson.Result) error {
		url = JO.Get("download_url").String()

		return nil
	})

	return url, err
}

func (this *TCTFile) fileShare(IsPublic bool, Ids []string) (TCTFileFolderFileShares, error) {
	var result TCTFileFolderFileShares

	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/share`, public_private(IsPublic)), map[string]interface{}{
		"session": this.token,
		"ids":     Ids,
	}, func(JO gjson.Result) error {
		JA := JO.Get("results").Array()
		if len(JA) > 0 {
			for i := 0; i < len(JA); i++ {
				result = append(result, TCTFileFolderFileShare{
					Key:        JA[i].Get("key").String(),
					Icon:       JA[i].Get("icon").String(),
					Name:       JA[i].Get("name").String(),
					Size:       JA[i].Get("size").Int(),
					Date:       JA[i].Get("date").Int(),
					Weblink:    JA[i].Get("weblink").String(),
					Xtlink:     JA[i].Get("xtlink").String(),
					Directlink: JA[i].Get("directlink").String(),
				})
			}
		}

		return nil
	})

	return result, err
}

func (this *TCTFile) fileMove(IsPublic bool, Folder_id string, Ids []string) error {
	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/move`, public_private(IsPublic)), map[string]interface{}{
		"session":   this.token,
		"ids":       Ids,
		"folder_id": Folder_id,
	}, nil)

	return err
}

func (this *TCTFile) fileDelete(IsPublic bool, Ids []string) error {
	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/delete`, public_private(IsPublic)), map[string]interface{}{
		"session": this.token,
		"ids":     Ids,
	}, nil)

	return err
}

func (this *TCTFile) fileSave(IsPublic bool, Ids []string) error {
	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/save`, public_private(IsPublic)), map[string]interface{}{
		"session": this.token,
		"ids":     Ids,
	}, nil)

	return err
}

func getFileSize(Filename string) int64 {
	fi, err := os.Stat(Filename)
	if err != nil {
		return 0
	}

	return fi.Size()
}

func getfileMD5(Filename string) string {
	file, inerr := os.Open(Filename)
	if inerr == nil {
		md5 := md5.New()
		if _, err := io.Copy(md5, file); err == nil {
			return hex.EncodeToString(md5.Sum(nil))
		}
		return ""
	}

	return ""
}

func extractFilename(path string) string {
	_, fileName := filepath.Split(path)
	return fileName
}

func file_upload(url, file_name string, size int64) (string, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	_ = writer.WriteField("name", extractFilename(file_name))
	_ = writer.WriteField("filesize", fmt.Sprintf(`%d`, size))

	file, errFile3 := os.Open(file_name)
	defer file.Close()

	part3, errFile3 := writer.CreateFormFile("file", filepath.Base(file_name))
	_, errFile3 = io.Copy(part3, file)
	if errFile3 != nil {
		return ``, errFile3
	}
	err := writer.Close()
	if err != nil {
		return ``, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(`POST`, url, payload)
	if err != nil {

		return ``, err
	}
	req.Header.Add("host", "upload.ctfile.com")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return ``, err
	}
	defer res.Body.Close()

	body, ep := ioutil.ReadAll(res.Body)
	if ep != nil {
		return ``, ep
	}

	if res.StatusCode != 200 {
		JO := gjson.ParseBytes(body)
		if JO.Exists() {
			if JO.Get("code").Int() != 200 {
				return ``, fmt.Errorf(JO.Get("message").String())
			}
		} ///  目前上传成功返回的东西很奇怪....
	}

	return string(body), ep
}

func (this *TCTFile) fileUpload(IsPublic bool, Folder_id, Filename string) (string, error) {
	Size := getFileSize(Filename)
	Data := map[string]interface{}{
		"session":   this.token,
		"folder_id": Folder_id,
		"checksum":  getfileMD5(Filename) + fmt.Sprintf(`-%d`, Size),
		"size":      Size,
		"name":      extractFilename(Filename),
	}

	upload_url := ""

	err := cTFilehttpPost(fmt.Sprintf(`https://rest.ctfile.com/v1/%s/file/upload`, public_private(IsPublic)), Data,
		func(JO gjson.Result) error {
			if JO.Get("exists").Int() > 0 {
				return fmt.Errorf("文件已经存在了,无需上传")
			} else {
				upload_url = JO.Get("upload_url").String()
			}
			return nil
		})

	Body := ""
	if err == nil {
		Body, err = file_upload(upload_url, Filename, Size)
	}

	return Body, err
}

func (this *TCTFile) Income() (TCTFileIncome, error) {
	Data := TCTFileIncome{}
	err := cTFilehttpPost(`https://rest.ctfile.com/v1/union/info/income`,
		map[string]interface{}{"session": this.token},
		func(JO gjson.Result) error {
			if JO.Get("code").Int() == 200 {
				Data.AccountMode = int(JO.Get("account_mode").Int())
				Data.AccountType = JO.Get("account_type").String()
				Data.AccountModeInfo = JO.Get("account_mode_info").String()
				Data.UserLevel = int(JO.Get("user_level").Int())
				Data.GroupType = int(JO.Get("group_type").Int())
				Data.TodayIncome = JO.Get("today_income").Float()
				Data.TodayClicked = int(JO.Get("today_clicked").Int())
				Data.AspireIncome = JO.Get("aspire_income").Float()
				Data.UnpaidIncome = JO.Get("unpaid_income").Float()
				Data.PaidIncome = JO.Get("paid_income").Float()
			} else {
				return fmt.Errorf(JO.Get("message").String())
			}

			return nil
		})

	return Data, err
}
