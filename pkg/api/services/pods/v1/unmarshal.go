package pods

import (
	"encoding/json"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// UnmarshalYaml reads v1 Pods data in YAML format and unmarshals it to v1 api model
func UnmarshalYaml(data []byte) ([]*Pod, error) {
	target := &[]*Pod{}

	unmarshalErr := yaml.Unmarshal(data, target)
	if unmarshalErr != nil {
		return []*Pod{}, errors.Wrapf(unmarshalErr, "Unable to parse Yaml data")
	}

	return Defaults(*target), nil
}

// UnmarshalJSON reads v1 Pods data in JSON format and unmarshals it to v1 api model
func UnmarshalJSON(data []byte) ([]*Pod, error) {
	target := &[]*Pod{}

	unmarshalErr := json.Unmarshal(data, target)
	if unmarshalErr != nil {
		log.Debugf("Unable to parse JSON: %s", string(data[:]))
		return []*Pod{}, errors.Wrapf(unmarshalErr, "Unable to parse JSON data")
	}

	return Defaults(*target), nil
}
