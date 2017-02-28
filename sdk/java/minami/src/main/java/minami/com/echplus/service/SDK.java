package minami.com.echplus.service;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Properties;

import javax.annotation.PostConstruct;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.hyperledger.fabric.sdk.Chain;
import org.hyperledger.fabric.sdk.ChainCodeResponse;
import org.hyperledger.fabric.sdk.ChaincodeLanguage;
import org.hyperledger.fabric.sdk.DeployRequest;
import org.hyperledger.fabric.sdk.FileKeyValStore;
import org.hyperledger.fabric.sdk.InvokeRequest;
import org.hyperledger.fabric.sdk.Member;
import org.hyperledger.fabric.sdk.QueryRequest;
import org.hyperledger.fabric.sdk.RegistrationRequest;
import org.hyperledger.fabric.sdk.ChainCodeResponse.Status;
import org.hyperledger.fabric.sdk.helper.SDKUtil;
import org.springframework.stereotype.Service;

@Service
public class SDK {

    private static final Logger logger = LogManager.getLogger(SDK.class);

    private static String keyValStore = null;
    private static String ccid = null;
    private static String cname = null;
    private static String membersrvc = null;
    private static String peer = null;
    private static boolean confidential = true;

    private static final Chain chain = new Chain("dummy chain");

    private static final Properties properties = new Properties();

    @PostConstruct
    private void init() throws Exception {
        properties.load(SDK.class.getClassLoader().getResourceAsStream("/application.properties"));
        keyValStore = properties.getProperty("keyValStore");
        ccid = properties.getProperty("ccid");
        cname = properties.getProperty("cname");
        membersrvc = properties.getProperty("membersrvc");
        peer = properties.getProperty("peer");
        confidential = Boolean.parseBoolean(properties.getProperty("confidential"));

        chain.setKeyValStore(new FileKeyValStore(keyValStore));
        chain.setMemberServicesUrl(membersrvc, null);
        chain.addPeer(peer, null);
    }

    public void Enroll(String username, String password) throws Exception {
        Member admin = chain.getMember(username);
        if (!admin.isEnrolled()) {
            admin.enroll(password);
        }
        chain.setRegistrar(admin);
        logger.info(username + " Enroll Success");
    }

    public String Registar(String user, String affiliation, String[] roles) throws Exception {
        Member minami = chain.getMember(user);
        if (!minami.isRegistered()) {
            RegistrationRequest req = new RegistrationRequest();
            req.setAffiliation(affiliation);
            req.setEnrollmentID(user);
            req.setRoles(new ArrayList<>(Arrays.asList(roles)));
            minami.register(req);
        }
        return minami.getEnrollmentSecret();
    }

    public String Deploy(String username, String path, String[] args) throws Exception {
        DeployRequest request = new DeployRequest();
        request.setChaincodePath(path);
        request.setArgs(new ArrayList<>(Arrays.asList(args)));
        request.setChaincodeName(cname);
        request.setChaincodeLanguage(ChaincodeLanguage.JAVA);
        request.setConfidential(confidential);

        Member minami = chain.getMember(username);
        ChainCodeResponse resp = minami.deploy(request);
        logger.info("Deploy CCID: " + resp.getChainCodeID());
        logger.info("Deploy Txid: " + resp.getTransactionID());
        logger.info("Deploy Message: " + resp.getMessage());
        logger.info("Deploy Status: " + resp.getStatus());

        return resp.getChainCodeID();
    }

    public String Invoke(String username, String[] args) throws Exception {
        InvokeRequest request = new InvokeRequest();
        request.setArgs(new ArrayList<>(Arrays.asList(args)));
        request.setChaincodeID(ccid);
        request.setChaincodeName(SDKUtil.generateUUID());
        request.setChaincodeLanguage(ChaincodeLanguage.JAVA);
        request.setConfidential(confidential);
        Member minami = chain.getMember(username);
        ChainCodeResponse resp = minami.invoke(request);
        if (resp.getStatus() != Status.SUCCESS) {
            throw new Exception(resp.getMessage());
        }
        logger.info("Invoke Txid: " + resp.getTransactionID());
        logger.info("Invoke Message: " + resp.getMessage());
        logger.info("Invoke Status: " + resp.getStatus());

        return resp.getTransactionID();
    }

    public String Query(String username, String[] args) throws Exception {
        QueryRequest request = new QueryRequest();
        request.setArgs(new ArrayList<>(Arrays.asList(args)));
        request.setChaincodeID(ccid);
        request.setChaincodeName(cname);
        request.setChaincodeLanguage(ChaincodeLanguage.JAVA);
        request.setConfidential(confidential);
        Member minami = chain.getMember(username);
        ChainCodeResponse resp = minami.query(request);
        if (resp.getStatus() != Status.SUCCESS) {
            throw new Exception(resp.getMessage());
        }
        logger.info("Query Txid: " + resp.getTransactionID());
        logger.info("Query Message: " + resp.getMessage());
        logger.info("Query Status: " + resp.getStatus());

        return resp.getMessage();
    }

}
