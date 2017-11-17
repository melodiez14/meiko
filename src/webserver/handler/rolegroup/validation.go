package rolegroup

import (
	"encoding/json"
	"fmt"
	"html"
	"strconv"
	"strings"

	rg "github.com/melodiez14/meiko/src/module/rolegroup"
	"github.com/melodiez14/meiko/src/util/helper"
)

func (params createParams) validate() (createArgs, error) {

	var args createArgs
	params = createParams{
		name:    html.EscapeString(params.name),
		modules: params.modules,
	}

	if helper.IsEmpty(params.name) {
		return args, fmt.Errorf("Name cannot be empty")
	}

	if !helper.IsAlphaSpace(params.name) {
		return args, fmt.Errorf("Name should only contain an alphabet and space")
	}

	name := helper.Trim(params.name)
	name = strings.Title(name)

	if helper.IsEmpty(params.modules) {
		return createArgs{name: name, modules: map[string][]string{}}, nil
	}

	modules := map[string][]string{}
	err := json.Unmarshal([]byte(params.modules), &modules)
	if err != nil {
		return args, nil
	}

	// validate all modules
	for mdl, roles := range modules {
		listModule := rg.GetModuleList()
		listAbility := rg.GetAbilityList()
		if !helper.IsStringInSlice(mdl, listModule) {
			return args, fmt.Errorf("Invalid Request")
		}

		for _, val := range roles {
			if !helper.IsStringInSlice(val, listAbility) {
				return args, fmt.Errorf("Invalid Request")
			}
		}
	}

	return createArgs{name: name, modules: modules}, nil
}

func (params readParams) validate() (readArgs, error) {
	var args readArgs

	page, err := strconv.ParseUint(params.page, 10, 8)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	total, err := strconv.ParseUint(params.total, 10, 8)
	if err != nil {
		return args, fmt.Errorf("Invalid request")
	}

	// should be positive number
	if page < 1 || total < 1 {
		return args, fmt.Errorf("Invalid request")
	}

	return readArgs{
		page:  uint8(page),
		total: uint8(total),
	}, nil
}

func (params readDetailParams) validate() (readDetailArgs, error) {

	var args readDetailArgs
	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, err
	}

	return readDetailArgs{id: id}, nil
}

func (params deleteParams) validate() (deleteArgs, error) {

	var args deleteArgs
	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, err
	}

	return deleteArgs{id: id}, nil
}

func (params updateParams) validate() (updateArgs, error) {

	var args updateArgs
	params = updateParams{
		id:      params.id,
		name:    html.EscapeString(params.name),
		modules: params.modules,
	}

	if helper.IsEmpty(params.id) {
		return args, fmt.Errorf("Invalid Request")
	}

	id, err := strconv.ParseInt(params.id, 10, 64)
	if err != nil {
		return args, fmt.Errorf("Invalid Request")
	}

	if helper.IsEmpty(params.name) {
		return args, fmt.Errorf("Name cannot be empty")
	}

	if !helper.IsAlphaSpace(params.name) {
		return args, fmt.Errorf("Name should only contain an alphabet and space")
	}

	name := helper.Trim(params.name)
	name = strings.Title(name)

	if helper.IsEmpty(params.modules) {
		return updateArgs{id: id, name: name, modules: map[string][]string{}}, nil
	}

	modules := map[string][]string{}
	err = json.Unmarshal([]byte(params.modules), &modules)
	if err != nil {
		return args, nil
	}

	// validate all modules
	for mdl, roles := range modules {
		listModule := rg.GetModuleList()
		listAbility := rg.GetAbilityList()
		if !helper.IsStringInSlice(mdl, listModule) {
			return args, fmt.Errorf("Invalid Request")
		}

		for _, val := range roles {
			if !helper.IsStringInSlice(val, listAbility) {
				return args, fmt.Errorf("Invalid Request")
			}
		}
	}

	return updateArgs{id: id, name: name, modules: modules}, nil
}
