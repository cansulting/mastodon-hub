package constants

import "github.com/cansulting/elabox-system-tools/foundation/app"

const TIMELINES_ENDPOINT = "/api/v1/timelines/public"

// ACTIONS
const AC_LOAD_DATA = "LOAD_DATA"
const AC_LOAD_TIMELINES = "LOAD_TIMELINES"
const AC_ADDCHANNEL = "ADD_CHANNEL"
const AC_RMCHANNEL = "REMOVE_CHANNEL"

const USERPREF = "user.pref"

var AppController *app.Controller
