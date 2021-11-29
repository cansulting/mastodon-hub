package pref

import (
	"encoding/json"
	"os"
	"social/backend/constants"
)

type Data struct {
	Channels []string `json:"channels"`
}

func getPath() string {
	return constants.AppController.Config.GetDataDir() + "/" + constants.USERPREF
}

func CreateNew() *Data {
	new := &Data{}
	new.Channels = make([]string, 0, 4)
	return new
}

func LoadPref() (*Data, error) {
	content, err := os.ReadFile(getPath())
	if err != nil {
		println(err)
		return nil, nil
	}
	data := Data{}
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func SavePref(data *Data) error {
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := os.WriteFile(getPath(), content, 0666); err != nil {
		return err
	}
	return nil
}

func (instance *Data) AddChannel(channel string) {
	for _, val := range instance.Channels {
		if val == channel {
			return
		}
	}
	instance.Channels = append(instance.Channels, channel)
}

func (instance *Data) RemoveChannel(channel string) {
	i := -1
	for index, val := range instance.Channels {
		if val == channel {
			i = index
			break
		}
	}
	if i >= 0 {
		instance.Channels = append(instance.Channels[:i], instance.Channels[i+1:]...)
	}
}

func (instance *Data) ToJson() []byte {
	content, err := json.Marshal(instance)
	if err != nil {
		println(err)
		return nil
	}
	return content
}

func (instance *Data) Save() error {
	return SavePref(instance)
}
