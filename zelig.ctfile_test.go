package ZeligCTFile

import (
	"fmt"
	"testing"
)

const cToken = "2dcd93bd3cf0e27b568225494d341ef9"

func TestTCTFilePublic(t *testing.T) {
	var ctfile TCTFile
	//	err := ctfile.Login("email", "password")

	err := ctfile.LoginFromToken(cToken)
	if err == nil {

		fmt.Println(ctfile.token)
		fmt.Println(ctfile.Quota)
		fmt.Println(ctfile.Bandwidth)
		fmt.Println(ctfile.Profile)
		/*
			data, _ := ctfile.PublicCloud().FileShare([]string{"f478659195"})
			t.Log(data)

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

				forders, err := ctfile.PublicCloud().FolderList(`d0`)
				if err == nil {
					for i := 0; i < len(forders); i++ {
						fmt.Println(forders[i])
					}
				} else {
					fmt.Println(err)
				}
		*/
		fid, _ := ctfile.PublicCloud().FileUpload(`d48182796`, `E:\Tools\Tools\CleanMaster.exe`)
		fmt.Println(fid)
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
