package pingconnection

import "github.com/elastic/go-elasticsearch/v8"

func PingElastic(db *elasticsearch.Client) error {
	_, err := db.Ping()
	if err != nil {
		return err
	}

	return nil
}
