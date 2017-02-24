package minami.com.echplus.domain;

import java.io.Serializable;

public class UserDomain implements Serializable {

    private static final long serialVersionUID = 5393908299329541777L;
    private int row;
    private String id;
    private String enrollmentId;
    private String token;

    public int getRow() {
        return row;
    }

    public void setRow(int row) {
        this.row = row;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getEnrollmentId() {
        return enrollmentId;
    }

    public void setEnrollmentId(String enrollmentId) {
        this.enrollmentId = enrollmentId;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }

}
