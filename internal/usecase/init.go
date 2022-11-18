package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"friends/pkg/cryptograph"
	"friends/pkg/execute"
	"os"
	"path"
)

var (
	ErrInitStorageError = errors.New("storage not specified")
)

type InitReq struct {
	Storage string
}

func (req *InitReq) Validate() error {
	if req.Storage == "" {
		return fmt.Errorf("%w[%s]", ErrInitStorageError, req.Storage)
	}

	return nil
}

type InitRes struct {
}

// Init command
// TODO: Rewrite using gocloud.dev package integration
func Init(req InitReq) (*InitRes, error) {

	uDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	settingsDir := path.Join(uDir, ".friends")

	if err = execute.CheckDir(settingsDir); err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil, err
		}
	}

	s := &Settings{
		Storage: req.Storage,
		UserKey: cryptograph.GenPrivateKey(),
	}

	if err = execute.MkDir(settingsDir); err != nil {
		return nil, err
	}

	settingsByt, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return nil, err
	}

	settingsPath := path.Join(settingsDir, "settings.json")

	f, err := os.Create(settingsPath)
	if err != nil {
		return nil, err
	}

	if _, err = f.Write(settingsByt); err != nil {
		return nil, err
	}

	return &InitRes{}, nil
}

type Settings struct {
	Storage string `json:"storage"`
	UserKey string `json:"userKey"`
}
