package bluebox

// Struct get details of each event
type EquipmentEventType struct {
	Name        string
	Description string
}

// Fix variable with data
var EquipmentEventTypes = map[string]EquipmentEventType{
	"LOAD": {"Load", "The action of lifting cargo or a container on board of the mode of transportation. Load is complete once the cargo or container has been lifted on board the mode of transport and secured."},
	"DISC": {"Discharge", "The action of lifting cargo or containers off a mode of transport. Discharge is the opposite of load."},
	"GTIN": {"Gate in", "The action when a container is introduced into a controlled area like a port - or inland terminal. Gate in has been completed once the operator of the area is legally in possession of the container."},
	"GTOT": {"Gate out", "The action when a container is removed from a controlled area like a port – or inland terminal. Gate-out has been completed once the possession of the container has been transferred from the operator of the terminal to the entity who is picking up the container."},
	"STUF": {"Stuffing", "The process of loading the cargo in a container or in/onto another piece of equipment."},
	"STRP": {"Stripping", "The action of unloading cargo from containers or equipment."},
	"PICK": {"Pick-up", "The action of collecting the container at customer location."},
	"DROP": {"Drop-off", "The action of delivering the container at customer location."},
	"INSP": {"Inspected", "Identifies that the seal on equipment has been inspected."},
	"RSEA": {"Resealed", "Identifies that the equipment has been resealed after inspection."},
	"RMVD": {"Removed", "Identifies that a Seal has been removed from the equipment for inspection."},
	"AVPU": {"Available for Pick-up", "Identifies that shipment/ Container is ready to be picked up / collection at a facility."},
	"AVDO": {"Available for Drop-off", "Identifies that shipment/ container is ready to be dropped off / delivered at a facility."},
	"CUSS": {"Customs Selected for Scan", "Identifies that Customs has selected the equipment for scanning."},
	"CUSI": {"Customs Selected for Inspection", "Identifies that that Customs has selected the equipment for inspection."},
	"CUSR": {"Customs Released", "Identifies that Customs has released the equipment for either export from or import into the country."},
	"WAYP": {"Way Point Crossed", "A waypoint is an intermediate point or place during transit of shipment, waypoint crossed indicates that the equipment has crossed the particular waypoint on its transit."},
}

// Struct para armazenar os detalhes do carrier
type Carrier struct {
	NMFTACode   string
	Description string
}

// Variável fixa com os dados dos carriers
var Carriers = map[string]Carrier{
	"CMA":  {"CMDU", "CMA CGM"},
	"EMC":  {"EGLV", "Evergreen Marine Corporation"},
	"HLC":  {"HLCU", "Hapag Lloyd"},
	"HMM":  {"HDMU", "Hyundai"},
	"MSK":  {"MAEU", "Maersk"},
	"MSC":  {"MSCU", "Mediterranean Shipping Company"},
	"ONE":  {"ONEY", "Ocean Network Express Pte. Ltd."},
	"YML":  {"YMLU", "Yang Ming Line"},
	"ZIM":  {"ZIMU", "Zim Israel Navigation Company"},
	"YMM":  {"YMLU", "Yang Ming Line"},
	"HAP":  {"HLCU", "Hapag Lloyd"},
	"EVG":  {"EGLV", "Evergreen Marine Corporation"},
	"CSC":  {"COSU", "Cosco Shipping"},
	"SAF":  {"SAFM", "Safmarine"},
	"SFM":  {"SAFM", "Safmarine"},
	"PIL":  {"PACU", "Pacific International Lines"},
	"ANR":  {"ANRM", "Aliança Navegação e Logística"},
	"OOCL": {"OOLU", "Orient Overseas Container Line Ltd."},
}
