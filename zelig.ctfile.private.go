package ZeligCTFile

/*
   	城通网盘的 API 实现  http://openapi.ctfile.com/
    作者 Icy
    Web zelig.cn

    私有云的方法

*/

type TCTFilePrivate struct {
	ctfile *TCTFile
}

func newCTFilePrivate(ctf *TCTFile) *TCTFilePrivate {
	return &TCTFilePrivate{
		ctfile: ctf,
	}
}

func (this *TCTFilePrivate) ForderCreate(Folder_id, Name, Description string, Is_Hidden int) (map[string]string, error) {
	return this.ctfile.folderCreate(false, Folder_id, Name, Description, Is_Hidden)
}

func (this *TCTFilePrivate) FolderMeta(Folder_id string) (TCTFileFolderMeta, error) {
	return this.ctfile.folderMeta(false, Folder_id)
}

func (this *TCTFilePrivate) FolderList(Folder_id string) (TCTFileFolders, error) {
	return this.ctfile.folderList(false, Folder_id)
}

func (this *TCTFilePrivate) FolderModifyMeta(Folder_id, Name, Description string, Is_Hidden int) (map[string]string, error) {
	return this.ctfile.folderModifyMeta(false, Folder_id, Name, Description, Is_Hidden)
}

func (this *TCTFilePrivate) FileList(Folder_id string, Start, Reload int, Orderby, Filter, Keyword string) (TCTFileFolderFiles, error) {
	return this.ctfile.fileList(false, Folder_id, Start, Reload, Orderby, Filter, Keyword)
}

func (this *TCTFilePrivate) FileIdsList(Ids []string) (TCTFileFolderFiles, error) {
	return this.ctfile.fileIdsList(false, Ids)
}

func (this *TCTFilePrivate) FileRecycle(Start, Reload int) (TCTFileFolderFileRecycles, error) {
	return this.ctfile.fileRecycle(false, Start, Reload)
}

func (this *TCTFilePrivate) FileRecycle_empty(Ids []string) error {
	return this.ctfile.fileRecycle_empty(false, Ids)
}

func (this *TCTFilePrivate) FileRecycle_empty_all() error {
	return this.ctfile.fileRecycle_empty_all(false)
}

func (this *TCTFilePrivate) FileDownload(Ids []string) (TCTFileFolderFileDownloads, error) {
	return this.ctfile.fileDownload(false, Ids)
}

func (this *TCTFilePrivate) FileFetch_urlb(File_id string) (string, error) {
	return this.ctfile.fileFetch_urlb(false, File_id)
}

func (this *TCTFilePrivate) FileShare(Ids []string) (TCTFileFolderFileShares, error) {
	return this.ctfile.fileShare(false, Ids)
}

func (this *TCTFilePrivate) FileMove(Folder_id string, Ids []string) error {
	return this.ctfile.fileMove(false, Folder_id, Ids)
}

func (this *TCTFilePrivate) FileDelete(Ids []string) error {
	return this.ctfile.fileDelete(false, Ids)
}

func (this *TCTFilePrivate) FileSave(Ids []string) error {
	return this.ctfile.fileSave(false, Ids)
}

func (this *TCTFilePrivate) FileUpload(Folder_id, Filename string) error {
	return this.ctfile.fileUpload(false, Folder_id, Filename)
}
