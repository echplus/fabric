package com.echplus.chaincode;

import java.util.ArrayList;
import java.util.List;

import org.hyperledger.fabric.sdk.shim.ChaincodeBase;
import org.hyperledger.fabric.sdk.shim.ChaincodeStub;
import org.hyperledger.protos.TableProto.Column;
import org.hyperledger.protos.TableProto.ColumnDefinition;
import org.hyperledger.protos.TableProto.Row;

import com.google.protobuf.InvalidProtocolBufferException;

/**
 *
 */
public class App extends ChaincodeBase {

	private static final String Ok = null;

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
			try {
				stub.createTable("dummyTable", new ArrayList<ColumnDefinition>() {
					private static final long serialVersionUID = -646985333253068937L;
					{
						add(ColumnDefinition.newBuilder().setName("id").setKey(true)
								.setType(ColumnDefinition.Type.STRING).build());
						add(ColumnDefinition.newBuilder().setName("name").setKey(false)
								.setType(ColumnDefinition.Type.STRING).build());
					}
				});
			} catch (Exception e) {
				System.err.println(e);
				e.printStackTrace();
				return e.getMessage();
			}

			List<Column> cols = new ArrayList<Column>() {
				private static final long serialVersionUID = 795900172831697917L;
				{
					add(Column.newBuilder().setString(args[0]).build());
					add(Column.newBuilder().setString(args[1]).build());
				}
			};

			try {
				stub.insertRow("dummyTable", Row.newBuilder().addAllColumns(cols).build());
			} catch (Exception e) {
				System.err.println(e);
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

		case "err":
			return "dummy error";

		default:
			System.err.println("No matching case for function:" + function);
			return "No matching case for function:" + function;
		}

		return Ok;
	}

	@Override
	public String query(ChaincodeStub stub, String function, String[] args) {
		System.out.println("In query, function:" + function);
		if (args.length > 0) {
			for (int i = 0; i < args.length; i++) {
				System.out.println("args[" + i + "] :  " + args[i]);
			}
		}

		switch (function) {
		case "kv":
			String result = stub.getState("result");
			System.out.println("result: " + result);
			return result;

		case "table":

			List<Column> key = new ArrayList<Column>() {
				private static final long serialVersionUID = 795900172831697917L;
				{
					add(Column.newBuilder().setString(args[0]).build());
				}
			};

			try {
				Row row = stub.getRow("dummyTable", key);
				if (row.getSerializedSize() > 0) {
					return row.getColumns(1).getString();
				} else {
					return "not found";
				}

			} catch (InvalidProtocolBufferException e) {
				System.err.println(e);
				e.printStackTrace();
				return e.getMessage();
			}

		default:
			System.err.println("No matching case for function:" + function);
			return "No matching case for function:" + function;
		}
	}

	@Override
	public String getChaincodeID() {
		return "hello";
	}

}
