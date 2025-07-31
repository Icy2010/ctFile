![image](https://github.com/Icy2010/ZeligCTFile/blob/main/zelig.ctfile.jpg)
# ctFile
#### 城通网盘的API Go语言实现 
#### 作者 Icy 
#### Web http://zelig.cn

## 目前实现的
1. EMail 登录 ✔
2. Token 登录  ✔
3. 获取用户信息 ✔
4. 获取网盘容量 ✔
5. 获取网盘直连流量 ✔
6. 创建文件夹 ✔
7. 获取文件夹信息 ✔
8. 获取文件夹列表 ✔
9. 修改文件夹信息 ✔
10. 获取文件列表 ✔
11. 获取指定文件列表 ✔
12. 文件上传完成 ✔
13. 获取文件信息 ✔ 
14. 获取收入信息 ✔

```golang
package main

import (
	"fmt"
	z "github.com/Icy2010/ctFile"
)

func FileList(ctfile z.TCTFile) {
	files, err := ctfile.PublicCloud().FileList("d41982115", 0, 0, "", "", "")
	if err == nil {
		for i := 0; i < len(files); i++ {
			fmt.Println(files[i])
		}
	} else {
		fmt.Println(err)
	}
}

func FileListFromIds(ctfile z.TCTFile) {
	files, err := ctfile.PublicCloud().FileIdsList([]string{"d41982115", "d39859968"})
	if err == nil {
		for i := 0; i < len(files); i++ {
			fmt.Println(files[i])
		}
	} else {
		fmt.Println(err)
	}
}

func main() {

	var ctfile z.TCTFile
	err := ctfile.LoginFromToken("d11e6b01e40daa604280000c199d4f16")
	if err == nil {

		fmt.Println(ctfile.Quota)
		fmt.Println(ctfile.Bandwidth)
		fmt.Println(ctfile.Profile)
		//基本信息

		FileList(ctfile)

		FileListFromIds(ctfile)
		//文件列取

		forders, e := ctfile.PublicCloud().FolderList(`d0`)
		if e == nil {
			for i := 0; i < len(forders); i++ {
				fmt.Println(forders[i])
			}
			//.........................
			
			// 其他的实现 
		} else {
			fmt.Println(e)
		}
		//列取文件夹
	} else {
		fmt.Println(err)
	}
}

```

* 2025.7.2x 城通网盘大升级 终于把上传的返回改了。
* 原来家在返回的内容内，有一个文件ID，最早未记得还会出现很多空格什么的
* 这次是修改了 希望别改了也。
* 这里声明的新的返回JSON如下：

```go
type TCTFileUploadResult struct {
Id          int64  `json:"id,omitempty"`
Name        string `json:"name,omitempty"`
Size        int    `json:"size,omitempty"`
Time        int    `json:"time,omitempty"`
Userid      int    `json:"userid,omitempty"`
FolderId    int    `json:"folderid,omitempty"`
WorkspaceId int    `json:"workspaceid,omitempty"`
}
```