package com.echplus.chaincode;

import java.util.ArrayList;
import java.util.List;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.hyperledger.fabric.sdk.shim.ChaincodeBase;
import org.hyperledger.fabric.sdk.shim.ChaincodeStub;
import org.hyperledger.protos.TableProto.Column;
import org.hyperledger.protos.TableProto.ColumnDefinition;
import org.hyperledger.protos.TableProto.Row;

import com.google.protobuf.InvalidProtocolBufferException;

public class App extends ChaincodeBase {

    private static final Logger logger = LogManager.getLogger(App.class);

    private static final String Ok = null;

    public static void main(String[] args) throws Exception {
        logger.info("starting");
        new App().start(args);
    }

    @Override
    // deploy and invoke will call this
    public String run(ChaincodeStub stub, String function, String[] args) {
        logger.info("In run, function:" + function);
        if (args.length > 0) {
            for (int i = 0; i < args.length; i++) {
                logger.info("args[" + i + "] :  " + args[i]);
            }
        }

        List<Column> cols = null;

        switch (function) {
        case "init":
            try {
                stub.createTable("dummyTable", new ArrayList<ColumnDefinition>() {
                    private static final long serialVersionUID = -646985333253068937L;
                    {
                        add(ColumnDefinition.newBuilder().setName("name").setKey(true)
                                .setType(ColumnDefinition.Type.STRING).build());
                        add(ColumnDefinition.newBuilder().setName("age").setKey(false)
                                .setType(ColumnDefinition.Type.STRING).build());
                        add(ColumnDefinition.newBuilder().setName("agent").setKey(false)
                                .setType(ColumnDefinition.Type.STRING).build());
                    }
                });
            } catch (Exception e) {
                logger.error(e);
                e.printStackTrace();
                return e.getMessage();
            }
            break;

        case "summary":
            stub.putState("result", String.valueOf(Integer.parseInt(args[0]) + Integer.parseInt(args[1])));
            break;

        case "subtraction":
            stub.putState("result", String.valueOf(Integer.parseInt(args[0]) - Integer.parseInt(args[1])));
            break;

        case "newUser":
            cols = new ArrayList<Column>() {
                private static final long serialVersionUID = 795900172831697917L;
                {
                    add(Column.newBuilder().setString(args[0]).build());
                    add(Column.newBuilder().setString(args[1]).build());
                    add(Column.newBuilder().setString(args[2]).build());
                }
            };

            try {
                stub.insertRow("dummyTable", Row.newBuilder().addAllColumns(cols).build());
            } catch (Exception e) {
                logger.error(e);
                e.printStackTrace();
                return e.getMessage();
            }
            break;

        case "modifyUser":
            cols = new ArrayList<Column>() {
                private static final long serialVersionUID = 795900172831697917L;
                {
                    add(Column.newBuilder().setString(args[0]).build());
                    add(Column.newBuilder().setString(args[1]).build());
                    add(Column.newBuilder().setString(args[2]).build());
                }
            };

            try {
                stub.replaceRow("dummyTable", Row.newBuilder().addAllColumns(cols).build());
            } catch (Exception e) {
                logger.error(e);
                e.printStackTrace();
                return e.getMessage();
            }
            break;

        case "removeUser":
            cols = new ArrayList<Column>() {
                private static final long serialVersionUID = 795900172831697917L;
                {
                    add(Column.newBuilder().setString(args[0]).build());
                }
            };

            try {
                stub.deleteRow("dummyTable", cols);
            } catch (Exception e) {
                logger.error(e);
                e.printStackTrace();
                return e.getMessage();
            }
            break;

        case "err":
            return "dummy error";

        default:
            logger.error("No matching case for function:" + function);
            return "No matching case for function:" + function;
        }

        return Ok;
    }

    @Override
    public String query(ChaincodeStub stub, String function, String[] args) {
        logger.info("In query, function:" + function);
        if (args.length > 0) {
            for (int i = 0; i < args.length; i++) {
                logger.info("args[" + i + "] :  " + args[i]);
            }
        }

        switch (function) {
        case "kv":
            String result = stub.getState("result");
            logger.info("result: " + result);
            return result;

        case "queryUser":

            List<Column> key = new ArrayList<Column>() {
                private static final long serialVersionUID = 795900172831697917L;
                {
                    add(Column.newBuilder().setString(args[0]).build());
                }
            };

            try {
                Row row = stub.getRow("dummyTable", key);
                if (row.getSerializedSize() > 0) {
                    return row.getColumns(0).getString() + "\t" + row.getColumns(1).getString() + "\t"
                            + row.getColumns(2).getString();
                } else {
                    return "not found";
                }

            } catch (InvalidProtocolBufferException e) {
                logger.error(e);
                e.printStackTrace();
                return e.getMessage();
            }

        default:
            logger.error("No matching case for function:" + function);
            return "No matching case for function:" + function;
        }
    }

    @Override
    public String getChaincodeID() {
        return "hello";
    }

}
