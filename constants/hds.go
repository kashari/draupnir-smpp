package constants

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ESME_ROK-0]
	_ = x[ESME_RINVMSGLEN-1]
	_ = x[ESME_RINVCMDLEN-2]
	_ = x[ESME_RINVCMDID-3]
	_ = x[ESME_RINVBNDSTS-4]
	_ = x[ESME_RALYBND-5]
	_ = x[ESME_RINVPRTFLG-6]
	_ = x[ESME_RINVREGDLVFLG-7]
	_ = x[ESME_RSYSERR-8]
	_ = x[ESME_RINVSRCADR-10]
	_ = x[ESME_RINVDSTADR-11]
	_ = x[ESME_RINVMSGID-12]
	_ = x[ESME_RBINDFAIL-13]
	_ = x[ESME_RINVPASWD-14]
	_ = x[ESME_RINVSYSID-15]
	_ = x[ESME_RCANCELFAIL-17]
	_ = x[ESME_RREPLACEFAIL-19]
	_ = x[ESME_RMSGQFUL-20]
	_ = x[ESME_RINVSERTYP-21]
	_ = x[ESME_RADDCUSTFAIL-25]
	_ = x[ESME_RDELCUSTFAIL-26]
	_ = x[ESME_RMODCUSTFAIL-27]
	_ = x[ESME_RENQCUSTFAIL-28]
	_ = x[ESME_RINVCUSTID-29]
	_ = x[ESME_RINVCUSTNAME-31]
	_ = x[ESME_RINVCUSTADR-33]
	_ = x[ESME_RINVADR-34]
	_ = x[ESME_RCUSTEXIST-35]
	_ = x[ESME_RCUSTNOTEXIST-36]
	_ = x[ESME_RADDDLFAIL-38]
	_ = x[ESME_RMODDLFAIL-39]
	_ = x[ESME_RDELDLFAIL-40]
	_ = x[ESME_RVIEWDLFAIL-41]
	_ = x[ESME_RLISTDLSFAIL-48]
	_ = x[ESME_RPARAMRETFAIL-49]
	_ = x[ESME_RINVPARAM-50]
	_ = x[ESME_RINVNUMDESTS-51]
	_ = x[ESME_RINVDLNAME-52]
	_ = x[ESME_RINVDLMEMBDESC-53]
	_ = x[ESME_RINVDLMEMBTYP-56]
	_ = x[ESME_RINVDLMODOPT-57]
	_ = x[ESME_RINVDESTFLAG-64]
	_ = x[ESME_RINVSUBREP-66]
	_ = x[ESME_RINVESMCLASS-67]
	_ = x[ESME_RCNTSUBDL-68]
	_ = x[ESME_RSUBMITFAIL-69]
	_ = x[ESME_RINVSRCTON-72]
	_ = x[ESME_RINVSRCNPI-73]
	_ = x[ESME_RINVDSTTON-80]
	_ = x[ESME_RINVDSTNPI-81]
	_ = x[ESME_RINVSYSTYP-83]
	_ = x[ESME_RINVREPFLAG-84]
	_ = x[ESME_RINVNUMMSGS-85]
	_ = x[ESME_RTHROTTLED-88]
	_ = x[ESME_RPROVNOTALLWD-89]
	_ = x[ESME_RINVSCHED-97]
	_ = x[ESME_RINVEXPIRY-98]
	_ = x[ESME_RINVDFTMSGID-99]
	_ = x[ESME_RX_T_APPN-100]
	_ = x[ESME_RX_P_APPN-101]
	_ = x[ESME_RX_R_APPN-102]
	_ = x[ESME_RQUERYFAIL-103]
	_ = x[ESME_RINVPGCUSTID-128]
	_ = x[ESME_RINVPGCUSTIDLEN-129]
	_ = x[ESME_RINVCITYLEN-130]
	_ = x[ESME_RINVSTATELEN-131]
	_ = x[ESME_RINVZIPPREFIXLEN-132]
	_ = x[ESME_RINVZIPPOSTFIXLEN-133]
	_ = x[ESME_RINVMINLEN-134]
	_ = x[ESME_RINVMIN-135]
	_ = x[ESME_RINVPINLEN-136]
	_ = x[ESME_RINVTERMCODELEN-137]
	_ = x[ESME_RINVCHANNELLEN-138]
	_ = x[ESME_RINVCOVREGIONLEN-139]
	_ = x[ESME_RINVCAPCODELEN-140]
	_ = x[ESME_RINVMDTLEN-141]
	_ = x[ESME_RINVPRIORMSGLEN-142]
	_ = x[ESME_RINVPERMSGLEN-143]
	_ = x[ESME_RINVPGALERTLEN-144]
	_ = x[ESME_RINVSMUSERLEN-145]
	_ = x[ESME_RINVRTDBLEN-146]
	_ = x[ESME_RINVREGDELLEN-147]
	_ = x[ESME_RINVMSGDISTLEN-148]
	_ = x[ESME_RINVPRIORMSG-149]
	_ = x[ESME_RINVMDT-150]
	_ = x[ESME_RINVPERMSG-151]
	_ = x[ESME_RINVMSGDIST-152]
	_ = x[ESME_RINVPGALERT-153]
	_ = x[ESME_RINVSMUSER-154]
	_ = x[ESME_RINVRTDB-155]
	_ = x[ESME_RINVREGDEL-156]
	_ = x[ESME_RINVOPTPARLEN-159]
	_ = x[ESME_RINVOPTPARSTREAM-192]
	_ = x[ESME_ROPTPARNOTALLWD-193]
	_ = x[ESME_RINVPARLEN-194]
	_ = x[ESME_RMISSINGOPTPARAM-195]
	_ = x[ESME_RINVOPTPARAMVAL-196]
	_ = x[ESME_RDELIVERYFAILURE-254]
	_ = x[ESME_RUNKNOWNERR-255]
	_ = x[ESME_LAST_ERROR-300]
}

const _CommandStatusType_name = "ESME_ROKESME_RINVMSGLENESME_RINVCMDLENESME_RINVCMDIDESME_RINVBNDSTSESME_RALYBNDESME_RINVPRTFLGESME_RINVREGDLVFLGESME_RSYSERRESME_RINVSRCADRESME_RINVDSTADRESME_RINVMSGIDESME_RBINDFAILESME_RINVPASWDESME_RINVSYSIDESME_RCANCELFAILESME_RREPLACEFAILESME_RMSGQFULESME_RINVSERTYPESME_RADDCUSTFAILESME_RDELCUSTFAILESME_RMODCUSTFAILESME_RENQCUSTFAILESME_RINVCUSTIDESME_RINVCUSTNAMEESME_RINVCUSTADRESME_RINVADRESME_RCUSTEXISTESME_RCUSTNOTEXISTESME_RADDDLFAILESME_RMODDLFAILESME_RDELDLFAILESME_RVIEWDLFAILESME_RLISTDLSFAILESME_RPARAMRETFAILESME_RINVPARAMESME_RINVNUMDESTSESME_RINVDLNAMEESME_RINVDLMEMBDESCESME_RINVDLMEMBTYPESME_RINVDLMODOPTESME_RINVDESTFLAGESME_RINVSUBREPESME_RINVESMCLASSESME_RCNTSUBDLESME_RSUBMITFAILESME_RINVSRCTONESME_RINVSRCNPIESME_RINVDSTTONESME_RINVDSTNPIESME_RINVSYSTYPESME_RINVREPFLAGESME_RINVNUMMSGSESME_RTHROTTLEDESME_RPROVNOTALLWDESME_RINVSCHEDESME_RINVEXPIRYESME_RINVDFTMSGIDESME_RX_T_APPNESME_RX_P_APPNESME_RX_R_APPNESME_RQUERYFAILESME_RINVPGCUSTIDESME_RINVPGCUSTIDLENESME_RINVCITYLENESME_RINVSTATELENESME_RINVZIPPREFIXLENESME_RINVZIPPOSTFIXLENESME_RINVMINLENESME_RINVMINESME_RINVPINLENESME_RINVTERMCODELENESME_RINVCHANNELLENESME_RINVCOVREGIONLENESME_RINVCAPCODELENESME_RINVMDTLENESME_RINVPRIORMSGLENESME_RINVPERMSGLENESME_RINVPGALERTLENESME_RINVSMUSERLENESME_RINVRTDBLENESME_RINVREGDELLENESME_RINVMSGDISTLENESME_RINVPRIORMSGESME_RINVMDTESME_RINVPERMSGESME_RINVMSGDISTESME_RINVPGALERTESME_RINVSMUSERESME_RINVRTDBESME_RINVREGDELESME_RINVOPTPARLENESME_RINVOPTPARSTREAMESME_ROPTPARNOTALLWDESME_RINVPARLENESME_RMISSINGOPTPARAMESME_RINVOPTPARAMVALESME_RDELIVERYFAILUREESME_RUNKNOWNERRESME_LAST_ERROR"

var _CommandStatusType_map = map[CommandStatusType]string{
	0:   _CommandStatusType_name[0:8],
	1:   _CommandStatusType_name[8:23],
	2:   _CommandStatusType_name[23:38],
	3:   _CommandStatusType_name[38:52],
	4:   _CommandStatusType_name[52:67],
	5:   _CommandStatusType_name[67:79],
	6:   _CommandStatusType_name[79:94],
	7:   _CommandStatusType_name[94:112],
	8:   _CommandStatusType_name[112:124],
	10:  _CommandStatusType_name[124:139],
	11:  _CommandStatusType_name[139:154],
	12:  _CommandStatusType_name[154:168],
	13:  _CommandStatusType_name[168:182],
	14:  _CommandStatusType_name[182:196],
	15:  _CommandStatusType_name[196:210],
	17:  _CommandStatusType_name[210:226],
	19:  _CommandStatusType_name[226:243],
	20:  _CommandStatusType_name[243:256],
	21:  _CommandStatusType_name[256:271],
	25:  _CommandStatusType_name[271:288],
	26:  _CommandStatusType_name[288:305],
	27:  _CommandStatusType_name[305:322],
	28:  _CommandStatusType_name[322:339],
	29:  _CommandStatusType_name[339:354],
	31:  _CommandStatusType_name[354:371],
	33:  _CommandStatusType_name[371:387],
	34:  _CommandStatusType_name[387:399],
	35:  _CommandStatusType_name[399:414],
	36:  _CommandStatusType_name[414:432],
	38:  _CommandStatusType_name[432:447],
	39:  _CommandStatusType_name[447:462],
	40:  _CommandStatusType_name[462:477],
	41:  _CommandStatusType_name[477:493],
	48:  _CommandStatusType_name[493:510],
	49:  _CommandStatusType_name[510:528],
	50:  _CommandStatusType_name[528:542],
	51:  _CommandStatusType_name[542:559],
	52:  _CommandStatusType_name[559:574],
	53:  _CommandStatusType_name[574:593],
	56:  _CommandStatusType_name[593:611],
	57:  _CommandStatusType_name[611:628],
	64:  _CommandStatusType_name[628:645],
	66:  _CommandStatusType_name[645:660],
	67:  _CommandStatusType_name[660:677],
	68:  _CommandStatusType_name[677:691],
	69:  _CommandStatusType_name[691:707],
	72:  _CommandStatusType_name[707:722],
	73:  _CommandStatusType_name[722:737],
	80:  _CommandStatusType_name[737:752],
	81:  _CommandStatusType_name[752:767],
	83:  _CommandStatusType_name[767:782],
	84:  _CommandStatusType_name[782:798],
	85:  _CommandStatusType_name[798:814],
	88:  _CommandStatusType_name[814:829],
	89:  _CommandStatusType_name[829:847],
	97:  _CommandStatusType_name[847:861],
	98:  _CommandStatusType_name[861:876],
	99:  _CommandStatusType_name[876:893],
	100: _CommandStatusType_name[893:907],
	101: _CommandStatusType_name[907:921],
	102: _CommandStatusType_name[921:935],
	103: _CommandStatusType_name[935:950],
	128: _CommandStatusType_name[950:967],
	129: _CommandStatusType_name[967:987],
	130: _CommandStatusType_name[987:1003],
	131: _CommandStatusType_name[1003:1020],
	132: _CommandStatusType_name[1020:1041],
	133: _CommandStatusType_name[1041:1063],
	134: _CommandStatusType_name[1063:1078],
	135: _CommandStatusType_name[1078:1090],
	136: _CommandStatusType_name[1090:1105],
	137: _CommandStatusType_name[1105:1125],
	138: _CommandStatusType_name[1125:1144],
	139: _CommandStatusType_name[1144:1165],
	140: _CommandStatusType_name[1165:1184],
	141: _CommandStatusType_name[1184:1199],
	142: _CommandStatusType_name[1199:1219],
	143: _CommandStatusType_name[1219:1237],
	144: _CommandStatusType_name[1237:1256],
	145: _CommandStatusType_name[1256:1274],
	146: _CommandStatusType_name[1274:1290],
	147: _CommandStatusType_name[1290:1308],
	148: _CommandStatusType_name[1308:1327],
	149: _CommandStatusType_name[1327:1344],
	150: _CommandStatusType_name[1344:1356],
	151: _CommandStatusType_name[1356:1371],
	152: _CommandStatusType_name[1371:1387],
	153: _CommandStatusType_name[1387:1403],
	154: _CommandStatusType_name[1403:1418],
	155: _CommandStatusType_name[1418:1431],
	156: _CommandStatusType_name[1431:1446],
	159: _CommandStatusType_name[1446:1464],
	192: _CommandStatusType_name[1464:1485],
	193: _CommandStatusType_name[1485:1505],
	194: _CommandStatusType_name[1505:1520],
	195: _CommandStatusType_name[1520:1541],
	196: _CommandStatusType_name[1541:1561],
	254: _CommandStatusType_name[1561:1582],
	255: _CommandStatusType_name[1582:1598],
	300: _CommandStatusType_name[1598:1613],
}

func (i CommandStatusType) String() string {
	if str, ok := _CommandStatusType_map[i]; ok {
		return str
	}
	return "CommandStatusType(" + strconv.FormatInt(int64(i), 10) + ")"
}

func (i CommandStatusType) Desc() string {
	switch i {
	case ESME_ROK:
		return "No Error"
	case ESME_RINVMSGLEN:
		return "Message Length is invalid"
	case ESME_RINVCMDLEN:
		return "Command Length is invalid"
	case ESME_RINVCMDID:
		return "Invalid Command ID"
	case ESME_RINVBNDSTS:
		return "Incorrect BIND Status for given command"
	case ESME_RALYBND:
		return "ESME Already in Bound State"
	case ESME_RINVPRTFLG:
		return "Invalid Priority Flag"
	case ESME_RINVREGDLVFLG:
		return "Invalid Registered Delivery Flag"
	case ESME_RSYSERR:
		return "System Error"
	case ESME_RINVSRCADR:
		return "Invalid Source Address"
	case ESME_RINVDSTADR:
		return "Invalid Dest Addr"
	case ESME_RINVMSGID:
		return "Message ID is invalid"
	case ESME_RBINDFAIL:
		return "Bind Failed"
	case ESME_RINVPASWD:
		return "Invalid Password"
	case ESME_RINVSYSID:
		return "Invalid System ID"
	case ESME_RCANCELFAIL:
		return "Cancel SM Failed"
	case ESME_RREPLACEFAIL:
		return "Replace SM Failed"
	case ESME_RMSGQFUL:
		return "Message Queue Full"
	case ESME_RINVSERTYP:
		return "Invalid Service Type"
	case ESME_RADDCUSTFAIL:
		return "Failed to Add Customer"
	case ESME_RDELCUSTFAIL:
		return "Failed to delete Customer"
	case ESME_RMODCUSTFAIL:
		return "Failed to modify customer"
	case ESME_RENQCUSTFAIL:
		return "Failed to Enquire Customer"
	case ESME_RINVCUSTID:
		return "Invalid Customer ID"
	case ESME_RINVCUSTNAME:
		return "Invalid Customer Name"
	case ESME_RINVCUSTADR:
		return "Invalid Customer Address"
	case ESME_RINVADR:
		return "Invalid Address"
	case ESME_RCUSTEXIST:
		return "Customer Exists"
	case ESME_RCUSTNOTEXIST:
		return "Customer does not exist"
	case ESME_RADDDLFAIL:
		return "Failed to Add DL"
	case ESME_RMODDLFAIL:
		return "Failed to modify DL"
	case ESME_RDELDLFAIL:
		return "Failed to Delete DL"
	case ESME_RVIEWDLFAIL:
		return "Failed to View DL"
	case ESME_RLISTDLSFAIL:
		return "Failed to list DLs"
	case ESME_RPARAMRETFAIL:
		return "Param Retrieve Failed"
	case ESME_RINVPARAM:
		return "Invalid Param"
	case ESME_RINVNUMDESTS:
		return "Invalid number of destinations"
	case ESME_RINVDLNAME:
		return "Invalid Distribution List name"
	case ESME_RINVDLMEMBDESC:
		return "Invalid DL Member Description"
	case ESME_RINVDLMEMBTYP:
		return "Invalid DL Member Type"
	case ESME_RINVDLMODOPT:
		return "Invalid DL Modify Option"
	case ESME_RINVDESTFLAG:
		return "Destination flag is invalid (submit_multi)"
	case ESME_RINVSUBREP:
		return "Invalid ‘submit with replace’ request (i.e. submit_sm with replace_if_present_flag set)"
	case ESME_RINVESMCLASS:
		return "Invalid esm_class field data"
	case ESME_RCNTSUBDL:
		return "Cannot Submit to Distribution List"
	case ESME_RSUBMITFAIL:
		return "submit_sm or submit_multi failed"
	case ESME_RINVSRCTON:
		return "Invalid Source address TON"
	case ESME_RINVSRCNPI:
		return "Invalid Source address NPI"
	case ESME_RINVDSTTON:
		return "Invalid Destination address TON"
	case ESME_RINVDSTNPI:
		return "Invalid Destination address NPI"
	case ESME_RINVSYSTYP:
		return "Invalid system_type field"
	case ESME_RINVREPFLAG:
		return "Invalid replace_if_present flag"
	case ESME_RINVNUMMSGS:
		return "Invalid number of messages"
	case ESME_RTHROTTLED:
		return "Throttling error (ESME has exceeded allowed message limits)"
	case ESME_RPROVNOTALLWD:
		return "Provisioning Not Allowed"
	case ESME_RINVSCHED:
		return "Invalid Scheduled Delivery Time"
	case ESME_RINVEXPIRY:
		return "Invalid message validity period (Expiry time)"
	case ESME_RINVDFTMSGID:
		return "Predefined Message Invalid or Not Found"
	case ESME_RX_T_APPN:
		return "ESME Receiver Temporary App Error Code"
	case ESME_RX_P_APPN:
		return "ESME Receiver Permanent App Error Code"
	case ESME_RX_R_APPN:
		return "ESME Receiver Reject Message Error Code"
	case ESME_RQUERYFAIL:
		return "query_sm request failed"
	case ESME_RINVPGCUSTID:
		return "Paging Customer ID Invalid No such subscriber"
	case ESME_RINVPGCUSTIDLEN:
		return "Paging Customer ID length Invalid"
	case ESME_RINVCITYLEN:
		return "City Length Invalid"
	case ESME_RINVSTATELEN:
		return "State Length Invalid"
	case ESME_RINVZIPPREFIXLEN:
		return "Zip Prefix Length Invalid"
	case ESME_RINVZIPPOSTFIXLEN:
		return "Zip Postfix Length Invalid"
	case ESME_RINVMINLEN:
		return "MIN Length Invalid"
	case ESME_RINVMIN:
		return "MIN Invalid (i.e. No such MIN)"
	case ESME_RINVPINLEN:
		return "PIN Length Invalid"
	case ESME_RINVTERMCODELEN:
		return "Terminal Code Length Invalid"
	case ESME_RINVCHANNELLEN:
		return "Channel Length Invalid"
	case ESME_RINVCOVREGIONLEN:
		return "Coverage Region Length Invalid"
	case ESME_RINVCAPCODELEN:
		return "Cap Code Length Invalid"
	case ESME_RINVMDTLEN:
		return "Message delivery time Length Invalid"
	case ESME_RINVPRIORMSGLEN:
		return "Priority Message Length Invalid"
	case ESME_RINVPERMSGLEN:
		return "Periodic Messages Length Invalid"
	case ESME_RINVPGALERTLEN:
		return "Paging Alerts Length Invalid"
	case ESME_RINVSMUSERLEN:
		return "int16 Message User Group Length Invalid"
	case ESME_RINVRTDBLEN:
		return "Real Time Data broadcasts Length Invalid"
	case ESME_RINVREGDELLEN:
		return "Registered Delivery Length Invalid"
	case ESME_RINVMSGDISTLEN:
		return "Message Distribution Length Invalid"
	case ESME_RINVPRIORMSG:
		return "Priority Message Length Invalid"
	case ESME_RINVMDT:
		return "Message delivery time Invalid"
	case ESME_RINVPERMSG:
		return "Periodic Messages Invalid"
	case ESME_RINVMSGDIST:
		return "Message Distribution Invalid"
	case ESME_RINVPGALERT:
		return "Paging Alerts Invalid"
	case ESME_RINVSMUSER:
		return "int16 Message User Group Invalid"
	case ESME_RINVRTDB:
		return "Real Time Data broadcasts Invalid"
	case ESME_RINVREGDEL:
		return "Registered Delivery Invalid"
	case ESME_RINVOPTPARLEN:
		return "Invalid Optional Parameter Length"
	case ESME_RINVOPTPARSTREAM:
		return "Error in the optional part of the PDU Body."
	case ESME_ROPTPARNOTALLWD:
		return "Optional Parameter not allowed"
	case ESME_RINVPARLEN:
		return "Invalid Parameter Length."
	case ESME_RMISSINGOPTPARAM:
		return "Expected Optional Parameter missing"
	case ESME_RINVOPTPARAMVAL:
		return "Invalid Optional Parameter Value"
	case ESME_RDELIVERYFAILURE:
		return "Delivery Failure (used for data_sm_resp)"
	case ESME_RUNKNOWNERR:
		return "Unknown Error"
	case ESME_LAST_ERROR:
		return "The value of the last error code"
	}
	return i.String()
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GENERIC_NACK - -2147483648]
	_ = x[BIND_RECEIVER-1]
	_ = x[BIND_RECEIVER_RESP - -2147483647]
	_ = x[BIND_TRANSMITTER-2]
	_ = x[BIND_TRANSMITTER_RESP - -2147483646]
	_ = x[QUERY_SM-3]
	_ = x[QUERY_SM_RESP - -2147483645]
	_ = x[SUBMIT_SM-4]
	_ = x[SUBMIT_SM_RESP - -2147483644]
	_ = x[DELIVER_SM-5]
	_ = x[DELIVER_SM_RESP - -2147483643]
	_ = x[UNBIND-6]
	_ = x[UNBIND_RESP - -2147483642]
	_ = x[REPLACE_SM-7]
	_ = x[REPLACE_SM_RESP - -2147483641]
	_ = x[CANCEL_SM-8]
	_ = x[CANCEL_SM_RESP - -2147483640]
	_ = x[BIND_TRANSCEIVER-9]
	_ = x[BIND_TRANSCEIVER_RESP - -2147483639]
	_ = x[OUTBIND-11]
	_ = x[ENQUIRE_LINK-21]
	_ = x[ENQUIRE_LINK_RESP - -2147483627]
	_ = x[SUBMIT_MULTI-33]
	_ = x[SUBMIT_MULTI_RESP - -2147483615]
	_ = x[ALERT_NOTIFICATION-258]
	_ = x[DATA_SM-259]
	_ = x[DATA_SM_RESP - -2147483389]
}

const (
	_CommandIDType_name_0 = "GENERIC_NACKBIND_RECEIVER_RESPBIND_TRANSMITTER_RESPQUERY_SM_RESPSUBMIT_SM_RESPDELIVER_SM_RESPUNBIND_RESPREPLACE_SM_RESPCANCEL_SM_RESPBIND_TRANSCEIVER_RESP"
	_CommandIDType_name_1 = "ENQUIRE_LINK_RESP"
	_CommandIDType_name_2 = "SUBMIT_MULTI_RESP"
	_CommandIDType_name_3 = "DATA_SM_RESP"
	_CommandIDType_name_4 = "BIND_RECEIVERBIND_TRANSMITTERQUERY_SMSUBMIT_SMDELIVER_SMUNBINDREPLACE_SMCANCEL_SMBIND_TRANSCEIVER"
	_CommandIDType_name_5 = "OUTBIND"
	_CommandIDType_name_6 = "ENQUIRE_LINK"
	_CommandIDType_name_7 = "SUBMIT_MULTI"
	_CommandIDType_name_8 = "ALERT_NOTIFICATIONDATA_SM"
)

var (
	_CommandIDType_index_0 = [...]uint8{0, 12, 30, 51, 64, 78, 93, 104, 119, 133, 154}
	_CommandIDType_index_4 = [...]uint8{0, 13, 29, 37, 46, 56, 62, 72, 81, 97}
	_CommandIDType_index_8 = [...]uint8{0, 18, 25}
)

func (i CommandIDType) String() string {
	switch {
	//go-staticcheck:ignore SA4003 redundant comparison for int32 range
	case -2147483648 <= i && i <= -2147483639:
		i -= -2147483648
		return _CommandIDType_name_0[_CommandIDType_index_0[i]:_CommandIDType_index_0[i+1]]
	case i == -2147483627:
		return _CommandIDType_name_1
	case i == -2147483615:
		return _CommandIDType_name_2
	case i == -2147483389:
		return _CommandIDType_name_3
	case 1 <= i && i <= 9:
		i -= 1
		return _CommandIDType_name_4[_CommandIDType_index_4[i]:_CommandIDType_index_4[i+1]]
	case i == 11:
		return _CommandIDType_name_5
	case i == 21:
		return _CommandIDType_name_6
	case i == 33:
		return _CommandIDType_name_7
	case 258 <= i && i <= 259:
		i -= 258
		return _CommandIDType_name_8[_CommandIDType_index_8[i]:_CommandIDType_index_8[i+1]]
	default:
		return "CommandIDType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
