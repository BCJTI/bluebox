package bluebox

/*
Booking{
receiptTypeAtOrigin*	receiptTypeAtOriginstring
maxLength: 3
example: CY

Indicates the type of service offered at Origin. Options are defined in the Receipt/Delivery entity.

	CY (Container yard (incl. rail ramp))
	SD (Store Door)
	CFS (Container Freight Station)

Enum:
Array [ 3 ]
deliveryTypeAtDestination*	deliveryTypeAtDestinationstring
maxLength: 3
example: CY

Indicates the type of service offered at Destination. Options are defined in the Receipt/Delivery entity.

	CY (Container yard (incl. rail ramp))
	SD (Store Door)
	CFS (Container Freight Station)

Enum:
Array [ 3 ]
cargoMovementTypeAtOrigin*	cargoMovementTypeAtOriginstring
maxLength: 3
example: FCL

Refers to the shipment term at the loading of the cargo into the container. Options are defined in the Cargo Movement Type entity. Possible values are:

	FCL (Full Container Load)
	LCL (Less than Container Load)

cargoMovementTypeAtDestination*	cargoMovementTypeAtDestinationstring
maxLength: 3
example: FCL

Refers to the shipment term at the unloading of the cargo out of the container. Options are defined in the Cargo Movement Type entity. Possible values are:

	FCL (Full Container Load)
	LCL (Less than Container Load)

serviceContractReference	serviceContractReferencestring
maxLength: 30
example: HHL51800000

Reference number for agreement between shipper and carrier through which the shipper commits to provide a certain minimum quantity of cargo over a fixed period, and the carrier commits to a certain rate or rate schedule.
freightPaymentTermCode	freightPaymentTermCodestring
example: PRE

An indicator of whether freight and ancillary fees for the main transport are prepaid (PRE) or collect (COL). When prepaid the charges are the responsibility of the shipper or the Invoice payer on behalf of the shipper (if provided). When collect, the charges are the responsibility of the consignee or the Invoice payer on behalf of the consignee (if provided).

	PRE (Prepaid)
	COL (Collect)

Enum:
Array [ 2 ]
originChargesPaymentTermCode	originChargesPaymentTermCodestring
example: PRE

An indicator of whether origin charges are prepaid (PRE) or collect (COL). When prepaid, the charges are the responsibility of the shipper or the Invoice payer on behalf of the shipper (if provided). When collect, the charges are the responsibility of the consignee or the Invoice payer on behalf of the consignee (if provided). Examples of origin charges are customs clearance fees, documentation fees, container packing and loading charges levied at the port of origin to cover the costs of preparing the cargo for shipment. They include the cost of inland transportation to the port, when applicable.

	PRE (Prepaid)
	COL (Collect)

Enum:
Array [ 2 ]
destinationChargesPaymentTermCode	destinationChargesPaymentTermCodestring
example: PRE

An indicator of whether destination charges are prepaid (PRE) or collect (COL). When prepaid, the charges are the responsibility of the shipper or the Invoice payer on behalf of the shipper (if provided). When collect, the charges are the responsibility of the consignee or the Invoice payer on behalf of the consignee (if provided). Examples of destination charges are customs clearance fees, documentation fees, terminal handling fees at the destination port and the costs of inland transportation from the port to the final delivery location, when applicable.

	PRE (Prepaid)
	COL (Collect)

Enum:
Array [ 2 ]
contractQuotationReference	contractQuotationReferencestring
pattern: ^\S+(\s+\S+)*$
maxLength: 35
example: HHL1401

Information provided by the shipper to identify whether pricing for the shipment has been agreed via a contract or a quotation reference. Mandatory if service contract (owner) is not provided.
vessel	{...}
carrierServiceName	carrierServiceNamestring
pattern: ^\S+(\s+\S+)*$
maxLength: 50
example: Great Lion Service

The name of a service as specified by the carrier
carrierServiceCode	carrierServiceCodestring
pattern: ^\S+(\s+\S+)*$
maxLength: 11
example: FE1

The carrier-specific code of the service for which the schedule details are published.
universalServiceReference	universalServiceReferencestring
pattern: ^SR\d{5}[A-Z]$
maxLength: 8
example: SR12345A

A global unique service reference, as per DCSA standard, agreed by VSA partners for the service. The service reference must match the regular expression pattern: SR\d{5}[A-Z]. The letters SR followed by 5 digits, followed by a checksum-character as a capital letter from A to Z.
carrierExportVoyageNumber	carrierExportVoyageNumberstring
pattern: ^\S+(\s+\S+)*$
maxLength: 50
example: 2103S

The identifier of an export voyage. The carrier-specific identifier of the export Voyage.
universalExportVoyageReference	universalExportVoyageReferencestring
pattern: ^\d{2}[0-9A-Z]{2}[NEWSR]$
example: 2103N

A global unique voyage reference for the export Voyage, as per DCSA standard, agreed by VSA partners for the voyage. The voyage reference must match the regular expression pattern: \d{2}[0-9A-Z]{2}[NEWSR]

	2 digits for the year
	2 alphanumeric characters for the sequence number of the voyage
	1 character for the direction/haul (North, East, West, South or Roundtrip).

declaredValue	declaredValuenumber($float)
minimum: 0
example: 1231.1

The value of the cargo that the shipper declares to avoid the carrier's limitation of liability and "Ad Valorem" freight, i.e. freight which is calculated based on the value of the goods declared by the shipper.
declaredValueCurrency	declaredValueCurrencystring
pattern: ^[A-Z]{3}$
maxLength: 3
example: DKK

The currency used for the declared value, using the 3-character code defined by ISO 4217.
isPartialLoadAllowed*	isPartialLoadAllowedboolean
example: true

Indication whether the shipper agrees to load part of the shipment in case where not all of the cargo is delivered within cut-off.
isExportDeclarationRequired*	isExportDeclarationRequiredboolean
example: true

Information provided by the shipper whether an export declaration is required for this particular shipment/commodity/destination.
exportDeclarationReference	exportDeclarationReferencestring
pattern: ^\S+(\s+\S+)*$
maxLength: 35
example: ABC123123

A government document permitting designated goods to be shipped out of the country. Reference number assigned by an issuing authority to an Export License. The export license must be valid at time of departure. Required if Export declaration required is ‘True’.
isImportLicenseRequired*	isImportLicenseRequiredboolean
example: true

Information provided by the shipper whether an import permit or license is required for this particular shipment/commodity/destination.
importLicenseReference	importLicenseReferencestring
pattern: ^\S+(\s+\S+)*$
maxLength: 35
example: ABC123123

A certificate, issued by countries exercising import controls, that permits importation of the articles stated in the license. Reference number assigned by an issuing authority to an Import License. The import license number must be valid at time of arrival. Required if import license required is ‘True’.
expectedDepartureDate	expectedDepartureDatestring($date)
example: 2021-05-17

The date when the shipment is expected to be loaded on board a vessel as provided by the shipper or its agent. If vessel/voyage or expected date of arrival is not provided, this is mandatory
expectedArrivalAtPlaceOfDeliveryStartDate	expectedArrivalAtPlaceOfDeliveryStartDatestring($date)
example: 2021-05-17

The start date (provided as a range together with expectedArrivalAtPlaceOfDeliveryEndDate) for when the shipment is expected to arrive at final destination. If vessel/voyage or expectedDepartureDate is not provided, this is mandatory together with expectedArrivalAtPlaceOfDeliveryEndDate
expectedArrivalAtPlaceOfDeliveryEndDate	expectedArrivalAtPlaceOfDeliveryEndDatestring($date)
example: 2021-05-19

The end date (provided as a range together with expectedArrivalAtPlaceOfDeliveryStartDate) for when the shipment is expected to arrive at final destination. If vessel/voyage or expectedDepartureDate is not provided, this is mandatory together with expectedArrivalAtPlaceOfDeliveryStartDate
transportDocumentTypeCode	transportDocumentTypeCodestring
example: SWB

Specifies the type of the transport document

	BOL (Bill of Lading)
	SWB (Sea Waybill)

Enum:
Array [ 2 ]
transportDocumentReference	string
pattern: ^\S+(\s+\S+)*$
maxLength: 20
example: reserved-HHL123

A unique reference allocated by the shipping line to the Transport Document that the booking concerns.
bookingChannelReference	bookingChannelReferencestring
pattern: ^\S+(\s+\S+)*$
maxLength: 20
example: Inttra reference

Identification number provided by the platform/channel used for booking request/confirmation, ex: Inttra booking reference, or GTNexus, other. Conditional on booking channel being used
incoTerms	incoTermsstring
maxLength: 3
example: FCA

Transport obligations, costs and risks as agreed between buyer and seller as defined by ICC. A list of possible values:

	EXW (Ex-Works)
	FCA (Free Carrier)
	FAS (Free Alongside Ship)
	FOB (Free On Board)
	CFR (Cost and Freight)
	CIF (Cost, Insurance and Freight)
	CPT (Carriage Paid To)
	CIP (Carriage And Insurance Paid To)
	DAP (Delivered At Place)
	DPU (Delivered At Place Unloaded)
	DDP (Delivered Duty Paid)

communicationChannelCode*	communicationChannelCodestring
maxLength: 2
example: AO

Specifying which communication channel is to be used for this booking e.g. Possible values are:

	EI (EDI transmission)
	EM (Email)
	AO (API)

isEquipmentSubstitutionAllowed*	isEquipmentSubstitutionAllowedboolean
example: true

Indicates if an alternate equipment type can be provided by the carrier.
invoicePayableAt	{...}
example: OrderedMap { "locationName": "Eiffel Tower", "locationType": "UNLO", "UNLocationCode": "FRPAR" }
placeOfBLIssue	{...}
example: OrderedMap { "locationName": "DCSA Headquarters", "locationType": "UNLO", "UNLocationCode": "NLAMS" }
references	[...]
documentParties	[...]
partyContactDetails	[...]
shipmentLocations*	[...]
requestedEquipments*	[...]
}
*/
type Booking struct {
	ReceiptTypeAtOrigin                       *string         `json:"receiptTypeAtOrigin,omitempty"`
	DeliveryTypeAtDestination                 *string         `json:"deliveryTypeAtDestination,omitempty"`
	CargoMovementTypeAtOrigin                 *string         `json:"cargoMovementTypeAtOrigin,omitempty"`
	CargoMovementTypeAtDestination            *string         `json:"cargoMovementTypeAtDestination,omitempty"`
	ServiceContractReference                  *string         `json:"serviceContractReference,omitempty"`
	FreightPaymentTermCode                    *string         `json:"freightPaymentTermCode,omitempty"`
	OriginChargesPaymentTermCode              *string         `json:"originChargesPaymentTermCode,omitempty"`
	DestinationChargesPaymentTermCode         *string         `json:"destinationChargesPaymentTermCode,omitempty"`
	ContractQuotationReference                *string         `json:"contractQuotationReference,omitempty"`
	Vessel                                    *Vessel         `json:"vessel,omitempty"`
	CarrierServiceName                        *string         `json:"carrierServiceName,omitempty"`
	CarrierServiceCode                        *string         `json:"carrierServiceCode,omitempty"`
	UniversalServiceReference                 *string         `json:"universalServiceReference,omitempty"`
	CarrierExportVoyageNumber                 *string         `json:"carrierExportVoyageNumber,omitempty"`
	UniversalExportVoyageReference            *string         `json:"universalExportVoyageReference,omitempty"`
	DeclaredValue                             *float64        `json:"declaredValue,omitempty"`
	DeclaredValueCurrency                     *string         `json:"declaredValueCurrency,omitempty"`
	IsPartialLoadAllowed                      *bool           `json:"isPartialLoadAllowed,omitempty"`
	IsExportDeclarationRequired               *bool           `json:"isExportDeclarationRequired,omitempty"`
	ExportDeclarationReference                *string         `json:"exportDeclarationReference,omitempty"`
	IsImportLicenseRequired                   *bool           `json:"isImportLicenseRequired,omitempty"`
	ImportLicenseReference                    *string         `json:"importLicenseReference,omitempty"`
	ExpectedDepartureDate                     *string         `json:"expectedDepartureDate,omitempty"`
	ExpectedArrivalAtPlaceOfDeliveryStartDate *string         `json:"expectedArrivalAtPlaceOfDeliveryStartDate,omitempty"`
	ExpectedArrivalAtPlaceOfDeliveryEndDate   *string         `json:"expectedArrivalAtPlaceOfDeliveryEndDate,omitempty"`
	TransportDocumentTypeCode                 *string         `json:"transportDocumentTypeCode,omitempty"`
	TransportDocumentReference                *string         `json:"transportDocumentReference,omitempty"`
	BookingChannelReference                   *string         `json:"bookingChannelReference,omitempty"`
	IncoTerms                                 *string         `json:"incoTerms,omitempty"`
	CommunicationChannelCode                  *string         `json:"communicationChannelCode,omitempty"`
	IsEquipmentSubstitutionAllowed            *bool           `json:"isEquipmentSubstitutionAllowed,omitempty"`
	InvoicePayableAt                          *Location       `json:"invoicePayableAt,omitempty"`
	PlaceOfBLIssue                            *Location       `json:"placeOfBLIssue,omitempty"`
	References                                []Reference     `json:"references,omitempty"`
	DocumentParties                           []DocumentParty `json:"documentParties,omitempty"`
	PartyContactDetails                       []ContactDetail `json:"partyContactDetails,omitempty"`
	ShipmentLocations                         []Location      `json:"shipmentLocations,omitempty"`
	//RequestedEquipments                       []Equipment     `json:"requestedEquipments,omitempty"`
}
