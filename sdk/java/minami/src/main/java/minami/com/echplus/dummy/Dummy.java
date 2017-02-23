package minami.com.echplus.dummy;

import java.util.Map;
import java.util.HashMap;
import java.util.List;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.google.gson.Gson;

import minami.com.echplus.domain.User;
import minami.com.echplus.mapper.UserMapper;

@RestController
public class Dummy {

    private static final Logger logger = LogManager.getLogger(Dummy.class);

    @Autowired
    private UserMapper userMapper;

    @RequestMapping(value = "/get", method = { RequestMethod.GET, RequestMethod.POST })
    public String hello(@RequestParam("row") int row) {

        logger.info(row);

        List<User> users = userMapper.getUsers0(row);
        for (User user : users) {
            logger.info(user.getRow() + "\t" + user.getId() + "\t" + user.getEnrollmentId() + "\t" + user.getToken());
        }

        logger.info("--------------------------------------");
        users = userMapper.getUsers1(row);
        for (User user : users) {
            logger.info(user.getRow() + "\t" + user.getId() + "\t" + user.getEnrollmentId() + "\t" + user.getToken());
        }

        return "Hello World! Minami " + row;
    }

    @SuppressWarnings("unchecked")
    @RequestMapping(value = "/post", method = RequestMethod.POST)
    public String post(@RequestParam("foo") String foo, @RequestBody String body) {

        logger.info(body);

        Map<String, String> map = (Map<String, String>) new Gson().fromJson(body,
                new HashMap<String, String>().getClass());

        for (String key : map.keySet()) {
            logger.info(key + "  " + map.get(key));
        }

        return "Hello World! Minami " + foo + "  " + body;
    }

}
