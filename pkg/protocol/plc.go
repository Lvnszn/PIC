package protocol

import (
	"encoding/hex"
	"fmt"
	"main/pkg/logger"
	"main/pkg/parser"
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
	entity := &PLCEntity{}
	entity.Result = parser.HexToInt16(hex.EncodeToString(msg[22:26]))
	entity.PartName = parser.HexToString(msg[26:50])
	entity.Code2D = parser.HexToString(msg[50:182])
	entity.RecipeCode = parser.HexToInt16(hex.EncodeToString(msg[182:186]))
	entity.Index = parser.HexToInt16(hex.EncodeToString(msg[186:190]))
	entity.PressResult = parser.HexToInt16(hex.EncodeToString(msg[190:194]))
	entity.ForceLimitMax = parser.HexToFloat32(hex.EncodeToString(msg[194:202]))
	entity.ForceMax = parser.HexToFloat32(hex.EncodeToString(msg[202:210]))
	entity.ForceMin = parser.HexToFloat32(hex.EncodeToString(msg[210:218]))
	entity.PsoitionResult = parser.HexToInt16(hex.EncodeToString(msg[218:222]))
	entity.PsoitionLimitMax = parser.HexToFloat32(hex.EncodeToString(msg[222:230]))
	entity.PsoitionMax = parser.HexToFloat32(hex.EncodeToString(msg[230:238]))
	entity.PsoitionMin = parser.HexToFloat32(hex.EncodeToString(msg[238:246]))
	entity.TotalAngleResult = parser.HexToInt16(hex.EncodeToString(msg[246:250]))
	entity.TotalAngleLimitMax = parser.HexToFloat32(hex.EncodeToString(msg[250:258]))
	entity.TotalAngle = parser.HexToFloat32(hex.EncodeToString(msg[258:266]))
	entity.TotalAngleLimitMin = parser.HexToFloat32(hex.EncodeToString(msg[266:274]))
	entity.MiddleAngleResult = parser.HexToInt16(hex.EncodeToString(msg[274:278]))
	entity.MiddleAngleLimitMax = parser.HexToFloat32(hex.EncodeToString(msg[278:286]))
	entity.BeforeMiddleAngle = parser.HexToFloat32(hex.EncodeToString(msg[286:294]))
	entity.AfterMiddleAngle = parser.HexToFloat32(hex.EncodeToString(msg[294:302]))
	entity.MiddleAngleLimitMin = parser.HexToFloat32(hex.EncodeToString(msg[302:310]))
	logger.Printf("decode msg show: %+v", entity)
	return entity
}

// GenSQL ret sql
func (p *PLCEntity) GenSQL() string {
	time.LoadLocation("Asia/Shanghai")
	sql := "insert into IPA01 values (DATE, RESULT, PShaft, Model,Index, DATA01, DATA02, DATA03, DATA04, DATA05, DATA06, DATA07, DATA08, DATA09, DATA10, DATA11, DATA12, DATA13, DATA14, DATA16, DATA37, DATA38) " +
		"values (%q, %v, %q, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)"
	return fmt.Sprintf(
		sql,
		time.Now().Format("2006-01-02 15:04:05"),
		p.Result,
		p.Code2D,
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
