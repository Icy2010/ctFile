package ZeligCTFile

/*
   	城通网盘的 API 实现  http://openapi.ctfile.com/
    作者 Icy
    Web zelig.cn

    公有云的方法

*/

type TCTFilePublic struct {
	ctfile *TCTFile
}

func newCTFilePublic(ctf *TCTFile) *TCTFilePublic {
	return &TCTFilePublic{
		ctfile: ctf,
	}
}

func (this *TCTFilePublic) ForderCreate(Folder_id, Name, Description string, Is_Hidden int) (map[string]string, error) {
	return this.ctfile.folderCreate(true, Folder_id, Name, Description, Is_Hidden)
}

func (this *TCTFilePublic) FolderMeta(Folder_id string) (TCTFileFolderMeta, error) {
	return this.ctfile.folderMeta(true, Folder_id)
}

func (this *TCTFilePublic) FolderList(Folder_id string) (TCTFileFolders, error) {
	return this.ctfile.folderList(true, Folder_id)
}

func (this *TCTFilePublic) FolderModifyMeta(Folder_id, Name, Description string, Is_Hidden int) (map[string]string, error) {
	return this.ctfile.folderModifyMeta(true, Folder_id, Name, Description, Is_Hidden)
}

func (this *TCTFilePublic) FileList(Folder_id string, Start, Reload int, Orderby, Filter, Keyword string) (TCTFileFolderFiles, error) {
	return this.ctfile.fileList(true, Folder_id, Start, Reload, Orderby, Filter, Keyword)
}

func (this *TCTFilePublic) FileIdsList(Ids []string) (TCTFileFolderFiles, error) {
	return this.ctfile.fileIdsList(true, Ids)
}

func (this *TCTFilePublic) FileRecycle(Start, Reload int) (TCTFileFolderFileRecycles, error) {
	return this.ctfile.fileRecycle(true, Start, Reload)
}

func (this *TCTFilePublic) FileRecycle_empty(Ids []string) error {
	return this.ctfile.fileRecycle_empty(true, Ids)
}

func (this *TCTFilePublic) FileRecycle_empty_all() error {
	return this.ctfile.fileRecycle_empty_all(true)
}

func (this *TCTFilePublic) FileDownload(Ids []string) (TCTFileFolderFileDownloads, error) {
	return this.ctfile.fileDownload(true, Ids)
}

func (this *TCTFilePublic) FileFetch_urlb(File_id string) (string, error) {
	return this.ctfile.fileFetch_urlb(true, File_id)
}

func (this *TCTFilePublic) FileShare(Ids []string) (TCTFileFolderFileShares, error) {
	return this.ctfile.fileShare(true, Ids)
}

func (this *TCTFilePublic) FileMove(Folder_id string, Ids []string) error {
	return this.ctfile.fileMove(true, Folder_id, Ids)
}

func (this *TCTFilePublic) FileDelete(Ids []string) error {
	return this.ctfile.fileDelete(true, Ids)
}

func (this *TCTFilePublic) FileSave(Ids []string) error {
	return this.ctfile.fileSave(true, Ids)
}

func (this *TCTFilePublic) FileUpload(Folder_id, Filename string) (string, error) {
	return this.ctfile.fileUpload(true, Folder_id, Filename)
}

func (this *TCTFilePublic) FileMeta(file_id string) (TCTFileFileMeta, error) {
	return this.ctfile.fileMeta(true, file_id)
}
