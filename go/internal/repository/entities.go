package repository

type Elasticsearch struct {
	ComputerName string `json:"computer_name"`
	Location     struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}
	IP struct {
		IPv4 string `json:"ipv4"`
		IPv6 string `json:"ipv6"`
	}
}
