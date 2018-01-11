package ont

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	ONT_RPC_GETVERSION          = "getversion"
	ONT_RPC_GETTRANSACTION      = "getrawtransaction"
	ONT_RPC_SENDTRANSACTION     = "sendrawtransaction"
	ONT_RPC_GETBLOCK            = "getblock"
	ONT_RPC_GETBLOCKCOUNT       = "getblockcount"
	ONT_RPC_GETBLOCKHASH        = "getblockhash"
	ONT_RPC_GETUNSPENDOUTPUT    = "getunspendoutput"
	ONT_RPC_GETCURRENTBLOCKHASH = "getbestblockhash"
	ONT_RPC_GETIDENTITY         = "getidentity"
	ONT_RPC_GETIDENTITYCLAIM    = "getidentityclaim"
)

const (
	ONT_API_GETCONNCOUNT     = "/api/v1/node/connectioncount"
	ONT_API_GETBLOCKBYHEIGHT = "/api/v1/block/details/height"
	ONT_API_GETBLOCKBYHASH   = "/api/v1/block/details/hash"
	ONT_API_GETBLOCKCOUNT    = "/api/v1/block/height"
	ONT_API_GETTRANSACTION   = "/api/v1/transaction"
	ONT_API_GETASSET         = "/api/v1/asset"
	ONT_API_SENDTRANSACTION  = "/api/v1/transaction"

	//Api_Getconnectioncount = "/api/v1/node/connectioncount"
	//Api_Getblockbyheight   = "/api/v1/block/details/height/:height"
	//Api_Getblockbyhash     = "/api/v1/block/details/hash/:hash"
	//Api_Getblockheight     = "/api/v1/block/height"
	//Api_Gettransaction     = "/api/v1/transaction/:hash"
	//Api_Getasset           = "/api/v1/asset/:hash"
	//Api_Restart            = "/api/v1/restart"
	//Api_SendRawTransaction = "/api/v1/transaction"
	//Api_OauthServerAddr    = "/api/v1/config/oauthserver/addr"
	//Api_NoticeServerAddr   = "/api/v1/config/noticeserver/addr"
	//Api_NoticeServerState  = "/api/v1/config/noticeserver/state"

	ONT_SENDTRANSACTION     = "sendrawtransaction"
	ONT_HEARTBEAT           = "heartbeat"
	ONT_SMARTCONTRACTINVOKE = "InvokeTransaction"
	ONT_NOTIFY              = "Notify"
)

type ONTJsonRpcRes struct {
	Id      string          `json:"id"`
	JsonRpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
}

const (
	DNA_ERR_OK                     = 0
	DNA_ERR_SESSION_EXPIRED        = 41001
	DNA_ERR_SERVICE_CEILING        = 41002
	DNA_ERR_ILLEGAL_DATAFORMAT     = 41003
	DNA_ERR_OAUTH_TIMEOUT          = 41004
	DNA_ERR_INVALID_METHOD         = 42001
	DNA_ERR_INVALID_PARAMS         = 42002
	DNA_ERR_INVALID_TOKEN          = 42003
	DNA_ERR_INVALID_TRANSACTION    = 43001
	DNA_ERR_INVALID_ASSET          = 43002
	DNA_ERR_INVALID_BLOCK          = 43003
	DNA_ERR_UNKNOWN_TRANSACTION    = 44001
	DNA_ERR_UNKNOWN_ASSET          = 44002
	DNA_ERR_UNKNOWN_BLOCK          = 44003
	DNA_ERR_INVALID_VERSION        = 45001
	DNA_ERR_INTERNAL_ERROR         = 45002
	DNA_ERR_OAUTH_INVALID_APPID    = 46001
	DNA_ERR_OAUTH_INVALID_CHECKVAL = 46002
	DNA_ERR_SMARTCODE_ERROR        = 47001

	ErrNoCode               = -2
	ErrNoError              = 0
	ErrUnknown              = -1
	ErrDuplicatedTx         = 1
	ErrDuplicateInput       = 45003
	ErrAssetPrecision       = 45004
	ErrTransactionBalance   = 45005
	ErrAttributeProgram     = 45006
	ErrTransactionContracts = 45007
	ErrTransactionPayload   = 45008
	ErrDoubleSpend          = 45009
	ErrTxHashDuplicate      = 45010
	ErrStateUpdaterVaild    = 45011
	ErrSummaryAsset         = 45012
	ErrXmitFail             = 45013
)

const (
	OntRpcInvalidHash        = "invalid hash"
	OntRpcInvalidBlock       = "invalid block"
	OntRpcInvalidTransaction = "invalid transaction"
	OntRpcInvalidParameter   = "invalid parameter"
	OntRpcUnknownBlock       = "unknown block"
	OntRpcUnknownTransaction = "unknown transaction"
	OntRpcNil                = "null"
	OntRpcUnsupported        = "Unsupported"
	OntRpcInternalError      = "internal error"
	OntRpcIOError            = "internal IO error"
	OntRpcAPIError           = "internal API error"
)

var ONTRpcError map[string]string = map[string]string{
	OntRpcInvalidHash:        "",
	OntRpcInvalidBlock:       "",
	OntRpcInvalidTransaction: "",
	OntRpcInvalidParameter:   "",
	OntRpcUnknownBlock:       "",
	OntRpcUnknownTransaction: "",
	OntRpcUnsupported:        "",
	OntRpcInternalError:      "",
	OntRpcNil:                "",
	OntRpcIOError:            "",
	OntRpcAPIError:           "",
}

func (this *ONTJsonRpcRes) HandleResult() ([]byte, error) {
	res := strings.Trim(string(this.Result), "\"")
	_, ok := ONTRpcError[res]
	if ok {
		return nil, fmt.Errorf(res)
	}
	return []byte(res), nil
}

