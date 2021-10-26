package bigdatacloud

import (
	"github.com/evalphobia/bigdatacloud-go/client"
	"github.com/evalphobia/bigdatacloud-go/config"
	"github.com/evalphobia/bigdatacloud-go/log"
)

// BigDataCloud is service struct for BigDataCloud API.
type BigDataCloud struct {
	client   *client.Client
	logger   log.Logger
	language string
}

// New creates BigDataCloud from Config data.
func New(conf config.Config) (*BigDataCloud, error) {
	cli, err := conf.Client()
	if err != nil {
		return nil, err
	}

	return &BigDataCloud{
		client:   cli,
		logger:   log.DefaultLogger,
		language: conf.GetLanguage(),
	}, nil
}

// SetLogger sets internal API logger.
func (s *BigDataCloud) SetLogger(logger log.Logger) {
	s.logger = logger
}

// IPGeolocationFull executes IP Address Geolocation with Confidence Area and Hazard Report API.
// ref: https://www.bigdatacloud.com/ip-geolocation-apis/ip-address-geolocation-with-confidence-area-and-hazard-report-api
func (s *BigDataCloud) IPGeolocationFull(ipaddr string) (IPResponse, error) {
	params := make(map[string]interface{})
	params["ip"] = ipaddr
	if s.language != "" {
		params["localityLanguage"] = s.language
	}

	resp := IPResponse{}
	err := s.client.CallGET("/data/ip-geolocation-full", params, &resp)
	return resp, err
}

// IPGeolocationConfidenceArea executes IP Address Geolocation with Confidence Area API.
// ref: https://www.bigdatacloud.com/ip-geolocation-apis/ip-address-geolocation-with-confidence-area-api
func (s *BigDataCloud) IPGeolocationConfidenceArea(ipaddr string) (IPResponse, error) {
	params := make(map[string]interface{})
	params["ip"] = ipaddr
	if s.language != "" {
		params["localityLanguage"] = s.language
	}

	resp := IPResponse{}
	err := s.client.CallGET("/data/ip-geolocation-with-confidence", params, &resp)
	return resp, err
}

// IPGeolocation executes IP Address Geolocation API.
// ref: https://www.bigdatacloud.com/ip-geolocation-apis/ip-address-geolocation-api
func (s *BigDataCloud) IPGeolocation(ipaddr string) (IPResponse, error) {
	params := make(map[string]interface{})
	params["ip"] = ipaddr
	if s.language != "" {
		params["localityLanguage"] = s.language
	}

	resp := IPResponse{}
	err := s.client.CallGET("/data/ip-geolocation", params, &resp)
	return resp, err
}
