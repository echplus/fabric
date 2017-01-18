package api

import (
	"encoding/json"
	"strconv"
	"time"

	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/hyperledger/fabric/sdk/Go/protos"
	"github.com/hyperledger/fabric/sdk/Go/utils"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const ttl = time.Second * 3
const errInfo = "error"

func (mbsrvc *memberImpl) gethost() (string, error) {
	var strNets []string
	strNets = append(strNets, viper.GetString("grpc.net1"), viper.GetString("grpc.net2"), viper.GetString("grpc.net3"), viper.GetString("grpc.net4"))
	var strConn string
	for _, r := range strNets {
		conn := &utils.Clienter{}
		if conn.Connect(r) == true {
			strConn = r
			break
		}
	}
	return strConn, nil
}

func (mbsrvc *memberImpl) Network() (string, error) {
	strConn, err := mbsrvc.gethost()
	if err != nil {
		logger.Error()
		return errInfo, err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithTimeout(ttl))

	conn, err := grpc.Dial(strConn, opts...)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	defer conn.Close()

	openChain := pb.NewOpenchainClient(conn)
	peers, err := openChain.GetPeers(context.Background(), &google_protobuf1.Empty{})
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}

	resp, err := json.Marshal(peers)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}

	return string(resp), err

}

func (mbsrvc *memberImpl) GetBlocks(id string) (string, error) {
	strConn, err := mbsrvc.gethost()
	if err != nil {
		logger.Error()
		return "", err
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithTimeout(ttl))

	conn, err := grpc.Dial(strConn, opts...)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	defer conn.Close()

	openChain := pb.NewOpenchainClient(conn)

	uIntID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	_block, err := openChain.GetBlockByNumber(context.Background(), &pb.BlockNumber{uIntID})
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}

	resp, err := json.Marshal(_block)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}

	return string(resp), err
}

func (mbsrvc *memberImpl) GetChain() (string, error) {
	strConn, err := mbsrvc.gethost()
	if err != nil {
		logger.Error()
		return errInfo, err
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithTimeout(ttl))
	conn, err := grpc.Dial(strConn, opts...)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	defer conn.Close()

	openChain := pb.NewOpenchainClient(conn)
	blockchainInfo, err := openChain.GetBlockchainInfo(context.Background(), &google_protobuf1.Empty{})
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	resp, err := json.Marshal(blockchainInfo)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	return string(resp), err
}

func (mbsrvc *memberImpl) GetTransactions(txid string) (string, error) {
	strConn, err := mbsrvc.gethost()
	if err != nil {
		logger.Error()
		return errInfo, err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithTimeout(ttl))
	conn, err := grpc.Dial(strConn, opts...)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	defer conn.Close()

	openChain := pb.NewOpenchainClient(conn)
	transaction, err := openChain.GetTransactionStrByID(context.Background(), &pb.Transaction{Txid: txid})
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	resp, err := json.Marshal(transaction)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	return string(resp), err
}

func (mbsrvc *memberImpl) GetBlocksCount() (string, error) {

	strConn, err := mbsrvc.gethost()
	if err != nil {
		logger.Error()
		return errInfo, err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithTimeout(ttl))
	conn, err := grpc.Dial(strConn, opts...)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	defer conn.Close()

	openChain := pb.NewOpenchainClient(conn)
	blockCount, err := openChain.GetBlockCount(context.Background(), &google_protobuf1.Empty{})
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}

	resp, err := json.Marshal(blockCount)
	if err != nil {
		logger.Error(err)
		return errInfo, err
	}
	return string(resp), err
}
