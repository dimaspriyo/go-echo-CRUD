package elasticsearch

type MyLog struct {
	HostName string `json:"host_name"`
	Location Location
	IP       IP
}

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type IP struct {
	IPv4 string `json:"ipv4"`
}
