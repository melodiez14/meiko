package rolegroup

import (
	"encoding/json"
	"fmt"
	"html"
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
