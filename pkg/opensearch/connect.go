package opensearch

import (
	"crypto/tls"
	"errors"
	"github.com/cenkalti/backoff/v4"
	"github.com/opensearch-project/opensearch-go/v2"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func connectionToOpenSearch() (bool, error) {
	var cert, err = ioutil.ReadFile(GetInstance().config.CaCert)
	if err != nil {
		return false, errors.New("[ERROR] Error when reading the certificate file. " + err.Error())
	}
	retryBackoff := backoff.NewExponentialBackOff()

	var client, errorClient = opensearch.NewClient(opensearch.Config{
		CACert:    cert,
		Addresses: GetInstance().config.Hosts,
		Username:  GetInstance().config.User,
		Password:  GetInstance().config.Password,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
		RetryOnStatus: []int{502, 503, 504, 429},
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},
		MaxRetries: 5,
	})
	if errorClient != nil {
		return false, errors.New("[ERROR] It is impossible to create connections with Opensearch (`" + strings.Join(GetInstance().config.Hosts, ",") + "`), the configuration is incorrectly configured.")
	}

	var information, errGetInfo = client.Info()
	if errGetInfo != nil {
		return false, errors.New("[ERROR] Error connecting to Opensearch, " + errGetInfo.Error())
	}
	defer func() {
		if err = information.Body.Close(); err != nil {
			log.Println("[ERROR]", err)
		}
	}()

	if information.IsError() {
		return false, errors.New("[ERROR] Error connecting to Opensearch, " + information.String())
	}

	GetInstance().SetClient(client)

	return true, nil
}
