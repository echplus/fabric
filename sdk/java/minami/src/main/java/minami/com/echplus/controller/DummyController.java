package minami.com.echplus.controller;

import java.util.Map;
import java.util.UUID;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import java.util.HashMap;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.google.gson.Gson;

import minami.com.echplus.service.Session;
import minami.com.echplus.service.DummySvc;

@RestController
public class DummyController {

    private static final Logger logger = LogManager.getLogger(DummyController.class);

    @Autowired
    private DummySvc svc;

    @Autowired
    private Session session;

    @RequestMapping(value = "/get", method = { RequestMethod.GET, RequestMethod.POST })
    public String hello(@RequestParam("row") int row, HttpServletRequest request, HttpServletResponse response) {

        logger.info(row);

        svc.query(row);

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

    @RequestMapping(value = "/login")
    public String login(@RequestParam("name") String name, @RequestParam("pwd") String pwd, HttpServletRequest request,
            HttpServletResponse response) {

        String token = null;

        if (request.getCookies() != null) {
            token = request.getCookies()[0].getValue();
        }

        logger.info(name + "   " + pwd + "   " + token);

        if (token == null) {
            token = UUID.randomUUID().toString();
        }

        if (session.Exists(token)) {
            logger.info("token exists");
        } else {
            session.Update(token);
            logger.info("new token");
        }

        Cookie cookie = new Cookie("token", token);
        cookie.setMaxAge(session.GeTimeToLiveSeconds());
        response.addCookie(cookie);

        return "Hello World! Minami " + name + "   " + pwd;
    }

}
