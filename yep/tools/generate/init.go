// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package generate

import (
	"fmt"
	"go/build"

	"github.com/npiganeau/yep/yep/tools/logging"
)

const (
	// YEPPath is the go import path of the base yep package
	YEPPath string = "github.com/npiganeau/yep"
	// ModelsPath is the go import path of the yep/models package
	ModelsPath string = "github.com/npiganeau/yep/yep/models"
	// TypesPath is the go import path of the yep/models/types package
	TypesPath string = "github.com/npiganeau/yep/yep/models/types"
	// GeneratePath is the go import path of this package
	GeneratePath string = "github.com/npiganeau/yep/yep/tools/generate"
	// PoolPath is the go import path of the autogenerated pool package
	PoolPath string = "github.com/npiganeau/yep/pool"
)

var (
	log *logging.Logger
	// YEPDir is the directory of the base yep package
	YEPDir string
)

func init() {
	log = logging.GetLogger("tools/generate")
	yepPack, err := build.Import(YEPPath, ".", build.FindOnly)
	if err != nil {
		panic(fmt.Errorf("Error while getting YEP root path: %s", err))
	}
	YEPDir = yepPack.Dir
}
