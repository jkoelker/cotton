//

package config

import (
	"context"
	"fmt"

	"github.com/omeid/uconfig"
	"github.com/omeid/uconfig/plugins/file"
	"sigs.k8s.io/yaml"
)

// UConfigYamlFile is a helper function to create a `uconfig.Files` entry.
func UConfigYamlFile(path string) struct {
	Path      string
	Unmarshal file.Unmarshal
	Optional  bool
} {
	return struct {
		Path      string
		Unmarshal file.Unmarshal
		Optional  bool
	}{
		Path: path,
		Unmarshal: func(data []byte, value any) error {
			// NOTE k8s Unmarshal takes optional JSONOpts, so we wrap it.
			return yaml.Unmarshal(data, value)
		},
		Optional: true,
	}
}

// UConfig returns a `uconfig.Config` that can be used to load configuration.
//
//nolint:ireturn // UConfig does not export the `config` type only the interface.
func UConfig(config any, files ...string) (uconfig.Config, error) {
	var configFiles uconfig.Files

	for _, file := range files {
		configFiles = append(configFiles, UConfigYamlFile(file))
	}

	conf, err := uconfig.Classic(config, configFiles)
	if err != nil {
		return nil, fmt.Errorf("`uconfig.Classic` initialization failed: %w", err)
	}

	return conf, nil
}

// LoadUConfig is a function that can be passed to `WithLoaders`.
// It uses `uconfig.Classic` to load configuration into the provided struct.
func LoadUConfig[C any](files ...string) func(context.Context, *C) error {
	return func(_ context.Context, config *C) error {
		if _, err := UConfig(config, files...); err != nil {
			return err
		}

		return nil
	}
}
