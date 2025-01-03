package store

// definition database table with class
type File struct {
	ModelBase
	Name        string
	Size        int64   `json:"size"`
	Type        string  `json:"type"`
	Tags        *string `json:"tags"`
	Ticket      string  `json:"ticket"`
	StorageName string  `json:"storageName"`
}

// create file
func (s *Store) Creatfile(file File) error {
	res := s.db.Create(&file)

	return res.Error
}

// delete file record by id
func (s *Store) Deletefile(id int) error {
	res := s.db.Delete(&File{}, id)

	return res.Error

}

// get file record by id
func (s *Store) FindfileById(id int) (*File, error) {
	var files []File
	res := s.db.Where("id", id).Find(&files)

	if len(files) == 0 {
		return nil, res.Error
	}
	return &files[0], res.Error
}

// query files by keyword and Limit & Offset
func (s *Store) Queryfile(keyWord *string, offset int, take int) (*[]File, error) {

	q := s.db.Model(&File{})

	if keyWord != nil && *keyWord != "" {
		q = q.Where("name like", *keyWord+"%")
	}

	var files []File
	rest := q.Order("created_at desc").Offset(offset).Limit(take).Scan(&files)

	return &files, rest.Error

}
