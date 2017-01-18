package com.echplus.blockchain;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Random;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.hyperledger.fabric.sdk.Chain;
import org.hyperledger.fabric.sdk.ChainCodeResponse;
import org.hyperledger.fabric.sdk.ChaincodeLanguage;
import org.hyperledger.fabric.sdk.DeployRequest;
import org.hyperledger.fabric.sdk.Endpoint;
import org.hyperledger.fabric.sdk.FileKeyValStore;
import org.hyperledger.fabric.sdk.InvokeRequest;
import org.hyperledger.fabric.sdk.Member;
import org.hyperledger.fabric.sdk.Peer;
import org.hyperledger.fabric.sdk.QueryRequest;
import org.hyperledger.fabric.sdk.RegistrationRequest;
import org.hyperledger.fabric.sdk.helper.SDKUtil;
import org.hyperledger.protos.Fabric.Transaction;
import org.hyperledger.protos.Api.BlockCount;
import org.hyperledger.protos.Api.BlockNumber;
import org.hyperledger.protos.Fabric.Block;
import org.hyperledger.protos.Fabric.BlockchainInfo;
import org.hyperledger.protos.Fabric.PeersMessage;
import org.hyperledger.protos.OpenchainGrpc;
import org.hyperledger.protos.OpenchainGrpc.OpenchainBlockingStub;

import com.google.protobuf.Empty;

import org.hyperledger.fabric.sdk.ChainCodeResponse.Status;

public class App {

    private static final Log logger = LogFactory.getLog(App.class);
    private static final Boolean Confidential = true;
    private static final String CCID = "f9a4f277e2c790dc7641c9b520fc077a5a4e1238bb92420a5891ff4c979a8729";
    private static final String CCName = "sayaka";

    public static void main(String[] args) throws Exception {

        App app = new App();

        Chain chain = new Chain("dummy chain");
        chain.setKeyValStore(new FileKeyValStore("D:/test.properties"));
        chain.setMemberServicesUrl("grpc://ayaka.cc:7054", null);
        chain.addPeer("grpc://ayaka.cc:4000", null);

        // app.Enroll(chain);

        // app.Registar(chain);

        // app.Deploy(chain);

        // app.Invoke(chain, "summary");
        // app.Invoke(chain, "subtraction");

        // app.Query(chain);

        // app.Odds(chain);
    }

    public void Enroll(Chain chain) throws Exception {
        Member admin = chain.getMember("admin");
        if (!admin.isEnrolled()) {
            admin.enroll("Xurw3yU9zI0l");
        }
        chain.setRegistrar(admin);
        logger.info("admin Enroll Success");
    }

    public Member Registar(Chain chain) throws Exception {
        Member minami = chain.getMember("minami");
        if (!minami.isRegistered()) {
            RegistrationRequest req = new RegistrationRequest();
            req.setAffiliation("institution_a");
            req.setEnrollmentID("minami");
            req.setRoles(new ArrayList<>(Arrays.asList("client")));
            minami.register(req);
            minami.enroll(minami.getEnrollmentSecret());
            logger.info("Registar Success");
        }
        logger.info("minami Enroll Success");
        return minami;
    }

    public void Deploy(Chain chain) throws Exception {
        DeployRequest request = new DeployRequest();
        request.setChaincodePath("D:/NaN/Demo/src/chaincode/target/");
        request.setArgs(new ArrayList<>(Arrays.asList("init", "3", "2")));
        request.setChaincodeName(CCName);
        request.setChaincodeLanguage(ChaincodeLanguage.JAVA);
        request.setConfidential(Confidential);

        Member minami = chain.getMember("minami");
        ChainCodeResponse resp = minami.deploy(request);
        logger.info("CCID: " + resp.getChainCodeID());
        logger.info("Txid: " + resp.getTransactionID());
    }

    public void Invoke(Chain chain, String oper) throws Exception {
        InvokeRequest request = new InvokeRequest();
        request.setArgs(new ArrayList<>(Arrays.asList(oper, "3", "2")));
        request.setChaincodeID(CCID);
        request.setChaincodeName(SDKUtil.generateUUID());
        request.setChaincodeLanguage(ChaincodeLanguage.JAVA);
        request.setConfidential(Confidential);
        Member minami = chain.getMember("minami");
        ChainCodeResponse resp = minami.invoke(request);
        if (resp.getStatus() != Status.SUCCESS) {
            throw new Exception(resp.getMessage());
        }
        logger.info("TXID: " + resp.getTransactionID());
    }

    public void Query(Chain chain) throws Exception {
        QueryRequest request = new QueryRequest();
        request.setArgs(new ArrayList<>(Arrays.asList("query")));
        request.setChaincodeID(CCID);
        request.setChaincodeName(CCName);
        request.setChaincodeLanguage(ChaincodeLanguage.JAVA);
        request.setConfidential(Confidential);
        Member minami = chain.getMember("minami");
        ChainCodeResponse resp = minami.query(request);
        if (resp.getStatus() != Status.SUCCESS) {
            throw new Exception(resp.getMessage());
        }
        logger.info("Result: " + resp.getMessage());
    }

    public void Odds(Chain chain) throws Exception {

        Peer peer = chain.getPeers().get(new Random().nextInt(chain.getPeers().size()));

        Endpoint ep = new Endpoint(peer.getUrl(), null);

        OpenchainBlockingStub client = OpenchainGrpc.newBlockingStub(ep.getChannelBuilder().build());

        BlockchainInfo info = client.getBlockchainInfo(Empty.newBuilder().build());
        System.out.println("current height is :\n" + info.getHeight());

        BlockCount count = client.getBlockCount(Empty.newBuilder().build());
        System.out.println("blocks count is :\n" + count.getCount());

        Block block = client.getBlockByNumber(BlockNumber.newBuilder().setNumber(2).build());
        System.out.println("blocks is :\n" + block.toString());

        PeersMessage peers = client.getPeers(Empty.newBuilder().build());
        System.out.println("peers is :\n" + peers.toString());

        Transaction transaction = client
                .getTransactionByID(Transaction.newBuilder().setTxid("39cd8a7c-da98-4a83-9467-4f9b5242e744").build());
        System.out.println("transaction is :\n" + transaction.toString());
    }
}
