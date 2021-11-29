package main

import (
	"social/backend/channel"
	"social/backend/constants"
	"social/backend/pref"

	"github.com/cansulting/elabox-system-tools/foundation/app/rpc"
	"github.com/cansulting/elabox-system-tools/foundation/event/data"
	"github.com/cansulting/elabox-system-tools/foundation/event/protocol"
)

type Activity struct {
	savedPref *pref.Data
}

// callback when activity started
func (instance *Activity) OnStart(action *data.Action) error {
	// recieved requests from client
	constants.AppController.RPC.OnRecieved(constants.AC_LOAD_TIMELINES, instance.OnAction_LoadTimeline)
	constants.AppController.RPC.OnRecieved(constants.AC_ADDCHANNEL, instance.OnAction_AddChannel)
	constants.AppController.RPC.OnRecieved(constants.AC_RMCHANNEL, instance.OnAction_RemoveChannel)
	constants.AppController.RPC.OnRecieved(constants.AC_LOAD_DATA, instance.OnAction_LoadData)
	saved, err := pref.LoadPref()
	if err != nil {
		return err
	}
	if saved == nil {
		saved = pref.CreateNew()
	}
	instance.savedPref = saved
	return nil
}

func (instance *Activity) IsRunning() bool {
	return true
}
func (instance *Activity) OnEnd() error {
	return nil
}

func (instance *Activity) OnAction_LoadTimeline(client protocol.ClientInterface, data data.Action) string {
	host := data.DataToString()
	handler := channel.Load(host)
	json, err := handler.RetrieveLocalTimelines(true)
	if err != nil {
		return rpc.CreateResponse(rpc.SYSTEMERR_CODE, err.Error())
	}
	return rpc.CreateJsonResponse(rpc.SUCCESS_CODE, json)
}

func (instance *Activity) OnAction_AddChannel(client protocol.ClientInterface, data data.Action) string {
	channel := data.DataToString()
	if channel != "" {
		instance.savedPref.AddChannel(channel)
		if err := instance.savedPref.Save(); err != nil {
			return rpc.CreateResponse(rpc.SYSTEMERR_CODE, err.Error())
		}
	}
	return rpc.CreateSuccessResponse("added")
}

func (instance *Activity) OnAction_RemoveChannel(client protocol.ClientInterface, data data.Action) string {
	channel := data.DataToString()
	if channel != "" {
		instance.savedPref.RemoveChannel(channel)
		if err := instance.savedPref.Save(); err != nil {
			return rpc.CreateResponse(rpc.SYSTEMERR_CODE, err.Error())
		}
	}
	return rpc.CreateSuccessResponse("removed")
}

func (instance *Activity) OnAction_LoadData(client protocol.ClientInterface, data data.Action) string {
	content := instance.savedPref.ToJson()
	return rpc.CreateJsonResponse(rpc.SUCCESS_CODE, string(content))
}
