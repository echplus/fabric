package minami.com.echplus.controller;

import java.util.Map;
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

import minami.com.echplus.service.DummySvc;

@RestController
public class DummyController {

    private static final Logger logger = LogManager.getLogger(DummyController.class);

    @Autowired
    private DummySvc svc;

    @RequestMapping(value = "/get", method = { RequestMethod.GET, RequestMethod.POST })
    public String hello(@RequestParam("row") int row) {

        logger.info(row);

        svc.query(row);
        logger.info("add " + svc.insert() + " users");

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
