package ngapType

// Need to import "github.com/free5gc/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct QosFlowPerTNLInformationList */
/* QosFlowPerTNLInformationItem */
type QosFlowPerTNLInformationList struct {
	List []QosFlowPerTNLInformationItem `aper:"valueExt,sizeLB:1,sizeUB:3"`
}
