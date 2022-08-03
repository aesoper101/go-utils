package archivex

type IArchive interface {
	Archive(src string, dest string) error
	UnArchived(src string, dest string) error
}
