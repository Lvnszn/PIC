package protocol

import (
	"fmt"
	"main/pkg/parser"
	"testing"
)

func TestParse(t *testing.T) {
	hexStr := "0401260000000000000500020A0454455431000000000000400054454C3151494D413137393132333536000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000B0007000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040123000000"
	entity := &PLCEntity{}
	base := 20
	entity.Result = parser.HexToInt16(hexStr[base : base+4])
	entity.PartName = parser.HexToString(hexStr[28:38])
	entity.Code2D = parser.HexToString(hexStr[52:116])
	entity.RecipeCode = parser.HexToInt16(hexStr[180:184])
	entity.Index = parser.HexToInt16(hexStr[184:188])
	entity.PressResult = parser.HexToInt16(hexStr[188:192])
	entity.ForceLimitMax = parser.HexToFloat32(hexStr[192:200])
	entity.ForceMax = parser.HexToFloat32(hexStr[200:208])
	entity.ForceMin = parser.HexToFloat32(hexStr[208:216])
	entity.PsoitionResult = parser.HexToInt16(hexStr[216:220])
	entity.PsoitionLimitMax = parser.HexToFloat32(hexStr[220:228])
	entity.PsoitionMax = parser.HexToFloat32(hexStr[228:236])
	entity.PsoitionMin = parser.HexToFloat32(hexStr[236:244])
	entity.TotalAngleResult = parser.HexToInt16(hexStr[244:248])
	entity.TotalAngleLimitMax = parser.HexToFloat32(hexStr[248:256])
	entity.TotalAngle = parser.HexToFloat32(hexStr[256:264])
	entity.TotalAngleLimitMin = parser.HexToFloat32(hexStr[264:272])
	entity.MiddleAngleResult = parser.HexToInt16(hexStr[272:276])
	entity.MiddleAngleLimitMax = parser.HexToFloat32(hexStr[276:284])
	entity.BeforeMiddleAngle = parser.HexToFloat32(hexStr[284:292])
	entity.AfterMiddleAngle = parser.HexToFloat32(hexStr[292:300])
	entity.MiddleAngleLimitMin = parser.HexToFloat32(hexStr[300:308])
	fmt.Printf("%+v", entity)
	t.Logf("%+v", entity)
}

