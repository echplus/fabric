package com.echplus.chaincode;

import org.hyperledger.fabric.sdk.shim.ChaincodeBase;
import org.hyperledger.fabric.sdk.shim.ChaincodeStub;

/**
 *
 */
public class App extends ChaincodeBase {

    public static void main(String[] args) throws Exception {
        System.out.println("starting");
        new App().start(args);
    }

    @Override
    // deploy and invoke will call this
    public String run(ChaincodeStub stub, String function, String[] args) {
        System.out.println("In run, function:" + function);
        if (args.length > 0) {
            for (int i = 0; i < args.length; i++) {
                System.out.println("args[" + i + "] :  " + args[i]);
            }
        }

        switch (function) {
        case "init":
            stub.putState("a", args[0]);
            stub.putState("b", args[1]);
            break;
        case "summary":
            stub.putState("result", String.valueOf(Integer.parseInt(args[0]) + Integer.parseInt(args[1])));
            break;
        case "subtraction":
            stub.putState("result", String.valueOf(Integer.parseInt(args[0]) - Integer.parseInt(args[1])));
            break;
        default:
            System.err.println("No matching case for function:" + function);
            return "No matching case for function:" + function;
        }

        return stub.getState("result");
    }

    @Override
    public String query(ChaincodeStub stub, String function, String[] args) {
        System.out.println("In query, function:" + function);
        String result = stub.getState("result");
        System.out.println("result: " + result);
        return result;
    }

    @Override
    public String getChaincodeID() {
        return "hello";
    }

}
