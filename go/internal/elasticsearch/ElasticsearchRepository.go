package elasticsearch

import (
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/labstack/echo/v4"

	"go/pkg/ipinfo"
)

type IElasticsearchRepository interface {
	Log(es *elasticsearch.Client) error
}

type ElasticRepository struct {
}

func (r ElasticRepository) Log(es *elasticsearch.Client, c echo.Context) {
	result, err := ipinfo.ForeignIP(c.Request().RemoteAddr)
	if err != nil {
		fmt.Println(err.Error())
	}

	var Loc = strings.Split(result.Loc, ",")

	var myLog = MyLog{
		HostName: result.Hostname,
		IP: IP{
			IPv4: result.IP,
		},
		Location: Location{
			Latitude:  Loc[0],
			Longitude: Loc[1],
		},
	}

	res, err := es.Index("MyLog")
}
