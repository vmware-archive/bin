package bindata

import "errors"

func AssetDir(string) ([]string, error) {
	return nil, errors.New("no assets")
}

func RestoreAssets(string, string) error {
	return errors.New("no assets")
}
