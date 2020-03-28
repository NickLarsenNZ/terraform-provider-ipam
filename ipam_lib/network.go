package ipam_lib

type Network struct {
	ID          *int64
	Description string
	Prefix      string
	ParentId    *int64
	Tags        map[string]string
}
