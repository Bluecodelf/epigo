package epigo

import (
	"encoding/json"
)

type Netsoul []NetsoulPeriod

type NetsoulPeriod struct {
	Timestamp        float64
	SecondsActive    float64
	SecondsIdle      float64
	SecondsOutActive float64
	SecondsOutIdle   float64
	Norm             float64
}

func (netsoul *NetsoulPeriod) UnmarshalJSON(data []byte) (err error) {
	var jsonArray [6]float64
	err = json.Unmarshal(data, &jsonArray)
	netsoul.Timestamp = jsonArray[0]
	netsoul.SecondsActive = jsonArray[1]
	netsoul.SecondsIdle = jsonArray[2]
	netsoul.SecondsOutActive = jsonArray[3]
	netsoul.SecondsOutIdle = jsonArray[4]
	netsoul.Norm = jsonArray[5]
	return
}

func (client *Client) GetNetsoul(login string) ([]NetsoulPeriod, error) {
	url := client.Host + "/user/" + login + "/netsoul/?format=json"
	data, err := client.GetData(url)
	if err != nil {
		return make([]NetsoulPeriod, 0, 0), err
	}
	var netsoul []NetsoulPeriod
	err = json.Unmarshal(data, &netsoul)
	return netsoul, err
}
