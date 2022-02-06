package main

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RouteOperator struct {
	ensOperator *EnsOperator
}

func (routeOperator RouteOperator) getName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name, err := routeOperator.ensOperator.ResolveAddress(ps.ByName("address"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// request was performed successfully but the address does not have any ens name registered (returned empty string)
	if name == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(name)
}

func (routeOperator RouteOperator) getAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	address, err := routeOperator.ensOperator.ResolveName(ps.ByName("name"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}

// for testing purposes - eth.ens has avatar set
func (routeOperator RouteOperator) getAvatar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	avatar, err := routeOperator.ensOperator.GetAvatarByName(ps.ByName("name"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// todo handle empty string

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(avatar)
}

type GetAllResponse struct {
	address string
	name    string
	avatar  string
}

// param could be either address or ENS name
func (routeOperator RouteOperator) getAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("param")
	if common.IsHexAddress(param) {
		name, err := routeOperator.ensOperator.ResolveAddress(param)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var avatar string
		avatar, err = routeOperator.ensOperator.GetAvatarByName(name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(GetAllResponse{address: param, name: name, avatar: avatar})
	} else {
		address, err := routeOperator.ensOperator.ResolveName(param)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var avatar string
		avatar, err = routeOperator.ensOperator.GetAvatarByName(param)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(GetAllResponse{address: address.Hex(), name: param, avatar: avatar})
	}
}
