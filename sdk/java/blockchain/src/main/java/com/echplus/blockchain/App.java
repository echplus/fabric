package com.echplus.blockchain;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Random;

import org.apache.commons.codec.binary.Base64;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.bouncycastle.crypto.digests.SHA256Digest;
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
import org.bouncycastle.util.encoders.Hex;

import com.google.protobuf.Empty;

import org.hyperledger.fabric.sdk.ChainCodeResponse.Status;

public class App {

    private static final Log logger = LogFactory.getLog(App.class);
    private static final Boolean Confidential = true;
    private static final String CCID = "9869536a90e844f9a0f9a74388432621334eede79432326e54e4c97d7a29bd44";
    private static final String CCName = "sayaka";

    public static void main(String[] args) throws Exception {

        byte[] raw = SDKUtil.hash("Hello World".getBytes(), new SHA256Digest());
        System.out.println(new String(Hex.encode(raw)));
        System.out.println(Base64.encodeBase64String(raw));

        App app = new App();

        Chain chain = new Chain("dummy chain");
        chain.setKeyValStore(new FileKeyValStore("D:/test.properties"));
        chain.setMemberServicesUrl("grpc://192.168.25.5:7054", null);
        chain.addPeer("grpc://192.168.25.5:4000", null);

        app.Enroll(chain);

        app.Registar(chain);

        // app.Deploy(chain);

        // app.Invoke(chain, new String[] { "summary", "3", "2" });
        // app.Invoke(chain, new String[] { "subtraction", "3", "2" });
        // app.Invoke(chain, new String[] { "err" });

        // app.Invoke(chain, new String[] { "newUser", "yibuki", "20", "nerv" });
        // app.Invoke(chain, new String[] { "newUser", "katsuraki", "20", "eproject" });
        // app.Invoke(chain, new String[] { "newUser", "ruizi", "20", "nerv" });

        // app.Invoke(chain, new String[] { "modifyUser", "katsuraki", "30", "nerv" });

        // app.Invoke(chain, new String[] { "removeUser", "katsuraki" });

        // app.Query(chain,new String[] { "kv" });

        // app.Query(chain, new String[] { "queryUser", "yibuki" });
        // app.Query(chain, new String[] { "queryUser", "katsuraki" });
        // app.Query(chain, new String[] { "queryUser", "ruizi" });

        // app.Odds(chain, "42721d4f-7428-43c8-8028-31549b0c3a1d");

        /** call golang chaincode begin **/
        // app.PutValue(chain,"5000","0");
        // app.TotalSize(chain);
        /** call golang chaincode end **/
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
        request.setChaincodePath("D:/repo/src/github.com/amamina/fabric/sdk/java/chaincode/target/");
        request.setArgs(new ArrayList<>(Arrays.asList("init", "index1", "value1")));
        request.setChaincodeName(CCName);
        request.setChaincodeLanguage(ChaincodeLanguage.JAVA);
        request.setConfidential(Confidential);

        Member minami = chain.getMember("minami");
        ChainCodeResponse resp = minami.deploy(request);
        logger.info("CCID: " + resp.getChainCodeID());
        logger.info("Txid: " + resp.getTransactionID());
        logger.info("Message: " + resp.getMessage());
        logger.info("Status: " + resp.getStatus());
    }

    public void Invoke(Chain chain, String[] args) throws Exception {
        InvokeRequest request = new InvokeRequest();
        request.setArgs(new ArrayList<>(Arrays.asList(args)));
        request.setChaincodeID(CCID);
        request.setChaincodeName(SDKUtil.generateUUID());
        request.setChaincodeLanguage(ChaincodeLanguage.JAVA);
        request.setConfidential(Confidential);
        Member minami = chain.getMember("minami");
        ChainCodeResponse resp = minami.invoke(request);
        if (resp.getStatus() != Status.SUCCESS) {
            throw new Exception(resp.getMessage());
        }
        logger.info("Txid: " + resp.getTransactionID());
        logger.info("Message: " + resp.getMessage());
        logger.info("Status: " + resp.getStatus());
    }

    public void Query(Chain chain, String[] args) throws Exception {
        QueryRequest request = new QueryRequest();
        request.setArgs(new ArrayList<>(Arrays.asList(args)));
        request.setChaincodeID(CCID);
        request.setChaincodeName(CCName);
        request.setChaincodeLanguage(ChaincodeLanguage.JAVA);
        request.setConfidential(Confidential);
        Member minami = chain.getMember("minami");
        ChainCodeResponse resp = minami.query(request);
        if (resp.getStatus() != Status.SUCCESS) {
            throw new Exception(resp.getMessage());
        }
        logger.info("Txid: " + resp.getTransactionID());
        logger.info("Message: " + resp.getMessage());
        logger.info("Status: " + resp.getStatus());
    }

    public void Odds(Chain chain, String txid) throws Exception {

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

        Transaction transaction = client.getTransactionByID(Transaction.newBuilder().setTxid(txid).build());
        System.out.println("transaction is :\n" + transaction.toString());
    }

    public void PutValue(Chain chain, String total, String start) throws Exception {
        InvokeRequest request = new InvokeRequest();
        request.setArgs(new ArrayList<>(Arrays.asList("putValue", total, start)));
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

    public void TotalSize(Chain chain) throws Exception {
        QueryRequest request = new QueryRequest();
        request.setArgs(new ArrayList<>(Arrays.asList("totalSize")));
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
}
