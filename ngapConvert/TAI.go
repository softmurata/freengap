package ngapConvert

import (
	"encoding/hex"

	"github.com/softmurata/freengap/logger"
	"github.com/softmurata/freengap/ngapType"
	"github.com/softmurata/freeopenapi/models"
)

func TaiToModels(tai ngapType.TAI) models.Tai {
	var modelsTai models.Tai

	plmnID := PlmnIdToModels(tai.PLMNIdentity)
	modelsTai.PlmnId = &plmnID
	modelsTai.Tac = hex.EncodeToString(tai.TAC.Value)

	return modelsTai
}

func TaiToNgap(tai models.Tai) ngapType.TAI {
	var ngapTai ngapType.TAI

	ngapTai.PLMNIdentity = PlmnIdToNgap(*tai.PlmnId)
	if tac, err := hex.DecodeString(tai.Tac); err != nil {
		logger.NgapLog.Warnf("Decode TAC failed: %+v", err)
	} else {
		ngapTai.TAC.Value = tac
	}
	return ngapTai
}
