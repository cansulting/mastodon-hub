
import { PKID } from "../config"
const AC_LOADTIMELINE = "LOAD_TIMELINES"
const AC_ADDCHANNEL = "ADD_CHANNEL"
const AC_LOADDATA = "LOAD_DATA"
const AC_RMCHANNEL = "REMOVE_CHANNEL"

export function retrieveTimeline(actionCenter, host = "") {
    return new Promise((resolve, reject) => {
        actionCenter.sendRPC(PKID, {id: AC_LOADTIMELINE, data:host}, (response) => {
            console.log(response)
            resolve(response)
        })
    })
}

export function addChannel(actionCenter, channel = "") {
    return new Promise((resolve, reject) => {
        actionCenter.sendRPC(PKID, {id: AC_ADDCHANNEL, data:channel}, (response) => {
            console.log(response)
            resolve(response)
        })
    })
}

export function removeChannel(actionCenter, channel = "") {
    return new Promise((resolve, reject) => {
        actionCenter.sendRPC(PKID, {id: AC_RMCHANNEL, data:channel}, (response) => {
            console.log(response)
            resolve(response)
        })
    })
}

export function retrieveInitData(actionCenter) {
    return new Promise((resolve, reject) => {
        actionCenter.sendRPC(PKID, {id: AC_LOADDATA}, (response) => {
            console.log(response)
            resolve(response)
        })
    })
}