package minami.com.echplus.service;

import java.util.List;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import minami.com.echplus.domain.UserDomain;
import minami.com.echplus.mapper.UserMapper;

@Service
public class DummySvc {

    private static final Logger logger = LogManager.getLogger(DummySvc.class);

    @Autowired
    private UserMapper userMapper;

    public void query(int row) {

        List<UserDomain> users = userMapper.getUsers0(row);
        for (UserDomain user : users) {
            logger.info(user.getRow() + "\t" + user.getId() + "\t" + user.getEnrollmentId() + "\t" + user.getToken());
        }

        logger.info("--------------------------------------");
        users = userMapper.getUsers1(row);
        for (UserDomain user : users) {
            logger.info(user.getRow() + "\t" + user.getId() + "\t" + user.getEnrollmentId() + "\t" + user.getToken());
        }

    }

    @Transactional
    public int insert() {

        UserDomain user0 = new UserDomain();
        user0.setId("minami");
        userMapper.newUser(user0);

        // UserDomain user1 = new UserDomain();
        // user1.setRow(7);// will panic
        // user1.setId("toma");
        // userMapper.newUser(user1);

        return 1;
    }

    @Scheduled(cron = "0/60 * * * * *")
    public void schedule() {
        logger.info("------dummy schedule job------");
    }

}
