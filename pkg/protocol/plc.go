package protocol

import (
	"encoding/hex"
	"fmt"
	"main/pkg/parser"
	"strings"
	"time"
)

// PLCEntity .
type PLCEntity struct {
	Result              int16
	PartName            string
	Code2D              string
	RecipeCode          int16
	Index               int16
	PressResult         int16
	ForceLimitMax       float32
	ForceMax            float32
	ForceMin            float32
	PsoitionResult      int16
	PsoitionLimitMax    float32
	PsoitionMax         float32
	PsoitionMin         float32
	TotalAngleResult    int16
	TotalAngleLimitMax  float32
	TotalAngle          float32
	TotalAngleLimitMin  float32
	MiddleAngleResult   int16
	MiddleAngleLimitMax float32
	BeforeMiddleAngle   float32
	AfterMiddleAngle    float32
	MiddleAngleLimitMin float32
}

// DecodeMsg .
func DecodeMsg(msg []byte) *PLCEntity {
	hexStr := hex.EncodeToString(msg)
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
	return entity
}

// GenSQL ret sql
func (p *PLCEntity) GenSQL() string {
	time.LoadLocation("Asia/Shanghai")
	n := time.Now()
	nDbDate := fmt.Sprintf("%v%.2d", n.Year(), int(n.Month()))
	sql := "insert into IPA_%v.dbo.IPA01 (DATE, RESULT, PShaft, Model, MC_Index, DATA01, DATA02, DATA03, DATA04, DATA05, DATA06, DATA07, DATA08, DATA09, DATA10, DATA11, DATA12, DATA13, DATA14, DATA16, DATA37, DATA38) " +
		"values ('%v', %v, '%v', '%v', %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)"
	return fmt.Sprintf(
		sql,
		nDbDate,
		time.Now().Format("2006-01-02 15:04:05"),
		p.Result,
		strings.TrimSpace(strings.Replace(p.Code2D, " ", "", -1)),
		p.RecipeCode,
		p.Index,
		p.PressResult,
		p.ForceMax,
		p.ForceLimitMax,
		p.ForceMin,
		p.PsoitionResult,
		p.PsoitionMax,
		p.PsoitionLimitMax,
		p.PsoitionMin,
		p.TotalAngleResult,
		p.TotalAngleLimitMax,
		p.TotalAngle,
		p.TotalAngleLimitMin,
		p.MiddleAngleResult,
		p.MiddleAngleLimitMax,
		p.MiddleAngleLimitMin,
		p.BeforeMiddleAngle,
		p.AfterMiddleAngle,
	)
}
