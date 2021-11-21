package ZeligCTFile

import (
	"fmt"
	"testing"
)

const cToken = "a67519c04f954499f4fab2e1817c8fcf"

func TestTCTFilePublic(t *testing.T) {
	var ctfile TCTFile
	//	err := ctfile.Login("email", "password")

	err := ctfile.LoginFromToken(cToken)
	if err == nil {

		fmt.Println(ctfile.token)
		fmt.Println(ctfile.Quota)
		fmt.Println(ctfile.Bandwidth)
		fmt.Println(ctfile.Profile) /*
					res := make(map[string]string)
					res, err = ctfile.FolderCreate("0", "测试创建的文件夹", "测试的文件夹咯", 0)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println(res)
					}

				folders, err := ctfile.PublicCloud().FolderList("0")
				if err == nil {
					for i := 0; i < len(folders); i++ {
						fmt.Println(folders[i])
					}
				}

				files, err := ctfile.PublicCloud().FileList("d41982115", 0, 0, "", "", "")
				if err == nil {
					for i := 0; i < len(files); i++ {
						fmt.Println(files[i])
					}
				} else {
					fmt.Println(err)
				}

			files, err := ctfile.PublicCloud().FileIdsList([]string{"d41982115", "d39859968"})
			if err == nil {
				for i := 0; i < len(files); i++ {
					fmt.Println(files[i])
				}
			} else {
				fmt.Println(err)
			}*/
		/*
			forders, err := ctfile.PublicCloud().FolderList(`d0`)
			if err == nil {
				for i := 0; i < len(forders); i++ {
					fmt.Println(forders[i])
				}
			} else {
				fmt.Println(err)
			}
		*/
		//err = ctfile.PublicCloud().FileUpload(`d44189303`, `D:\\TortoiseSVN-1.14.1.29085-x64-svn-1.14.1.msi`)
		//fmt.Println(err)
	} else {
		fmt.Println(err)
	}
}

func TestTCTFilePrivate(t *testing.T) {
	var ctfile TCTFile
	err := ctfile.LoginFromToken(cToken)
	if err == nil {

		fmt.Println(ctfile.token)
		fmt.Println(ctfile.Quota)
		fmt.Println(ctfile.Bandwidth)
		fmt.Println(ctfile.Profile)

		forders, e := ctfile.PrivateCloud().FolderList(`d0`)
		if e == nil {
			for i := 0; i < len(forders); i++ {
				fmt.Println(forders[i])
			}
		} else {
			fmt.Println(e)
		}
	} else {
		fmt.Println(err)
	}
}
