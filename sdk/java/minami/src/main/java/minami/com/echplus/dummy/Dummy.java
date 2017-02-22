package minami.com.echplus.dummy;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class Dummy {

    private static final Logger logger = LogManager.getLogger(Dummy.class);

    @RequestMapping(value = "/hello", method = { RequestMethod.GET, RequestMethod.POST })
    public String hello(@RequestParam("foo") String foo) {

        logger.info(foo);
        return "Hello World! Minami " + foo;
    }

}
